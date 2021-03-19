package mgotable_load

import (
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"

	"dsp-template/api/backend"
	"dsp-template/api/enum"
	"dsp-template/internal/pkg/base-dbstruct/mgologger"
	"dsp-template/internal/pkg/base-dbstruct/mgotables"
	bifrost "dsp-template/pkg2/helper-bifrost"
	"dsp-template/pkg2/helper-bifrost/container"
)

func DefaultUnitCfg(parser *UnitParser) BiFrostConfig {

	return BiFrostConfig{
		BaseInterval: 2592000, //数据表全量更新周期
		IncInterval:  60,      //数据表增量更新周期
		MongoConfig: backend.MongoConfig{
			URI:        "",             //数据表mongo url, 因各集群地址不同，需要外部输入
			DB:         NewAdnDB,       //数据表存储数据库
			Collection: UnitCollection, //数据表name
		},
		DataParser: parser,
	}
}

//--------------------------- BiFrostExtractor ------------------------------------

func NewUnitParser() *UnitParser {
	return &UnitParser{
		Parser: func(info *mgotables.UnitInfo) bool { return true },
	}
}

type UnitParser struct {
	MaxUpdated int64

	Parser func(*mgotables.UnitInfo) bool // 提供给外部的 Parser 方法，用于 解析非标字段 + 判断当前doc是否可用状态，返回 true 保留记录，返回 false 删除记录。
}

func (parser *UnitParser) Parse(data []byte, userData interface{}) []bifrost.ParserResult {
	table := &mgotables.UnitInfo{}

	if err := bson.Unmarshal(data, &table); err != nil {
		mgologger.MgoLog.Error("MgoExtractor Parse Unit error! bson.Unmarshal error =" + err.Error())
		return []bifrost.ParserResult{{container.DataModeDel, nil, nil, errors.New("Parse Unit error ")}}
	}

	if table.Updated > parser.MaxUpdated {
		parser.MaxUpdated = table.Updated
	}

	table.Parser()
	reserved := parser.Parser(table)

	if reserved && table.Unit.Status == enum.MongoStatus_ACTIVE {
		body := ""
		if DebugMgoExtractor {
			j, _ := json.Marshal(table)
			body = string(j)
		}
		mgologger.MgoLog.Info("MgoExtractor Add Unit [key]=", table.UnitId, body)
	} else {
		mgologger.MgoLog.Info("MgoExtractor Delete Unit [key]=", table.UnitId)
		return []bifrost.ParserResult{{container.DataModeDel, container.I64Key(table.UnitId), nil, nil}}
	}

	return []bifrost.ParserResult{{container.DataModeAdd, container.I64Key(table.UnitId), table, nil}}
}
