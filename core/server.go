package core

import "sync"

const MaxClientsNum = 512
const Version = "0.1.1"

type Server struct {
	sync.RWMutex
	Cmds       map[string]*Command
	Clients    []*Client
	ClientsNum int
}

func NewServer() *Server {
	cmds := make(map[string]*Command)
	clients := make([]*Client, MaxClientsNum)
	return &Server{
		Cmds:       cmds,
		Clients:    clients,
		ClientsNum: 0,
	}
}

func (s *Server) RegisterService(cmd *Command) {
	s.Cmds[cmd.Name] = cmd
}

func (s *Server) LookupService(name string) (*Command, error) {
	if cmd, ok := s.Cmds[name]; ok {
		return cmd, nil
	}
	return nil, ErrNotFoundCmd
}

func (s *Server) Commands() map[string]*Command {
	return s.Cmds
}

func (s *Server) AddClient(client *Client) error {
	s.Lock()
	defer s.Unlock()
	if s.ClientsNum >= MaxClientsNum {
		return ErrClientsTooMany
	}
	s.Clients = append(s.Clients, client)
	s.ClientsNum++
	return nil
}

func (s *Server) DeleteClient(client *Client) {
	s.Lock()
	defer s.Unlock()
	for k, c := range s.Clients {
		if c == client {
			s.Clients = append(s.Clients[0:k], s.Clients[k+1:]...)
			break
		}
	}
	s.ClientsNum--
}

func (s *Server) AcceptRequest(c *Client) error {
	var err error
	d := NewDecoder(*c.Conn)
	c.Req, err = d.DecodeReq()
	return err
}

func (s *Server) SendResponse(c *Client) {
	e := NewEncoder(*c.Conn)
	e.EncodeResp(c.Resp)
	return
}
