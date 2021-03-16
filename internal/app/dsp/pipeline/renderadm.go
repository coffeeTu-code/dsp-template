package pipeline

import (
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

func NewRenderAdmPipeline() *RenderAdmPipeline {
	return &RenderAdmPipeline{}
}

type RenderAdmPipeline struct {
}

func (pipe *RenderAdmPipeline) Description() string {
	return "RenderAdm"
}

func (pipe *RenderAdmPipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {
	return
}
