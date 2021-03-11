package base

type RequestBase struct {
	RequestId string
	Ip        string            // 上游服务ip
	ABTest    map[string]string // k:v
	Debug     bool
	OpenLog   bool
	TMax      int64 // ms
}

type ResponseBase struct {
	Ip       string // 本机ip
	Status   StatusBase
	Elapsed  int64 // ms
	ABTested map[string]string
	DebugMsg map[string]string
	ErrorMsg string
}

type StatusBase int

const (
	SuccessStatusBase StatusBase = 1
	FailStatusBase    StatusBase = 2
)
