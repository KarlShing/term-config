package main

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"socks5-demo/core"
)

const listenAddr = "localhost:1080"
const Timeout = 5000

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Printf("Failed to start listener on %s", listenAddr)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Failed to get connection: %s", err)
		}
		go handleConn(conn)
	}
}

func handleConn(source net.Conn) {
	defer source.Close()
	remoteAddr, err := core.Handshake(source)
	if err != nil {
		log.Printf("failed to get remote address from source: %s", err)
		return
	}
	log.Printf("handling proxy %s <=> %s", source.RemoteAddr(), remoteAddr)

	remote, err := net.Dial("tcp", remoteAddr.String())
	if err != nil {
		log.Printf("failed to connect to remote[%s]: %s", remoteAddr, err)
		return
	}
	defer remote.Close()

	if err := relay(source, remote); err != nil {
		log.Printf("failed to process proxy: %s", err)
	}
}

func relay(left, right net.Conn) error {
	var err, err1 error
	wg := sync.WaitGroup{}
	wait := Timeout * time.Millisecond
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err1 = io.Copy(right, left)
		right.SetReadDeadline(time.Now().Add(wait))
	}()
	_, err = io.Copy(left, right)
	left.SetReadDeadline(time.Now().Add(wait))
	wg.Wait()
	if err1 != nil && !errors.Is(err1, os.ErrDeadlineExceeded) {
		return err1
	}
	if err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
		return err
	}
	return nil
}
