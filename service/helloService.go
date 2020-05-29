package service

import "github.com/znddzxx112/minids/common"

func HelloService(cmd *common.Command) error {
	cmd.ResInfo = cmd.ArgInfo
	return nil
}
