package transport

import (
	"net"
)

type Socket struct {
	Listener *net.UnixListener
	Conn     *net.UnixConn
	Buf      [4108]byte
	// should be enough space for 8 file descriptors
	Oob [128]byte
}

func NewSocket(path string) (*Socket, error) {
	addr := &net.UnixAddr{Name: path, Net: "unix"}
	l, err := net.ListenUnix("unix", addr)
	if err != nil {
		return nil, err
	}

	return &Socket{
		Listener: l,
	}, nil
}

func (s *Socket) Accept() error {
	conn, err := s.Listener.AcceptUnix()
	if err != nil {
		return err
	}
	s.Conn = conn
	return nil
}
