package tests

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"minids/common"
	"net"
	"net/http"
	"testing"
)

func TestHttppingService(t *testing.T) {
	resp, _ := http.Get("http://apigateway/ping")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("TestHttppingService:", string(body))

}

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

	conn.Write([]byte("tcpping\n"))
	line, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("TestTcppingService:", line)
}
