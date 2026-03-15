package util

import (
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"vhost-go/blk"
	"vhost-go/transport"
)

func CleanExit(socket *transport.Socket, path string, dev *atomic.Pointer[blk.Device]) {
	if d := dev.Load(); d != nil && d.MemTable != nil {
		d.MemTable.Unmap()
	}

	if err := socket.Listener.Close(); err != nil {
		log.Printf("close Listener: %v", err)
	}

	if err := os.Remove(path); err != nil {
		log.Printf("remove socket: %v", err)
	}
}

func SetCleanExit(socket *transport.Socket, path string, dev *atomic.Pointer[blk.Device]) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		CleanExit(socket, path, dev)
		os.Exit(0)
	}()
}
