package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samuel/go-zookeeper/zk"
	"minids/zkcommon"
	"net/http"
	"time"
)

var serverIp string
var listenPort = ":80"

func init() {

}

func main() {
	zkConn := zookeeperRegister()
	defer zkConn.Close()

	// do something
	router := gin.Default()
	// print "172.0.0.1:80:say hello"
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, serverIp+listenPort+":say hello")
	})
	router.Run(listenPort)
}

func zookeeperRegister() *zk.Conn {
	conn, zkconnError := zkcommon.GetConnect([]string{"zookeeper"}, time.Second)
	if zkconnError != nil {
		panic("zookeeper connect failed")
	}

	serverIp, _ = zkcommon.LocalIp()
	fmt.Println(serverIp)
	registerErr := zkcommon.RegistServer(conn, "httpping", serverIp+listenPort)
	if registerErr != nil {
		panic("zookeeper register server failed")
	}

	return conn
}
