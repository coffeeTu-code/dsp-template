package service

import (
	"net/http"
	"time"

	"github.com/astaxie/beego"

	"dsp-template/internal/app/dsp/service/controllers"
)

func NewHttpServer() *HTTPServer {
	return &HTTPServer{
		httpAddr: "httpAddr",
		httpServer: &http.Server{
			Addr:           "httpAddr",
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Minute,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

type HTTPServer struct {
	httpAddr       string
	httpServer     *http.Server
	PolarisHandler http.Handler
}

func (hs *HTTPServer) initRouter() {
	beego.Router("/madx_request", controllers.NewAdxController(), "post:BidServer")
}
