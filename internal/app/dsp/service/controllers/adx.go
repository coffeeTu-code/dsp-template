package controllers

import (
	"time"

	"github.com/astaxie/beego"

	"dsp-template/api/adx/madx"
	"dsp-template/api/base"
	dsp_context "dsp-template/internal/app/dsp/dsp-context"
	dsp_metrics "dsp-template/internal/app/dsp/dsp-metrics"
	dsp_status "dsp-template/internal/app/dsp/dsp-status"
	"dsp-template/internal/app/dsp/pipeline"
	"dsp-template/internal/app/dsp/service/current"
	base_feature "dsp-template/internal/pkg/base-feature"
)

func NewAdxController() *AdxController {
	return &AdxController{}
}

type AdxController struct {
	beego.Controller
}

func (c *AdxController) BidServer() {
	// current limit
	cl := current.NewCurrentLimiter()
	if !cl.Get() {
		circuit()
		return
	}
	defer cl.Put()

	// run
	c.run()
}

func circuit() {
	// status

	// metrics

	// log

}

func (c *AdxController) run() {
	var err error
	defer func() {
		if err == nil {
			return
		}
		// run error
	}()

	// decode
	var request = madx.NewDspRequest()
	var response = madx.NewDspResponse()
	err = request.FromJson(c.Ctx.Input.RequestBody)
	if err != nil {
		return
	}

	// service
	response, err = runDspService(request)
	if err != nil {
		return
	}

	// encode
	var body string
	body, err = response.String()
	if err != nil {
		return
	}
	err = c.Ctx.Output.Body([]byte(body))
	if err != nil {
		return
	}

}

func runDspService(request *madx.MOrtbRequest) (response *madx.MOrtbResponse, err error) {
	ctx := &dsp_context.DspContext{
		Request:      request,
		Response:     madx.NewDspResponse(),
		RequestBase:  base.NewBaseRequest(),
		ResponseBase: base.NewBaseResponse(),
		MetricsCtx:   dsp_metrics.NewDspMetricsContext(),
		Feature:      base_feature.FeatureFormation(request),
	}

	pStart := time.Now()
	p := pipeline.DistributePipeline(ctx)
	pErr := p.Process(ctx)
	if pErr != nil {
		ctx.ResponseBase.Errormsg = pErr.Error()
		ctx.ResponseBase.Status = base.StatusBase_Fail
	} else {
		ctx.ResponseBase.Status = base.StatusBase_Success
		ctx.ModelStatus = dsp_status.DspStatusOk
	}
	ctx.ResponseBase.Ip = ctx.MetricsCtx.Base.Ip
	ctx.ResponseBase.Elapsed = time.Since(pStart).Milliseconds()
	p.Log(ctx)

	return ctx.Response, pErr
}
