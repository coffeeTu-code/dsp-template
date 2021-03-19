package mgotables

import (
	"strconv"
	"strings"

	"dsp-template/api/enum"
)

type Campaign struct {
	CampaignId                int64                `bson:"campaignId,omitempty" json:"campaignId"`
	PackageName               string               `bson:"packageName,omitempty" json:"packageName"`
	AdvertiserId              int64                `bson:"advertiserId,omitempty" json:"advertiserId"`
	AppName                   string               `bson:"appName,omitempty" json:"appName"` // [render]Advertiser
	CountryCode               []string             `bson:"countryCode,omitempty" json:"countryCode"`
	CityCode                  map[string][]int     `bson:"cityCode,omitempty" json:"citycode"` //城市
	PublisherId               int64                `bson:"publisherId,omitempty" json:"publisherId"`
	Platform                  uint8                `bson:"platform,omitempty" json:"platform"`
	Status                    int8                 `bson:"status,omitempty" json:"status"`
	Updated                   int64                `bson:"updated,omitempty" json:"updated"`
	TryNew                    map[string]TryNew    `bson:"tryNew,omitempty" json:"tryNew"`
	Name                      string               `bson:"name,omitempty" json:"name"`
	Price                     float64              `bson:"price,omitempty" json:"price"`
	DeviceAndIpuaRetarget     int8                 `bson:"gaid_idfa_needs,omitempty" json:"gaid_idfa_needs"`
	AdvUserId                 int64                `bson:"advUserId,omitempty" json:"advUserId"`
	IabCategory               map[string][]string  `bson:"iabCategory,omitempty" json:"iabCategory"`
	EffectiveCountryCode      map[string]int       `bson:"effectiveCountryCode,omitempty" json:"effectiveCountryCode"` // 分地区预算控制。key：为ALL或具体国家编码，value为1或其他，为 1 表示预算充足。key["ALL"] =1 or 2;key[country="IN"] = 1 or 2
	Tag                       int                  `bson:"tag,omitempty" json:"tag"`
	SamePackageCreative       int                  `bson:"samePackageCreative,omitempty" json:"samePackageCreative"` //同包名素材开关，枚举值=2 禁止召回同包名素材，枚举值=1 允许召回同包名素材，默认情况 允许召回
	CampaignType              int8                 `bson:"campaignType,omitempty" json:"campaignType"`
	IconUrl                   string               `bson:"iconUrl,omitempty" json:"iconUrl"`
	AppDesc                   string               `bson:"appDesc,omitempty" json:"appDesc"`
	AppSize                   string               `bson:"appSize,omitempty" json:"appSize"`
	AppScore                  float64              `bson:"appScore,omitempty" json:"appScore"`
	AppInstall                uint32               `bson:"appInstall,omitempty" json:"appInstall"`
	Category                  uint8                `bson:"category,omitempty" json:"category"`
	OsVersionMin              int                  `bson:"osVersionMinV2,omitempty" json:"osVersionMinV2"`
	OsVersionMax              int                  `bson:"osVersionMaxV2,omitempty" json:"osVersionMaxV2"`
	AdSourceId                int8                 `bson:"adSourceId,omitempty" json:"adSourceId"`
	LandingPageUrl            string               `bson:"landingPageUrl,omitempty" json:"landingPageUrl"`
	Domain                    string               `bson:"appDomain,omitempty" json:"appDomain"`
	Developer                 string               `bson:"developer,omitempty" json:"developer"`
	DirectPkg                 string               `bson:"directPackageName,omitempty" json:"directPackageName"`
	ApkUrl                    string               `bson:"apkUrl,omitempty" json:"apkUrl"`
	System                    []int                `bson:"system,omitempty" json:"system"`             //3 -m系统 5 dsp
	DeviceTypeV2              []uint8              `bson:"deviceTypeV2,omitempty" json:"deviceTypeV2"` //4 phone 5 tablet
	RetargetOffer             uint8                `bson:"retargetOffer,omitempty" json:"retargetOffer"`
	Network                   uint8                `bson:"network,omitempty" json:"network"`
	NetworkTypeV2             []uint8              `bson:"networkTypeV2,omitempty" json:"networkTypeV2"` //0 all 9 wifi
	MobileCode                []string             `bson:"mobileCode,omitempty" json:"mobileCode"`       //运营商all
	DeviceModelV3             map[string][]string  `bson:"deviceModelV3,omitempty" json:"deviceModelV3"`
	TrackingUrl               string               `bson:"trackingUrl,omitempty" json:"trackingurl"`
	FrequencyCap              int                  `bson:"frequencyCap,omitempty" json:"frequencyCap"`
	AdUrlList                 []string             `bson:"adUrlList,omitempty" json:"adUrlList"`
	AdvImp                    []TrackUrl           `bson:"advImp,omitempty" json:"advImp"`
	UserInterest              map[string][]string  `bson:"userInterest,omitempty" json:"userInterest"`
	GenderV2                  []int                `bson:"genderV2,omitempty" json:"genderV2"`
	InventoryV2               InventoryV2          `bson:"inventoryV2,omitempty" json:"inventoryV2"`
	DeviceId                  string               `bson:"deviceId,omitempty"  json:"deviceId"`
	TrafficType               []uint8              `bson:"trafficType,omitempty" json:"trafficType"`
	VbaTrackingLink           string               `bson:"vbaTrackingLink,omitempty" json:"vbaTrackingLink"`
	VbaConnecting             int                  `bson:"vbaConnecting,omitempty" json:"vbaConnecting"`
	RetargetingDevice         int                  `bson:"retargetingDevice,omitempty" json:"retargetingDevice"` // 1.Yes 2.No   默认 No。
	AdSchedule                map[string][]int     `bson:"adSchedule,omitempty" json:"adSchedule"`
	UserAgeV2                 []int                `bson:"userAgeV2,omitempty" json:"userAgeV2"`
	StartTime                 int64                `bson:"startTime,omitempty" json:"startTime"`
	EndTime                   int64                `bson:"endTime,omitempty" json:"endTime"`
	BudgetFirst               bool                 `bson:"budgetFirst,omitempty" json:"budgetFirst"`
	ExcludeRule               ExcludeRule          `bson:"excludeRule,omitempty" json:"excludeRule"`
	InventoryBlackList        []string             `bson:"inventoryBlackList,omitempty" json:"inventoryBlackList"`
	PreviewUrl                string               `bson:"previewUrl,omitempty" json:"previewUrl"` //storeUrl
	IsCampaignCreative        int                  `bson:"isCampaignCreative,omitempty" json:"isCampaignCreative"`
	OriPrice                  float64              `bson:"oriPrice,omitempty" json:"oriPrice"`
	CostType                  int                  `bson:"costType,omitempty" json:"costType"`
	Source                    int                  `bson:"source,omitempty" json:"source"`
	Direct                    int                  `bson:"direct,omitempty" json:"direct"`                   // 单子类型，direct - 1 直单，直接客户单子；direct - 2 非直单，二手单
	ContentRating             int                  `bson:"contentRating,omitempty" json:"contentRating"`     // 年龄分级
	ContentRatingV2           map[string]int       `bson:"contentRatingV2,omitempty" json:"contentRatingV2"` // 年龄分级按照国家区分
	ThirdParty                string               `bson:"thirdParty,omitempty" json:"thirdParty"`           // 单子类型，
	SubCategoryId             int                  `bson:"subCategoryId,omitempty" json:"subCategoryId"`
	SubCategoryV2             []int                `bson:"subCategoryV2,omitempty" json:"subCategoryV2"`
	CategoryStr               string               //应用市场的分类
	SubCategoryStr            string               //应用市场的二级分类
	PackageDevice             string               `bson:"packageDevice,omitempty" json:"packageDevice"`
	SubCategoryName           []string             `bson:"subCategoryName,omitempty"` //过滤market分类，目前仅用于京东exclude_category
	InstallApps               []int                `bson:"installApps,omitempty" json:"installApps"`
	ExcludeInstalledApps      []int                `bson:"excludeInstalledApps,omitempty" json:"excludeInstalledApps"`
	UserInterestV2            [][]int              `bson:"userInterestV2,omitempty" json:"userInterestV2"`
	DeviceLanguage            map[string][]string  `bson:"deviceLanguage,omitempty" json:"deviceLanguage"`
	AdxWhiteBlack             map[string][]string  `bson:"adxWhiteBlack,omitempty" json:"adxWhiteBlack"`
	SChnlPrice                []ChannelPrice       `bson:"dspChanlPrice,omitempty" json:"dspChanlPrice"`
	IsEcAdv                   int                  `bson:"isEcAdv,omitempty" json:"isEcAdv"`
	RetargetVisitorType       int                  `bson:"retargetVisitorType,omitempty" json:"retargetVisitorType"`
	AdxInclude                []string             `bson:"adxInclude,omitempty" json:"adxInclude"`
	IndustryId                int64                `bson:"industryId,omitempty" json:"industryId"`         //广告主行业Id
	RedirectStatus            int                  `bson:"redirectStatus,omitempty" json:"redirectStatus"` //重定向cap控制
	DeepLink                  string               `bson:"deepLink,omitempty" json:"deepLink"`
	IsRetarget                int                  `bson:"isRetargeting,omitempty" json:"isRetargeting"` //1表示拉活, 2表示其他
	BtV3                      BtV3Detail           `bson:"btV3,omitempty" json:"btV3"`
	SpecialType               []int64              `bson:"specialType,omitempty" json:"specialType"`
	WxAppId                   string               `bson:"wxAppId,omitempty" json:"wxAppId"`
	MCountryChanlPrice        map[string]float64   `bson:"mCountryChanlPrice,omitempty" json:"mCountryChanlPrice"`
	MInventoryV2              MInventoryV2Detail   `bson:"mInventoryV2,omitempty" json:"mInventoryV2"`
	JmIcon                    uint8                `bson:"jmIcon,omitempty" json:"jmIcon"`
	TargetSdkVersion          TargetSdkVersionInfo `bson:"targetSdkVersion,omitempty" json:"targetSdkVersion"`
	BtBlackList               BtBlackListInfo      `bson:"btBlackList,omitempty" json:"btBlackList"`
	DeviceImei                uint8                `bson:"deviceImei,omitempty" json:"deviceImei"`
	DeviceAndroidId           uint8                `bson:"deviceAndroidId,omitempty" json:"deviceAndroidId"`
	DeviceOaid                uint8                `bson:"deviceOaid,omitempty" json:"deviceOaid"`
	DeviceGaid                uint8                `bson:"deviceGaid,omitempty" json:"deviceGaid"`
	DeviceSupportMd5          uint8                `bson:"deviceSupportMd5,omitempty" json:"deviceSupportMd5"`
	TryNewCap                 int                  `bson:"tryNewCap,omitempty" json:"tryNewCap"`
	UserActivation            int                  `bson:"userActivation,omitempty" json:"userActivation"` //是否是拉活单子标记，为 1 表示拉活单子，其他枚举值扩展业务含义暂无。
	Audience                  []int64              `bson:"audience,omitempty" json:"audience"`
	AudienceExclude           []int64              `bson:"audienceExclude,omitempty" json:"audienceExclude"`
	CType                     int                  `bson:"ctype,omitempty" json:"ctype"`         //支付类型： 'cpi' => 1,'cpc' => 2,'cpm' => 3,'cpa' => 4,'cpe' => 5
	OfferType                 int                  `bson:"offerType,omitempty" json:"offerType"` //
	Created                   int64                `bson:"created,omitempty" json:"created"`
	DSPCountryChannelPrice    map[string]float64   `bson:"dspCountryChanlPrice,omitempty" json:"dspCountryChanlPrice"`
	DSPCountryChannelOriPrice map[string]float64   `bson:"dspCountryChanlOriPrice,omitempty" json:"dspCountryChanlOriPrice"`
	CountryChannelOriPrice    map[string]float64   `bson:"countryChanlOriPrice,omitempty" json:"countryChanlOriPrice"`
	CrAbTestStart             int8                 `bson:"crAbTestStart,omitempty" json:"crAbTestStart"`             //该单子是否开启素材abtest开关 默认关闭，1:开启 非1:关闭
	CrAbTestTrafficRate       int32                `bson:"crAbTestTrafficRate,omitempty" json:"crAbTestTrafficRate"` //该单子素材abtest的比例控制 默认为0，范围: [0,100]

	TemplateConf      map[string]DspTemplateConfElement `bson:"dspTemplateConf,omitempty" json:"dspTemplateConf"` //模板定向，key无任何意义，是RDS保存的唯一值，见 http://confluence.mobvista.com/pages/viewpage.action?pageId=6984952#id-%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3-ETL-V1.8.15-Creativesv3%EF%BC%88%E7%B4%A0%E6%9D%90%E4%B8%89%E6%9C%9F%EF%BC%89-2.1.1Campaign%E8%A1%A8
	AdTypeTemplateMap map[string]map[string]int

	SdkDetect            SdkDetect              `bson:"sdkDetect,omitempty" json:"sdkDetect"`
	TrackingUrlHttps     string                 `bson:"trackingUrlHttps,omitempty" json:"trackingurlHttps"`
	TotalBudget          int64                  `bson:"totalBudget,omitempty" json:"totalBudget"` //总预算
	DailyBudget          int64                  `bson:"dailyBudget,omitempty" json:"dailyBudget"` //日预算
	OpenType             int8                   `bson:"openType,omitempty" json:"openType"`
	EventCountryOriPrice map[string]*EventPrice `bson:"eventCountryOriPrice,omitempty" json:"eventCountryOriPrice,omitempty"` // 事件price
	MultiSettlement      int                    `bson:"multiSettlement,omitempty" json:"multiSettlement,omitempty"`           // 多事件标记
	LocalPrice           *LocalPrice            `bson:"localPrice,omitempty" json:"localPrice,omitempty"`                     // local 币种
	CreateSource         int                    `bson:"createSrc,omitempty" json:"createSrc,omitempty"`
	StatusUpdated        int64                  `bson:"statusUpdated,omitempty" json:"statusUpdated,omitempty"`
	JumpTypeConfig2      map[string]int32       `bson:"JUMP_TYPE_CONFIG_2,omitempty" json:"JUMP_TYPE_CONFIG_2,omitempty"`
	Ext                  interface{}            // 提供给外部包的扩展字段，非通用的。如果是多个模块都需要的扩展字段，建议还是新增字段好一些
	DeepTarget           []DeepTarget           `bson:"deepTarget,omitempty" json:"deepTarget"` //深度优化

	/*add for as 模板召回配置*/
	AsTemplateConf             map[string]AsTemplateConfInfo `bson:"asTemplateConf,omitempty" json:"asTemplateConf"`
	asTemplateGroupWhiteList   map[int]map[int]bool          //map[adtype]map[group]bool
	asVideoTemplateWhiteList   map[int]map[string]bool       //map[adtype]map[template]bool
	asEndCardTemplateWhiteList map[int]map[string]bool       //map[adtype]map[template]bool
}

//DeepTarget 深度优化
type DeepTarget struct {
	TargetName   string             `bson:"targetName,omitempty" json:"targetName"`     //深度优化目标
	DefaultPrice float64            `bson:"defaultPrice,omitempty" json:"defaultPrice"` //深度优化目标出价 （USD）, 如果某个国家地区没有配置出价 ，就用这个价格。
	GeoPrice     []GeoPrice         `bson:"geoPrice,omitempty" json:"geoPrice"`         //分国家地区出价
	GeoPriceMap  map[string]float64 //分国家地区出价的map形式
}

//广告主定向设置
type MInventoryV2 struct {
	// 广告主定向adtype。@sample=[1, 2, 3, 4, 6, 7, 8, 9, 11, 10]
	AdnAdType    []int `bson:"adnAdtype,omitempty" json:"adnAdtype"`
	adnAdTypeMap map[int]bool
}

//GeoPrice 分国家地区出价
type GeoPrice struct {
	Geo                        string                            `bson:"geo,omitempty" json:"geo"`     //出价地区
	Price                      float64                           `bson:"price,omitempty" json:"price"` //地区出价价格
	CampaignId                 int64                             `bson:"campaignId,omitempty" json:"campaignId"`
	PackageName                string                            `bson:"packageName,omitempty" json:"packageName"`
	AdvertiserId               int64                             `bson:"advertiserId,omitempty" json:"advertiserId"`
	PublisherId                int64                             `bson:"publisherId,omitempty" json:"publisherId"`
	AppName                    string                            `bson:"appName,omitempty" json:"appName"`
	CountryCode                []string                          `bson:"countryCode,omitempty" json:"countryCode"`
	Status                     int8                              `bson:"status,omitempty" json:"status"`
	Updated                    int64                             `bson:"updated,omitempty" json:"updated"`
	SamePackageCreative        int                               `bson:"samePackageCreative,omitempty" json:"samePackageCreative"` //SamePackageCreative，同包名素材开关，枚举值=2 禁止召回同包名素材，枚举值=1 允许召回同包名素材，默认情况 允许召回
	UserActivation             int                               `bson:"userActivation,omitempty" json:"userActivation"`           //是否是拉活单子标记，为 1 表示拉活单子，其他枚举值扩展业务含义暂无。
	Tag                        int                               `bson:"tag,omitempty" json:"tag"`
	CampaignType               int8                              `bson:"campaignType,omitempty" json:"campaignType"`
	TemplateConf               map[string]DspTemplateConfElement `bson:"dspTemplateConf,omitempty" json:"dspTemplateConf"` //模板定向，key无任何意义，是RDS保存的唯一值，见 http://confluence.mobvista.com/pages/viewpage.action?pageId=6984952#id-%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3-ETL-V1.8.15-Creativesv3%EF%BC%88%E7%B4%A0%E6%9D%90%E4%B8%89%E6%9C%9F%EF%BC%89-2.1.1Campaign%E8%A1%A8
	AdTypeTemplateMap          map[string]map[string]int         //定向模板
	MInventoryV2               MInventoryV2                      `bson:"mInventoryV2,omitempty" json:"mInventoryV2"`
	Direct                     int                               `bson:"direct,omitempty" json:"direct"`         // 单子类型，direct - 1 直单，直接客户单子；direct - 2 非直单，二手单
	AdSourceId                 int8                              `bson:"adSourceId,omitempty" json:"adSourceId"` //广告来源id 1)如果 campaign_list.tag = CAMPAIGN_TAG_MY_OFFER (3)  时， 为 AD_SOURCE_MY_OFFER (2) , 否则为 AD_SOURCE_API_OFFER (1) .
	SdkDetect                  SdkDetect                         `bson:"sdkDetect,omitempty" json:"sdkDetect"`
	CType                      int                               `bson:"ctype,omitempty" json:"ctype"` //支付类型： 'cpi' => 1,'cpc' => 2,'cpm' => 3,'cpa' => 4,'cpe' => 5
	CostType                   int                               `bson:"costType,omitempty" json:"costType"`
	OfferType                  int                               `bson:"offerType,omitempty" json:"offerType"`             //
	VbaConnecting              int                               `bson:"vbaConnecting,omitempty" json:"vbaConnecting"`     //表示VTA单子
	VbaTrackingLink            string                            `bson:"vbaTrackingLink,omitempty" json:"vbaTrackingLink"` //展示回传广告主url
	StartTime                  int64                             `bson:"startTime,omitempty" json:"startTime"`
	Created                    int64                             `bson:"created,omitempty" json:"created"`
	OriPrice                   float64                           `bson:"oriPrice,omitempty" json:"oriPrice"`
	CountryChannelPrice        map[string]float64                `bson:"dspCountryChanlPrice,omitempty" json:"dspCountryChanlPrice"`
	CountryChannelOriPrice     map[string]float64                `bson:"dspCountryChanlOriPrice,omitempty" json:"dspCountryChanlOriPrice"`
	CrAbTestStart              int8                              `bson:"crAbTestStart,omitempty" json:"crAbTestStart"`             //该单子是否开启素材abtest开关 默认关闭，1:开启 非1:关闭
	CrAbTestTrafficRate        int32                             `bson:"crAbTestTrafficRate,omitempty" json:"crAbTestTrafficRate"` //该单子素材abtest的比例控制 默认为0，范围: [0,100]
	InstallApps                []int                             `bson:"installApps,omitempty" json:"installApps"`                 //人群包定向
	RedirectStatus             int                               `bson:"redirectStatus,omitempty" json:"redirectStatus"`           //重定向cap控制
	AsTemplateConf             map[string]AsTemplateConfInfo     `bson:"asTemplateConf,omitempty" json:"asTemplateConf"`
	asTemplateGroupWhiteList   map[int]map[int]bool              //map[adtype]map[group]bool
	asVideoTemplateWhiteList   map[int]map[string]bool           //map[adtype]map[template]bool
	asEndCardTemplateWhiteList map[int]map[string]bool           //map[adtype]map[template]bool
}

type AsTemplateConfInfo struct {
	Status          int            `bson:"status,omitempty" json:"status"`
	Adtype          string         `bson:"adtype,omitempty" json:"adtype"`
	TemplateGroup   []TemplateInfo `bson:"templateGroup,omitempty" json:"templateGroup"`
	VideoTemplate   []TemplateInfo `bson:"videoTemplate,omitempty" json:"videoTemplate"`
	EndcardTemplate []TemplateInfo `bson:"endcardTemplate,omitempty" json:"endcardTemplate"`
}

type TemplateInfo struct {
	Type   int64 `bson:"type,omitempty" json:"type"`
	Weight int64 `bson:"weight,omitempty" json:"weight"`
}

type TryNew struct {
	CampId   int64    `bson:"campaignId,omitempty" json:"campaignId"`
	Status   int64    `bson:"status,omitempty" json:"status"`
	Stime    int64    `bson:"stime,omitempty" json:"stime"`
	Etime    int64    `bson:"etime,omitempty" json:"etime"`
	Rate     float32  `bson:"rate,omitempty" json:"rate"`
	IsM      int64    `bson:"is_m,omitempty" json:"is_m"`
	Area     string   `bson:"area,omitempty" json:"area"`
	IsDsp    int64    `bson:"is_dsp,omitempty" json:"is_dsp"`
	Adx      []string `bson:"adx,omitempty" json:"adx"`
	AdType   []int32  `bson:"adType,omitempty" json:"adType"`
	AppId    []string `bson:"appId,omitempty" json:"appId"`
	TagId    []string `bson:"tagId,omitempty" json:"tagId"`
	BidRaise float32  `bson:"bidRaise,omitempty" json:"bidRaise"`
	MTime    int64    `bson:"mTime,omitempty" json:"mTime"`
	Priority int32    `bson:"priority,omitempty" json:"priority"`
}

type TrackUrl struct {
	Sec int
	Url string
}

// 库存类别
type InventoryV2 struct {
	IabCategory map[string][]string `bson:"iabCategory,omitempty" json:"iabCategory"`
	AdnAdtype   []int               `bson:"adnAdtype,omitempty" json:"adnAdtype"`
	AppSite     []string            `bson:"app_site,omitempty" json:"app_site"`
}

type ExcludeRule struct {
	RuleType      int64   `bson:"type,omitempty" json:"type"`
	IncludeAppIds []int64 `bson:"includeAppIds,omitempty" json:"includeAppIds"`
	ExcludeAppIds []int64 `bson:"excludeAppIds,omitempty" json:"excludeAppIds"`
	Status        int64   `bson:"status,omitempty" json:"status"`
	Grades        []int64 `bson:"grades,omitempty" json:"grades"`
}

type ChannelPrice struct {
	Chanl string  `bson:"chanl,omitempty" json:"chanl"` //渠道
	Price float64 `bson:"price,omitempty" json:"price"` //价格
}

type (
	BtV3Detail struct {
		BtClass map[string]BtClassDetail `bson:"btClass,omitempty" json:"btClass"`
	}
	BtClassDetail struct {
		Status uint8 `bson:"status,omitempty" json:"status"`
	}
)

type MInventoryV2Detail struct {
	Units       []int64             `bson:"units,omitempty" json:"units"`
	IabCategory map[string][]string `bson:"iabCategory,omitempty" json:"iabCategory"`
	SubCategory []string            `bson:"subCategory,omitempty" json:"subCategory"`

	AdnAdType    []int64 `bson:"adnAdtype,omitempty" json:"adnAdtype"`
	adnAdTypeMap map[int]bool
}

type TargetSdkVersionInfo struct {
	Include []map[string]string `bson:"include,omitempty" json:"include"`
	Exclude []map[string]string `bson:"exclude,omitempty" json:"exclude"`
}

type SdkVersionItem struct {
	Min    int64
	Max    int64
	Prefix string
}

type BtBlackListInfo struct {
	Status uint8   `bson:"status,omitempty" json:"status"`
	Apps   []int64 `bson:"apps,omitempty" json:"apps"`
}

type (
	DspTemplateConfElement struct {
		Adtype string          `bson:"adtype,omitempty" json:"adtype"`               //adtype 类型
		TGroup []TemplateGroup `bson:"templateGroup,omitempty" json:"templateGroup"` //模板类型
		Status int             `bson:"status,omitempty" json:"status"`               //status类型 1.active 2.pause
	}
	TemplateGroup struct {
		Type   int `bson:"type,omitempty" json:"type"`     //具体的模板组合枚举类型
		Weight int `bson:"weight,omitempty" json:"weight"` //对应的模板组合权重，为0的话，算法自由组合。
	}
)

type SdkDetect struct {
	JumpCount int `bson:"jumpCount,omitempty" json:"jumpCount"`
	Success   int `bson:"success,omitempty" json:"success"`
}

type LocalPrice struct {
	Currency             int                    `bson:"currency,omitempty" json:"currency,omitempty"`
	ExchangeRate         float64                `bson:"exchangeRate,omitempty" json:"exchangeRate,omitempty"`
	OriPrice             float64                `bson:"oriPrice,omitempty" json:"oriPrice,omitempty"`
	CountryChanlOriPrice map[string]float64     `bson:"countryChanlOriPrice,omitempty" json:"countryChanlOriPrice,omitempty"` // 按 国家 + 渠道的  receive price
	EventCountryOriPrice map[string]*EventPrice `bson:"eventCountryOriPrice,omitempty" json:"eventCountryOriPrice,omitempty"`
}

type EventPrice struct {
	EventPrice float64            `bson:"eventPrice,omitempty" json:"eventPrice,omitempty"`
	MatchType  int                `bson:"matchType,omitempty" json:"matchType,omitempty"`
	GeoPrice   map[string]float64 `bson:"geoPrice,omitempty" json:"geoPrice,omitempty"` // 按 国家 + price
}

// ----------------------------------------------- Get Attr ------------------------------------------------------------

func (camp *Campaign) AllowSamePackageCreative() bool {
	//SamePackageCreative，同包名素材开关，枚举值=2 禁止召回同包名素材，枚举值=1 允许召回同包名素材，默认情况 允许召回
	return camp.SamePackageCreative != 2
}

func (camp *Campaign) GetAdTypeTemplateMap() map[string]map[string]int {
	if camp.AdTypeTemplateMap == nil {
		return map[string]map[string]int{}
	}
	// key = adtype(as) - templateid : value = weight
	return camp.AdTypeTemplateMap
}

func (camp *Campaign) GetAdnAdTypeMap() map[int]bool {
	if camp.MInventoryV2.adnAdTypeMap == nil {
		return map[int]bool{}
	}
	// key = adnAdtype : value = exist
	return camp.MInventoryV2.adnAdTypeMap
}

func (camp *Campaign) GetGivenType() int {
	if camp.CostType == 0 {
		camp.CostType = camp.CType
	}

	return int(camp.CostType)
}

func (camp *Campaign) GetRecvType() int {
	return camp.CType
}

//GetGivenPirce  获取单子渠道价格，按照
func (camp *Campaign) GetRecvPrice(country, appId string) float64 {
	if len(camp.DSPCountryChannelOriPrice) == 0 {
		return camp.OriPrice
	}
	appId = strings.Replace(appId, "-", "--", -1)
	appId = strings.Replace(appId, ".", "-", -1)
	for _, key := range []string{
		country + "_" + appId,
		country,
	} {
		if price, ok := camp.DSPCountryChannelOriPrice[key]; ok {
			return price
		}
	}
	return camp.OriPrice
}

//GetGivenPirce  获取单子渠道价格，按照
func (camp *Campaign) GetGivenPrice(country, exchange, appId string) float64 {
	if len(camp.DSPCountryChannelPrice) == 0 {
		return camp.Price
	}
	appId = strings.Replace(appId, "-", "--", -1)
	appId = strings.Replace(appId, ".", "-", -1)
	for _, key := range []string{
		country + "_" + exchange + "_" + appId,
		country + "_all_" + appId,
		country + "_" + appId,
		country + "_" + exchange + "_all",
		country,
	} {
		if price, ok := camp.DSPCountryChannelPrice[key]; ok {
			return price
		}
	}
	return camp.Price
}

// ----------------------------------------------- Parser --------------------------------------------------------------

func (camp *Campaign) Parser() {
	{
		//单子定向流量/模版，及模板投放权重配置
		if len(camp.TemplateConf) > 0 {
			camp.AdTypeTemplateMap = make(map[string]map[string]int)
		}
		for _, template := range camp.TemplateConf {
			if template.Status != 1 || len(template.TGroup) == 0 {
				continue
			}
			adType := strconv.Itoa(int(enum.GetAdTypeDspToAs(template.Adtype)))
			camp.AdTypeTemplateMap[adType] = make(map[string]int)
			for _, item := range template.TGroup {
				camp.AdTypeTemplateMap[adType][strconv.Itoa(item.Type)] = item.Weight
			}
		}
	}
	{
		//广告主模版定向
		if len(camp.MInventoryV2.AdnAdType) > 0 {
			camp.MInventoryV2.adnAdTypeMap = make(map[int]bool)
		}
		for _, adnAdtype := range camp.MInventoryV2.AdnAdType {
			camp.MInventoryV2.adnAdTypeMap[int(adnAdtype)] = true
		}
	}
	{
		//单子结算类型
		if camp.CostType == 0 {
			camp.CostType = camp.CType
		}
		camp.CType = camp.CostType
	}
	{
		//For as, groupTemplate, videoTemplate, endcardTemplate,需要注意201
		for _, templateConf := range camp.AsTemplateConf {

			if templateConf.Status != 1 {
				continue
			}

			adtype, err := strconv.Atoi(templateConf.Adtype)
			if err != nil {
				continue
			}

			if camp.asTemplateGroupWhiteList == nil {
				camp.asTemplateGroupWhiteList = make(map[int]map[int]bool)
			}
			camp.asTemplateGroupWhiteList[adtype] = make(map[int]bool)

			for _, templateGroup := range templateConf.TemplateGroup {

				camp.asTemplateGroupWhiteList[adtype][int(templateGroup.Type)] = true
				if mappingGroup, ok := enum.GroupMapping[int(templateGroup.Type)]; ok {
					for _, groupItem := range mappingGroup {
						camp.asTemplateGroupWhiteList[adtype][groupItem] = true
					}
				}
			}

			for _, videoTemplate := range templateConf.VideoTemplate {
				if camp.asVideoTemplateWhiteList == nil {
					camp.asVideoTemplateWhiteList = make(map[int]map[string]bool)
				}
				if camp.asVideoTemplateWhiteList[adtype] == nil {
					camp.asVideoTemplateWhiteList[adtype] = make(map[string]bool)
				}

				if _, ok := enum.VideoTemplateTypeToVideoTemplateId[int(videoTemplate.Type)]; ok {
					for _, templateId := range enum.VideoTemplateTypeToVideoTemplateId[int(videoTemplate.Type)] {
						camp.asVideoTemplateWhiteList[adtype][templateId] = true
					}
				}
			}

			for _, endcardTemplate := range templateConf.EndcardTemplate {
				if camp.asEndCardTemplateWhiteList == nil {
					camp.asEndCardTemplateWhiteList = make(map[int]map[string]bool)
				}
				if camp.asEndCardTemplateWhiteList[adtype] == nil {
					camp.asEndCardTemplateWhiteList[adtype] = make(map[string]bool)
				}

				if _, ok := enum.EndcardTemplateTypeToEndcardTemplateId[int(endcardTemplate.Type)]; ok {
					for _, templateId := range enum.EndcardTemplateTypeToEndcardTemplateId[int(endcardTemplate.Type)] {
						camp.asEndCardTemplateWhiteList[adtype][templateId] = true
					}
				}
			}

		}
	}

}

//add for as
func (this *Campaign) CheckAsTemplateGroup(adtype int, group int) bool {

	if templateGroupMap, ok := this.asTemplateGroupWhiteList[adtype]; ok {
		if len(templateGroupMap) != 0 {
			if _, groupOk := templateGroupMap[group]; groupOk {
				return true
			}
		} else {
			if _, groupOk := AsDefaultTemplateGroupMap[group]; groupOk {
				return true
			}
		}
	} else {
		if _, groupOk := AsDefaultTemplateGroupMap[group]; groupOk {
			return true
		}
	}
	return false
}

func (this *Campaign) GetAsTemplate(adtype int) (map[string]bool, map[string]bool) {
	return this.asVideoTemplateWhiteList[adtype], this.asEndCardTemplateWhiteList[adtype]
}
func CheckAsVirtualTemplateStr(templateId string) bool {
	if templateId == "7002001" || templateId == "8002001" || templateId == "7002000" || templateId == "0" ||
		templateId == "-5002000__Low" || templateId == "-5002000__High" {
		return true
	}
	return false
}

func (this *Campaign) GetAsRecalledTemplate(adtype int, templateIds map[string]bool) (map[string]bool, bool) {
	videoTemplateIds, endcardTemplateIds := this.GetAsTemplate(adtype)
	intersecRecallTemplateIds := make(map[string]bool)

	for templateId, _ := range templateIds {
		//minicard && 封面模板不受约束
		/*
			if templateId == "7002001" || templateId == "8002001" || templateId == "7002000" || templateId == "0" {
				intersecRecallTemplateIds[templateId] = true
				continue
			}*/

		if CheckAsVirtualTemplateStr(templateId) {
			intersecRecallTemplateIds[templateId] = true
			continue
		}
		if len(videoTemplateIds) != 0 {
			if _, ok := videoTemplateIds[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		} else {
			if _, ok := AsCampaignDefaultVideoTemplateSet[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		}

		if len(endcardTemplateIds) != 0 {
			if _, ok := endcardTemplateIds[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		} else {
			if _, ok := AsCampaignDefaultEndcardTemplateSet[templateId]; ok {
				intersecRecallTemplateIds[templateId] = true
				continue
			}
		}
	}
	return intersecRecallTemplateIds, len(videoTemplateIds) != 0 && len(endcardTemplateIds) != 0
}

var AsCampaignDefaultVideoTemplateSet = map[string]bool{
	"5002001__Low": true, "5002001__High": true,
	"5002002__Low": true, "5002002__High": true,
	"5002003__Low": true, "5002003__High": true,
	"5002004__Low": true, "5002004__High": true,
	"5002007__Low": true, "5002007__High": true,
	"5002009__Low": true, "5002009__High": true,
	"5002008__Low": true, "5002008__High": true,
}

var AsCampaignDefaultEndcardTemplateSet = map[string]bool{
	"6004001": true, "6003001": true,
	"6002001": true, "6002002": true,
	"6002003": true, "6002007": true,
}

var AsDefaultTemplateGroupMap = map[int]bool{
	1: true, 2: true, 3: true,
	6: true, 7: true, 18: true,
	19: true, 20: true, 21: true,
}

//分国家地区出价结构体转换成map k是国家，v是出价
func (campaign *Campaign) ChangeGeoMap(geoPrice []GeoPrice) map[string]float64 {
	geoPrices := make(map[string]float64)
	for _, geoPriceVal := range geoPrice {
		geoPrices[geoPriceVal.Geo] = geoPriceVal.Price
	}
	return geoPrices
}

//end add for as
