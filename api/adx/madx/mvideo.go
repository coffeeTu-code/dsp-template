package madx

import (
	"errors"
)

// Validation errors
var (
	ErrInvalidVideoNoMimes       = errors.New("openrtb: video has no mimes")
	ErrInvalidVideoNoLinearity   = errors.New("openrtb: video linearity missing")
	ErrInvalidVideoNoMinDuration = errors.New("openrtb: video min-duration missing")
	ErrInvalidVideoNoMaxDuration = errors.New("openrtb: video max-duration missing")
	ErrInvalidVideoNoProtocols   = errors.New("openrtb: video protocols missing")
)

// OpenRTB 2.5: The following table lists the various types of video placements
// derived largely from the IAB Digital Video Guidelines.
type VideoPlacementType int32

const (
	// The video placement is not defined.
	// Default value.
	UNDEFINED_VIDEO VideoPlacementType = 0
	// Played before, during or after the streaming video content
	// that the consumer has requested.
	// E.G.: Pre-roll, Mid-roll, Post-roll.
	IN_STREAM VideoPlacementType = 1
	// Exists within a web banner that leverages the banner space
	// to deliver a video experience as opposed to another static
	// or rich media format.
	// The format relies on the existence of display ad inventory
	// on the page for its delivery.
	IN_BANNER VideoPlacementType = 2
	// Loads and plays dynamically between paragraphs of editorial content;
	// existing as a standalone branded message.
	IN_ARTICLE VideoPlacementType = 3
	// In-Feed - Found in content, social, or product feeds.
	IN_FEED VideoPlacementType = 4
	// Interstitial/Slider/Floating.
	// Covers the entire or a portion of screen area,
	// but is always on screen while displayed
	// (i.e. cannot be scrolled out of view).
	// Note that a full-screen interstitial (e.g., in mobile)
	// can be distinguished from a floating/slider unit by the imp.instl field.
	INTERSTITIAL VideoPlacementType = 5
)

//PlacementType 枚举
const (
	RewardedVideo     = "rewarded video"
	InterstitialVideo = "interstitial video"
	InstreamVideo     = "instream video"
	NativeVideo       = "native video"
)

// 这个对象表示一个流式视频展示。 许多属性对于最小可用功能不是必须的，但是为了在需要的时候提供更好的控制能力会被使用。
// OpenRTB中的视频通常都是与标准一致的，复合广告的概念也是支持的， 可以包含用于定义复合广告的一组Banner对象。
// Video作为Imp的子对象出现表示它是一个具有视频类型的展示对象。 同样的展示也可以是一个Banner或者Native广告，
// 只要包含Banner对象或者Native对象。然而， 任何为展示给定的竞价请求必须符合提供类型中的一个。
type MVideo struct {
	Mimes          []string   `json:"mimes,omitempty"`          // 支持的内容mime-type， 常用的类型包括用于windows媒体的video/x-ms-wmv以及用于Flash视频的video/x-flv.
	MinDuration    int        `json:"minduration,omitempty"`    // 最小的视频长度， 以秒为单位
	MaxDuration    int        `json:"maxduration,omitempty"`    // 最大的视频长度， 以秒为单位
	Protocols      []int      `json:"protocols,omitempty"`      // 支持的视频竞价响应协议数组。参考表5.8.至少一个支持的协议必须在protocol或者protocols属性中被指定
	W              int        `json:"w,omitempty"`              // 视频播放器的宽度，像素为单位
	H              int        `json:"h,omitempty"`              // 视频播放器的高度，像素为单位
	StartDelay     int        `json:"startdelay,omitempty"`     // 视频前，中及之后的广告位中视频广告的启动延时，以秒为单位, 参考5.10
	Linearity      int        `json:"linearity,omitempty"`      // 展示是否必须是线性的， 如果没有指定，则标识都是被允许的，参考5.11
	Skip           int        `json:"skip,omitempty"`           // Indicates if the player will allow the video to be skipped, where 0 = no, 1 = yes.
	SkipMin        int        `json:"skipmin,omitempty"`        // Videos of total duration greater than this number of seconds can be skippable
	SkipAfter      int        `json:"skipafter,omitempty"`      // Number of seconds a video must play before skipping is enabled
	Sequence       int        `json:"sequence,omitempty"`       // 如果在同一个竞价请求中提供了多个展示， 则需要考虑多个物料传输的顺序 Default: 1
	BAttr          []int      `json:"battr,omitempty"`          // 限制的物料属性，参考5.3
	MaxExtended    int        `json:"maxextended,omitempty"`    // 最大的视频广告延长时间长度（如果支持延长）。如果为空或者0，表示不允许延长， 如果为-1，表示允许延时，且没有时间限制， 如果为大于0的数字， 则表示可以延长的时间长度比maxduration大的值
	MinBitrate     int        `json:"minbitrate,omitempty"`     // 最小的比特率，以Kbps为单位。交易平台可以动态的设置这个值或者为所有发布者统一设置该值
	MaxBitrate     int        `json:"maxbitrate,omitempty"`     // 最大的比特率，以Kbps为单位。交易平台可以动态的设置这个值或者为所有发布者统一设置该值
	BoxingAllowed  int        `json:"boxingallowed,omitempty"`  // 是否允许将4：3的内容展示在16：9的窗口， 0表示不允许，1表示允许
	PlaybackMethod []int      `json:"playbackmethod,omitempty"` // 允许的播放方式， 如果没有指定，表示支持全部，参考5.9
	PlaybackEnd    int        `json:"playbackend,omitempty"`    // The event that causes playback to end. Refer to List 5.11.
	Delivery       []int      `json:"delivery,omitempty"`       // 支持的传输方式 （例如流式传输，逐步传输），如果没有指定，表示全部支持，参考5.13
	Pos            int        `json:"pos,omitempty"`            // 广告在屏幕上的位置，参考5.4
	CompanionAd    []*MBanner `json:"companionad,omitempty"`    // 如果支持复合广告，表示一组Banner对象
	Api            []int      `json:"api,omitempty"`            // 本次展示支持的API框架列表， 参考5.6. 如果一个API没有被显式在列表中指明，则表示不支持
	CompanionType  []int      `json:"companiontype,omitempty"`  // 支持的VAST companion 广告类型， 参考5.12。 如果在companionad中填充了Banner对象则推荐使用
	Placement      int        `json:"placement,omitempty"`      // Video placement type 参考5.9
	Ext            *MVideoExt `json:"ext,omitempty"`            // 特定交易的OpenRTB协议的扩展信息占位符
}

type MVideoExt struct {
	PlacementType string `json:"placementtype,omitempty"` // 如果是激励广告，本属性值为字符串rewarded，默认为空.
	Orientation   int    `json:"orientation,omitempty"`   // empty or 0 is unknow, 1 PORTRAIT, 2 LANDSCAPE
}
