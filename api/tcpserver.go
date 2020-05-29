package api

import (
	"github.com/golang/protobuf/proto"
	"github.com/znddzxx112/minids/common"
	"github.com/znddzxx112/minids/protos"
	"github.com/znddzxx112/minids/service"
	"net"
)

var serviceRegistry *common.ServiceRegistry

func init() {
	serviceRegistry = common.NewServiceRegistry()

	helloCmd := common.NewCommand("hello", service.HelloService)
	serviceRegistry.Registry(helloCmd)
}

func Process() error {

	ln, err := net.Listen("tcp", ":9736")
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, acceptErr := ln.Accept()
		if acceptErr != nil {
			continue
		}

		go tcpserver(conn)
	}

}

func tcpserver(conn net.Conn) {
	transCoder := core.NewTransCoder(conn)
	defer conn.Close()

	for {
		bs, ReceiveErr := transCoder.Receive()
		if ReceiveErr != nil {
			break
		}

		cmdpb := &protos.Cmd{}
		unmarshalErr := proto.Unmarshal(*bs, cmdpb)
		if unmarshalErr != nil {
			break
		}

		command := core.NewCommandArgInfo(cmdpb.Name, &cmdpb.ArgInfo)
		if processErr := serviceRegistry.Process(command); processErr != nil {
			// add log
		}

		responseCmdPb := &protos.Cmd{}
		responseCmdPb.Name = command.Name
		responseCmdPb.ResInfo = *command.ResInfo

		respbf, marshalErr := proto.Marshal(responseCmdPb)
		if marshalErr != nil {
			break
		}

		SendErr := transCoder.Send(&respbf)
		if SendErr != nil {
			break
		}
	}

}
