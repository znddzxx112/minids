package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"minids/zkcommon"
	"net/http"
	"time"
)

func main() {
	serverList := GetServerList()

	serverIp, randErr := randServerIp(serverList)
	if randErr != nil {
		panic(randErr.Error())
	}

	router := gin.Default()

	// forward httpingservice
	router.GET("/ping", func(c *gin.Context) {
		resp, _ := httpGet(serverIp, "/ping")
		resp.Write(c.Writer)
	})

	router.Run(":80")

}

func GetServerList() []string {
	conn, err := zkcommon.GetConnect([]string{"zookeeper"}, time.Second)
	if err != nil {
		panic("zookeeper connect failed")
	}
	defer conn.Close()

	serverList, getErr := zkcommon.GetServerList(conn, "httpping")
	if getErr != nil {
		panic("get serverList failed")
	}
	return serverList
}

func randServerIp(serverList []string) (string, error) {
	count := len(serverList)
	if count == 0 {
		return "", errors.New("serverList is empty")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	serverIp := serverList[r.Intn(count)]
	return serverIp, nil
}

func httpGet(host string, location string) (resp *http.Response, err error) {
	url := fmt.Sprintf("http://%s%s", host, location)
	resp, err = http.Get(url)
	return resp, err
	/**
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
	**/
}
