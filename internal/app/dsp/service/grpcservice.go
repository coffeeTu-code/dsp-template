package service

import (
	"time"

	mortb "dsp-template/api/adx/madx/go"
	"dsp-template/api/base"
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	"dsp-template/internal/app/dsp/pipeline"
)

func NewDspService(request *mortb.MOrtbRequest) (response *mortb.MOrtbResponse, err error) {
	ctx := &dsp_context.DspContext{
		Request:      request,
		Response:     new(mortb.MOrtbResponse),
		RequestBase:  new(base.RequestBase),
		ResponseBase: new(base.ResponseBase),
	}

	pStart := time.Now()
	p := pipeline.DistributePipeline(ctx)
	pErr := p.Process(ctx)
	if pErr != nil {
		ctx.ResponseBase.ErrorMsg = pErr.Error()
		ctx.ResponseBase.Status = base.FailStatusBase
	} else {
		ctx.ResponseBase.Status = base.SuccessStatusBase
	}
	ctx.ResponseBase.Elapsed = time.Since(pStart).Milliseconds()

	return
}
