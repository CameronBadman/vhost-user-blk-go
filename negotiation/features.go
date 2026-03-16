package negotiation

import (
	"encoding/binary"
	"fmt"

	"vhost-go/blk"
	"vhost-go/memory"
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

func setMemTable(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	t, err := memory.ParseAndMap(msg.Payload, msg.Fds)
	if err != nil {
		return fmt.Errorf("SET_MEM_TABLE: %w", err)
	}
	dev.MemTable = t
	return socket.SendAck(msg)
}

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

func getConfig(dev *blk.Device, socket *transport.Socket, msg *wires.VhostUserMsg) error {
	offset := binary.LittleEndian.Uint32(msg.Payload[0:4])
	size := binary.LittleEndian.Uint32(msg.Payload[4:8])
	flags := binary.LittleEndian.Uint32(msg.Payload[8:12])

	var raw [64]byte
	dev.Config.ToBinary(raw[:])

	resp := wires.VirtioDevice{
		Offset:  offset,
		Size:    size,
		Flags:   flags,
		Payload: raw[offset : offset+size],
	}

	payload := make([]byte, 12+size)
	resp.ToBinary(payload)

	reply := &wires.VhostUserMsg{
		Request: msg.Request,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply,
		Size:    12 + size,
		Payload: payload,
	}
	return socket.Send(reply)
}
