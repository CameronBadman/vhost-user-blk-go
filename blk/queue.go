package blk

type Queue struct {
	Index uint32
	Num   uint32

	DescAddr  uint64
	AvailAddr uint64
	UsedAddr  uint64
	LogAddr   uint64

	LastAvailIdx uint16
	LastUsedIdx  uint16

	Enabled bool

	KickFD int
	CallFD int
	ErrFD  int
}
