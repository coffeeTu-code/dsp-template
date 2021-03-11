package dsp_context

import (
	mortb "dsp-template/api/adx/madx/go"
	"dsp-template/api/base"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

type DspContext struct {
	Request  *mortb.MOrtbRequest
	Response *mortb.MOrtbResponse

	RequestBase  *base.RequestBase
	ResponseBase *base.ResponseBase

	ModelStatus dsp_status.DspStatus
}
