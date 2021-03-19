package bifrost

type LocalFileStreamerCfg struct {
	Name         string
	Path         string
	UpdateMode   string
	Interval     int
	IsSync       bool
	DataParser   DataParser
	UserData     interface{}
	Logger       BiLogger
	OnBeforeBase func(streamer Streamer) error
	OnFinishBase func(streamer Streamer)
}
