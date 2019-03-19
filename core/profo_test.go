package core

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestNewReq(t *testing.T) {
	req := NewReq([]string{"foo", "bar"})
	if req.Argc != 2 {
		t.Error("Error req argc", req.Argc)
	}
	if req.Argv[0] != "foo" {
		t.Error("Error req Argv", req.Argv[0])
	}
	if req.Argv[1] != "bar" {
		t.Error("Error req Argv", req.Argv[1])
	}
}

func TestNewEncoder(t *testing.T) {
	var str string
	buf := bytes.NewBufferString(str)
	e := NewEncoder(buf)
	if _, ok := e.wt.(io.Writer); !ok {
		t.Error("Error NewEncoder")
	}
}

func TestNewDecoder(t *testing.T) {
	var str string
	rd := strings.NewReader(str)
	d := NewDecoder(rd)
	if _, ok := d.rd.(io.Reader); !ok {
		t.Error("Error NewDecoder")
	}
}

func TestNewResp(t *testing.T) {
	resp := NewResp(ErrRespError, []string{"foo", "bar"})
	if resp.Err != ErrRespError {
		t.Error("Error resp err", resp.Err)
	}
	if resp.Argc != 2 {
		t.Error("Error resp Argc", resp.Argc)
	}
	if resp.Argv[0] != "foo" && resp.Argv[1] != "bar" {
		t.Error("Error resp Argv", resp.Argv)
	}
}

func TestDecoder_DecodeResp(t *testing.T) {
	rd := strings.NewReader("+2\r\nfoo\r\nbar\r\n")
	decoder := NewDecoder(rd)
	resp, err := decoder.DecodeResp()
	if err != nil {
		t.Error("Error decode resp", err)
	}
	if resp.Err != nil {
		t.Error("Error decode resp.err", resp.Err)
	}
	if resp.Argc != 2 {
		t.Error("Error decode resp.Argc", resp.Argc)
	}
	if resp.Argv[0] != "foo" {
		t.Error("Error decode resp.Argv[0]", resp.Argv[0])
	}
	if resp.Argv[1] != "bar" {
		t.Error("Error decode resp.Argv[1]", resp.Argv[1])
	}
}

func TestEncoder_EncodeResp(t *testing.T) {
	resp := NewResp(ErrRespError, []string{"foo", "bar"})
	var str string
	buf := bytes.NewBufferString(str)
	encoder := NewEncoder(buf)
	encoder.EncodeResp(resp)

	if buf.String() != "-2\r\nfoo\r\nbar\r\n" {
		t.Error("Error encode str", str)
	}
}

func TestDecoder_DecodeReq(t *testing.T) {
	var str string = "$2\r\nfoo\r\nbar\r\n"
	rd := strings.NewReader(str)
	d := NewDecoder(rd)
	req, err := d.DecodeReq()
	if err != nil || req.Argc != 2 || req.Argv[0] != "foo" || req.Argv[1] != "bar" {
		t.Error("Error DecodeReq", req)
	}
}

func TestEncoder_EncodeReq(t *testing.T) {
	var str string
	buf := bytes.NewBufferString(str)
	e := NewEncoder(buf)
	req := NewReq([]string{"foo", "bar"})
	e.EncodeReq(req)

	if buf.String() != "$2\r\nfoo\r\nbar\r\n" {
		t.Error("Error encode", buf.String())
	}
}
