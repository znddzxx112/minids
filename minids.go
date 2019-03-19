package main

import (
	"errors"
	"fmt"
	"io"
	"minids/core"
	"minids/services/info"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// this is minids's server
// todo:
// 1. config file
// 2. cmd config

func main() {
	// port from config
	IPPort := "127.0.0.1:21115"
	server := core.NewServer()

	/*---- 监听信号 平滑退出 ----*/
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go sigHandler(c)

	// register Service from config
	server.RegisterService(core.NewCommand("info", info.InfoService))
	server.RegisterService(core.NewCommand("exit", info.ExitService))

	ln, err := net.Listen("tcp", IPPort)
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(&conn, server)
	}
}

func handle(conn *net.Conn, server *core.Server) {
	defer (*conn).Close()

	client := core.NewClient(conn)
	server.AddClient(client)
	defer server.DeleteClient(client)
	for {
		if err := server.AcceptRequest(client); err != nil {
			if err == io.EOF {
				return
			}
			client.SetResponse(core.NewResp(err, []string{err.Error()}))
			server.SendResponse(client)
			continue
		}

		client.SetCommand(core.NewCommand(client.Requst().Argv[0], nil))
		if err := core.ProcessService(client, server); err != nil {
			if err == core.ErrClientExit {
				return
			}
			client.SetResponse(core.NewResp(err, []string{err.Error()}))
			server.SendResponse(client)
			continue
		}

		if client.Response() == nil {
			client.SetResponse(core.NewResp(errors.New("service don't set Respon"), []string{"service don't set Respon"}))
		}
		server.SendResponse(client)
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
