package dsp_metrics

import "dsp-template/api/base"

func NewDspMetricsContext() *DspMetricsContext {
	return &DspMetricsContext{
		Base: dspMetricsBase,
	}
}

type DspMetricsContext struct {
	Base        *base.BaseMetrics
	Exchange    string
	AdType      string
	Platform    string
	ModelStatus string
	OfferID     string
	Price       string
}

var dspMetricsBase = &base.BaseMetrics{
	Namespace: "Dsp",
	Subsystem: "Dsp",
	Region:    region(),
	Ip:        ip(),
}

func region() string {
	return ""
}

func ip() string {
	return ""
}
