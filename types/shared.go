package types

type VHostUserShared struct {
	UUID [16]byte
}

func (payload *VHostUserShared) ToBinary(buf []byte) (int, error) {
	b := castToBytes(payload)
	return copy(buf, b), nil
}
