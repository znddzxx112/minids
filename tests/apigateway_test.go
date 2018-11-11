package tests

import (
	"fmt"
	"minids/common"
	"net"
	"testing"
)

/**
func TestHttppingService(t *testing.T) {
	resp, _ := http.Get("http://apigateway/ping")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("TestHttppingService:", string(body))

}
**/

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
