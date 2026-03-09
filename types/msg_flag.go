package types

type MsgFlag uint32

const (
	MsgFlagVersion MsgFlag = 0x1
	MsgFlagReply   MsgFlag = 0x4
	MsgFlagErrMsg  MsgFlag = 0x8
)
