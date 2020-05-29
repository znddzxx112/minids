/**
tcp proto
*/
package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
)

const headerLength = 4

var headerLenthErr error = errors.New("headbuf length less " + strconv.Itoa(headerLength))
var bodyLenthErr error = errors.New("body read bytes not enough")

type TransCoder struct {
	rw io.ReadWriter
}

func NewTransCoder(rw io.ReadWriter) *TransCoder {
	t := new(TransCoder)
	t.rw = rw
	return t
}

func (t *TransCoder) Send(body *[]byte) error {
	header, int32TobytesErr := int32Tobytes(int32(len(*body)))
	if int32TobytesErr != nil {
		return int32TobytesErr
	}

	if _, writeErr := t.rw.Write(bytesCombine(header, *body)); writeErr != nil {
		return writeErr
	}

	return nil
}

func (t *TransCoder) Receive() (*[]byte, error) {
	headbuf := make([]byte, headerLength)
	readHeadN, ReadFullErr := io.ReadFull(t.rw, headbuf)
	if ReadFullErr != nil {
		return nil, ReadFullErr
	}

	if readHeadN < headerLength {
		return nil, headerLenthErr
	}

	bodyLen, bytesToint32Err := bytesToint32(headbuf)
	if bytesToint32Err != nil {
		return nil, bytesToint32Err
	}

	body := make([]byte, bodyLen)
	readBodyN, readFullerr := io.ReadFull(t.rw, body)
	if readFullerr != nil {
		return nil, readFullerr
	}

	if int32(readBodyN) < bodyLen {
		return nil, bodyLenthErr
	}

	return &body, nil
}

func bytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func int32Tobytes(len int32) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, len); err == nil {
		return buf.Bytes(), nil
	} else {
		return nil, err
	}
}

func bytesToint32(buf []byte) (int32, error) {
	b_buf := bytes.NewBuffer(buf)
	var i32 int32
	if err := binary.Read(b_buf, binary.BigEndian, &i32); err == nil {
		return i32, nil
	} else {
		return 0, err
	}
}
