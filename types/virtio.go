// Package types with serialisation
package types

type VirtioDevice struct {
	Offset  uint32
	Size    uint32
	Flags   uint32
	Payload []byte
}

func (payload *VirtioDevice) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}
