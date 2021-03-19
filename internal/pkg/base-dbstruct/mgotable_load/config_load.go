package mgotable_load

import (
	"encoding/json"
	"errors"
	"github.com/Mintegral-official/mtggokit/bifrost/container"
	"github.com/Mintegral-official/mtggokit/bifrost/streamer"
	"gitlab.mobvista.com/voyager/mgoextractor/mgoconfig"
	"gitlab.mobvista.com/voyager/mgoextractor/mgologger"
	"gitlab.mobvista.com/voyager/mgoextractor/mgotables"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"sync"
)

var getConfigStreamerError = errors.New("unit streamer not found key")

const (
	SINGLE_SLOT                     = 101
	SINGLE_SLOT_GOOGLEPLAY          = 102
	SELECTABLE_TWO_VIDEO            = 201
	SELECTABLE_TWO_VIDEO_PLAY_FIRST = 301
	REWARDPLUS_TWO_SLOT             = 401
)

const (
	BIG_TEMPLATE_MAP_RULE_LIST = "BIG_TEMPLATE_MAP_RULE_LIST"
)

func DefaultConfigParserCfg(parser *ConfigParser) mgoconfig.BiFrostConfig {
	return mgoconfig.BiFrostConfig{
		BaseInterval: 2592000,                    //数据表全量更新周期
		IncInterval:  60,                         //数据表增量更新周期
		URI:          "",                         //数据表mongo url, 因各集群地址不同，需要外部输入
		DB:           mgoconfig.NewAdnDB,         //数据表存储数据库
		Collection:   mgoconfig.ConfigCollection, //数据表name
		DataParser:   parser,
	}
}

//--------------------------- BiFrostExtractor ------------------------------------

func NewConfigParser() *ConfigParser {
	return &ConfigParser{
		Parser: func(config *mgotables.Config) bool { return true },
	}
}

type ConfigParser struct {
	MaxUpdated         int64
	Parser             func(config *mgotables.Config) bool // 提供给外部的 Parser 方法，用于 解析非标字段 + 判断当前doc是否可用状态，返回 true 保留记录，返回 false 删除记录。
	bigTemplateRuleMap sync.Map
	bigTemplateSet     sync.Map
}

type BigTemplateRuleInfo struct {
	IsBlack     bool
	PublisherId map[int64]bool
	AppId       map[int64]bool
	UnitId      map[int64]bool
}

func (parser *ConfigParser) Parse(data []byte, userData interface{}) []streamer.ParserResult {
	table := &mgotables.Config{}

	//这里使用mgo.v2的反序列化函数，不实用mongo.driver的反序列化函数的原因是：对部分字段的解析出错导致数据读取失败，数据丢失。
	if err := bson.Unmarshal(data, &table); err != nil {
		mgologger.MgoLog.Error("MgoExtractor Parse Config error! bson.Unmarshal error =" + err.Error())
		return []streamer.ParserResult{{container.DataModeDel, nil, nil, errors.New("Parse Config error ")}}
	}

	key := table.Key
	//table.Parser()
	reserved := parser.Parser(table)

	//add for as ,大模板召回规则配置,定义了大模板召回模板的配置
	parser.fillAsBigTemplateRule(data)

	if reserved {
		body := ""
		if mgoconfig.DebugMgoExtractor {
			j, _ := json.Marshal(table)
			body = string(j)
		}
		mgologger.MgoLog.Info("MgoExtractor Add Config [key]=", key, body)
	} else {
		mgologger.MgoLog.Info("MgoExtractor Delete Config [key]=", key)
		return []streamer.ParserResult{{container.DataModeDel, container.StrKey(key), nil, nil}}
	}
	return []streamer.ParserResult{{
		container.DataModeAdd, container.StrKey(table.Key), table.Value, nil,
	}}
}

func (this *ConfigParser) fillAsBigTemplateRule(data []byte) {
	var bigTemplateRuleConfig = map[string]interface{}{}
	if err := bson.Unmarshal(data, &bigTemplateRuleConfig); err != nil {
		mgologger.MgoLog.Warn("Parse bigTemplatRuleConfig error! bson.Unmarshal error =" + err.Error())
		return
	}

	if key, keyOk := bigTemplateRuleConfig["key"]; keyOk == false || key != BIG_TEMPLATE_MAP_RULE_LIST {
		return
	}

	value, ok := bigTemplateRuleConfig["value"]
	if !ok {
		mgologger.MgoLog.Warn("Parse bigTemplateRuleConfig error! can't find value field")
		return
	}

	for bigTemplateIdStr, ruleObj := range value.(map[string]interface{}) {
		if bigTemplateId, err := strconv.Atoi(bigTemplateIdStr); err == nil {
			var bigTemplateRule = BigTemplateRuleInfo{}

			bigTemplateRuleMap := ruleObj.(map[string]interface{})
			ruleType, _ := bigTemplateRuleMap["type"]
			ruleTypeInt, err := GetInt64(ruleType)
			if err != nil {
				mgologger.MgoLog.Warn(err.Error())

			}
			if int32(ruleTypeInt) == 2 {
				bigTemplateRule.IsBlack = true
			} else {
				bigTemplateRule.IsBlack = false
			}

			if appIds, ok := bigTemplateRuleMap["appId"]; ok {
				if bigTemplateRule.AppId == nil {
					bigTemplateRule.AppId = map[int64]bool{}
				}
				for _, appId := range appIds.([]interface{}) {
					appIdInt, err := GetInt64(appId)
					if err != nil {
						mgologger.MgoLog.Warn(err.Error())
					}
					bigTemplateRule.AppId[appIdInt] = true
				}
			}

			if publisherIds, ok := bigTemplateRuleMap["publisherId"]; ok {
				if bigTemplateRule.PublisherId == nil {
					bigTemplateRule.PublisherId = map[int64]bool{}
				}
				for _, publisherId := range publisherIds.([]interface{}) {
					publisherIdInt, err := GetInt64(publisherId)
					if err != nil {
						mgologger.MgoLog.Warn(err.Error())
					}
					bigTemplateRule.PublisherId[publisherIdInt] = true
				}
			}

			if unitIds, ok := bigTemplateRuleMap["unitId"]; ok {
				if bigTemplateRule.UnitId == nil {
					bigTemplateRule.UnitId = map[int64]bool{}
				}
				for _, unitId := range unitIds.([]interface{}) {
					unitIdInt, err := GetInt64(unitId)
					if err != nil {
						mgologger.MgoLog.Warn(err.Error())
					}
					bigTemplateRule.UnitId[unitIdInt] = true
				}
			}

			this.bigTemplateRuleMap.Store(int64(bigTemplateId), &bigTemplateRule)
			this.bigTemplateSet.Store(int64(bigTemplateId), true)
		}
	}

	mgologger.MgoLog.Debug("Parse bigTemplateRuleMap success, bigTemplate size is=")

}

func (this *ConfigParser) FilterBigTemplateRule(requestId string, appId int64,
	unitId int64, publisherId int64, bigTemplateId int64) bool {

	//模板集合表里压根没有该大模板，
	_, ok := this.bigTemplateSet.Load(bigTemplateId)
	if !ok {
		return true
	}

	//101大模板不受限制
	if bigTemplateId == SINGLE_SLOT {
		return true
	}

	//大模板列表中查不到，返回false
	bigTemplateRuleObj, ok := this.bigTemplateRuleMap.Load(bigTemplateId)
	if !ok {
		mgologger.MgoLog.Warn(requestId, "BigTemplateId ", bigTemplateId, "can't be found in config")
		return false
	}

	//黑白名单+appId, publisherId, unitId取并
	bigTemplateRule := bigTemplateRuleObj.(*BigTemplateRuleInfo)

	_, inAppIds := bigTemplateRule.AppId[appId]
	_, inUnitIds := bigTemplateRule.UnitId[unitId]
	_, inPublisherIds := bigTemplateRule.PublisherId[publisherId]

	if bigTemplateRule.IsBlack {
		return !(inAppIds || inUnitIds || inPublisherIds)
	} else {
		return inAppIds || inUnitIds || inPublisherIds
	}

}

func GetInt64(intVal interface{}) (int64, error) {
	if intVal == nil {
		return -1, nil
	}
	switch s := intVal.(type) {
	case float64:
		return int64(s), nil
	case int64:
		return s, nil
	case int32:
		return int64(s), nil
	case float32:
		return int64(s), nil
	case int:
		return int64(s), nil
	case uint8:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint64:
		return int64(s), nil
	default:
		err := errors.New("Not allowed type value!")
		return -1, err
	}

}
