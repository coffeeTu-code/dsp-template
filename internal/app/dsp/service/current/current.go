package current

import "sync/atomic"

func NewCurrentLimiter() *CurrentLimit {
	return &CurrentLimit{
		max: max(),
	}
}

type CurrentLimit struct {
	current int64
	max     int64
}

func (cl *CurrentLimit) Get() bool {
	n := atomic.AddInt64(&cl.current, 1)
	if n > cl.max {
		cl.Put()
		return false
	}
	return true
}

func (cl *CurrentLimit) Put() {
	if cl.current == 0 {
		return
	}
	atomic.AddInt64(&cl.current, -1)
}

func max() int64 {
	return 100
}
