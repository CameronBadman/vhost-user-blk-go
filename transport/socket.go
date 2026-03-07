package transport

import "net"

type Socket struct {
	Listener net.Listener
	Conn     net.Conn
	Buf      [4108]byte
}

func NewSocket(path string) (*Socket, error) {
	l, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}

	return &Socket{
		Listener: l,
	}, nil
}

func (s *Socket) Accept() error {
	conn, err := s.Listener.Accept()
	if err != nil {
		return err
	}
	s.Conn = conn
	return nil
}
