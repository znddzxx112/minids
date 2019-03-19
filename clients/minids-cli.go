package main

import (
	"bufio"
	"fmt"
	"minids/core"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	IPPort := "127.0.0.1:21115"

	/*---- 监听信号 平滑退出 ----*/
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go sigHandler(c)

	intputBuf := bufio.NewReader(os.Stdin)
	fmt.Println("Hi minids")
	conn, err := net.Dial("tcp", IPPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer conn.Close()

	client := core.NewClient(&conn)
	for {
		fmt.Print(IPPort, ">")
		queryString, err := intputBuf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		if strings.TrimRight(queryString, "\r\n") == "exit" {
			exitHandler()
			return
		}
		client.SetCommandAndArgInfo(queryString)
		client.SendRequst()
		client.RerviceResponse()
		resp := client.Response()
		if resp.Argc > 0 {
			for _, res := range resp.Argv {
				fmt.Println(res)
			}
		}
	}
}

func sigHandler(c chan os.Signal) {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			exitHandler()
		default:
			fmt.Println("signal ", s)
		}
	}
}

func exitHandler() {
	fmt.Println("exiting minids ... bye !")
	os.Exit(0)
}
