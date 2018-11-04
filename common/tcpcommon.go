package common

import (
	"bufio"
	"net"
)

func GetStringFromTcpConn(conn net.Conn) []byte {
	//var buffer bytes.Buffer
	//for {
	//	var buff [1024]byte
	//	readnum, readErr := conn.Read(buff[:])
	//	if readnum < 4 || readErr != nil {
	//		break
	//	}
	//	buffer.Write(buff[:readnum])
	//	if buff[readnum] == '\n' {
	//		break
	//	}
	//}
	//return buffer.Bytes()
	line, _ := bufio.NewReader(conn).ReadBytes('\n')
	return line
}
