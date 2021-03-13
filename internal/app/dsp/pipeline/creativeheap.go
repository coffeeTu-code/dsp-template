package pipeline

import (
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

func NewCreativeHeapPipeline() *CreativeHeapPipeline {
	return &CreativeHeapPipeline{}
}

type CreativeHeapPipeline struct {
}

func (pipe *CreativeHeapPipeline) Description() string {
	return "CreativeHeap"
}

func (pipe *CreativeHeapPipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {
	return
}
