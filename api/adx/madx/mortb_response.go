package madx

import (
	"bytes"
	"strings"
)

type MOrtbResponse struct {
	ID       string      `json:"id"`              // *竞价请求的标识
	SeatBid  []*MSeatBid `json:"seatbid"`         // *一组SeatBid对象， 如果出价，则至少应该填充一个
	BidID    string      `json:"bidid,omitempty"` // *竞拍者生成的响应ID, 辅助日志或者交易追踪
	Currency string      `json:"cur,omitempty"`   // *使用ISO-4217码表标识货币类型
	NBR      int         `json:"nbr,omitempty"`   // 不出价的原因， 参考表5.19
}

type MSeatBid struct {
	Bid  []*MBid `json:"bid"`            // *至少一个Bid对象的数组，每个对象关联一个展示。多个出价可以关联同一个展示
	Seat string  `json:"seat,omitempty"` // 出价者席位标识， 代表本次出价的出价人
}
type MBid struct {
	ID             string   `json:"id"`                       // *竞拍者生成的竞价ID，用于记录日志或行为追踪
	ImpID          string   `json:"impid"`                    // *关联的竞价请求中的Imp对象的ID
	Price          float64  `json:"price"`                    // *虽然本次只是对某一个展示的出价，但是竞拍价格是以CPM表示。需要注意数据类型是float,所以在处理货币的时候强烈建议使用相关的数学处理对象（比如，Java中的BigDecimal)
	AdID           string   `json:"adid,omitempty"`           // 预加载的广告ID, 可以在交易胜出的时候使用
	NURL           string   `json:"nurl,omitempty"`           // *胜出通知地址
	BURL           string   `json:"burl,omitempty"`           // 胜出通知地址， 如果竞价胜出的时候由交易平台调用； 可选标识serving ad markup
	AdMarkup       string   `json:"adm,omitempty"`            // *Actual ad markup. XHTML if a response to a banner object, or VAST XML if a response to a video object.
	AdvDomain      []string `json:"adomain,omitempty"`        // *Advertiser’s primary or top-level domain for advertiser checking; or multiple if imp rotating.
	Bundle         string   `json:"bundle,omitempty"`         // *A platform-specific application identifier intended to be unique to the app and independent of the exchange.
	IURL           string   `json:"iurl,omitempty"`           // *Sample image URL.
	CampaignID     string   `json:"cid,omitempty"`            // *Campaign ID that appears with the Ad markup.
	CreativeID     string   `json:"crid,omitempty"`           // *Creative ID for reporting content issues or defects. This could also be used as a reference to a creative ID that is posted with an exchange.
	Cat            []string `json:"cat,omitempty"`            // creative的IAB内容类型，参考表5.1
	Attr           []int    `json:"attr,omitempty"`           // 描述creative的属性集合，参考表5.3
	API            int      `json:"api,omitempty"`            // 本次展示支持的API框架列表， 参考5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	Protocol       int      `json:"protocol,omitempty"`       // 支持的视频竞价响应协议。参考5.8.至少一个支持的协议必须在protocol或者protocols属性中被指定
	QAGMediaRating int      `json:"qagmediarating,omitempty"` // Creative media rating per IQG guidelines.
	H              int      `json:"h,omitempty"`              // *creative 的高度， 以像素为单位
	W              int      `json:"w,omitempty"`              // *creative 的宽度， 以像素为单位
	Exp            int      `json:"exp,omitempty"`            // Advisory as to the number of seconds the bidder is willing to wait between the auction and the actual impression.
	Ext            *MBidExt `json:"ext,omitempty"`            // ext扩展
}

type MBidExt struct {
	CrType              string            `json:"crtype,omitempty"`          //*required.枚举 "VAST 2.0", "VAST 3.0", "MRAID 1.0", "MRAID 2.0", "MRAID playable", "Image Ad", "HTML", "HTML5", "JS", "native"
	Duration            int               `json:"duration,omitempty"`        // 返回为视频时必传，Length/duration of the video in seconds
	AppName             string            `json:"appname,omitempty"`         //*app name
	StoreUrl            string            `json:"storeurl,omitempty"`        //*商店地址
	CreativeGroupId     string            `json:"creativegroupId,omitempty"` // 审核通过的素材组ID，请求ext中有ADUnit时必须返回。
	CreativeSpec        string            `json:"creativespec,omitempty"`    // 国内adx等指定投递广告位，请求ext中ADUnit的其中之一，与上面CreativeGroupId关联。
	Imptrackers         []string          `json:"imptrackers,omitempty"`     // 展示跟踪地址
	ImptrackerMap       map[string]string `json:"imptrackermap,omitempty"`   // 展示跟踪地址,通过key区分不同的url, value为对应的url.
	Clicktrackers       []string          `json:"clicktrackers,omitempty"`   // 点击跟踪地址
	ClickParam          string            `json:"clickparam,omitempty"`      // 替换DSP 点击监测 URL 里的宏__CLICK_PARAM__ 请求为true时必须返回，false无需返回
	ImpressionParam     string            `json:"impressionparam,omitempty"` // 替换DSP 曝光监测 URL（monitor_url1 ）里的宏__IMPRESSION_PARAM__ 请求为true时必须返回，false无需返回
	IsApk               bool              `json:"isapk,omitempty"`           // 单子是否为apk
	LinkType            int               `json:"link_type,omitempty"`
	Deeplink            string            `json:"deeplink,omitempty"`            // deeplink url
	DeeplinkFallbackUrl string            `json:"deeplinkfallbackurl,omitempty"` // deeplink fallback url
	AdvertiserId        int               `json:"advertiserid,omitempty"`        // 广告主ID
	LandingPage         bool              `json:"landingpage,omitempty"`         // 是否投放的是landing page类型的下载单
	VAST                *VAST             `json:"vast,omitempty"`                // 返回里的vast结构体
}

type AdmNative struct {
	Native *MNativeRes `json:"native,omitempty"` // native
}

type AdmVast struct {
	Vast *VAST `json:"vasttag,omitempty"` // vast
}

type MNativeRes struct {
	Ver         string       `json:"ver,omitempty"`         // Version of the Native Markup
	Assets      []*MResAsset `json:"assets"`                // An array of Asset Objects
	Link        *MLink       `json:"link"`                  // Destination Link. This is default link object for the ad
	ImpTrackers []string     `json:"imptrackers,omitempty"` // Array of impression tracking URLs, expected to return a 1x1 image or 204 response
}
type MLink struct {
	URL           string   `json:"url"`                     // Landing URL of the clickable link
	ClickTrackers []string `json:"clicktrackers,omitempty"` // List of third-party tracker URLs to be fired on click of the URL
}

type MResAsset struct {
	Id       int        `json:"id,omitempty"`
	Required int        `json:"required,omitempty"`
	Title    *MResTitle `json:"title,omitempty"`
	Image    *MResImage `json:"img,omitempty"`
	Data     *MResData  `json:"data,omitempty"`
	Video    *MResVideo `json:"video,omitempty"`
	Link     *MLink     `json:"link,omitempty"` // Link object for call to actions. The link object applies if the asset item is activated (clicked)
	//Type     string     `json:"type,omitempty"`
	//AssetType int        `json:"AssetType,omitempty"` //IsTitle 1, IsImg 2, IsData 3, IsVideo 4
}

type MResTitle struct {
	Text string `json:"text"`
}

type MResImage struct {
	Url    string `json:"url"`
	Height int    `json:"h"`
	Width  int    `json:"w"`
}

type MResData struct {
	//Type  uint8  `json:"type"`
	Label string `json:"label,omitempty"` // The optional formatted string name of the data type to be displayed
	Value string `json:"value"`           // The formatted string of data to be displayed. Can contain a formatted value such as “5 stars” or “$10” or “3.4 stars out of 5”
}

type MResVideo struct {
	VastTag string `json:"vasttag"`
}

func (this *MOrtbResponse) FromJson(data []byte) error {
	return jsonit.Unmarshal(data, &this)
}

func (this *MOrtbResponse) String() (string, error) {
	if str, err := jsonit.Marshal(this); err != nil {
		return "", err
	} else {
		return string(str), nil
	}
}

func (this *MOrtbResponse) ToUnescapeJson() (jsonstr string, err error) {

	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := jsonit.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false) // 防止<、>被替换成\u003c、\u003e
	err = jsonEncoder.Encode(this)
	//fmt.Println(bf.String())
	jsonstr = bf.String()
	return strings.TrimRight(jsonstr, "\n"), err
	/*
		var str []byte

		if str, err = jsonit.Marshal(this); err != nil {
			return "", err
		}
		jsonstr = string(str)
		jsonstr = strings.Replace(jsonstr, "\\u003c", "<", -1)
		jsonstr = strings.Replace(jsonstr, "\\u003e", ">", -1)
		jsonstr = strings.Replace(jsonstr, "\\u0026", "&", -1)
		//rubicon not allow return blank nurl ,it is for video
		jsonstr = strings.Replace(jsonstr, `"nurl":"",`, "", -1)
		jsonstr = strings.Replace(jsonstr, `"adm":"",`, "", -1)
		jsonstr = strings.Replace(jsonstr, `"adviewnurl":`, `"nurl":`, -1)
		return jsonstr, err*/
}

func (this *AdmNative) ToString() (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := jsonit.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false) // 防止<、>被替换成\u003c、\u003e
	if err := jsonEncoder.Encode(this); err != nil {
		return "", err

	} else {
		//fmt.Println(bf.String())
		return strings.TrimRight(bf.String(), "\n"), nil
	}

	/*if str, err := jsonit.Marshal(this); err != nil {
		return "", err
	} else {
		return string(str), nil
	}*/
}

func (this *AdmNative) ToAdmNative(nativeStr string) (*AdmNative, error) {
	err := jsonit.UnmarshalFromString(nativeStr, this)
	return this, err
}
