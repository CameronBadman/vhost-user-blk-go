package types

import "unsafe"

type InflightDesc struct {
	MmapSize   uint64
	MmapOffset uint64
	NumQueues  uint16
	QueueSize  uint16
}

var _ = [1]struct{}{}[unsafe.Sizeof(InflightDesc{})-24]
