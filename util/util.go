package util

import (
	"os"
	"os/signal"
	"syscall"

	"vhost-go/transport"
)

func CleanExit(socket *transport.Socket, path string) error {
	err := socket.Listener.Close()
	if err != nil {
		return err
	}

	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func SetCleanExit(socket *transport.Socket, path string) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		CleanExit(socket, path)
		os.Exit(0)
	}()
}
