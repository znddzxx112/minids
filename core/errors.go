package core

import "errors"

var (
	ErrRespError      = errors.New("resp has error")
	ErrCmdIsEmpty     = errors.New("cmd is empty")
	ErrNotFoundCmd    = errors.New("not found cmd")
	ErrDecodeRespRead = errors.New("DecodeResp fail")
	ErrDecodeReqRead  = errors.New("DecodeReq fail")
	ErrReqFormat      = errors.New("req format error")
	ErrClientsTooMany = errors.New("too many clients")
	ErrClientExit     = errors.New("client exit")
)
