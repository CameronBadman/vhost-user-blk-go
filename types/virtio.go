// Package types with serialisation
package types

import "encoding/binary"

type VirtioDevice struct {
	Offset  uint32
	Size    uint32
	Flags   uint32
	Payload []byte
}

func (payload *VirtioDevice) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.Offset)
	binary.LittleEndian.PutUint32(buf[4:], payload.Size)
	binary.LittleEndian.PutUint32(buf[8:], payload.Flags)
	copy(buf[16:], payload.Payload)
	// 4 + 4 + 4 + x
	return 12 + len(payload.Payload), nil
}
