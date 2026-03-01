package types

type MemRegionDesc struct {
	GuestAdd   uint64
	Size       uint64
	UserAdd    uint64
	MmapOffset uint64
}

type MemRegion struct {
	Padding uint64
	Region  MemRegionDesc
}

type MultiMemRegions struct {
	Num     uint32
	Padding uint32
	Regions [8]MemRegionDesc
}
