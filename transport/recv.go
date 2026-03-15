package transport

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"syscall"

	"vhost-go/types"
	"vhost-go/wires"
)

func (s *Socket) Recv() (*wires.VhostUserMsg, error) {
	_, oobn, _, _, err := s.Conn.ReadMsgUnix(s.Buf[:12], s.Oob[:])
	if err != nil {
		return nil, err
	}

	request := binary.LittleEndian.Uint32(s.Buf[0:4])
	flags := binary.LittleEndian.Uint32(s.Buf[4:8])
	size := binary.LittleEndian.Uint32(s.Buf[8:12])

	payloadSize := size + 12
	if size > 4096 {
		return nil, fmt.Errorf("msg was too large")
	}
	if size > 0 {
		_, err = io.ReadFull(s.Conn, s.Buf[12:12+size])
		if err != nil {
			return nil, err
		}
	}

	log.Printf("recv: %x", s.Buf[:12+size])

	var fds []int
	if oobn > 0 {
		msgs, _ := syscall.ParseSocketControlMessage(s.Oob[:oobn])
		for _, msg := range msgs {
			gotFds, _ := syscall.ParseUnixRights(&msg)
			fds = append(fds, gotFds...)
		}
	}

	return &wires.VhostUserMsg{
		Request: types.MsgType(request),
		Flags:   types.MsgFlag(flags),
		Size:    size,
		Payload: s.Buf[12:payloadSize],
		Fds:     fds,
	}, nil
}
