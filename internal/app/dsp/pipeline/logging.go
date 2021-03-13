package pipeline

import (
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

func NewLoggingPipeline() *LoggingPipeline {
	return &LoggingPipeline{}
}

type LoggingPipeline struct {
}

func (pipe *LoggingPipeline) Description() string {
	return "Logging"
}

func (pipe *LoggingPipeline) Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error) {
	return
}
