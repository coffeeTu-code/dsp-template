package madx

import (
	"errors"
	"strings"
)

const NativeImage = "native image"
const (
	_ = iota
	IsTitle
	IsImg
	IsData
	IsVideo
)

// Image type. OpenRTB Native Ads 7.4
const (
	_ = iota
	Icon
	Logo
	Main
)

// Data type. OpenRTB Native Ads 7.3
const (
	_ = iota
	Sponsored
	Desc
	Rating
	Likes
	Downloads
	Price
	Saleprice
	Phone
	Address
	Desc2
	Displayurl
	Ctatext
)

type MNativeReq struct {
	Request string `json:"request"`         // 遵守Native Ad规范的请求体
	Ver     string `json:"ver,omitempty"`   // Native Ad规范的版本， 为了高效解析强烈推荐
	API     []int  `json:"api,omitempty"`   // 本次展示支持的API框架列表， 参考表5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	BAttr   []int  `json:"battr,omitempty"` // 限制的物料属性，参考表5.3
	//Ext     *MNativeReqExt `json:"ext,omitempty"`
}
type MNativeReqExt struct {
	//NativeRequest *MNativeRequest `json:"nativerequest,omitempty"` // request string对应的结构体
	Native *MNativeRequestObj `json:"native,omitempty"`
}

//对应MNativeReq中Request stirng json 解析
/*type MNativeRequest struct {
	Native *MNativeRequestObj `json:"native,omitempty"`
}*/
type MNativeRequestObj struct {
	Ver              string       `json:"ver,omitempty"`            // Version of the Native Markup
	ContextTypeID    int          `json:"context,omitempty"`        // The context in which the ad appears
	ContextSubTypeID int          `json:"contextsubtype,omitempty"` // A more detailed context in which the ad appears
	PlacementTypeID  int          `json:"plcmttype,omitempty"`      // The design/format/layout of the ad unit being offered
	PlacementCount   int          `json:"plcmtcnt,omitempty"`       // The number of identical placements in this Layout
	Sequence         int          `json:"seq,omitempty"`            // 0 for the first ad, 1 for the second ad, and so on
	ReqAssets        []*MReqAsset `json:"assets"`                   // An array of Asset Objects
	//Ext              *MNativeExt `json:"ext,omitempty"`
}

//type MNativeExt struct {
//	mainImageSizeExactWH int `json:"mainimagesizeexactwh,omitempty"` //1 如果main image中的w,h不为空，优先使用image中的w和h size. 0或无 w,h, wmin,hmin都可以，默认0。
//}
// The main container object for each asset requested or supported by Exchange
// on behalf of the rendering client.  Only one of the {title,img,video,data}
// objects should be present in each object.  The id is to be unique within the
// AssetObject array so that the response can be aligned.
type MReqAsset struct {
	ID       int             `json:"id"`                 // Unique asset ID, assigned by exchange
	Required int             `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *MReqAssetTitle `json:"title,omitempty"`    // Title object for title assets
	Image    *MReqAssetImage `json:"img,omitempty"`      // Image object for image assets
	Video    *MReqAssetVideo `json:"video,omitempty"`    // Video object for video assets
	Data     *MReqAssetData  `json:"data,omitempty"`     // Data object for brand name, description, ratings, prices etc.
	//AssetType int             `json:"assettype,omitempty"` // mortb，标识该ReqAsset的类型，枚举Title 1,Img 2,Data 3,Video 4
}
type MReqAssetTitle struct {
	Length int `json:"len"` // Maximum length of the text in the title element
}
type MReqAssetImage struct {
	TypeID    int `json:"type,omitempty"` // Type ID of the image element supported by the publisher
	Width     int `json:"w,omitempty"`    // Width of the image in pixels
	WidthMin  int `json:"wmin,omitempty"` // The minimum requested width of the image in pixels
	Height    int `json:"h,omitempty"`    // Height of the image in pixels
	HeightMin int `json:"hmin,omitempty"` // The minimum requested height of the image in pixels
	// Either h/w or hmin/wmin should be transmitted. If only h/w is included, it
	// should be considered an exact requirement
	Mimes []string `json:"mimes,omitempty"` // Whitelist of content MIME types supported
}
type MReqAssetData struct {
	TypeID int `json:"type"` // Type ID of the element supported by the publisher. The publisher can display this information in an appropriate format
	Length int `json:"len"`  // Maximum length of the text in the element’s response
}

type MReqAssetVideo struct {
	H              int                `json:"h"`
	W              int                `json:"w"`
	Mimes          []string           `json:"mimes,omitempty"`          // Whitelist of content MIME types supported
	MinDuration    int                `json:"minduration,omitempty"`    // Minimum video ad duration in seconds
	MaxDuration    int                `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds
	Protocols      []int              `json:"protocols,omitempty"`      // Video bid response protocols
	Linearity      int                `json:"linearity,omitempty"`      // 展示是否必须是线性的， 如果没有指定，则标识都是被允许的，参考5.11
	Skip           int                `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	BAttr          []int              `json:"battr,omitempty"`          // 限制的物料属性，参考5.3
	MinBitrate     int                `json:"minbitrate,omitempty"`     // 最小的比特率，以Kbps为单位。交易平台可以动态的设置这个值或者为所有发布者统一设置该值
	MaxBitrate     int                `json:"maxbitrate,omitempty"`     // 最大的比特率，以Kbps为单位。交易平台可以动态的设置这个值或者为所有发布者统一设置该值
	PlaybackMethod []int              `json:"playbackmethod,omitempty"` // 允许的播放方式， 如果没有指定，表示支持全部，参考5.9
	Delivery       []int              `json:"delivery,omitempty"`       // 支持的传输方式 （例如流式传输，逐步传输），如果没有指定，表示全部支持，参考5.13
	Api            []int              `json:"api,omitempty"`            // 本次展示支持的API框架列表， 参考5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	Ext            *MReqAssetVideoExt `json:"ext,omitempty"`
}

type MReqAssetVideoExt struct {
	Orientation int `json:"orientation,omitempty"`
}

//解析MNativeReq的request string 到 NativeRequest 属性（结构体）
func (mnativeReq *MNativeReq) ParseRequest() (mnativeReqExt *MNativeReqExt, err error) {

	if mnativeReq.Request == "" {
		return nil, errors.New("native request string is empty.")
	}

	//如果请求的native string没有native:{}外层对象，这里补齐native对象。
	if strings.Index(mnativeReq.Request, `"native"`) == -1 {
		mnativeReq.Request = `{"native":` + mnativeReq.Request + `}`
	}
	//mnativeReq.Request = `{"native":{"assets":[{"id":1,"required":0,"data":{"len":90,"type":1}},{"id":2,"required":1,"title":{"len":90}},{"id":3,"required":1,"img":{"type":1,"w":128,"h":128,"mimes":["image/jpg","image/jpeg","image/png"]}}],"plcmtcnt":5,"ver":"1.2","context":1,"contextsubtype":31,"plcmttype":4,"urlsupport":0,"eventtrackers":[{"event":1,"method":1}],"privacy":1}}`
	//mnativeRequest := new(MNativeRequest)
	mnativeReqExt = new(MNativeReqExt)
	//解析到mortb native request
	err = jsonit.Unmarshal([]byte(mnativeReq.Request), mnativeReqExt)
	if err != nil {
		return nil, err
	}
	return
}

//解析MNativeReq的request string 中 MReqAsset 到 map[id]*MReqAsset
func (mnativeReq *MNativeReq) ParseRequest2AssetsMap(mnativeReqExt *MNativeReqExt) (map[int]*MReqAsset, error) {

	//记录当前native的所有assets
	assetsMap := make(map[int]*MReqAsset)
	for _, asset := range mnativeReqExt.Native.ReqAssets {
		if asset.Title != nil {
			//asset.AssetType = IsTitle
			assetsMap[asset.ID] = asset
			continue
		}
		if asset.Image != nil {
			//asset.AssetType = IsImg
			assetsMap[asset.ID] = asset
			continue
		}
		if asset.Data != nil {
			//asset.AssetType = IsData
			assetsMap[asset.ID] = asset
			continue
		}
		if asset.Video != nil {
			//asset.AssetType = IsVideo
			assetsMap[asset.ID] = asset
		}
	}

	return assetsMap, nil
}
