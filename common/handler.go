package common

import (
	"fmt"
	"os"
	"syscall"
)

var DefaultHandler *Handler

func init()  {
	DefaultHandler = &Handler{}
}

type Handler struct {

}

func (h *Handler)SigHandler(c chan os.Signal) {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			h.ExitHandler()
		default:
			fmt.Println("signal ", s)
		}
	}
}

func (h *Handler)ExitHandler() {
	fmt.Println("exiting minids ... bye !")
	os.Exit(0)
}

func (h *Handler)HelpClientHandler() {
	fmt.Println("minids is a mini distributed system")
	fmt.Println("Usage: minids [OPTIONS]")
	fmt.Println("OPTIONS:")
	fmt.Println("\t-h, --host=name     Connect to host.Default is 127.0.0.1.")
	fmt.Println("\t-p, --port=port     Port number to use for connection.Default is 9736.")
}