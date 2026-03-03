package types

import "encoding/binary"

type VhostUserMsg struct {
	VhostUserRequest uint32
	Flags            uint32
	Size             uint32
	Payload          []byte
}

func (payload *VhostUserMsg) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.VhostUserRequest)
	binary.LittleEndian.PutUint32(buf[4:], payload.Flags)
	binary.LittleEndian.PutUint32(buf[8:], payload.Size)
	copy(buf[12:], payload.Payload)
	return 12 + len(payload.Payload), nil
}
