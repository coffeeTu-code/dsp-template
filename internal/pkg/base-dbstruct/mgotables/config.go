package mgotables

type Config struct {
	Key     string      `bson:"key,omitempty" json:"key"`
	Value   interface{} `bson:"value,omitempty" json:"value"`
	Updated int64       `bson:"updated,omitempty" json:"updated"`
}

type BigTemplateRuleInfo struct {
	IsBlack     bool
	PublisherId map[int64]bool
	AppId       map[int64]bool
	UnitId      map[int64]bool
}
