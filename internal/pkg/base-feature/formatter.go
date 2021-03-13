package base_feature

import (
	"dsp-template/api/adx/madx"
	"dsp-template/api/dbstruct"
)

type Formatter interface {
	FeatureFormation(adx *madx.MOrtbRequest, feature *dbstruct.Feature)
}
