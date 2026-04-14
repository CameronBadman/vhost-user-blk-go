package negotiation

import (
	"encoding/binary"

	"vhost-go/blk"
	"vhost-go/transport"
	"vhost-go/types"
	"vhost-go/wires"
)

func getProtocolFeatures(socket *transport.Socket) error {
	reply := &wires.VhostUserMsg{
		Request: types.MsgGetProtocolFeatures,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply,
		Size:    8,
		Payload: make([]byte, 8),
	}
	binary.LittleEndian.PutUint64(reply.Payload, blk.SupportedProtocolFeatures())
	return socket.Send(reply)
}

func setProtocolFeatures(dev *blk.Device, msg *wires.VhostUserMsg) error {
	dev.ProtocolFreatures = binary.LittleEndian.Uint64(msg.Payload)
	return nil
}
