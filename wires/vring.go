package wires

import "encoding/binary"

type VringState struct {
	Index uint32
	Num   uint32
}

func (payload *VringState) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

func ParseVringState(buf []byte) (*VringState, error) {
	return &VringState{
		Index: binary.LittleEndian.Uint32(buf[0:4]),
		Num:   binary.LittleEndian.Uint32(buf[5:8]),
	}, nil
}

type VringDesc struct {
	Index      uint32
	AvailIndex uint32
}

func (payload *VringDesc) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

func ParseVringDesc(buf []byte) (*VringDesc, error) {
	return &VringDesc{
		Index:      binary.LittleEndian.Uint32(buf[0:4]),
		AvailIndex: binary.LittleEndian.Uint32(buf[5:8]),
	}, nil
}

type VringDescIndices struct {
	Index       uint32
	DescIndices uint32
}

func (payload *VringDescIndices) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}

func ParseVringDescIndices(buf []byte) (*VringDescIndices, error) {
	return &VringDescIndices{
		Index:       binary.LittleEndian.Uint32(buf[0:4]),
		DescIndices: binary.LittleEndian.Uint32(buf[5:8]),
	}, nil
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

func ParseVringAddDesc(buf []byte) (*VringAddDesc, error) {
	return &VringAddDesc{
		Index:      binary.LittleEndian.Uint32(buf[0:4]),
		Flags:      binary.LittleEndian.Uint32(buf[5:8]),
		Descriptor: binary.LittleEndian.Uint64(buf[9:17]),
		Used:       binary.LittleEndian.Uint64(buf[17:25]),
		Available:  binary.LittleEndian.Uint64(buf[26:34]),
		Log:        binary.LittleEndian.Uint64(buf[34:42]),
	}, nil
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
