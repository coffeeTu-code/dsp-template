package dbstruct

func NewFeature() *Feature {
	return &Feature{
		Imp:       &FImp{},
		DeviceIds: &FDeviceIds{},
		Device:    &FDevice{},
		User:      &FUser{},
		App:       &FApp{},
		Source:    &FSource{},
		Regs:      &FRegs{},
		Video:     &FVideo{},
	}
}

type Feature struct {
	Exchange    string
	AdType      string
	Platform    string
	CountryCode string
	Size        []string
	Imp         *FImp
	DeviceIds   *FDeviceIds
	Device      *FDevice
	User        *FUser
	App         *FApp
	Source      *FSource
	Regs        *FRegs
	Video       *FVideo
}

type FImp struct {
	BannerBType []int32 // 物料限制属性

	Skadn        string `json:"skadn,omitempty"`        // Ext
	SessionDepth int    `json:"sessiondepth,omitempty"` // Ext
	SourceID     int    `json:"source_id,omitempty"`    // Ext
	SourceURL    string `json:"source_url,omitempty"`   // Ext
}

type FDeviceIds struct {
	GoogleAdId     string //1
	Idfa           string //2
	GoogleAdIdMD5  string //3
	GoogleAdIdSHA1 string //4
	Imei           string //5
	ImeiMD5        string //6
	ImeiSHA1       string //7
	AndroidId      string //8
	AndroidIdMD5   string //9
	AndroidIdSHA1  string //10
	OAId           string //11
	OAIdMD5        string //12
}

type FDevice struct {
	IP          string  `json:"ip,omitempty"`
	IPv6        string  `json:"ipv6,omitempty"`
	UserAgent   string  `json:"useragent,omitempty"`
	ScreenW     int32   `json:"screenw,omitempty"`
	ScreenH     int32   `json:"screenh,omitempty"`
	DeviceType  string  `json:"devicetype,omitempty"`
	Pxratio     float64 `json:"pxratio,omitempty"`
	Lat         float64 `json:"lat,omitempty"` // Geo
	Lon         float64 `json:"lon,omitempty"` // Geo
	Orientation int     `json:"orientation,omitempty"`
	PIDMD5      string  `json:"dpidmd5,omitempty"`

	Atts            int    `json:"atts,omitempty"`             // Ext
	Ifv             string `json:"ifv,omitempty"`              // Ext
	TotalDisk       int    `json:"totaldisk,omitempty"`        // Ext
	CbId            string `json:"cb_id,omitempty"`            // Ext chartboost
	DiskSpace       int    `json:"disk_space,omitempty"`       // Ext chartboost
	LastBootUp      int    `json:"last_bootup,omitempty"`      // Ext chartboost
	SessionDuration int    `json:"session_duration,omitempty"` // Ext chartboost

}

type FUser struct {
	ID string `json:"id,omitempty"`

	LastBundle      string `json:"lastbundle,omitempty"`      // Ext  inneractive
	LastADomain     string `json:"lastadomain,omitempty"`     // Ext  inneractive
	ImpDepth        int    `json:"impdepth,omitempty"`        // Ext  inneractive
	SessionDuration int    `json:"sessionduration,omitempty"` // Ext  inneractive
}

type FApp struct {
	Cat       []string   `json:"cat,omitempty"`
	Keywords  string     `json:"keywords,omitempty"`
	Publisher FPublisher `json:"publisher,omitempty"`

	DevUserId string `json:"devuserid,omitempty"` // Ext inneractive
}

type FPublisher struct {
	Name string `json:"name,omitempty"`
}

type FVideo struct {
	PlaybackEnd    int32   `json:"playbackend,omitempty"`
	PlaybackMethod []int32 `json:"playbackmethod,omitempty"`
	Skip           int32   `json:"skip,omitempty"`
}

type FRegs struct {
	Gdpr      int `json:"gdpr,omitempty"`      // Ext
	UsPrivacy int `json:"usprivacy,omitempty"` // Ext chartboost
}
type FSource struct {
	Schain FSchain `json:"schain,omitempty"` // 特定交易的OpenRTB协议的扩展信息占位符
}

type FSchain struct {
	Ver      string  `json:"ver"`
	Complete int     `json:"complete"`
	Nodes    []FNode `json:"nodes,omitempty"`
}

type FNode struct {
	Asi string `json:"asi,omitempty"`
	Sid string `json:"sid,omitempty"`
	Rid string `json:"rid,omitempty"`
	Hp  int    `json:"hp,omitempty"`
}
