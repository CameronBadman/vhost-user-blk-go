package main

import (
	"fmt"

	"vhost-go/transport"
)

func main() {
	path := "/tmp/vhost.sock"
	socket, err := transport.NewSocket(path)
	if err != nil {
		panic(err)
	}

	err = socket.Accept()
	if err != nil {
		panic(err)
	}

	for {
		n, err := socket.Conn.Read(socket.Buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", socket.Buf[:n])
	}
}
