package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttppingService(t *testing.T) {
	resp, _ := http.Get("http://apigateway/ping")

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("TestHttppingService:", string(body))

}
