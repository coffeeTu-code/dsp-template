package madx

import (
	"bytes"
	"strings"
)

type MOrtbRequest struct {
	ID          string        `json:"id"`                // 竞价请求的唯一ID, 由广告交易平台提供
	Imp         []*MImp       `json:"imp,omitempty"`     // 代表提供的展示信息的数组， 要求至少有一个
	Site        *MSite        `json:"site,omitempty"`    // Site对象，表示发布者站点相关的详细信息， 仅仅对站点适用且推荐填充
	App         *MApp         `json:"app,omitempty"`     // App对象，表示发布应用的详细信息。仅对应用适用且推荐填充
	Device      *MDevice      `json:"device,omitempty"`  // Device对象， 表示展示将要被发送到的用户设备的信息
	User        *MUser        `json:"user,omitempty"`    // User对象， 表示使用设备的对象， 广告的受众
	AuctionType int           `json:"at"`                // 拍卖类型（胜出策略）1表示第一价格 ，2标识第二价格。交易特定的拍卖类型可以用大于500的值定义
	TMax        int           `json:"tmax,omitempty"`    // 用于在提交竞价请求时避免超时的最大时间，以毫秒为单位，这个值通常是线下沟通的
	AllImps     int           `json:"allimps,omitempty"` // 用于标识交易平台是否可以验证当前的展示列表包含了当前上下文中所有展示。（例如，一个页面上的所有广告位，所有的视频广告点（视频前，视频中，时候后））用于支持路由封锁。 0表示不支持或未知， 1表示提供的展示列表代表所有可用的展示。 Default: 0
	BCat        []string      `json:"bcat,omitempty"`    // 被封锁的广告主类别，使用IAB 内容类别，参考5.1。
	BAdv        []string      `json:"badv,omitempty"`    // 域名封锁列表（比如 ford.com)
	BApp        []string      `json:"bapp,omitempty"`    // Block list of applications by their platform-specific exchange-independent application identifiers. On Android, these should be bundle or package names (e.g., com.foo.mygame).  On iOS, these are numeric IDs.
	BSeat       []string      `json:"bseat,omitempty"`
	WSeat       []string      `json:"wseat,omitempty"`
	Source      *MSource      `json:"source,omitempty"` // A Source object (Section 3.2.2) that provides data about the inventory source and which entity makes the final decision.
	Regs        *MRegulations `json:"regs,omitempty"`   // Reg对象， 指明对本次请求有效的工业，法律或政府条例
	Test        int           `json:"test,omitempty"`   // Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	Ext         *MReqExt      `json:"ext,omitempty"`    // 特定ADX: 特定交易的OpenRTB协议的扩展信息占位符
}

type MReqExtBidFeedback struct {
	RequestId       string  `protobuf:"bytes,1,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	MinimumBidToWin float64 `protobuf:"fixed64,6,opt,name=minimum_bid_to_win,json=minimumBidToWin" json:"minimum_bid_to_win,omitempty"`
}

//MOrtbRequest ext
type MReqExt struct {
	Exchange     string               `json:"exchange,omitempty"` // adx name
	Targetpkgids []string             `json:"targetpkgids,omitempty"`
	ABTest       map[string]string    `json:"abtest,omitempty"`       //adx与dsp的上下游test标记
	UpstreamTags map[string]string    `json:"upsteam_tags,omitempty"` // 扩展字段，用于存储来自外部ADX的原始内容
	BidFeedback  []MReqExtBidFeedback `protobuf:"bytes,1,rep,name=bid_feedback,json=bidFeedback" json:"bid_feedback,omitempty"`
	//OnlyNURL int    `json:"onlynurl,omitempty"` // 当值为1时，获胜仅通知nurl，不会通知impression tracking url
	//VastRecall      int                          `json:"vastrecall,omitempty"`      // bid时返回Vast url地址,获胜后访问vast地址，获取完整vast协议。默认 0 否， 1 是
	//ADUnit          []string                     `json:"adunit,omitempty"`          // 腾讯adx等本次请求指定投递广告位
	//NeedAudit       int                          `json:"needaudit,omitempty"`       // 素材是否需要审核过 默认0 否   1 返回的素材是送审通过的素材
	//MDSPParams      map[string]map[string]string `json:"mdspparams,omitempty"`      // m dsp 各处理模块需要的参数
}

// This object describes the nature and behavior of the entity that is the source of the bid request
// upstream from the exchange. The primary purpose of this object is to define post-auction or upstream
// decisioning when the exchange itself does not control the final decision. A common example of this is
// header bidding, but it can also apply to upstream server entities such as another RTB exchange, a
// mediation platform, or an ad server combines direct campaigns with 3rd party demand in decisioning
type MSource struct {
	FD     int         `json:"id,omitempty"`     // Entity responsible for the final impression sale decision, where 0 = exchange, 1 = upstream source.
	TID    string      `json:"tid,omitempty"`    // Transaction ID that must be common across all participants in this bid request (e.g., potentially multiple exchanges).
	PChain string      `json:"pchain,omitempty"` // Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol v1.0.
	Ext    *MSourceExt `json:"ext,omitempty"`    // 特定交易的OpenRTB协议的扩展信息占位符
}

type MSourceExt struct {
	Schain *MSchain `json:"schain,omitempty"`
	Omidpn string   `json:"omidpn,omitempty"`
	Omidpv string   `json:"omidpv,omitempty"`
}

type MSchain struct {
	Ver      string   `json:"ver"`
	Complete int      `json:"complete"`
	Nodes    []*MNode `json:"nodes,omitempty"`
}

type MNode struct {
	Asi    string `json:"asi,omitempty"`
	Sid    string `json:"sid,omitempty"`
	Rid    string `json:"rid,omitempty"`
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Hp     int    `json:"hp,omitempty"`
}

func (this *MOrtbRequest) FromJson(data []byte) error {
	return jsonit.Unmarshal(data, &this)
}

func (this *MOrtbRequest) String() string {
	if str, err := jsonit.Marshal(this); err != nil {
		return "err content"
	} else {
		return string(str)
	}
}

func (this *MOrtbRequest) ToUnescapeJson() (jsonstr string, err error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := jsonit.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false) // 防止<、>被替换成\u003c、\u003e
	err = jsonEncoder.Encode(this)
	//fmt.Println(bf.String())
	jsonstr = bf.String()
	return strings.TrimRight(jsonstr, "\n"), err
}
