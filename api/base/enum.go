package base

const (
	ExchangeOppo  = "oppocn"
	ExchangeIqiyi = "iqiyi"
)

const (
	AdTypeVideoRewarded     = "vre"
	AdTypeVideoInterstitial = "vin"
)

const (
	PlatformAndroid = "android"
	PlatformIos     = "ios"
)

const (
	DeviceTypeMobileTablet = "mobiletablet"
	DeviceTypeComputer     = "computer"
	DeviceTypeTV           = "tv"
	DeviceTypePhone        = "phone"
	DeviceTypeTablet       = "tablet"
)

// Banner.Btype; from open rtb 2.0 final 6.2
const (
	BannerBTypeTEXT       = 1
	BannerBTypeBANNER     = 2
	BannerBTypeJAVASCRIPT = 3 //bType 不为 3 时可召回 only playable
	BannerBTypeIFRAME     = 4
)
