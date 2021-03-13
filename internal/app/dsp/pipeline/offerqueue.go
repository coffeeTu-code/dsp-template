package pipeline

import (
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

func NewOfferQueuePipeline() *OfferQueuePipeline {
	return &OfferQueuePipeline{}
}

type OfferQueuePipeline struct {
}

func (pipe *OfferQueuePipeline) Description() string {
	return "OfferQueue"
}

func (pipe *OfferQueuePipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {
	return
}
