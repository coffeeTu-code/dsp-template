package mgotable_load

import (
	"context"
	"errors"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"dsp-template/api/backend"
	"dsp-template/api/enum"
	"dsp-template/internal/pkg/base-dbstruct/mgologger"
	"dsp-template/internal/pkg/base-dbstruct/mgometrics"
	"dsp-template/internal/pkg/base-dbstruct/mgotables"
	bifrost "dsp-template/pkg2/helper-bifrost"
	"dsp-template/pkg2/helper-bifrost/container"
)

var (
	BiFrostIndex      = bifrost.NewBifrost()
	DebugMgoExtractor bool
)

var (
	NewAdnDB = "new_adn"

	CampaignCollection        = "campaign"
	CreativeCollection        = "creative"
	UnitCollection            = "unit"
	AppPackageMtgIdCollection = "app_package_mtg_id"
	CreativePackageCollection = "creative_package"
	ConfigCollection          = "config"
	AppCollection             = "app"
)

var (
	DspAuditDB = "dsp_audit"

	CampaignAdxCreativeAuditCollection = "campaign_adx_creative_audit"
)

func NewBiFrostStreamer(biCfg BiFrostConfig) (*bifrost.MongoStreamer, error) {
	if biCfg.QueryAll == nil || biCfg.QueryInc == nil {
		return nil, errors.New("no set query sql")
	}
	if biCfg.Projection == nil {
		biCfg.Projection = func() bson.M { return nil }
	}
	ms, err := bifrost.NewMongoStreamer(&bifrost.MongoStreamerCfg{
		BaseInterval: biCfg.BaseInterval,
		IncInterval:  biCfg.IncInterval,
		MongoConfig:  biCfg.MongoConfig,
		Name:         biCfg.MongoConfig.Collection,

		BaseParser:   biCfg.DataParser,
		IncParser:    biCfg.DataParser,
		OnBeforeBase: func(userData interface{}) interface{} { return biCfg.QueryAll() },
		OnBeforeInc:  func(userData interface{}) interface{} { return biCfg.QueryInc() },
		FindOpt:      options.Find().SetProjection(biCfg.Projection()),

		IsSync:       true,
		Logger:       mgologger.MgoLog,
		OnFinishBase: OnFinish,
		OnFinishInc:  OnFinish,
	})

	if ms == nil || err != nil {
		mgologger.MgoLog.Error("Init streamer error! table=", biCfg.MongoConfig, "error=", err)
		return nil, err
	}

	switch biCfg.MongoConfig.Collection {
	case AppPackageMtgIdCollection:
		ms.SetContainer(CreateMtgIdContainer(20, 1))
	default:
		ms.SetContainer(container.CreateBlockingMapContainer(20, 1))
	}
	ctx, _ := context.WithCancel(context.Background())

	if err := ms.UpdateData(ctx); err != nil {
		mgologger.MgoLog.Error("update mongo streamer error! table=", biCfg.MongoConfig, "error=", err)
		return nil, err
	}

	return ms, nil
}

// metrics
func OnFinish(streamer bifrost.Streamer) {
	if streamer != nil && streamer.GetInfo() != nil {
		mgometrics.SetMetrics(mgometrics.DataNumber, mgometrics.Labels{TableName: streamer.GetInfo().Name}, float64(streamer.GetInfo().TotalNum))
		mgometrics.DataUpdatedTime(streamer.GetInfo().Name)

		switch streamer.GetInfo().Name {
		case CreativeCollection:

			var crNum = map[int]map[int]int{} // key=creative type, key=format type
			for _, val := range CreativeIndex.CreativeKeyIndexMap {
				for _, cr := range val {
					if crNum[cr.CreativeType] == nil {
						crNum[cr.CreativeType] = map[int]int{}
					}
					crNum[cr.CreativeType][cr.FormatType]++
				}
			}
			for crtype, val := range crNum {
				for formattype, num := range val {
					mgometrics.SetMetrics(mgometrics.CreativeNum, mgometrics.Labels{CreativeType: strconv.Itoa(crtype), FormatType: enum.FormatType(formattype).String()}, float64(num))
				}
			}

		case CampaignAdxCreativeAuditCollection:

			groupCrMap := map[string]map[string]map[string]int{} // key=adx, key=specid, key=status, value=num
			streamer.GetContainer().Range(func(key, value interface{}) bool {
				switch vt := value.(type) {
				case *mgotables.CampaignAdxCreativeAudit:
					for adx, val := range vt.AdxCreativeSpecGroupCreativesIdx {
						if groupCrMap[adx] == nil {
							groupCrMap[adx] = map[string]map[string]int{}
						}
						for spec, groups := range val.CreativeSpecMap {
							if groupCrMap[adx][spec] == nil {
								groupCrMap[adx][spec] = map[string]int{}
							}
							groupCrMap[adx][spec]["matched"] += len(groups.CreativeGroup)
							groupCrMap[adx][spec]["no activate"] += len(groups.NoActiveGroup)
							groupCrMap[adx][spec]["loss creativeid"] += len(groups.NoCompleteGroup)
						}
					}
				}
				return true
			})

			for adx, val := range groupCrMap {
				for spec, statusmap := range val {
					for status, groupNum := range statusmap {
						if groupNum == 0 {
							continue
						}
						mgometrics.SetMetrics(mgometrics.GroupCreativeNum, mgometrics.Labels{Adx: adx, CreativeSpec: spec, Status: status}, float64(groupNum))
					}
				}
			}
		}
	}
}

func GetStoreKey(campId int64, pkgName string, country string) (key string) {
	if campId != 0 {
		return strconv.FormatInt(campId, 10) + "-" + country
	}
	return pkgName + "-" + country
}

func GetCampaignRecallKeys(camp *mgotables.Campaign) map[string]struct{} {
	recallKeys := make(map[string]struct{})
	platform := strconv.FormatInt(int64(camp.Platform), 10)
	for _, countryCode := range camp.CountryCode {
		recallKeys[platform+countryCode] = struct{}{}
	}
	return recallKeys
}

type BiFrostConfig struct {
	bifrost.DataParser //解析mongo表数据

	BaseInterval int `toml:"base_interval" json:"base_interval"`
	IncInterval  int `toml:"inc_interval" json:"inc_interval"`
	MongoConfig  backend.MongoConfig

	QueryAll, QueryInc, Projection func() bson.M
}
