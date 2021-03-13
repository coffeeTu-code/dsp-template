package base

import jsoniter "github.com/json-iterator/go"

var jsonit = jsoniter.ConfigCompatibleWithStandardLibrary

func NewBaseRequest() *BaseRequest {
	return &BaseRequest{
		Abtest: map[string]string{},
	}
}

func NewBaseResponse() *BaseResponse {
	return &BaseResponse{
		Abtested: map[string]string{},
		Debugmsg: map[string]string{},
	}
}
