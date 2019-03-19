package gateway

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"minids/core/common"
	"net"
	"net/http"
)

func main() {

	router := gin.Default()

	// forward httpingservice
	router.GET("/ping", func(c *gin.Context) {
		serverIp, serverErr := common.GetServerList("httpping")
		if serverErr != nil {
			panic(serverErr.Error())
		}
		resp, _ := httpGet(serverIp, "/ping")
		resp.Write(c.Writer)
	})

	// forward tcppingservice
	router.GET("/tcpping", func(c *gin.Context) {
		serverIp, serverErr := common.GetServerList("tcpping")
		if serverErr != nil {
			panic(serverErr.Error())
		}
		fmt.Println(serverIp)
		conn, dialErr := net.Dial("tcp", serverIp)
		if dialErr != nil {
			panic(dialErr.Error())
		}
		defer conn.Close()

		conn.Write([]byte("tcpping\n"))
		line, _ := bufio.NewReader(conn).ReadString('\n')

		c.String(http.StatusOK, string(line))
	})

	router.Run(":80")

}

func httpGet(host string, location string) (resp *http.Response, err error) {
	url := fmt.Sprintf("http://%s%s", host, location)
	resp, err = http.Get(url)
	return resp, err
}
