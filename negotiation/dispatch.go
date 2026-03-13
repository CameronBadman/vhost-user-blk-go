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

	}
	return nil
}
