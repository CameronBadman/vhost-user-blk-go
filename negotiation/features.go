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
	err := socket.Send(reply)
	return err
}

func setFeatures(dev *blk.Device, msg *wires.VhostUserMsg) error {
	dev.Features = binary.LittleEndian.Uint64(msg.Payload)
	return nil
}

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
