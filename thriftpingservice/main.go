package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"minids/common"
	"minids/services/tpingservice/gen-go/tpingservice"
)

var serverIp string
var listenPort = ":8181"
var serverIPAndPort string

func init() {
	serverIp, _ = common.LocalIp()
	serverIPAndPort = serverIp + listenPort
}

type tpingserviceImp struct {
}

func (self tpingserviceImp) Ping(ctx context.Context, action string, content string) (r string, err error) {
	return "", nil
}

func main() {

	zkConn := common.RegisterServiceName("thriftping", serverIPAndPort)
	defer zkConn.Close()

	handler := tpingserviceImp{}
	processor := tpingservice.NewTpingServiceProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(serverIPAndPort)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", serverIPAndPort)
	server.Serve()

}
