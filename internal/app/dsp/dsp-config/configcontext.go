package dsp_config

import "dsp-template/api/base"

type DspConfig struct {
	ConsulConfig base.ConsulConfig
}

type (
	DebugConfig struct {
		BidForce *BidForce
		LogRate  float64
	}

	BidForce struct {
	}
)

type ServiceConfig struct {
	Port string
}
