package blk

import "vhost-go/memory"

type Device struct {
	Features          uint64
	ProtocolFreatures uint64
	MemTable          *memory.Table
	NumQueues         uint64
}

func NewDevice() *Device {
	return &Device{
		NumQueues: 1,
	}
}
