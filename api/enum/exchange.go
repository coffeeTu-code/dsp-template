package enum

const (
	Exchange_InnerActive = "inneractive"
	Exchange_Mopub       = "mopub"
	Exchange_DoubleClick = "doubleclick"
	Exchange_AppLovin    = "applovin"
	Exchange_Unity       = "unity"
	Exchange_SmartAds    = "smartyads"
	Exchange_Mobfox      = "mobfox"
	Exchange_Nexage      = "nexage"
	Exchange_Epom        = "epom"
	Exchange_Zplay       = "zplay"
	Exchange_Pubnative   = "pubnative"
	Exchange_Avocarrot   = "avocarrot"
	Exchange_Appodeal    = "appodeal"
	Exchange_Smaato      = "smaato"
	Exchange_Adcolony    = "adcolony"
	Exchange_Jingdong    = "jingdong"
	Exchange_ChartBoost  = "chartboost"
	Exchange_Tencent     = "tencent"
	Exchange_Samsung     = "samsung"
	Exchange_Mintegral   = "mintegral"
	Exchange_LoopMe      = "loopme"
	Exchange_OppoCN      = "oppocn"
	Exchange_Bytedance   = "bytedance"
	Exchange_Iqiyi       = "iqiyi"
)

var exchange_to_request_type_value = map[string]int32{
	Exchange_InnerActive: 100,
	Exchange_Mopub:       101,
	Exchange_DoubleClick: 102,
	Exchange_AppLovin:    103,
	Exchange_Unity:       104,
	Exchange_SmartAds:    105,
	Exchange_Mobfox:      106,
	Exchange_Nexage:      107,
	Exchange_Epom:        108,
	Exchange_Zplay:       109,
	Exchange_Pubnative:   110,
	Exchange_Avocarrot:   111,
	Exchange_Appodeal:    112,
	Exchange_Smaato:      113,
	Exchange_Adcolony:    114,
	Exchange_Jingdong:    115,
	Exchange_ChartBoost:  116,
	Exchange_Tencent:     117,
	Exchange_Samsung:     118,
	Exchange_Mintegral:   119,
	Exchange_LoopMe:      120,
	Exchange_OppoCN:      121,
	Exchange_Bytedance:   122,
	Exchange_Iqiyi:       123,
}

func GetPublisherTypeFromExchange(exchange string) (requestType int32) {
	return exchange_to_request_type_value[exchange]
}
