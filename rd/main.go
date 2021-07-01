package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"rd/models"
	"syscall"

	"github.com/gin-gonic/gin"

	"rd/pkg/logging"
	"rd/pkg/setting"
	"rd/pkg/util"
	"rd/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	// gredis.Setup()
	util.Setup()
}


func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	log.Printf("[info] start http server listening %s", endPoint)


	//server := &http.Server{
	//	Addr:           endPoint,
	//	Handler:        routersInit,
	//	ReadTimeout:    readTimeout,
	//	WriteTimeout:   writeTimeout,
	//	MaxHeaderBytes: maxHeaderBytes,
	//}
	//server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
