package common

import "errors"

var SERVICE_EXIST_ERR = errors.New("service has exist")
var SERVICE_NOT_FOUND = errors.New("service not find")

type ServiceRegistry struct {
	Cmds map[string]*Command
}

type Command struct {
	Name    string
	Service ServiceFunc
	ArgInfo *[]byte // protobuf's data
	ResInfo *[]byte // result protobuf's data
}

type ServiceFunc func(com *Command) error

func NewCommand(name string, service ServiceFunc) *Command {
	c := new(Command)
	c.Name = name
	c.Service = service
	return c
}

func NewCommandArgInfo(name string, argInfo *[]byte) *Command {
	c := new(Command)
	c.Name = name
	c.ArgInfo = argInfo
	return c
}

func NewServiceRegistry() *ServiceRegistry {
	s := new(ServiceRegistry)
	cmds := make(map[string]*Command)
	s.Cmds = cmds
	return s
}

func (s *ServiceRegistry) Registry(c *Command) error {
	if _, ok := s.Cmds[c.Name]; ok {
		return SERVICE_EXIST_ERR
	}
	s.Cmds[c.Name] = c
	return nil
}

func (s *ServiceRegistry) Lookup(c *Command) (*Command, error) {
	if c, ok := s.Cmds[c.Name]; ok {
		return c, nil
	} else {
		return nil, errors.New("service not find")
	}
}

func (s *ServiceRegistry) Process(c *Command) error {
	if command, ok := s.Cmds[c.Name]; !ok {
		return SERVICE_NOT_FOUND
	} else {
		return command.Service(c)
	}
}
