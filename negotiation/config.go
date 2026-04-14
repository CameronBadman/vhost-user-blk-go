package negotiation

import (
	"fmt"

	"vhost-go/blk"
	"vhost-go/transport"
	"vhost-go/types"

	"vhost-go/wires"
)

func getConfig(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	req := wires.ParseVirtioDevice(msg.Payload)
	raw := dev.Config.AsBytes()
	if uint32(len(raw)) < req.Offset+req.Size {
		return fmt.Errorf("config request out of bounds: offset=%d size=%d", req.Offset, req.Size)
	}
	reply := &wires.VhostUserMsg{
		Request: msg.Request,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply,
		Size:    req.Size,
		Payload: raw[req.Offset : req.Offset+req.Size],
	}
	return socket.Send(reply)
}
