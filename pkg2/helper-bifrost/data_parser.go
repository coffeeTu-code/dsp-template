package bifrost

import "dsp-template/pkg2/helper-bifrost/container"

type ParserResult struct {
	DataMode container.DataMode
	Key      container.MapKey
	Value    interface{}
	Err      error
}

type DataParser interface {
	Parse([]byte, interface{}) []ParserResult
}
