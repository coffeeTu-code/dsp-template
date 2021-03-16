package dsp_status

type DspStatus string

const (
	DspStatusDefault                    = "default"
	DspStatusOk                         = "ok" // Bid
	DspStatusMediaHardFilterSize        = "size_unuseful"
	DspStatusMediaHardFilterTraffic     = "unuseful_traffic"
	DspStatusMediaHardFilterCountry     = "country_unrecognized"
	DspStatusMediaHardFilterUA          = "devid_ua_empty"
	DspStatusMediaHardFilterDeviceType  = "not_phone_or_pad"
	DspStatusMediaHardFilterBannerBType = "block_img"
	DspStatusMediaHardFilterPlatform    = "os_unuseful"
	DspStatusMediaHardFilterDeviceID    = "device_in_blacklist"
	DspStatusMediaHardFilterBidCache    = "bid_cache_filter"
	DspStatusMediaSoftFilterPackageName = "no_valid_pkgname"
)
