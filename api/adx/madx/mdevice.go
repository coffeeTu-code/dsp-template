package madx

type MDevice struct {
	UA         string      `json:"ua,omitempty"`             // 浏览器User-Agent字符串
	Geo        *MGeo       `json:"geo,omitempty"`            // Geo对象，用用户当前位置表示设备位置
	DNT        int         `json:"dnt,omitempty"`            // 浏览器在HTTP头中设置的标准的 “Do Not Track"标识， 0表示不限制追踪， 1表示限制（不允许）追踪
	LMT        int         `json:"lmt,omitempty"`            // "1": Limit Ad Tracking
	IP         string      `json:"ip,omitempty"`             // 最接近设备的IPv4地址
	IPv6       string      `json:"ipv6,omitempty"`           // 最接近设备的IPV6地址
	DeviceType int         `json:"devicetype,omitempty"`     // 设备类型，参考5.17
	Make       string      `json:"make,omitempty"`           // 设备制造商，例如 "Apple"
	Model      string      `json:"model,omitempty"`          // 设备型号，例如 "iphone"
	OS         string      `json:"os,omitempty"`             // 设备操作系统， 例如 "ios"
	OSVer      string      `json:"osv,omitempty"`            // 设备操作系统版本号， 例如 “3.1.2”
	HwVer      string      `json:"hwv,omitempty"`            // 设备硬件版本， 例如 “5S” (e.g., "5S" for iPhone 5S).
	H          int         `json:"h,omitempty"`              // 屏幕的物理高度， 以像素为单位
	W          int         `json:"w,omitempty"`              // 屏幕的物理宽度，以像素为单位
	Language   string      `json:"language,omitempty"`       // 浏览器语言，使用ISO-639-1-alpha-2
	Carrier    string      `json:"carrier,omitempty"`        // Carrier or ISP (e.g., “VERIZON”) using exchange curated string names which should be published to bidders a priori.
	ConnType   int         `json:"connectiontype,omitempty"` // 网络连接类型， 参考表5.18
	IFA        string      `json:"ifa,omitempty"`            // ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IDSHA1     string      `json:"didsha1,omitempty"`        // 硬件设备ID(例如 IMEI),使用SHA1哈希算法
	IDMD5      string      `json:"didmd5,omitempty"`         // 硬件设备ID(例如 IMEI),使用md5哈希算法
	PIDSHA1    string      `json:"dpidsha1,omitempty"`       // 设备平台ID(例如 Android ID),使用SHA1哈希算法
	PIDMD5     string      `json:"dpidmd5,omitempty"`        // 设备平台ID(例如 Android ID),使用md5哈希算法
	OAID       string      `json:"oaid,omitempty"`
	Pxratio    float64     `json:"pxratio,omitempty"` // The ratio of physical pixels to device independent pixels
	Js         int         `json:"js,omitempty"`      // The ratio of physical pixels to device independent pixels
	Ext        *MDeviceExt `json:"ext,omitempty"`     // 设备扩展
}

type MDeviceExt struct {
	OaidMd5      string   `json:"oaidmd5,omitempty"`
	Imei         string   `json:"imei,omitempty"`
	AndroidId    string   `json:"androidid,omitempty"`
	IdfaMd5      string   `json:"idfamd5,omitempty"`
	IdfaSha1     string   `json:"idfasha1,omitempty"`
	TotalDisk    int      `json:"total_disk,omitempty"` // inneractive
	InstalledApp []uint64 `json:"installed_app,omitempty"`
	Atts         int      `json:"atts,omitempty"`           // mopub: (iOS Only) An integer passed to represent the app's app tracking authorization status, where 0 = not determined  1 = restricted  2 = denied  3 = authorized
	Ifv          string   `json:"ifv,omitempty"`            // mopub、inneractive: IDFV of device in that publisher. Only passed when IDFA is unavailable or all zeros. Currently passed for iOS only.
	CbId         string   `json:"cb_id,omitempty"`          //chartboost
	DiskSpace    int      `json:"disk_space,omitempty"`     //设备总磁盘空间 chartboost 和 iqiyi ios有值
	LastBootUp   int      `json:"last_bootup,omitempty"`    // chartboost
	TimeZone     string   `json:"time_zone,omitempty"`      // 本地时区, 只iqiyi-ios有值
	OsUpdateTime int      `json:"os_update_time,omitempty"` // 系统更新时间，ts, 只iqiyi ios有值
	StartupTime  int      `json:"startup_time,omitempty"`   // 开机时间 ，ts, 只iqiyi ios 有值
	CpuNum       int      `json:"cpu_num,omitempty"`        // 设备CPU数 , 只iqiyi ios有值
}
