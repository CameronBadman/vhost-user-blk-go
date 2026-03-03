package wires

type VringState struct {
	Index uint32
	Num   uint32
}

func (payload *VringState) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

type VringDesc struct {
	Index      uint32
	AvailIndex uint32
}

func (payload *VringDesc) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

type VringDescIndices struct {
	Index       uint32
	DescIndices uint32
}

func (payload *VringDescIndices) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

type VringAddDesc struct {
	Index      uint32
	Flags      uint32
	Descriptor uint64
	Used       uint64
	Available  uint64
	Log        uint64
}

func (payload *VringAddDesc) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

type VringAreaDesc struct {
	U64    uint64
	Size   uint64
	Offset uint64
}

func (payload *VringAreaDesc) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}
