package base_feature

import (
	"dsp-template/api/adx/madx"
	"dsp-template/api/base"
	"dsp-template/api/enum"
)

func init() {
	formatterList[enum.ExchangeIqiyi] = &IqiyiFormatter{}
}

type IqiyiFormatter struct {
}

func (f *IqiyiFormatter) FeatureFormation(adx *madx.MOrtbRequest, feature *base.Feature) {
	f.setImpFeature(adx, feature)
}

// SetImpExtFeature 新增imp.ext.SourceID/imp.ext.SourceURL 两个特性，标识爱奇艺站点信息.
func (f *IqiyiFormatter) setImpFeature(adx *madx.MOrtbRequest, feature *base.Feature) {
	if len(adx.Imp) == 0 || adx.Imp[0].Ext == nil {
		return
	}
	impext := adx.Imp[0].Ext
	feature.Imp.SourceID = impext.SourceID
	if impext.SourceURL != "" {
		feature.Imp.SourceURL = impext.SourceURL
	}
}
