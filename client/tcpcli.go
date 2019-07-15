package client

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/znddzxx112/minids/core"
	"github.com/znddzxx112/minids/protos"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func Process() {
	ip := "127.0.0.1"
	port := "9736"
	addr := ip + ":" + port

	/*---- 监听信号 平滑退出 ----*/
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go sigHandler(c)

	stdInRd := bufio.NewReader(os.Stdin)
	fmt.Println("Hi minids")

	conn, dialErr := net.Dial("tcp", addr)
	if dialErr != nil {
		fmt.Println(dialErr.Error())
		os.Exit(1)
	}
	defer conn.Close()

	for {
		fmt.Printf(addr + ">")

		iBytes, err := stdInRd.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		if strings.EqualFold(string(bytes.TrimRight(iBytes, "\r\n")), "exit") {
			exitHandler()
			return
		}

		resBytes, tcpcliProcess := Tcpclient(conn, &iBytes)
		if tcpcliProcess != nil {
			fmt.Println(tcpcliProcess.Error())
		}
		fmt.Println(string(*resBytes))
	}
}

func Tcpclient(conn io.ReadWriter, iBytes *[]byte) (*[]byte, error) {

	iBytesSlice := bytes.Split(*iBytes, []byte(" "))
	reqCmd := &protos.Cmd{}
	reqCmd.Name = string(iBytesSlice[0])
	reqCmd.ArgInfo = bytes.Join(iBytesSlice[1:], []byte(" "))

	reqBytes, marshalErr := proto.Marshal(reqCmd)
	if marshalErr != nil {
		return nil, marshalErr
	}

	transCoder := core.NewTransCoder(conn)
	sendErr := transCoder.Send(&reqBytes)
	if sendErr != nil {
		return nil, sendErr
	}

	respBytes, receiveErr := transCoder.Receive()
	if receiveErr != nil {
		return nil, receiveErr
	}

	unmarshalErr := proto.Unmarshal(respBytes, reqCmd)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &reqCmd.ResInfo, nil
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
