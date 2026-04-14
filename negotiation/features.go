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

func sendNotImplemented(socket *transport.Socket, msg *wires.VhostUserMsg) error {
	reply := &wires.VhostUserMsg{
		Request: msg.Request,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply | types.MsgFlagErrMsg,
		Size:    0,
	}
	return socket.Send(reply)
}

func setOwner(dev *blk.Device) error {
	dev.Owned = true
	return nil
}
