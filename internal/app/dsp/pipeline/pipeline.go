package pipeline

import (
	"errors"
	"sync/atomic"
	"time"

	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
)

type Processer interface {
	Description() string
	Process(ctx *dsp_context.DspContext) (modelStatus dsp_status.DspStatus, modelErr error)
}

func DistributePipeline(ctx *dsp_context.DspContext) *WallTimePipeline {
	return &WallTimePipeline{
		Description: "Dsp Serve",
		Filters: []Processer{
			NewFilteringListPipeline(),
		},
	}
}

type WallTimePipeline struct {
	Description string
	LocalTime   bool
	Filters     []Processer
	Logging     Processer   // bidlog; requestlog; filterlog;
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
		err error
	)
	wtp.TimeElapsed = make([]AtomicInt, len(wtp.Filters))
	for i, filter := range wtp.Filters {
		// process
		mStart := time.Now()
		ctx.ModelStatus, err = filter.Process(ctx)
		elapsed := time.Since(mStart).Milliseconds()
		wtp.TimeElapsed[i].Add(elapsed)

		// metrics
		filter.Description()

		if err != nil {
			break
		}
	}

	return err
}

func (wtp *WallTimePipeline) Log(ctx *dsp_context.DspContext) {
	if wtp == nil || wtp.Logging == nil {
		return
	}

	go wtp.Logging.Process(ctx)
}
