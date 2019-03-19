package info

import (
	"minids/core"
	"strconv"
)

// talk something about server
// server infomation
func InfoService(c *core.Client, s *core.Server) error {
	var info []string

	if c.Req.Argc > 1 {
		switch c.Req.Argv[1] {
		case "service":
			info = append(info, "service list:")
			cmds := s.Commands()
			for cmdName, _ := range cmds {
				info = append(info, cmdName)
			}
		case "version":
			info = append(info, "version"+core.Version)
		case "clientnum":
			info = append(info, strconv.Itoa(s.ClientsNum))
		}
	} else {
		info = []string{"commands list:", "info service", "info version", "info clientnum"}
	}
	resp := core.NewResp(nil, info)
	c.SetResponse(resp)
	return nil
}

// client exit
func ExitService(c *core.Client, s *core.Server) error {
	resp := core.NewResp(core.ErrClientExit, []string{})
	c.SetResponse(resp)
	return core.ErrClientExit
}
