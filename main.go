package main

import (
	"fmt"

	"vhost-go/transport"
	"vhost-go/util"
)

func main() {
	path := "/tmp/vhost-blk.sock"
	socket, err := transport.NewSocket(path)
	if err != nil {
		panic(err)
	}

	// sets the signal clean
	util.SetCleanExit(socket, path)
	// sets clean exit on panic
	defer util.CleanExit(socket, path)

	err = socket.Accept()
	if err != nil {
		panic(err)
	}

	for {
		n, err := socket.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%x\n", n)
	}
}
