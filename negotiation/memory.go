package negotiation

import (
	"fmt"

	"vhost-go/blk"
	"vhost-go/memory"
	"vhost-go/transport"
	"vhost-go/wires"
)

func setMemTable(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	t, err := memory.ParseAndMap(msg.Payload, msg.Fds)
	if err != nil {
		return fmt.Errorf("SET_MEM_TABLE: %w", err)
	}
	dev.MemTable = t
	return socket.SendAck(msg)
}

