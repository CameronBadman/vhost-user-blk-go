package types

type VhostUserMsg struct {
	VhostUserRequest uint32
	Flags            uint32
	Size             uint32
	Payload          []byte
}
