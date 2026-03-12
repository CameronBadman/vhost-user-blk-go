package negotiation

import (
	"encoding/binary"

	"vhost-go/blk"
	"vhost-go/transport"
	"vhost-go/types"
	"vhost-go/wires"
)

func Dispatch(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	switch msg.Request {
	case types.MsgGetFeatures:
		reply := &wires.VhostUserMsg{
			Request: types.MsgGetFeatures,
			Flags:   0x5,
			Size:    8,
			Payload: make([]byte, 8),
		}
		binary.LittleEndian.PutUint64(reply.Payload, 0)
		return socket.Send(reply)
	}
	return nil
}
