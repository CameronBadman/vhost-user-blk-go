package wires

import "encoding/binary"

type VhostUserMsg struct {
	Request uint32
	Flags   uint32
	Size    uint32
	Payload []byte
}

func (payload *VhostUserMsg) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.Request)
	binary.LittleEndian.PutUint32(buf[4:], payload.Flags)
	binary.LittleEndian.PutUint32(buf[8:], payload.Size)
	copy(buf[12:], payload.Payload)
	return 12 + len(payload.Payload), nil
}
