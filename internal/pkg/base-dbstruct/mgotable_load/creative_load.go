package mgotable_load

import (
	"encoding/json"
	"errors"

	"github.com/Mintegral-official/mtggokit/bifrost/container"
	"github.com/Mintegral-official/mtggokit/bifrost/streamer"
	"gitlab.mobvista.com/voyager/common/enum"
	"gopkg.in/mgo.v2/bson"

	"gitlab.mobvista.com/voyager/mgoextractor/mgoconfig"
	"gitlab.mobvista.com/voyager/mgoextractor/mgologger"
	"gitlab.mobvista.com/voyager/mgoextractor/mgotables"
)

func DefaultCreativeCfg(parser *CreativeParser) mgoconfig.BiFrostConfig {
	return mgoconfig.BiFrostConfig{
		BaseInterval: 2592000,                      //数据表全量更新周期
		IncInterval:  60,                           //数据表增量更新周期
		URI:          "",                           //数据表mongo url, 因各集群地址不同，需要外部输入
		DB:           mgoconfig.NewAdnDB,           //数据表存储数据库
		Collection:   mgoconfig.CreativeCollection, //数据表name
		DataParser:   parser,
	}
}

//--------------------------- BiFrostExtractor ------------------------------------

func NewCreativeParser() *CreativeParser {
	return &CreativeParser{
		Parser: func(creative *mgotables.Creative) bool { return true },
	}
}

type CreativeParser struct {
	MaxUpdated int64

	Parser func(*mgotables.Creative) bool // 提供给外部的 Parser 方法，用于 解析非标字段 + 判断当前doc是否可用状态，返回 true 保留记录，返回 false 删除记录。
}

func (parser *CreativeParser) Parse(data []byte, userData interface{}) []streamer.ParserResult {
	table := &mgotables.Creative{}

	if err := bson.Unmarshal(data, &table); err != nil {
		mgologger.MgoLog.Error("MgoExtractor Parse Creative error! bson.Unmarshal error =" + err.Error())
		return []streamer.ParserResult{{container.DataModeDel, nil, nil, errors.New("Parse Creative error ")}}
	}

	if table.Updated > parser.MaxUpdated {
		parser.MaxUpdated = table.Updated
	}

	key := GetStoreKey(table.CampaignId, table.PackageName, table.CountryCode)

	table.Parser()
	reserved := parser.Parser(table)

	// - CreativeIndex -
	UpdateCrIndex(table)

	if reserved && table.Status == enum.MongoStatus_ACTIVE {
		body := ""
		if mgoconfig.DebugMgoExtractor {
			j, _ := json.Marshal(table)
			body = string(j)
		}
		mgologger.MgoLog.Info("MgoExtractor Add Creative [key]=", key, body)
	} else {
		mgologger.MgoLog.Info("MgoExtractor Delete Creative [key]=", key)
		return []streamer.ParserResult{{container.DataModeDel, container.StrKey(key), nil, nil}}
	}

	return []streamer.ParserResult{{container.DataModeAdd, container.StrKey(key), table, nil}}
}

func GetCreativeDetail(campId int64, pkgName string, country string) (*mgotables.Creative, error) {
	key := GetStoreKey(campId, pkgName, country)
	value, err := BiFrostIndex.Get(mgoconfig.CreativeCollection, container.StrKey(key))
	if err != nil {
		return nil, err
	}
	creative, ok := value.(*mgotables.Creative)
	if !ok {
		return nil, errors.New("Creative streamer value type bad ")
	}
	return creative, nil
}
