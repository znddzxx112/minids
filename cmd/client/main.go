package main

import (
	client "github.com/znddzxx112/minids/cli"
	"github.com/znddzxx112/minids/common"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go common.DefaultHandler.SigHandler(c)

	client.Process("127.0.0.1", "9736")


}