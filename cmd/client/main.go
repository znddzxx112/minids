package main

import (
	"github.com/znddzxx112/minids/client"
	"github.com/znddzxx112/minids/services/core"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go core.DefaultHandler.SigHandler(c)

	client.Process("127.0.0.1", "9736")


}