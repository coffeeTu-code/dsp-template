package microservice_helper

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

func ExampleCallDependentService_WithoutFallback() (int, error) {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout: 1000,
	})
	ret, err := CallDependentService("my_command", func() (interface{}, error) {
		//Put your logic here
		time.Sleep(time.Millisecond * 1200)
		return 1, nil
	}, nil /*no fallback method*/)
	if err != nil {
		return -1, err
	}
	return ret.(int), err
	//Output : -1 ,hystrix.ErrTimeout
}

func ExampleCallDependentService_WithFallback() (int, error) {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout: 1000,
	})
	ret, err := CallDependentService("my_command", func() (interface{}, error) {
		//Put your logic here
		time.Sleep(time.Millisecond * 1200)
		return 1, nil
	}, func(e error) (interface{}, error) { return 2, nil })
	return ret.(int), err
	//Output : 2,nil
}
