package mgotables

import (
	"strconv"

	"dsp-template/api/enum"
	"dsp-template/internal/pkg/base-dbstruct/mgotables/helper"
)

type UnitInfo struct {
	UnitId  int64   `bson:"unitId,omitempty" json:"unitId"`
	Unit    Unit    `bson:"unit,omitempty" json:"unit"`
	Setting Setting `bson:"setting,omitempty" json:"setting"`
	AppId   int64   `bson:"appId,omitempty" json:"appId"`

	TryNewRate   float64                     `bson:"tryNewRate,omitempty" json:"tryNewRate"`
	Updated      int64                       `bson:"updated,omitempty" json:"updated"`
	TemplateConf map[string]TemplateConfInfo `bson:"templateConf,omitempty" json:"templateConf"`

	Ext interface{} // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
}

type SdkVersionInfo struct {
	Include []MinMax `bson:"include,omitempty" json:"include"`
	Exclude []MinMax `bson:"exclude,omitempty" json:"exclude"`
}

type MinMax struct {
	Min string `bson:"min,omitempty" json:"min"`
	Max string `bson:"max,omitempty" json:"max"`
}

type TemplateConfInfo struct {
	OsMax           int64          `bson:"osMax,omitempty" json:"osMax"`
	OsMin           int64          `bson:"osMin,omitempty" json:"osMin"`
	SdkVersion      SdkVersionInfo `bson:"sdkVersion,omitempty" json:"sdkVersion"`
	Orientation     int64          `bson:"orientation,omitempty" json:"orientation"`
	VideoTemplate   []TemplateInfo `bson:"videoTemplate,omitempty" json:"videoTemplate"`
	EndcardTemplate []TemplateInfo `bson:"endcardTemplate,omitempty" json:"endcardTemplate"`
}

type Unit struct {
	UnitID      int64 `bson:"unitId,omitempty" json:"unitId"`
	Status      int   `bson:"status,omitempty" json:"status"`
	AdType      int   `bson:"adType,omitempty" json:"adType"`
	Orientation int32 `bson:"orientation,omitempty" json:"orientation,omitempty"`

	//模版白名单 @sample=[9002001, 9002002]
	WhiteTemplate              []int                               `bson:"whiteTemplate,omitempty" json:"whiteTemplate"`
	whiteTemplateMap           map[string]bool                     //内部使用
	asUnitRecalledTemplateConf map[int64][]*AsUnitRecalledTemplate //内部使用

}
type SdkVersion struct {
	Min    int64
	Max    int64
	Prefix string
}

func (this *SdkVersion) Match(sdkPrefix *string, sdkVersionCode int64) bool {
	if *sdkPrefix != this.Prefix {
		return false
	}

	return sdkVersionCode >= this.Min && sdkVersionCode <= this.Max
}

type AsUnitRecalledTemplate struct {
	Orientation        int64
	IncludeSdkVersions []*SdkVersion
	ExcludeSdkVersions []*SdkVersion
	osMin              int64
	osMax              int64

	VideoTemplate   map[string]bool
	EndCardTemplate map[string]bool
}

type Setting struct {
	//网络状态 @sample=[4, 9]
	RecallNet []int `bson:"recallNet,omitempty" json:"recallNet"`
}

func getSdkVersion(minMax *MinMax) *SdkVersion {
	var maxSdkVersion int64
	var minSdkVersion int64
	var prefixSdk string

	if minMax.Max == "" {
		maxSdkVersion = int64(^uint64(0) >> 1)
	} else {
		maxSdkVersion = int64(helper.GetSdkVersionNum(minMax.Max))
		prefixSdk = helper.GetSdkVersionPrefix(minMax.Max)
	}

	if minMax.Min == "" {
		minSdkVersion = 0
	} else {
		minSdkVersion = int64(helper.GetSdkVersionNum(minMax.Min))
		prefixSdk = helper.GetSdkVersionPrefix(minMax.Min)
	}

	return &SdkVersion{Max: maxSdkVersion, Min: minSdkVersion, Prefix: prefixSdk}
}

func getTemplateId(templateInfo []TemplateInfo, templateType string) map[string]bool {

	var hitTemplateMapping *map[int][]string
	switch templateType {
	case "video":
		hitTemplateMapping = &enum.VideoTemplateTypeToVideoTemplateId
	case "endcard":
		hitTemplateMapping = &enum.EndcardTemplateTypeToEndcardTemplateId
	default:
		return nil
	}

	templateIds := map[string]bool{}
	for _, template := range templateInfo {
		if _, ok := (*hitTemplateMapping)[int(template.Type)]; ok {
			for _, templateId := range (*hitTemplateMapping)[int(template.Type)] {
				templateIds[templateId] = true
			}

		}

	}
	return templateIds
}

// ----------------------------------------------- Get Attr ------------------------------------------------------------

func (this *UnitInfo) GetWhiteTemplateMap() map[string]bool {
	if this.Unit.whiteTemplateMap == nil {
		return map[string]bool{}
	}
	return this.Unit.whiteTemplateMap
}

// ----------------------------------------------- Parser --------------------------------------------------------------

func (u *UnitInfo) Parser() {
	{
		if u.Unit.whiteTemplateMap == nil {
			u.Unit.whiteTemplateMap = make(map[string]bool, len(u.Unit.WhiteTemplate))
		}
		for _, templateId := range u.Unit.WhiteTemplate {
			u.Unit.whiteTemplateMap[strconv.Itoa(templateId)] = true
		}
	}

	//for as unit 召回模板，按orientation进行组织
	u.Unit.asUnitRecalledTemplateConf = map[int64][]*AsUnitRecalledTemplate{}
	//polaris_logger.Logger.Debug("config Unit [unitId]=", this.UnitId, "templateConf size is ", len(this.TemplateConf))

	for _, conf := range u.TemplateConf {
		var asUnitRecalledTemplate AsUnitRecalledTemplate

		asUnitRecalledTemplate.Orientation = conf.Orientation

		asUnitRecalledTemplate.osMax = conf.OsMax
		if asUnitRecalledTemplate.osMax == 0 {
			asUnitRecalledTemplate.osMax = int64(^uint64(0) >> 1)
		}

		asUnitRecalledTemplate.osMin = conf.OsMin
		//polaris_logger.Logger.Debug("config Unit [unitId]=", this.UnitId, "will set sdkVersion")

		sdkVersion := conf.SdkVersion
		includeSdkVersions := []*SdkVersion{}
		excludeSdkVersions := []*SdkVersion{}

		for _, includeItem := range sdkVersion.Include {
			includeSdkVersions = append(includeSdkVersions, getSdkVersion(&includeItem))
		}
		asUnitRecalledTemplate.IncludeSdkVersions = includeSdkVersions

		for _, excludeItem := range sdkVersion.Exclude {
			excludeSdkVersions = append(excludeSdkVersions, getSdkVersion(&excludeItem))
		}
		asUnitRecalledTemplate.ExcludeSdkVersions = excludeSdkVersions

		//	polaris_logger.Logger.Debug("config Unit [unitId]=", this.UnitId, "before fill endcardTemplates")

		//fill endcardTemplates
		if conf.EndcardTemplate != nil {
			asUnitRecalledTemplate.EndCardTemplate = getTemplateId(conf.EndcardTemplate, "endcard")
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, "endcardTemplates: ", asUnitRecalledTemplate.EndCardTemplate)

		}
		//	polaris_logger.Logger.Debug("config Unit [unitId]=", this.UnitId, "before fill videoTemplates")

		//fill videoTemplates
		if conf.VideoTemplate != nil {
			asUnitRecalledTemplate.VideoTemplate = getTemplateId(conf.VideoTemplate, "video")
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, "videoTemplates: ", asUnitRecalledTemplate.VideoTemplate)

		}
		//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, "asUnitRecalledTemplate: ", asUnitRecalledTemplate)
		orientationTempateList := u.Unit.asUnitRecalledTemplateConf[asUnitRecalledTemplate.Orientation]
		u.Unit.asUnitRecalledTemplateConf[asUnitRecalledTemplate.Orientation] = append(orientationTempateList, &asUnitRecalledTemplate)
		//	polaris_logger.Logger.Debug("config Unit [unitId]=", this.UnitId, "after fill videoTemplates")

	}
}

func (this *UnitInfo) GetAsUnitRecalledTempalte(osVersion int64, sdkVersion int32,
	sdkPrefix string, orientation int64, recalledTempalteIds map[string]bool) map[string]bool {
	videoTemplateIds, endcardTemplateIds := this.GetAsUnitTemplate(osVersion, sdkVersion, sdkPrefix, orientation)

	intersecRecallTemplateIds := make(map[string]bool)

	for templateId, _ := range recalledTempalteIds {
		//minicard && 封面模板不受约束
		/*if templateId == "7002001" || templateId == "8002001" || templateId == "7002000" || templateId == "0" {
			intersecRecallTemplateIds[templateId] = true
			continue
		}*/
		if CheckAsVirtualTemplateStr(templateId) {
			intersecRecallTemplateIds[templateId] = true
			continue
		}

		if len(videoTemplateIds) != 0 {
			if _, ok := videoTemplateIds[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		} else {
			if _, ok := AsUnitDefaultVideoTemplateSet[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		}

		if len(endcardTemplateIds) != 0 {
			if _, ok := endcardTemplateIds[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		} else {
			if _, ok := AsUnitDefaultEndcardTemplateSet[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		}
	}

	return intersecRecallTemplateIds /*, len(videoTemplateIds) != 0 || len(endcardTemplateIds) != 0*/

}

//返回值： 根据流量获取可以在该unit上投放tempalte的集合
func (this *UnitInfo) GetAsUnitTemplate(osVersion int64, sdkVersion int32,
	sdkPrefix string, orientation int64) (map[string]bool, map[string]bool) {
	//portal侧保证同一个orientation下的对应的多组配置不会有交集，因此符合定向条件，循环即终止
	rules := this.Unit.asUnitRecalledTemplateConf[orientation]
	//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, "osVersion:", osVersion, "sdkVersion:", sdkVersion,
	//	"sdkPrefix", sdkPrefix, "orientation:", orientation)

	for _, rule := range rules {
		//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, "rule", *rule)
		if osVersion < rule.osMin || osVersion > rule.osMax {
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, " osVersion filter!")
			continue
		}

		//portal 设置时，如果有excludeSdkVersion，那么includeSdkVersion必须有
		if len(rule.IncludeSdkVersions) == 0 {
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, " includeSdkVersions filter!")
			//continue
			return rule.VideoTemplate, rule.EndCardTemplate
		}

		bFound := false
		for _, includeConf := range rule.IncludeSdkVersions {
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, " includeConf :", *includeConf)

			if includeConf.Match(&sdkPrefix, int64(sdkVersion)) {
				bFound = true
				break
			}
		}

		if !bFound {
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, " includeSdkVersions filter!")

			continue
		}

		//portal 设置的时候，全部excludeSdkVersion必须是全部includeSdkVersion的子集
		bFound = false
		for _, excludeConf := range rule.ExcludeSdkVersions {
			//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, "excludeConf:", *excludeConf)

			if excludeConf.Match(&sdkPrefix, int64(sdkVersion)) {
				//polaris_logger.Logger.Debug("Unit [unitId]=", this.UnitId, " excludeSdkVersions filter!")

				bFound = true
				break
			}
		}

		if !bFound {
			return rule.VideoTemplate, rule.EndCardTemplate
		}
	}

	return nil, nil
}

var AsUnitDefaultVideoTemplateSet = map[string]bool{
	"5002001__Low": true, "5002001__High": true,
	"5002002__Low": true, "5002002__High": true,
	"5002003__Low": true, "5002003__High": true,
	"5002004__Low": true, "5002004__High": true,
	"5002007__Low": true, "5002007__High": true,
	"5002008__Low": true, "5002008__High": true,
}

var AsUnitDefaultEndcardTemplateSet = map[string]bool{
	"6004001": true, "6003001": true,
	"6002001": true, "6002002": true,
	"6002003": true, "6002007": true,
	"6002005": true, "6002006": true,
}
