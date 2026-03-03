package types

import "encoding/binary"

type MemRegionDesc struct {
	GuestAdd   uint64
	Size       uint64
	UserAdd    uint64
	MmapOffset uint64
}

func (payload *MemRegionDesc) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint64(buf[0:], payload.GuestAdd)
	binary.LittleEndian.PutUint64(buf[8:], payload.Size)
	binary.LittleEndian.PutUint64(buf[16:], payload.UserAdd)
	binary.LittleEndian.PutUint64(buf[24:], payload.MmapOffset)
	// 8 + 8 + 8 + 8 
	return 32, nil
}

type MemRegion struct {
	Padding uint64
	Region  MemRegionDesc
}

func (payload *MemRegion) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint64(buf[0:], payload.Padding)
	copy(buf[8:], payload.Region.ToBinary(buf[8:]))
	// 8 + 32
	return 40, nil
}

type MultiMemRegions struct {
	Num     uint32
	Padding uint32
	Regions [8]MemRegionDesc
}

func (payload *VirtioDevice) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.Offset)
	binary.LittleEndian.PutUint32(buf[4:], payload.Size)
	binary.LittleEndian.PutUint32(buf[8:], payload.Flags)
	copy(buf[16:], payload.Payload)
	//4 + 4 + 4 + x
	return 12 + len(payload.Payload), nil
}
