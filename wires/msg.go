package wires

import (
	"encoding/binary"

	"vhost-go/types"
)

type VhostUserMsg struct {
	Request types.MsgType
	Flags   types.MsgFlag
	Size    uint32
	Payload []byte
	Fds     []int
}

func (payload *VhostUserMsg) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], uint32(payload.Request))
	binary.LittleEndian.PutUint32(buf[4:], uint32(payload.Flags))
	binary.LittleEndian.PutUint32(buf[8:], payload.Size)
	copy(buf[12:], payload.Payload)
	return 12 + len(payload.Payload), nil
}
