package bifrost

import (
	"go.mongodb.org/mongo-driver/mongo/options"

	"dsp-template/api/backend"
)

type MongoStreamerCfg struct {
	Name           string
	UpdateMode     string
	IncInterval    int
	BaseInterval   int
	IsSync         bool
	MongoConfig    backend.MongoConfig
	BaseParser     DataParser
	IncParser      DataParser
	BaseQuery      interface{}
	IncQuery       interface{}
	UserData       interface{}
	FindOpt        *options.FindOptions
	OnBeforeBase   func(interface{}) interface{}
	OnBeforeInc    func(interface{}) interface{}
	OnFinishBase   func(streamer Streamer)
	OnFinishInc    func(streamer Streamer)
	Logger         BiLogger
}
