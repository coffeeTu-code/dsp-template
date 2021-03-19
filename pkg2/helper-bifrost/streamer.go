package bifrost

import (
	"context"
	"time"

	"dsp-template/pkg2/helper-bifrost/container"
)

type Info struct {
	Name         string        `json:"name"`
	TotalNum     int           `json:"total_num"`
	AddNum       int           `json:"add_num"`
	ErrorNum     int           `json:"error_num"`
	LastBaseTime time.Time     `json:"last_base_time"`
	LastIncTime  time.Time     `json:"last_inc_time"`
	BaseTimeUsed time.Duration `json:"base_time_used"`
	IncTimeUsed  time.Duration `json:"inc_time_used"`
}

type Streamer interface {
	SetContainer(container.Container)
	GetContainer() container.Container
	GetScheduleInfo() *ScheduleInfo
	UpdateData(ctx context.Context) error

	GetInfo() *Info
}
