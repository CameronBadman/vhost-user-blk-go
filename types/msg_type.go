package types

type MsgType uint32

const (
	MsgNone                MsgType = 0
	MsgGetFeatures         MsgType = 1
	MsgSetFeatures         MsgType = 2
	MsgSetOwner            MsgType = 3
	MsgSetMemTable         MsgType = 5
	MsgSetVringNum         MsgType = 8
	MsgSetVringAddr        MsgType = 9
	MsgSetVringBase        MsgType = 10
	MsgGetVringBase        MsgType = 11
	MsgSetVringKick        MsgType = 12
	MsgSetVringCall        MsgType = 13
	MsgGetProtocolFeatures MsgType = 15
	MsgSetProtocolFeatures MsgType = 16
	MsgGetQueueNum         MsgType = 17
	MsgSetBackendReqFd     MsgType = 21
	MsgGetConfig           MsgType = 24
	MsgSetConfig           MsgType = 25
)
