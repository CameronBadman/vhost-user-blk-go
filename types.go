package types

type VhostUserMsg struct {
	VhostUserRequest string
	Flags            uint32
	Size             uint32
	QemuPacked       []byte
}

type VringState struct {
	Index uint32
	Num   uint32
}

func (v VringState) AvailIndex() uint16     { return uint16(v.Num & 0x7FFF) }
func (v VringState) AvailWrapCounter() bool { return v.Num&(1<<15) != 0 }
func (v VringState) UsedIndex() uint16      { return uint16((v.Num >> 16) & 0x7FFF) }
func (v VringState) UsedWrapCounter() bool  { return v.Num&(1<<31) != 0 }

type VringAddr struct {
	Index      uint32
	Flags      uint32
	Descriptor uint64
	Used       uint64
	Available  uint64
	Log        uint64
}

type MemoryRegion struct {
	GuestAddress uint64
	Size         uint64
	UserAddress  uint64
	MmapOffset   uint64
}

type MemoryRegionXen struct {
	MemoryRegion
	XenMmapFlags uint32
	Domid        uint32
}

type SingleMemoryRegion struct {
	Padding uint64
	Region  MemoryRegion
}

type Memory struct {
	NumRegions uint32
	Padding    uint32
	Regions    [8]MemoryRegion
}

type Log struct {
	LogSize   uint64
	LogOffset uint64
}

type IOTLB struct {
	IOVA        uint64
	Size        uint64
	UserAddress uint64
	Permissions uint8
	Type        uint8
}

type Config struct {
	Offset  uint32
	Size    uint32
	Flags   uint32
	Payload []byte
}

type VringArea struct {
	U64    uint64
	Size   uint64
	Offset uint64
}

type Inflight struct {
	MmapSize   uint64
	MmapOffset uint64
	NumQueues  uint16
	QueueSize  uint16
}

type UUID struct {
	Bytes [16]byte
}

type DeviceState struct {
	TransferDirection uint32
	MigrationPhase    uint32
}
