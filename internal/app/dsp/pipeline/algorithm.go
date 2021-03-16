package pipeline

import (
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

func NewAlgorithmPipeline() *AlgorithmPipeline {
	return &AlgorithmPipeline{}
}

type AlgorithmPipeline struct {
}

func (pipe *AlgorithmPipeline) Description() string {
	return "Algorithm"
}

func (pipe *AlgorithmPipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {
	return
}
