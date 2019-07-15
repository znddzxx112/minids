/**
tcp proto
 */
package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var headerLenthErr error = errors.New("headbuf length less 4")

const headerLength = 4

type TransCoder struct {
	rw io.ReadWriter
}

func NewTransCoder(rw io.ReadWriter) *TransCoder {
	t := new(TransCoder)
	t.rw = rw
	return t
}

func (t *TransCoder) Send(body *[]byte) error {
	header, berr := int32Tobytes(int32(len(*body)))
	if berr != nil {
		return berr
	}
	t.rw.Write(header)
	t.rw.Write(*body)
	return nil
}

func (t *TransCoder) Receive() ([]byte, error) {
	headbuf := make([]byte, headerLength)
	rlen, err := io.ReadFull(t.rw, headbuf)
	if err != nil {
		return nil, err
	}
	if rlen < headerLength {
		return nil, headerLenthErr
	}

	bodyLen, cbytesToint32err := bytesToint32(headbuf)
	if cbytesToint32err != nil {
		return nil, cbytesToint32err
	}

	body := make([]byte, bodyLen)
	_, readFullerr := io.ReadFull(t.rw, body)
	if readFullerr != nil {
		return nil, readFullerr
	}
	return body, nil
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