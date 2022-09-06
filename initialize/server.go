package initialize

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"spike-frame/config"
	"time"
)

func RunServer() {
	r := initRouter()
	system := config.Cfg.System
	server := initServer(system.Port, r)
	log.Errorf(server.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

type server interface {
	ListenAndServe() error
}
