package dsp_context

import (
	"dsp-template/api/adx/madx"
	"dsp-template/api/base"
	"dsp-template/api/dbstruct"
	dsp_metrics "dsp-template/internal/app/dsp/dsp-metrics"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

type DspContext struct {
	Request  *madx.MOrtbRequest
	Response *madx.MOrtbResponse

	RequestBase  *base.BaseRequest
	ResponseBase *base.BaseResponse

	ModelStatus dsp_status.DspStatus

	MetricsCtx *dsp_metrics.DspMetricsContext

	Feature *dbstruct.Feature
}
