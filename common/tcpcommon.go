package common

import (
	"bytes"
	"net"
	"strconv"
)

func DecodeMsgEnd(conn net.Conn) []byte {
	allbuffer := make([]byte, 0)
	var allbufferLen int = 0
	for {
		var buffer [1024]byte
		readSize, err := conn.Read(buffer[:])
		if err != nil {
			break
		}
		for i := 0; i < readSize; i++ {
			allbuffer = append(allbuffer, buffer[i])
			allbufferLen += 1
			// endline flag is "\r\n\r\n"
			if allbufferLen >= 4 &&
				buffer[allbufferLen-1] == '\n' &&
				buffer[allbufferLen-2] == '\r' &&
				buffer[allbufferLen-3] == '\n' &&
				buffer[allbufferLen-4] == '\r' {
				return allbuffer[0 : allbufferLen-4]
			}
		}
	}
	return nil
}

func EncodeMsgEnd(body []byte) []byte {
	var bodyBuffer bytes.Buffer
	bodyBuffer.Write(body)
	bodyBuffer.Write([]byte("\r\n\r\n"))
	return bodyBuffer.Bytes()
}

func DecodeFrame(conn net.Conn) []byte {
	allbuffer := make([]byte, 0)
	var bodyLen int = 0
	var headLen int = 0
	var readLen int = 0
	var strErr error
	for {
		var buffer [1024]byte
		readSize, err := conn.Read(buffer[:])
		if err != nil {
			break
		}
		for i := 0; i < readSize; i++ {
			allbuffer = append(allbuffer, buffer[i])
			readLen += 1
			if allbuffer[0] != '*' {
				return nil
			}
			if allbuffer[readLen-1] == '$' {
				bodyLen, strErr = strconv.Atoi(string(allbuffer[1 : readLen-1]))
				if strErr != nil {
					return nil
				}
				headLen = readLen
			}
			if headLen != 0 && bodyLen != 0 && headLen+bodyLen <= readLen {
				return allbuffer[headLen : bodyLen+headLen]
			}
		}
	}
	return nil
}

func EncodeFrame(body []byte) []byte {
	bodyLen := len(body)

	var frameBuffer bytes.Buffer
	frameBuffer.Write([]byte("*" + strconv.Itoa(bodyLen) + "$"))
	frameBuffer.Write(body)
	return frameBuffer.Bytes()
}
