package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"golang.org/x/sys/unix"
)

var (
	once   sync.Once
	connFD int = -1
)

func safeClose() {
	once.Do(func() {
		if connFD != -1 {
			unix.Shutdown(connFD, unix.SHUT_RDWR)
			unix.Close(connFD)
			connFD = -1
		}
	})
}

func StartSignalHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		safeClose()
		os.Exit(0)
	}()
}

func ConnectSecureDbus(path string, expectedUID int) (int, error) {
	fd, err := unix.Socket(unix.AF_UNIX, unix.SOCK_STREAM, 0)
	if err != nil { return 0, err }
	if err := unix.Connect(fd, &unix.SockaddrUnix{Name: path}); err != nil {
		return 0, err
	}
	ucred, err := unix.GetsockoptUcred(fd, unix.SOL_SOCKET, unix.SO_PEERCRED)
	if err != nil { return 0, err }
	if int(ucred.Uid) != expectedUID {
		return 0, fmt.Errorf("security breach: unauthorized daemon UID")
	}
	connFD = fd
	return fd, nil
}

func main() {
    StartSignalHandler()
    fmt.Println("gocore-ipc-resolver v1.0 active")
    select {}
}
