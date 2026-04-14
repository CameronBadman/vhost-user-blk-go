package negotiation

import (
	"encoding/binary"

	"vhost-go/blk"
	"vhost-go/transport"
	"vhost-go/types"
	"vhost-go/wires"
)

func getQueueNum(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	reply := &wires.VhostUserMsg{
		Request: msg.Request,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply,
		Size:    8,
		Payload: make([]byte, 8),
	}
	binary.LittleEndian.PutUint64(reply.Payload, 1)
	return socket.Send(reply)
}
