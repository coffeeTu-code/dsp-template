package enum

//  *****  campaign 枚举定义  *****
//  http://confluence.mobvista.com/pages/viewpage.action?pageId=2363102

//********************* Campaign Type *********************
const (
	Campaign_Type_AppStore     = 1
	Campaign_Type_GooglePlay   = 2
	Campaign_Type_APK          = 3
	Campaign_Type_Other        = 4
	Campaign_Type_IPA          = 5
	Campaign_Type_SubScription = 6
	Campaign_Type_WebSite      = 7
	Campaign_Type_RT           = 10
	Campaign_Type_ForBit       = 99
)

//支付类型
//********************* Cost Type *********************
type CampaignCostType int32

const (
	Campaign_CostType_CPI = 1
	Campaign_CostType_CPC = 2
	Campaign_CostType_CPM = 3
	Campaign_CostType_CPA = 4
	Campaign_CostType_CPE = 5
)

var CampaignCostType_name = map[int32]string{
	Campaign_CostType_CPI: "cpi",
	Campaign_CostType_CPC: "cpc",
	Campaign_CostType_CPM: "cpm",
	Campaign_CostType_CPA: "cpad",
	Campaign_CostType_CPE: "cpe",
}

func (x CampaignCostType) String() string {
	return CampaignCostType_name[int32(x)]
}

//campaign status
//参看：SS Offer状态整理  https://confluence.mobvista.com/pages/viewpage.action?pageId=16695287
//********************* campaign status *********************
const (
	MongoStatus_UNKNOWN             = 0
	MongoStatus_ACTIVE              = 1
	MongoStatus_PAUSED              = 2
	MongoStatus_DELETED             = 3
	MongoStatus_PENDING             = 4
	MongoStatus_REJECTED            = 5
	MongoStatus_UNFINISHED          = 6
	MongoStatus_REVIEW              = 7
	MongoStatus_OUT_OF_DAILY_BUDGET = 8
	MongoStatus_OUT_OF_TOTAL_BUDGET = 9
	MongoStatus_BALACNE             = 10
	MongoStatus_DAILY_CAP           = 11
	MongoStatus_PAUSED_FROM_ACTIVE  = 12 // Advertiser Paused
	MongoStatus_PAUSED_FROM_PENDING = 13 // Advertiser Pending
	// 20 广告主手动 pending
	// 21 广告主手动stopped
	// 22 Out of Advertiser Daily Budget
	// 23 Out of Advertiser Total Budget
	// 24 Insufficient Account Balance // 由于Adv的余额消耗完，停停掉的offer
)

//广告来源
//********************* Campaign Tag *********************
const (
	Campaign_Tag_UNKNOWN    = 0
	Campaign_Tag_ADMIN      = 1
	Campaign_Tag_ADVERTISER = 2
	Campaign_Tag_M_MYOFFER  = 3
	Campaign_Tag_BRANDOFFER = 4
)
