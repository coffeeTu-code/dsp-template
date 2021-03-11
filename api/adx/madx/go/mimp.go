package mortb

// 这个对象描述了一个广告位或者将要参与竞拍的展现。一个竞价请求可以包含多个Imp对象，
// 这种状况的一个示例是一个交易平台支持售卖一个页面的所有广告位。 为了便于竞拍者区分， 每一个Imp对象都要有一个唯一标识（ID).
// Banner, Video以及Native对象都属于Imp对象，只是标明了各自的展示类型。 展示者可以选择其中的一种类型或者混合使用多种类型。
// 但是，对于展示的任何给定的竞价请求必须属于提供的类型之一。
type MImp struct {
	ID                string      `json:"id"`                          // 在当前竞价请求上下文中唯一标识本次展示的标识（通常从1开始并以此递增)
	Metric            []*MMetric  `json:"metric,omitempty"`            // An array of Metric object (Section 3.2.5).
	Banner            *MBanner    `json:"banner,omitempty"`            // Banner对象，如果展示需要以banner的形式提供则需要填充
	Video             *MVideo     `json:"video,omitempty"`             // Video对象，如果展示需要以视频的形式提供则需要填充
	Native            *MNativeReq `json:"native,omitempty"`            // Native对象，如果展示需要以Native广告的形式提供则需要填充
	Pmp               *MPmp       `json:"pmp,omitempty"`               // Pmp对象 包含对本次展示生效的任何私有市场交易
	DisplayManager    string      `json:"displaymanager,omitempty"`    // 广告媒体合作伙伴的名字，用于渲染广告的SDK技术或者播放器（通常是视频或者移动广告）某些广告服务需要根据合作伙伴定制广告代码， 推荐在视频广告或应用广告中填充
	DisplayManagerVer string      `json:"displaymanagerver,omitempty"` // 广告媒体合作伙伴， 用于渲染广告的SDK技术或者播放器（通常是视频或者移动广告）的版本号。 某些广告服务需要根据合作伙伴定制广告代码， 推荐在视频广告或应用广告中填充
	Instl             int         `json:"instl,omitempty"`             // 1标识广告是插屏或者全屏广告，0表示不是插屏广告
	TagID             string      `json:"tagid,omitempty"`             // 某个特定广告位或者广告标签的标识，用于发起竞拍。为了方便调试问题或者进行买方优化
	BidFloor          float64     `json:"bidfloor,omitempty"`          // 本次展示的最低竞拍价格，以CPM表示
	BidFloorCurrency  string      `json:"bidfloorcur,omitempty"`       // ISO-4217规定的字母码表标识的货币类型。如果交易平台允许，可能与从竞拍者返回的竞价货币不同
	Secure            int         `json:"secure,omitempty"`            // 标识展示请求是否需要使用HTTPS加密物料信息以及markup以保证安全，0标识不需要使用安全链路，1标识需要使用安全链路，如果不填充，则表示未知，可以认为是不需要使用安全链路
	IFrameBuster      []string    `json:"iframebuster,omitempty"`      // 特定交易支持的iframe buster的名字数组
	Exp               int         `json:"exp,omitempty"`
	Ext               *MImpExt    `json:"ext,omitempty"` // 特定交易的OpenRTB协议的扩展信息占位符
}

type MImpExt struct {
	BlockingKeyword             []string          `json:"blockingkeyword,omitempty"`             // 广告位过滤的关键字
	BlockedIndustryId           []int64           `json:"blockedindustryid,omitempty"`           // 广告位过滤的行业ID 列表
	BlockedIndustryType         []int32           `json:"blockedindustrytype,omitempty"`         // 广告行业类型过滤
	IndustryFloorPrice          map[int64]float64 `json:"industryfloorprice,omitempty"`          // iqiyi 不同行业低价
	CreativeSpecs               []string          `json:"creativespecs,omitempty"`               // 可投放的分组素材
	ForbidApk                   bool              `json:"forbidapk,omitempty"`                   // 默认false，表示允许投apk单子, true表示不允许投apk单子
	StoreURLRequired            bool              `json:"storeurlrequired,omitempty"`            // 是否要求返回storeURL字段； true表示要求，false表示不返回
	ClickParamRequired          bool              `json:"clickparamrequired,omitempty"`          // 替换DSP 点击监测 URL 里的宏__CLICK_PARAM__ true为必须返回，false无需返回
	ImpressionParamRequired     bool              `json:"impressionparamrequired,omitempty"`     // 替换DSP 曝光监测 URL（monitor_url1 ）里的宏__IMPRESSION_PARAM__ true为必须返回，false无需返回
	CreativeReviewRequired      bool              `json:"creativereviewrequired,omitempty"`      // 是否要求素材审核
	QualificationReviewRequired bool              `json:"qualificationreviewrequired,omitempty"` // 是否要求资质审核
	DeepLink                    int               `json:"deeplink,omitempty"`                    // 1代表流量支持deeplink
	Reward                      int               `json:"reward,omitempty"`                      // 日本line-adx拓展字段，1-奖励视频
	// admType: 返回中adm的类型，枚举值：banner/video/native, 在有传这个值时，mvdsp对非审核的广告位会基于这里的admType来返回对应的adm类型。
	// eg: 请求中有 banner结构体， admType="native", mvdsp会作为native来处理,返回json而非banner html
	AdmType                string   `json:"admtype,omitempty"`
	CreativeFeatures       []string `json:"creativefeatures,omitempty"`
	TargetTemplate         string   `json:"target_template,omitempty"`        // 模板定向，如果有传，固定使用这个模板
	Videovolume            int      `json:"videovolume,omitempty"`            // 对视频素材的容积限制， 单位byte 0：无限制
	Imagevolume            int      `json:"imagevolume,omitempty"`            // 对图片素材的容积限制， 单位byte 0：无限制
	ApkLandingpageRequired bool     `json:"apklandingpagerequired,omitempty"` // 这个请求是否要求二跳页
	Skadn                  *Skadn   `json:"skadn,omitempty"`
	PlacementType          string   `json:"placementtype,omitempty"` //流量类型，如rewarded video【ps：前期为腾讯使用】
	SessionDepth           int      `json:"sessiondepth,omitempty"`  //当前session已经投放了几条广告。The number of times an ad has been delivered for the current session.
	SourceID               int      `json:"source_id,omitempty"`
	SourceURL              string   `json:"source_url,omitempty"`
}

type Skadn struct {
	Version    string    `json:"version,omitempty"`
	Sourceapp  string    `json:"sourceapp,omitempty"`
	Skadnetids []string  `json:"skadnetids,omitempty"`
	Ext        *SkadnExt `json:"ext,omitempty"`
}

type SkadnExt struct{}

// 用于封装一个地理位置信息的多种不同属性。 当作为Device对象的子节点的时候，标识设备的地理位置或者用户当前的地理位置。
// 当作为User的子节点的时候，标识用户家的位置（也就是说，不必是用户的当前位置）。
// 设备的位置或者用户家的位置，由其父对象决定
type MGeo struct {
	Lat           float64 `json:"lat,omitempty"`           // 纬度信息，取值范围-90.0到+90.0， 负值表示南方
	Lon           float64 `json:"lon,omitempty"`           // 经度信息， 取值返回-180.0到+180.0， 负值表示西方
	Type          int     `json:"type,omitempty"`          // 位置信息的源， 当传递lat/lon的时候推荐填充， 参考表5.16
	Accuracy      int     `json:"accuracy,omitempty"`      // Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services
	LastFix       int     `json:"lastfix,omitempty"`       // Number of seconds since this geolocation fix was established.
	IPService     int     `json:"ipservice,omitempty"`     // Service or provider used to determine geolocation from IP address if applicable
	Country       string  `json:"country,omitempty"`       // 国家码， 使用 ISO-3166-1-alpha-3
	Region        string  `json:"region,omitempty"`        // 区域码， 使用ISO-3166-2; 如果美国则使用2字母区域码
	RegionFIPS104 string  `json:"regionFIPS104,omitempty"` // 国家的区域，使用 FIPS 10-4 表示。 虽然OpenRTB支持这个属性，它已经与2008年被NIST撤销了
	Metro         string  `json:"metro,omitempty"`         // 谷歌metro code; 与Nielsen DMA相似但不完全相同， 参见附录A
	City          string  `json:"city,omitempty"`          // 城市名，使用联合国贸易与运输位置码， 参见附录A
	Zip           string  `json:"zip,omitempty"`           // 邮政编码或者邮递区号
	UTCOffset     int     `json:"utcoffset,omitempty"`     // 使用UTC加或者减分钟数的方式表示的本地时间
	//Ext           Extension `json:"ext,omitempty"`           // 特定交易的OpenRTB协议的扩展信息占位符
}

// 描述了解或者持有设备的用户的信息（也就是广告的受众）。
// 用户id是一个exchange artifact, 可能随着屏幕旋转或者其他的隐私策略改变。
// 尽管如此，用户id必须在足够长的一段时间内保持不变，以为目标用户定向和用户访问频率限制提供合理的服务。
// 设备的用户， 广告的受众
type MUser struct {
	ID          string   `json:"id,omitempty"`           // 交易特定的用户标识， 推荐id和buyeruid中至少提供一个
	BuyerUID    string   `json:"buyeruid,omitempty"`     // 买方为用户指定的ID，由交易平台为买方映射。推荐id和buyeruid中至少提供一个.
	YOB         int      `json:"yob,omitempty"`          // 生日年份，使用4位数字表示
	Gender      string   `json:"gender,omitempty"`       // 性别， M表示男性， F表示女性， O标识其他类型，不填充表示未知
	Keywords    string   `json:"keywords,omitempty"`     // 逗号分隔的关键字， 兴趣或者意向列表
	CustomData  string   `json:"customdata,omitempty"`   // 可选特性， 用于传递给竞拍者信息，在交易平台的cookie中设置。字符串必须使用base85编码的 cookie，可以是任意格式。 JSON加密的时候必须包括转义的引号
	Geo         *MGeo    `json:"geo,omitempty"`          // Geo对象， 用户家的位置信息。不必是用户的当前位置
	Data        []*MData `json:"data,omitempty"`         // 附加的用户信息， 每个 Data对象表示一个不同的数据源
	AudienceIds []uint32 `json:"audience_ids,omitempty"` //用户受众ids
	Ext         *UserExt `json:"ext,omitempty"`
}

type UserExt struct {
	UserCategory    []int64 `json:"user_category,omitempty"`
	Consent         string  `json:"consent,omitempty"`
	LastBundle      string  `json:"lastbundle,omitempty"`      // inneractive
	LastADomain     string  `json:"lastadomain,omitempty"`     // inneractive
	ImpDepth        int     `json:"impdepth,omitempty"`        // inneractive
	SessionDuration int     `json:"sessionduration,omitempty"` // inneractive
}

// Data和Segment对象一起允许指定用户附加信息。数据可能来自多个数据源， 可能来自交易平台自身或者第三方提供的信息， 可以使用id属性区分。
// 一个竞价请求可以混合来自多个提供者的数据信息。 交易平台应该优先提供正在使用的数据提供者的信息
type MData struct {
	ID      string      `json:"id,omitempty"`      // 交易特定的数据提供者标识
	Name    string      `json:"name,omitempty"`    // 交易特定的数据提供者名称
	Segment []*MSegment `json:"segment,omitempty"` // 包含数据内容的一组Segment对象
	//Ext     Extension `json:"ext,omitempty"`     // 特定交易的OpenRTB协议的扩展信息占位符
}

// 数据字段， 描述用户信息数据的键值对。 其父对象Data是某个给定数据提供者的数据字段的集合。
// 交易平台必须优先将字段的名称和值传递给竞拍者
type MSegment struct {
	ID    string `json:"id,omitempty"`    // 数据提供者的特定数据段的ID
	Name  string `json:"name,omitempty"`  // 数据提供者的特定数据段的名称
	Value string `json:"value,omitempty"` // 表示数据字段值的字符串
	//Ext   Extension `json:"ext,omitempty"`   // 特定交易的OpenRTB协议的扩展信息占位符
}

// 描述任何适用于该请求的法律，政府或者工业管控条例。
// coppa(Children’s Online Privacy Protection Act)标志着是否该请求是否符合美国联邦贸易委员会颁布的美国儿童在线隐私保护法案，详情可参考7.1节
type MRegulations struct {
	Coppa int      `json:"coppa,omitempty"` // 标志着该请求是否遵从COPPA法案， 0表示不遵从， 1表示遵从
	Ext   *MRegExt `json:"ext,omitempty"`
}

type MRegExt struct {
	Gdpr      int `json:"gdpr,omitempty"`
	UsPrivacy int `json:"us_privacy,omitempty"` // chartboost
}

type MPmp struct {
	Private int      `json:"private_auction,omitempty"` // 标识在Deal对象中指明的席位的竞拍合格标准， 0标识接受所有竞拍， 1标识竞拍受deals属性中描述的规则的限制
	Deals   []*MDeal `json:"deals,omitempty"`           // 一组Deal对象， 用于传输适用于本次展示的交易信息
	//Ext     Extension `json:"ext,omitempty"`             // 特定交易的OpenRTB协议的扩展信息占位符
}

// PMP Deal
type MDeal struct {
	ID               string   `json:"id,omitempty"`          // 直接交易的唯一ID
	BidFloor         float64  `json:"bidfloor,omitempty"`    // 本次展示的最低竞价，以CPM为单位
	BidFloorCurrency string   `json:"bidfloorcur,omitempty"` // 使用ISO-4217码表指定的货币。 如果交易平台允许，这可能与竞价者返回的竞价货币类型不一致
	WSeat            []string `json:"wseat,omitempty"`       // 允许参与本次交易竞价的买方席位白名单。 席位ID需要交易平台和竞拍者提前协商， 忽略本属性标示没有席位限制
	WAdvDomain       []string `json:"wadomain,omitempty"`    // Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal. Omission implies no advertiser restrictions.
	AuctionType      int      `json:"at,omitempty"`          // 允许参与本次交易竞价的广告主域名列表（例如， advertiser.com). 忽略本属性标示没有广告主限制。
	//Ext              Extension `json:"ext,omitempty"`         // 特定交易的OpenRTB协议的扩展信息占位符
}

type MPublisher struct {
	ID     string   `json:"id,omitempty"`
	Name   string   `json:"name,omitempty"`
	Cat    []string `json:"cat,omitempty"` // Array of IAB content categories
	Domain string   `json:"domain,omitempty"`
	//Ext    Extension `json:"ext,omitempty"`
}

// request.Imp.Metric.Type枚举
type MetricType string

const (
	ClickThroughRate    MetricType = "click_through_rate"    // 广告位的历史点击率
	VideoCompletionRate MetricType = "video_completion_rate" // 广告位的视频完全播放，没有跳过的比率
	Viewability         MetricType = "viewability"           // 素材的百分之多少被用户看见
)

// This object is associated with an impression as an array of metrics. These metrics can offer insight into
// the impression to assist with decisioning such as average recent viewability, click-through rate, etc. Each
// metric is identified by its type, reports the value of the metric, and optionally identifies the source or
// vendor measuring the value.
type MMetric struct {
	Type   string  `json:"type,omitempty"`   // Type of metric being presented using exchange curated string names which should be published to bidders a priori.
	Value  float64 `json:"value,omitempty"`  // Number representing the value of the metric. Probabilities must be in the range 0.0 – 1.0.
	Vendor string  `json:"vendor,omitempty"` // Source of the value using exchange curated string names which should be published to bidders a priori. If the exchange itself is the source versus a third party, “EXCHANGE” is recommended.
	//Ext    Extension `json:"ext,omitempty"`    // 特定交易的OpenRTB协议的扩展信息占位符
}

const (
	AdmTypeNative = "native"
	AdmTypeBanner = "banner"
	AdmTypeVideo  = "video"
)
