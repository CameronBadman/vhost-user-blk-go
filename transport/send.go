package transport

import (
	"encoding/binary"
	"log"

	"vhost-go/types"
	"vhost-go/wires"
)

func (s *Socket) Send(msg *wires.VhostUserMsg) error {
	binary.LittleEndian.PutUint32(s.Buf[0:4], uint32(msg.Request))
	binary.LittleEndian.PutUint32(s.Buf[4:8], uint32(msg.Flags))
	binary.LittleEndian.PutUint32(s.Buf[8:12], msg.Size)
	copy(s.Buf[12:], msg.Payload)

	_, err := s.Conn.Write(s.Buf[:12+msg.Size])

	log.Printf("sending bytes: %x", s.Buf[:12+msg.Size])

	if err != nil {
		return err
	}

	return nil
}

func (s *Socket) SendAck(msg *wires.VhostUserMsg) error {
	return s.Send(&wires.VhostUserMsg{
		Request: msg.Request,
		Flags:   types.MsgFlagVersion | types.MsgFlagReply,
		Size:    msg.Size,
	})
}
