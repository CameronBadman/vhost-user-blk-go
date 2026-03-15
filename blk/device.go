package blk

import "vhost-go/memory"

type Device struct {
	Features          uint64
	ProtocolFreatures uint64
	MemTable          *memory.Table
}

func NewDevice() *Device {
	return &Device{}
}
