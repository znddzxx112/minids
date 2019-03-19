package main

import (
	"bytes"
	"fmt"
	"minids/core/common"
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
	go tcpMsgFunc()

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

	line := common.DecodeMsgEnd(conn)
	fmt.Println(line)

	var bodyBuffer bytes.Buffer
	bodyBuffer.Write([]byte(serverIPAndPort))
	bodyBuffer.Write([]byte(" recevie:"))
	bodyBuffer.Write(line)

	conn.Write(common.EncodeMsgEnd(bodyBuffer.Bytes()))
}

func tcpMsgFunc() {
	tcpMsgListenPort := ":8082"
	tcpMsgServerIPAndPort := serverIp + tcpMsgListenPort
	zkConn := common.RegisterServiceName("tcpmsgping", tcpMsgServerIPAndPort)
	defer zkConn.Close()

	ln, lnErr := net.Listen("tcp", tcpMsgListenPort)
	if lnErr != nil {
		panic("listen failed")
	}
	defer ln.Close()

	for {
		conn, connErr := ln.Accept()
		if connErr != nil {
			continue
		}
		go tcpMsgHandleFunc(conn, tcpMsgListenPort)
	}
}

func tcpMsgHandleFunc(conn net.Conn, tcpMsgListenPort string) {
	defer conn.Close()

	line := common.DecodeFrame(conn)
	fmt.Println(line)

	var bodyBuffer bytes.Buffer
	bodyBuffer.Write([]byte(tcpMsgListenPort))
	bodyBuffer.Write([]byte("recevie:"))
	bodyBuffer.Write(line)

	conn.Write(common.EncodeFrame(bodyBuffer.Bytes()))
}
