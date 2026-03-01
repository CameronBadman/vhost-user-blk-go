package types

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

type VringAreaDesc struct {
	U64    uint64
	Size   uint64
	Offset uint64
}
