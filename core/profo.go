package core

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type ArgsInfo struct {
	Argv []string
	Argc int
}

func NewArgsInfo(queryString string) *ArgsInfo {
	argsinfo := new(ArgsInfo)
	queryString = strings.TrimRight(queryString, "\r\n")
	argsinfo.Argv = strings.Split(queryString, " ")
	argsinfo.Argc = len(argsinfo.Argv)
	return argsinfo
}

type Req struct {
	ArgsInfo
}

func NewReq(argv []string) *Req {
	req := new(Req)
	req.Argv = argv
	req.Argc = len(argv)
	return req
}

type Resp struct {
	Err error
	ArgsInfo
}

func NewResp(err error, argv []string) *Resp {
	resp := new(Resp)
	resp.Err = err
	resp.Argc = len(argv)
	resp.Argv = argv
	return resp
}

type Encoder struct {
	wt io.Writer
}

func NewEncoder(wt io.Writer) *Encoder {
	e := new(Encoder)
	e.wt = wt
	return e
}

// output:"+2\r\nfoo\r\nbar\r\n"
func (e *Encoder) EncodeResp(resp *Resp) {
	bf := bufio.NewWriter(e.wt)
	if resp.Err != nil {
		bf.WriteString("-")
	} else {
		bf.WriteString("+")
	}
	bf.WriteString(strconv.Itoa(resp.Argc))
	bf.WriteString("\r\n")
	for _, str := range resp.Argv {
		bf.WriteString(str)
		bf.WriteString("\r\n")
	}
	bf.Flush()
	return
}

// output:"$2\r\nfoo\r\nbar\r\n"
func (e *Encoder) EncodeReq(req *Req) {
	bf := bufio.NewWriter(e.wt)
	bf.WriteString("$")
	bf.WriteString(strconv.Itoa(req.Argc))
	bf.WriteString("\r\n")
	for _, str := range req.Argv {
		bf.WriteString(str)
		bf.WriteString("\r\n")
	}
	bf.Flush()
	return
}

type Decoder struct {
	rd io.Reader
}

func NewDecoder(rd io.Reader) *Decoder {
	d := new(Decoder)
	d.rd = rd
	return d
}

func (d *Decoder) DecodeResp() (*Resp, error) {
	resp := new(Resp)

	bf := bufio.NewReader(d.rd)
	str, err := bf.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, ErrDecodeRespRead
	}
	str = strings.TrimRight(str, "\r\n")
	switch str[:1] {
	case "+":
		resp.Err = nil
	case "-":
		resp.Err = ErrRespError
	}

	resp.Argc, _ = strconv.Atoi(str[1:])
	resp.Argv = make([]string, resp.Argc)
	for i := 0; i < resp.Argc; i++ {
		str, err := bf.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, ErrDecodeRespRead
		}
		resp.Argv[i] = strings.TrimRight(str, "\r\n")
	}
	return resp, nil
}

func (d *Decoder) DecodeReq() (*Req, error) {
	req := new(Req)

	bf := bufio.NewReader(d.rd)
	str, err := bf.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return nil, io.EOF
		}
		return nil, ErrDecodeReqRead
	}
	str = strings.TrimRight(str, "\r\n")
	if str[:1] != "$" {
		return nil, ErrReqFormat
	}
	req.Argc, _ = strconv.Atoi(str[1:])

	req.Argv = make([]string, req.Argc)
	for i := 0; i < req.Argc; i++ {
		str, err := bf.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, ErrDecodeRespRead
		}
		req.Argv[i] = strings.TrimRight(str, "\r\n")
	}
	return req, nil
}
