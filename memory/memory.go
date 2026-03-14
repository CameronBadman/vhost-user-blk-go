package memory

type Region struct {
	GuestPhysAddr uint64
	Size          uint64
	UserspaceAddr uint64
	MmapOffset    uint64
	HostAddr      uintptr
}

type Table struct {
	Regions []Region
}
