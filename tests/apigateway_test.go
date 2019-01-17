package tests

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"minids/common"
	"minids/services/tpingservice/gen-go/tpingservice"
	"testing"
	"time"
)

/**
func TestHttppingService(t *testing.T) {
	resp, _ := http.Get("http://apigateway/ping")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("TestHttppingService:", string(body))

}
**/

/**
func TestTcppingService(t *testing.T) {
	serverIp, serverErr := common.GetServerList("tcpping")
	if serverErr != nil {
		panic(serverErr.Error())
	}
	conn, dialErr := net.Dial("tcp", serverIp)
	if dialErr != nil {
		panic(dialErr.Error())
	}
	defer conn.Close()

	conn.Write(common.EncodeMsgEnd([]byte("tcpping")))
	line := common.DecodeMsgEnd(conn)
	fmt.Println("TestTcppingService:", string(line))
}
**/
/**
func TestTcpMsgHandleFunc(t *testing.T) {
	serverIp, serverErr := common.GetServerList("tcpmsgping")
	if serverErr != nil {
		panic(serverErr.Error())
	}
	conn, dialErr := net.Dial("tcp", serverIp)
	if dialErr != nil {
		panic(dialErr.Error())
	}
	defer conn.Close()

	conn.Write(common.EncodeFrame([]byte("tcpmsgping")))
	line := common.DecodeFrame(conn)
	fmt.Println("TestTcpMsgHandleFunc:", string(line))
}
**/
func TestThriftHandleFunc(t *testing.T) {
	serverIp, serverErr := common.GetServerList("thriftping")
	if serverErr != nil {
		panic(serverErr.Error())
	}
	tSocket, err := thrift.NewTSocket(serverIp)
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, _ := transportFactory.GetTransport(tSocket)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := tpingservice.NewTpingServiceClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", serverIp)
	}
	defer transport.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	d, err := client.Ping(ctx, "print", "hello thrift")
	fmt.Println(d)
}
