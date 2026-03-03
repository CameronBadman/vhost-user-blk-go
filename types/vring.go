package types

import "encoding/binary"

type VringState struct {
	Index uint32
	Num   uint32
}

func (payload *VringState) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.Index)
	binary.LittleEndian.PutUint32(buf[4:], payload.Num)
	// 4 + 4
	return 8, nil
}

type VringDesc struct {
	Index      uint32
	AvailIndex uint32
}

func (payload *VringDesc) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.Index)
	binary.LittleEndian.PutUint32(buf[4:], payload.AvailIndex)
	// 4 + 4
	return 8, nil
}

type VringDescIndices struct {
	Index       uint32
	DescIndices uint32
}

func (payload *VringDescIndices) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint32(buf[0:], payload.Index)
	binary.LittleEndian.PutUint32(buf[4:], payload.DescIndices)
	// 4 + 4
	return 8, nil
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
	binary.LittleEndian.PutUint32(buf[0:], payload.Index)
	binary.LittleEndian.PutUint32(buf[4:], payload.Flags)
	binary.LittleEndian.PutUint64(buf[8:], payload.Descriptor)
	binary.LittleEndian.PutUint64(buf[16:], payload.Used)
	binary.LittleEndian.PutUint64(buf[24:], payload.Available)
	binary.LittleEndian.PutUint64(buf[32:], payload.Log)
	// 4 + 4 + 8 + 8 + 8 + 8 
	return 40, nil
}


type VringAreaDesc struct {
	U64    uint64
	Size   uint64
	Offset uint64
}


func (payload *VringAreaDesc) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint64(buf[0:], payload.U64)
	binary.LittleEndian.PutUint64(buf[8:], payload.Size)
	binary.LittleEndian.PutUint64(buf[16:], payload.Offset)
	//  8 + 8 + 8
	return 24, nil
}
