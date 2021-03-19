package mgotable_load

import (
	"encoding/json"
	"errors"
	"github.com/Mintegral-official/mtggokit/bifrost/container"
	"github.com/Mintegral-official/mtggokit/bifrost/streamer"
	"gitlab.mobvista.com/voyager/common/enum"
	"gitlab.mobvista.com/voyager/mgoextractor/mgoconfig"
	"gitlab.mobvista.com/voyager/mgoextractor/mgologger"
	"gitlab.mobvista.com/voyager/mgoextractor/mgotables"
	"gopkg.in/mgo.v2/bson"
)

func DefaultCreativePackageCfg(parser *CreativePackageParser) mgoconfig.BiFrostConfig {
	return mgoconfig.BiFrostConfig{
		BaseInterval: 2592000,                             //数据表全量更新周期
		IncInterval:  60,                                  //数据表增量更新周期
		URI:          "",                                  //数据表mongo url, 因各集群地址不同，需要外部输入
		DB:           mgoconfig.NewAdnDB,                  //数据表存储数据库
		Collection:   mgoconfig.CreativePackageCollection, //数据表name
		DataParser:   parser,
	}
}

//--------------------------- BiFrostExtractor ------------------------------------

func NewCreativePackageParser() *CreativePackageParser {
	return &CreativePackageParser{
		Parser: func(creativePackage *mgotables.CreativePackage) bool { return true },
	}
}

type CreativePackageParser struct {
	MaxUpdated int64

	Parser func(*mgotables.CreativePackage) bool // 提供给外部的 Parser 方法，用于 解析非标字段 + 判断当前doc是否可用状态，返回 true 保留记录，返回 false 删除记录。
}

func (parser *CreativePackageParser) Parse(data []byte, userData interface{}) []streamer.ParserResult {
	table := &mgotables.CreativePackage{}

	if err := bson.Unmarshal(data, &table); err != nil {
		mgologger.MgoLog.Error("Parse CreativePackage error! bson.Unmarshal error =" + err.Error())
		return []streamer.ParserResult{{container.DataModeDel, nil, nil, errors.New("Parse CreativePackage error ")}}
	}

	if table.Updated > parser.MaxUpdated {
		parser.MaxUpdated = table.Updated
	}

	table.Parser()
	reserved := parser.Parser(table)

	if reserved && table.Status == enum.MongoStatus_ACTIVE {
		body := ""
		if mgoconfig.DebugMgoExtractor {
			j, _ := json.Marshal(table)
			body = string(j)
		}
		mgologger.MgoLog.Info("MgoExtractor Add CreativePackage [key]=", table.CpId, body)
	} else {
		mgologger.MgoLog.Info("MgoExtractor Delete CreativePackage [key]=", table.CpId)
		return []streamer.ParserResult{{container.DataModeDel, container.I64Key(table.CpId), nil, nil}}
	}

	return []streamer.ParserResult{{container.DataModeAdd, container.I64Key(table.CpId), table, nil}}
}
