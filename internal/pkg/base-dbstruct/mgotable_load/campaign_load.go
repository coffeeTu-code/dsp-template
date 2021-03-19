package mgotable_load

import (
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"

	"dsp-template/api/enum"
	"dsp-template/internal/pkg/base-dbstruct/mgologger"
	"dsp-template/internal/pkg/base-dbstruct/mgotables"
	bifrost "dsp-template/pkg2/helper-bifrost"
	"dsp-template/pkg2/helper-bifrost/container"
)

func DefaultCampaignCfg(parser *CampaignParser) BiFrostConfig {
	return BiFrostConfig{
		BaseInterval: 2592000,                      //数据表全量更新周期
		IncInterval:  60,                           //数据表增量更新周期
		URI:          "",                           //数据表mongo url, 因各集群地址不同，需要外部输入
		DB:           mgoconfig.NewAdnDB,           //数据表存储数据库
		Collection:   mgoconfig.CampaignCollection, //数据表name
		DataParser:   parser,
	}
}

//--------------------------- BiFrostExtractor ------------------------------------

func NewCampaignParser() *CampaignParser {
	return &CampaignParser{
		Parser: func(campaign *mgotables.Campaign) bool { return true },
	}
}

type CampaignParser struct {
	MaxUpdated int64

	Parser func(*mgotables.Campaign) bool // 提供给外部的 Parser 方法，用于 解析非标字段 + 判断当前doc是否可用状态，返回 true 保留记录，返回 false 删除记录。
}

func (parser *CampaignParser) Parse(data []byte, userData interface{}) []bifrost.ParserResult {
	table := &mgotables.Campaign{}

	//这里使用mgo.v2的反序列化函数，不实用mongo.driver的反序列化函数的原因是：对部分字段的解析出错导致数据读取失败，数据丢失。
	if err := bson.Unmarshal(data, &table); err != nil {
		mgologger.MgoLog.Error("MgoExtractor Parse Campaign error! bson.Unmarshal error =" + err.Error())
		return []bifrost.ParserResult{{container.DataModeDel, nil, nil, errors.New("Parse Campaign error ")}}
	}

	if table.Updated > parser.MaxUpdated {
		parser.MaxUpdated = table.Updated
	}

	key := table.CampaignId
	if key == 0 {
		mgologger.MgoLog.Info("MgoExtractor Illegal Campaign [key]=", key)
		return []bifrost.ParserResult{{container.DataModeDel, nil, nil, errors.New("Illegal Campaign ")}}
	}

	table.Parser()
	reserved := parser.Parser(table)

	// - CampaignInvertedIndex -
	UpdateCampInvertedIndex(table)

	if reserved && table.Status == enum.MongoStatus_ACTIVE {
		body := ""
		if DebugMgoExtractor {
			j, _ := json.Marshal(table)
			body = string(j)
		}
		mgologger.MgoLog.Info("MgoExtractor Add Campaign [key]=", key, body)
	} else {
		mgologger.MgoLog.Info("MgoExtractor Delete Campaign [key]=", key)
		return []bifrost.ParserResult{{container.DataModeDel, container.I64Key(key), nil, nil}}
	}

	return []bifrost.ParserResult{{container.DataModeAdd, container.I64Key(key), table, nil}}
}
