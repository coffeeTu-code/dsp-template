package madx

import jsoniter "github.com/json-iterator/go"

var jsonit = jsoniter.ConfigCompatibleWithStandardLibrary

func NewDspRequest() *MOrtbRequest {
	return &MOrtbRequest{}
}

func NewDspResponse() *MOrtbResponse {
	return &MOrtbResponse{}
}
