package types

type VHostUserShared struct {
	UUID [16]byte
}

func (payload *VHostUserShared) ToBinary(buf []byte) (int, error) {
	copy(buf[0:], payload.UUID[:])
	//16
	return 16, nil
}
