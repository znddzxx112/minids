package core

import (
	"net"
)

type Client struct {
	Cmd  *Command
	Conn *net.Conn
	Req  *Req
	Resp *Resp
}

func NewClient(Conn *net.Conn) *Client {
	c := new(Client)
	c.Conn = Conn
	return c
}

func (c *Client) SetCommand(cmd *Command) {
	c.Cmd = cmd
}

func (c *Client) Command() *Command {
	return c.Cmd
}

func (c *Client) SetRequst(req *Req) {
	c.Req = req
}

func (c *Client) Requst() *Req {
	return c.Req
}

func (c *Client) SendRequst() {
	encoder := NewEncoder(*c.Conn)
	encoder.EncodeReq(c.Req)
}

func (c *Client) RerviceResponse() error {
	var err error
	decode := NewDecoder(*c.Conn)
	if c.Resp, err = decode.DecodeResp(); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Client) SetResponse(resp *Resp) {
	c.Resp = resp
}

func (c *Client) Response() *Resp {
	return c.Resp
}

func (c *Client) SetCommandAndArgInfo(queryString string) {
	argsinfo := NewArgsInfo(queryString)
	req := NewReq(argsinfo.Argv)
	c.SetRequst(req)

	c.Cmd = NewCommand(req.Argv[0], nil)
}
