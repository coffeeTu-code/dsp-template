package dsp_config

import "dsp-template/api/base"

type DspConfig struct {
	ConsulConfig  base.BaseConsul
	DebugConfig   DebugConfig
	ServiceConfig ServiceConfig
}

type (
	DebugConfig struct {
		BidForce *BidForce
		LogRate  float64
	}

	BidForce struct {
		BidForceDevice map[string]*BidForceDevice `toml:"BidForceDevice"` // key=user
	}

	BidForceDevice struct {
		DeviceId       []string `toml:"DeviceId"`
		DeviceIdMd5    []string `toml:"DeviceIdMd5"`
		Adx            []string `toml:"Adx"`
		TargetCampaign int64    `toml:"Campaign"`
		TargetTemplate int32    `toml:"Template"`
		TargetPrice    float64  `toml:"Price"`
		User           string
	}
)

func (force *BidForceDevice) IsForce() bool {
	return force != nil
}

type ServiceConfig struct {
	Port string
}
