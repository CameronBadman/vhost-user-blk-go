package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"vhost-go/transport"
)

func main() {
	path := "/tmp/vhost-blk.sock"
	socket, err := transport.NewSocket(path)
	if err != nil {
		panic(err)
	}

	SetCleanExit(socket, path)

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

func SetCleanExit(socket *transport.Socket, path string) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		err := socket.Listener.Close()
		if err != nil {
			panic(err)
		}

		err = os.Remove(path)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}()
}
