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
	}
	return nil
}
