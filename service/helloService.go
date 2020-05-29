package services

import "github.com/znddzxx112/minids/services/core"

func HelloService(cmd *core.Command) error {
	cmd.ResInfo = cmd.ArgInfo
	return nil
}
