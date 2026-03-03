package wires

type MemRegionDesc struct {
	GuestAdd   uint64
	Size       uint64
	UserAdd    uint64
	MmapOffset uint64
}

func (payload *MemRegionDesc) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

type MemRegion struct {
	Padding uint64
	Region  MemRegionDesc
}

func (payload *MemRegion) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

type MultiMemRegions struct {
	Num     uint32
	Padding uint32
	Regions [8]MemRegionDesc
}

func (payload *MultiMemRegions) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}
