package main

import (
	"github.com/gin-gonic/gin"
	"minids/core/common"
	"net/http"
)

func init() {

}

var serverIp string
var listenPort = ":80"
var serverIPAndPort string

func init() {
	serverIp, _ = common.LocalIp()
	serverIPAndPort = serverIp + listenPort
}

func main() {
	zkConn := common.RegisterServiceName("httpping", serverIp+listenPort)
	defer zkConn.Close()

	// do something
	router := gin.Default()
	// print "172.0.0.1:80:say hello"
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, serverIp+listenPort+":say hello")
	})
	router.Run(listenPort)
}
