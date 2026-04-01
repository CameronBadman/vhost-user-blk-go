package blk

import (
	"vhost-go/memory"
	"vhost-go/wires"
)

type Device struct {
	Features          uint64
	ProtocolFreatures uint64
	MemTable          *memory.Table
	NumQueues         uint64
	Config            wires.VirtioBlkConfig
	Owned             bool
	Queue             *Queue
}

func NewDevice() *Device {
	return &Device{
		NumQueues: 1,
		Config: wires.VirtioBlkConfig{
			Capacity: 64 * 1024 * 1024 / 512,
		},
	}
}
