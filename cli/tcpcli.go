package client

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/znddzxx112/minids/services/core"
	"github.com/znddzxx112/minids/services/protos"
	"io"
	"net"
	"os"
)

func Process(ip, port string) {

	addr := fmt.Sprintf("%s:%s", ip, port)
	conn, DialErr := net.Dial("tcp", addr)
	if DialErr != nil {
		fmt.Println(DialErr.Error())
		os.Exit(1)
	}
	defer conn.Close()

	stdInRd := bufio.NewReader(os.Stdin)
	fmt.Println("Hi minids")
	for {
		fmt.Printf(addr + ">")

		iBytes, ReadBytesErr := stdInRd.ReadBytes('\n')
		if ReadBytesErr != nil {
			fmt.Println(ReadBytesErr)
			continue
		}
		inputBytes := bytes.TrimRight(iBytes, "\r\n")

		inputCmd := string(inputBytes)
		switch inputCmd {
		case "exit":
			core.DefaultHandler.ExitHandler()
			return
		case "help":
			core.DefaultHandler.HelpClientHandler()
			continue
		default:
			resBytes, inputBytesErr := tcpclient(conn, &inputBytes)
			if inputBytesErr != nil {
				fmt.Println(inputBytesErr.Error())
				continue
			}
			fmt.Println(string(*resBytes))
		}

	}
}

func tcpclient(conn io.ReadWriter, iBytes *[]byte) (*[]byte, error) {
	iBytesSlice := bytes.Split(*iBytes, []byte(" "))

	reqCmd := &protos.Cmd{}
	reqCmd.Name = string(iBytesSlice[0])
	reqCmd.ArgInfo = bytes.Join(iBytesSlice[1:], []byte(" "))

	reqBinary, MarshalErr := proto.Marshal(reqCmd)
	if MarshalErr != nil {
		return nil, MarshalErr
	}

	transCoder := core.NewTransCoder(conn)

	SendErr := transCoder.Send(&reqBinary)
	if SendErr != nil {
		return nil, SendErr
	}

	respBinary, ReceiveErr := transCoder.Receive()
	if ReceiveErr != nil {
		return nil, ReceiveErr
	}

	unmarshalErr := proto.Unmarshal(*respBinary, reqCmd)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &reqCmd.ResInfo, nil
}




