package wires

import "unsafe"

type InflightDesc struct {
	MmapSize   uint64
	MmapOffset uint64
	NumQueues  uint16
	QueueSize  uint16
	_          [4]byte
}

func (payload *InflightDesc) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

var _ = [1]struct{}{}[unsafe.Sizeof(InflightDesc{})-24]
