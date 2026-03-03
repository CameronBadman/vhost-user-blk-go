package transport

import "net"

type Socket struct {
	listener net.Listener
	Conn     net.Conn
	// 12 (header) + 4096 (max payload)
	Buf [4108]byte
}

func New(path string) (*Socket, error) {
	l, err := net.Listen("unix", "/tmp/vhost.sock")
	if err != nil {
		panic(err)
	}

	return &Socket{
		listener: l,
	}, nil
}
