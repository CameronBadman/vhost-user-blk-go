package transport

import (
	"encoding/binary"
	"fmt"
	"io"

	"vhost-go/types"
	"vhost-go/wires"
)

func (s *Socket) Recv() (*wires.VhostUserMsg, error) {
	_, err := io.ReadFull(s.Conn, s.Buf[:12])
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

	// check for bad payload
	if _, err = io.ReadFull(s.Conn, s.Buf[12:payloadSize]); err != nil {
		return nil, err
	}

	return &wires.VhostUserMsg{
		Request: types.MsgType(request),
		Flags:   types.MsgFlag(flags),
		Size:    size,
		Payload: s.Buf[12:payloadSize],
	}, nil
}
