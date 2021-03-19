package enum

//广告主侧 AdType 体系
//********************* Campaign.mInventoryV2.AdnAdtype *********************
const (
	AdType_Adn_Banner             = 1
	AdType_Adn_Interstitial       = 2
	AdType_Adn_DisPlayNative      = 3
	AdType_Adn_Appwalle           = 4
	AdType_Adn_Rewarded_Appwall   = 5
	AdType_Adn_Video_Interstitial = 6
	AdType_Adn_NativeVideo        = 7
	AdType_Adn_Video_Rewarded     = 8
	AdType_Adn_Video_Instream     = 9
	AdType_Adn_Interactive_Ads    = 10
	AdType_Adn_More_Offer         = 11
	AdType_Adn_Splash             = 12
)

//DSP AdType 体系
//********************* AdType *********************
const (
	AdType_DSP_Banner             = "b"
	AdType_DSP_Interstitial       = "i"
	AdType_DSP_Splash             = "s"
	AdType_DSP_Native             = "n"
	AdType_DSP_Appwall            = "a"
	AdType_DSP_Rewarded_Appwall   = "rap"
	AdType_DSP_Video_Interstitial = "vin"
	AdType_DSP_Video_Native       = "vna"
	AdType_DSP_Video_Rewarded     = "vre"
	AdType_DSP_Video_Instream     = "vis"
	AdType_DSP_Video_Overlay      = "vol" //banner 广告
	AdType_DSP_Video_Infeed       = "vif"
)

//AS AdType 体系
//********************* AdType *********************
const (
	AdType_AS_Appwall_STR        = "appwall"
	AdType_AS_Interstital_STR    = "interstitial"
	AdType_AS_Native_STR         = "native"
	AdType_AS_Reward_Video_STR   = "rewarded_video"
	AdType_AS_FeedsVideo_STR     = "feeds_video"
	AdType_AS_Offerwall_STR      = "offerwall"
	AdType_AS_InterstitalSdk_STR = "interstitial_sdk"
	AdType_AS_OnlineVideo_STR    = "online_video"
	AdType_AS_JsNativeVideo_STR  = "js_native_video"
	AdType_AS_JsBannerVideo_STR  = "js_banner_video"
	AdType_AS_Interstitial_STR   = "interstitial_video"
	AdType_AS_interactive_STR    = "interactive"
	AdType_AS_JmIcon_STR         = "jm_icon"
	AdType_AS_WxNative_STR       = "wx_native"
	AdType_AS_WxAppwall_STR      = "wx_appwall"
	AdType_AS_WxBanner_STR       = "wx_banner"
	AdType_AS_WxRewardImage_STR  = "wx_reward_image"
	AdType_AS_Moreoffer_STR      = "more_offer"
	AdType_AS_Banner_STR         = "sdk_banner"
	AdType_AS_Splash_STR         = "splash"
	AdType_AS_NativeH5_STR       = "native_h5"
)

//AS AdType 体系
//********************* AdType *********************

type AsAdType int32

func (adtype AsAdType) String() string {
	return adtypeAsName[adtype]
}

func (adtype AsAdType) Enum() int32 {
	return int32(adtype)
}

const (
	AdType_As_UNKNOWN            = 0
	AdType_As_BANNER             = 2  // AdType_Adn_Banner
	AdType_As_APPWALL            = 3  //AdType_As_APPWALL
	AdType_As_INTERSTITIAL       = 29 // AdType_Adn_Interstitial
	AdType_As_NATIVE             = 42 // AdType_Adn_DisPlayNative
	AdType_As_REWARDED_VIDEO     = 94 // AdType_Adn_Video_Rewarded
	AdType_As_FEEDS_VIDEO        = 95
	AdType_As_Instream_VIDEO     = 99 // AdType_Adn_Video_Instream
	AdType_As_RECTANGULAR        = 215
	AdType_As_REFFERLINK         = 216
	AdType_As_OFFERWALL          = 278
	AdType_As_INTERSTITIAL_SDK   = 279 // AdType_Adn_Video_Interstitial
	AdType_As_MOBPOWER           = 280
	AdType_As_ONLINE_VIDEO       = 284
	AdType_As_JS_NATIVE_VIDEO    = 285 // AdType_Adn_NativeVideo
	AdType_As_JS_BANNER_VIDEO    = 286
	AdType_As_INTERSTITIAL_VIDEO = 287 // AdType_Adn_Video_Interstitial
	AdType_As_INTERACTIVE        = 288
	AdType_As_JM_ICON            = 289
	AdType_As_WX_NATIVE          = 291 // AdType_Adn_DisPlayNative
	AdType_As_WX_APPWALL         = 292 // AdType_Adn_Appwalle
	AdType_As_WX_BANNER          = 293 // AdType_Adn_Banner
	AdType_As_WX_REWARD_IMAGE    = 294 //
	AdType_As_MORE_OFFER         = 295 // AdType_Adn_More_Offer
	AdType_As_SDK_BANNER         = 296 // AdType_Adn_Banner
	AdType_As_Splash             = 297 // AdType_Adn_Splash
	AdType_As_NATIVE_H5          = 298 // AdType_Adn_Banner
)

var adtypeAsName = map[AsAdType]string{
	AdType_As_SDK_BANNER:         "sdk_banner",
	AdType_As_NATIVE_H5:          "native_h5",
	AdType_As_Splash:             "splash",
	AdType_As_NATIVE:             "native",
	AdType_As_JS_NATIVE_VIDEO:    "js_native_video",
	AdType_As_REWARDED_VIDEO:     "rewarded_video",
	AdType_As_INTERSTITIAL_VIDEO: "interstitial_video",
}

var adtypeDspToAsValue = map[string]int32{
	AdType_DSP_Banner:             AdType_As_SDK_BANNER,
	AdType_DSP_Video_Overlay:      AdType_As_SDK_BANNER,
	AdType_DSP_Interstitial:       AdType_As_INTERSTITIAL,
	AdType_DSP_Splash:             AdType_As_Splash,
	AdType_DSP_Native:             AdType_As_NATIVE,
	AdType_DSP_Video_Native:       AdType_As_JS_NATIVE_VIDEO,
	AdType_DSP_Video_Rewarded:     AdType_As_REWARDED_VIDEO,
	AdType_DSP_Video_Interstitial: AdType_As_INTERSTITIAL_VIDEO,
	AdType_DSP_Video_Instream:     AdType_As_Instream_VIDEO,
	AdType_DSP_Video_Infeed:       AdType_As_FEEDS_VIDEO,
	AdType_AS_Appwall_STR:         AdType_As_APPWALL,
	AdType_AS_Interstital_STR:     AdType_As_INTERSTITIAL,
	AdType_AS_Native_STR:          AdType_As_NATIVE,
	AdType_AS_Reward_Video_STR:    AdType_As_REWARDED_VIDEO,
	AdType_AS_FeedsVideo_STR:      AdType_As_FEEDS_VIDEO,
	AdType_AS_Offerwall_STR:       AdType_As_OFFERWALL,
	AdType_AS_InterstitalSdk_STR:  AdType_As_INTERSTITIAL_SDK,
	AdType_AS_OnlineVideo_STR:     AdType_As_ONLINE_VIDEO,
	AdType_AS_JsNativeVideo_STR:   AdType_As_JS_NATIVE_VIDEO,
	AdType_AS_JsBannerVideo_STR:   AdType_As_JS_BANNER_VIDEO,
	AdType_AS_Interstitial_STR:    AdType_As_INTERSTITIAL_VIDEO,
	AdType_AS_interactive_STR:     AdType_As_INTERACTIVE,
	AdType_AS_JmIcon_STR:          AdType_As_JM_ICON,
	AdType_AS_WxNative_STR:        AdType_As_WX_NATIVE,
	AdType_AS_WxAppwall_STR:       AdType_As_WX_APPWALL,
	AdType_AS_WxBanner_STR:        AdType_As_WX_BANNER,
	AdType_AS_WxRewardImage_STR:   AdType_As_WX_REWARD_IMAGE,
	AdType_AS_Moreoffer_STR:       AdType_As_MORE_OFFER,
	AdType_AS_Banner_STR:          AdType_As_SDK_BANNER,
	AdType_AS_Splash_STR:          AdType_As_Splash,
	AdType_AS_NativeH5_STR:        AdType_As_NATIVE_H5,
}

var adtypeAsToAdnValue = map[int32]int32{
	AdType_As_APPWALL:            AdType_Adn_Appwalle,
	AdType_As_SDK_BANNER:         AdType_Adn_Banner,
	AdType_As_INTERSTITIAL:       AdType_Adn_Interstitial,
	AdType_As_Splash:             AdType_Adn_Splash,
	AdType_As_NATIVE:             AdType_Adn_DisPlayNative,
	AdType_As_JS_NATIVE_VIDEO:    AdType_Adn_NativeVideo,
	AdType_As_REWARDED_VIDEO:     AdType_Adn_Video_Rewarded,
	AdType_As_INTERSTITIAL_VIDEO: AdType_Adn_Video_Interstitial,
	AdType_As_Instream_VIDEO:     AdType_Adn_Video_Instream,
}

func GetAdTypeDspToAs(dspAdtype string) (asAdtype int32) {
	return adtypeDspToAsValue[dspAdtype]
}

func GetAdnAdtype(asAdType int32) (adnAdType int32) {
	return adtypeAsToAdnValue[asAdType]
}

func Transfer3SAdType(adtype int32) []int32 {
	switch adtype {
	case AdType_Adn_Banner:
		return []int32{AdType_As_BANNER, AdType_As_WX_BANNER, AdType_As_SDK_BANNER}
	case AdType_Adn_Interstitial:
		return []int32{AdType_As_INTERSTITIAL, AdType_As_INTERSTITIAL_SDK}
	case AdType_Adn_DisPlayNative:
		return []int32{AdType_As_NATIVE, AdType_As_WX_NATIVE}
	case AdType_Adn_Appwalle:
		return []int32{AdType_As_APPWALL, AdType_As_WX_APPWALL}
	case AdType_Adn_Video_Interstitial:
		return []int32{AdType_As_INTERSTITIAL_VIDEO}
	case AdType_Adn_NativeVideo:
		return []int32{AdType_As_NATIVE, AdType_As_JS_NATIVE_VIDEO}
	case AdType_Adn_Video_Rewarded:
		return []int32{AdType_As_REWARDED_VIDEO}
	case AdType_Adn_Video_Instream:
		return []int32{AdType_As_Instream_VIDEO}
	case AdType_Adn_More_Offer:
		return []int32{AdType_As_MORE_OFFER}
	case AdType_Adn_Splash:
		return []int32{AdType_As_Splash}
	}
	return []int32{}
}
