// Package types with serialisation
package wires

import (
	"encoding/binary"
	"unsafe"
)

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
	copy(buf[12:], payload.Payload)
	// 4 + 4 + 4 + x
	return 12 + len(payload.Payload), nil
}

type VirtioBlkConfig struct {
	Capacity           uint64
	SizeMax            uint32
	SegMax             uint32
	CylCount           uint16
	Heads              uint8
	Sectors            uint8
	BlockSize          uint32
	TopoPhysExp        uint8
	TopoAlignOff       uint8
	TopoMinIoSz        uint16
	TopoOptIoSz        uint32
	Writeback          uint8
	Unused0            [3]uint8
	NumQueues          uint16
	MaxDiscardSectors  uint32
	MaxDiscardSeg      uint32
	DiscardSectorAlign uint32
	MaxZeroSectors     uint32
	MaxZeroSeg         uint32
	ZeroMayUnmap       uint8
	Unused1            [3]uint8
}

var _ [64]struct{} = [unsafe.Sizeof(VirtioBlkConfig{})]struct{}{}

func (data *VirtioBlkConfig) ToBinary(buf []byte) (int, error) {
	b := castToBytes(data)
	return copy(buf, b), nil
}
