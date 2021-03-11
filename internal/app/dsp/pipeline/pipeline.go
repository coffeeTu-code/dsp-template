package pipeline

import (
	"errors"
	"sync/atomic"

	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

type Processer interface {
	Description() string
	Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelElapsed int64, modelErr error)
}

func DistributePipeline(ctx *dsp_context.DspContext) *WallTimePipeline {
	return &WallTimePipeline{
		Description: "Dsp Serve",
		Filters: []Processer{

		},
	}
}

type WallTimePipeline struct {
	Description string
	LocalTime   bool
	Filters     []Processer
	Logging     Processer
	TimeElapsed []AtomicInt // ms
}

type AtomicInt int64

func (a *AtomicInt) Add(i int64) {
	atomic.AddInt64((*int64)(a), i)
}

func (a *AtomicInt) Val() int64 {
	return *(*int64)(a)
}

func (wtp *WallTimePipeline) Process(ctx *dsp_context.DspContext) error {
	if wtp == nil || len(wtp.Filters) == 0 {
		return errors.New("wtp is nil")
	}

	var (
		elapsed int64
		err     error
	)
	wtp.TimeElapsed = make([]AtomicInt, len(wtp.Filters))
	for i, filter := range wtp.Filters {
		ctx.ModelStatus, elapsed, err = filter.Process(ctx)

		wtp.TimeElapsed[i].Add(elapsed)
		if err != nil {
			break
		}
	}

	return err
}
