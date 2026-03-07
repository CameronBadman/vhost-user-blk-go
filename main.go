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

	SetCleanExit(socket)

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

func SetCleanExit(socket *transport.Socket) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		socket.Listener.Close()
		os.Remove("/tmp/vhost-blk.sock")
		os.Exit(0)
	}()
}
