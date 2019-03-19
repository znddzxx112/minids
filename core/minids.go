package core

type Command struct {
	Name string
	Proc ServiceFunc
}

func NewCommand(name string, proc ServiceFunc) *Command {
	cmd := new(Command)
	cmd.Name = name
	cmd.Proc = proc
	return cmd
}

type ServiceFunc func(c *Client, s *Server) error

func ProcessService(c *Client, s *Server) error {
	if c.Cmd == nil {
		return ErrCmdIsEmpty
	}
	cmd, err := s.LookupService(c.Cmd.Name)
	if err != nil {
		return err
	}
	return cmd.Proc(c, s)
}
