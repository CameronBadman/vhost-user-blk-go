package negotiation

import (
	"vhost-go/blk"
	"vhost-go/transport"
	"vhost-go/types"
	"vhost-go/wires"
)

func Dispatch(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	switch msg.Request {
	case types.MsgGetFeatures:
		return getFeatures(socket)
	case types.MsgSetFeatures:
		return setFeatures(dev, msg)
	case types.MsgGetProtocolFeatures:
		return getProtocolFeatures(socket)
	case types.MsgSetProtocolFeatures:
		return setProtocolFeatures(dev, msg)
	case types.MsgSetMemTable:
		return setMemTable(dev, socket, msg)
	case types.MsgGetQueueNum:
		return getQueueNum(dev, socket, msg)
	case types.MsgGetConfig:
		return getConfig(dev, socket, msg)
	case types.MsgSetOwner:
		return setOwner(dev)
	case types.MsgSetVringNum:
		return sendNotImplemented(socket, msg)
	case types.MsgSetVringAddr:
		return sendNotImplemented(socket, msg)
	case types.MsgSetVringBase:
		return sendNotImplemented(socket, msg)
	case types.MsgSetVringKick:
		return sendNotImplemented(socket, msg)
	case types.MsgSetVringCall:
		return sendNotImplemented(socket, msg)
	case types.MsgSetVringEnable:
		return sendNotImplemented(socket, msg)
	default:
		return sendNotImplemented(socket, msg)
	}
}
