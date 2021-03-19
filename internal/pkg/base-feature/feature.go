package base_feature

import (
	"dsp-template/api/adx/madx"
	"dsp-template/api/base"
)

func FeatureFormation(adx *madx.MOrtbRequest) *base.Feature {
	feature := base.NewFeature()

	if adx.Ext == nil || len(adx.Ext.Exchange) == 0 {
		return feature
	}
	feature.Exchange = adx.Ext.Exchange

	setDeviceFeature(adx, feature)
	setUserFeature(adx, feature)
	setVideoFeature(adx, feature)
	setImpFeature(adx, feature)

	if len(formatterList) > 0 {
		formatter, ok := formatterList[feature.Exchange]
		if !ok {
			return feature
		}
		formatter.FeatureFormation(adx, feature)
	}

	return feature
}

func setDeviceFeature(adx *madx.MOrtbRequest, feature *base.Feature) {
	if adx.Device == nil {
		return
	}
	feature.Device.IP = adx.Device.IP
	feature.Device.IPv6 = adx.Device.IPv6
	feature.Device.ScreenW = int32(adx.Device.W)
	feature.Device.ScreenH = int32(adx.Device.H)

	if adx.Device.Ext == nil {
		return
	}
	feature.Device.Ifv = adx.Device.Ext.Ifv
	feature.Device.TotalDisk = adx.Device.Ext.TotalDisk
}

func setUserFeature(adx *madx.MOrtbRequest, feature *base.Feature) {
	if adx.User == nil {
		return
	}
	feature.User.ID = adx.User.ID

	if adx.User.Ext == nil {
		return
	}
	feature.User.LastBundle = adx.User.Ext.LastBundle
	feature.User.LastADomain = adx.User.Ext.LastADomain
	feature.User.ImpDepth = adx.User.Ext.ImpDepth
	feature.User.SessionDuration = adx.User.Ext.SessionDuration
}

func setVideoFeature(adx *madx.MOrtbRequest, feature *base.Feature) {
	if len(adx.Imp) == 0 || adx.Imp[0].Video == nil {
		return
	}
	video := adx.Imp[0].Video
	feature.Video.PlaybackEnd = int32(video.PlaybackEnd)
	feature.Video.Skip = int32(video.Skip)
	for _, val := range video.PlaybackMethod {
		feature.Video.PlaybackMethod = append(feature.Video.PlaybackMethod, int32(val))
	}
}

func setImpFeature(adx *madx.MOrtbRequest, feature *base.Feature) {
	if len(adx.Imp) == 0 || adx.Imp[0].Ext == nil {
		return
	}
	if skadn := adx.Imp[0].Ext.Skadn; skadn != nil {
		feature.Imp.Skadn, _ = json.MarshalToString(skadn)
	}
}
