package types

import "encoding/binary"

type IOTLBMessage struct {
	Iova        uint64
	Size        uint64
	UserAddress uint64
	PermFlags   uint8
	Type        uint8
}

func (payload *IOTLBMessage) ToBinary(buf []byte) (int, error) {
	binary.LittleEndian.PutUint64(buf[0:], payload.Iova)
	binary.LittleEndian.PutUint64(buf[8:], payload.Size)
	binary.LittleEndian.PutUint64(buf[16:], payload.UserAddress)
	binary.LittleEndian.PutUint8(buf[18:], payload.PermFlags)
	binary.LittleEndian.PutUint8(buf[20:], payload.Type)
	// 8 + 8 + 8 + 8 
	return 20, nil
}
