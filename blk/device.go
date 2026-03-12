package blk

type Device struct {
	Features          uint64
	ProtocolFreatures uint64
}

func NewDevice() *Device {
	return &Device{}
}
