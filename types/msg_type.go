package types

import "fmt"

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
	MsgSetVringErr         MsgType = 14
	MsgGetProtocolFeatures MsgType = 15
	MsgSetProtocolFeatures MsgType = 16
	MsgGetQueueNum         MsgType = 17
	MsgSetVringEnable      MsgType = 18
	MsgSetBackendReqFd     MsgType = 21
	MsgGetConfig           MsgType = 24
	MsgSetConfig           MsgType = 25
)

func (m MsgType) String() string {
	switch m {
	case MsgGetFeatures:
		return "GET_FEATURES"
	case MsgSetFeatures:
		return "SET_FEATURES"
	case MsgSetOwner:
		return "SET_OWNER"
	case MsgSetMemTable:
		return "SET_MEM_TABLE"
	case MsgSetVringNum:
		return "SET_VRING_NUM"
	case MsgSetVringAddr:
		return "SET_VRING_ADDR"
	case MsgSetVringBase:
		return "SET_VRING_BASE"
	case MsgGetVringBase:
		return "GET_VRING_BASE"
	case MsgSetVringKick:
		return "SET_VRING_KICK"
	case MsgSetVringCall:
		return "SET_VRING_CALL"
	case MsgGetProtocolFeatures:
		return "GET_PROTOCOL_FEATURES"
	case MsgSetProtocolFeatures:
		return "SET_PROTOCOL_FEATURES"
	case MsgGetQueueNum:
		return "GET_QUEUE_NUM"
	case MsgSetBackendReqFd:
		return "SET_BACKEND_REQ_FD"
	case MsgGetConfig:
		return "GET_CONFIG"
	case MsgSetConfig:
		return "SET_CONFIG"
	case MsgSetVringErr:
		return "SET_VRING_ERR"
	case MsgSetVringEnable:
		return "SET_VRING_ENABLE"

	default:
		return fmt.Sprintf("UNKNOWN(0x%x)", uint32(m))
	}
}
