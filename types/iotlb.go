package types

type IOTLBMessage struct {
	Iova        uint64
	Size        uint64
	UserAddress uint64
	PermFlags   uint8
	Type        uint8
}

func (payload *IOTLBMessage) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}
