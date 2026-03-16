package main

import (
	"log"
	"sync/atomic"

	"vhost-go/blk"
	"vhost-go/negotiation"
	"vhost-go/transport"
	"vhost-go/util"
)

func main() {
	path := "/tmp/vhost-blk.sock"
	socket, err := transport.NewSocket(path)
	var currentDevice atomic.Pointer[blk.Device]
	if err != nil {
		panic(err)
	}

	// sets the signal clean
	util.SetCleanExit(socket, path, &currentDevice)
	// sets clean exit on panic
	defer util.CleanExit(socket, path, &currentDevice)
	device := blk.NewDevice()
	currentDevice.Store(device)
	for {
		err = socket.Accept()
		if err != nil {
			panic(err)
		}

		for {
			n, err := socket.Recv()
			if err != nil {
				log.Printf("connection closed: %v", err)
				break
			}
			log.Printf("dispatch: %s", n.Request)
			err = negotiation.Dispatch(device, socket, n)
			if err != nil {
				panic(err)
			}
		}
	}
}
