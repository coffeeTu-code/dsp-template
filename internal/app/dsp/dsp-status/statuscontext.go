package dsp_status

func NewDspStatusContext() *DspStatusContext {
	return &DspStatusContext{
		InRequest: inRequest,
		InBid:     inBid,
	}
}

type DspStatusContext struct {
	InRequest map[DspStatus]bool
	InBid     map[DspStatus]bool
}

var inRequest = map[DspStatus]bool{

}

var inBid = map[DspStatus]bool{
	DspStatusOk: true,
}
