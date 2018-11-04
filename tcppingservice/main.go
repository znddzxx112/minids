package main

import (
	"bytes"
	"fmt"
	"minids/common"
	"net"
)

var serverIp string
var listenPort = ":8081"
var serverIPAndPort string

func init() {
	serverIp, _ = common.LocalIp()
	serverIPAndPort = serverIp + listenPort
}

func main() {
	zkConn := common.RegisterServiceName("tcpping", serverIPAndPort)
	defer zkConn.Close()

	ln, lnErr := net.Listen("tcp", listenPort)
	if lnErr != nil {
		panic("listen failed")
	}
	defer ln.Close()

	for {
		conn, connErr := ln.Accept()
		if connErr != nil {
			continue
		}
		go tcpHandleFunc(conn)
	}
}

func tcpHandleFunc(conn net.Conn) {
	defer conn.Close()

	line := common.GetStringFromTcpConn(conn)
	fmt.Println(line)

	var buffer bytes.Buffer
	buffer.Write([]byte(serverIPAndPort))
	buffer.Write([]byte(" recevie:"))
	buffer.Write(line)
	buffer.Write([]byte("\n"))

	conn.Write(buffer.Bytes())
}
