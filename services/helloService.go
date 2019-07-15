package services

import "github.com/znddzxx112/minids/core"

func HelloService(cmd *core.Command) error {
	cmd.ResInfo = cmd.ArgInfo
	return nil
}
