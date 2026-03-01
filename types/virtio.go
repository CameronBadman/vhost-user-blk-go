package types

type VirtioDevice struct {
	Offset  uint32
	Size    uint32
	Flags   uint32
	Payload []byte
}
