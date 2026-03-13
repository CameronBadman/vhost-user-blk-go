package negotiation

import (
	"encoding/binary"

	"vhost-go/blk"
	"vhost-go/transport"
	"vhost-go/types"
	"vhost-go/wires"
)

func getFeatures(socket *transport.Socket) error {
	reply := &wires.VhostUserMsg{
		Request: types.MsgGetFeatures,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply,
		Size:    8,
		Payload: make([]byte, 8),
	}
	binary.LittleEndian.PutUint64(reply.Payload, blk.SupportedFeatures())
	return socket.Send(reply)
}
