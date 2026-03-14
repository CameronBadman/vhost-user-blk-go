package main

import (
	"fmt"

	"vhost-go/blk"
	"vhost-go/negotiation"
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

	for {
		err = socket.Accept()
		if err != nil {
			panic(err)
		}
		device := blk.NewDevice()
		for {
			n, err := socket.Recv()

			fmt.Printf("%x\n", n)
			if err != nil {
				panic(err)
			}
			err = negotiation.Dispatch(device, socket, n)
			if err != nil {
				panic(err)
			}
		}
	}
}
