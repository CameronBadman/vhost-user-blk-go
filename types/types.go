// Packages types
package types

type VhostUserMsg struct {
	VhostUserRequest string
	Flags            uint32
	Size             uint32
	Payload          []byte
}

type VringState struct {
	Index uint32
	Num   uint32
}

type VringDesc struct {
	Index      uint32
	AvailIndex uint32
}

type VringDescIndices struct {
	Index       uint32
	DescIndices uint32
}

type VringAddDesc struct {
	Index      uint32
	Flags      uint32
	Descriptor uint64
	Used       uint64
	Available  uint64
	Log        uint64
}

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

type IOTLBMessage struct {
	Iova        uint32
	Size        uint64
	UserAddress uint64
	PermFlags   uint8
	Type        uint8
}

type VirtioDevice struct {
	Offset  uint32
	Size    uint32
	Flags   uint32
	Payload []byte
}

type VringAreaDesc struct {
	U64    uint64
	Size   uint64
	Offset uint64
}

type InflightDesc struct {
	MmapSize   uint64
	MmapOffset uint64
	NumQueues  uint16
	QueueSize  uint16
}

type VHostUserShared struct {
	UUID uint16
}
