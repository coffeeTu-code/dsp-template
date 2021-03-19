package base_feature

import (
	"dsp-template/api/adx/madx"
	"dsp-template/api/base"
)

type Formatter interface {
	FeatureFormation(adx *madx.MOrtbRequest, feature *base.Feature)
}
