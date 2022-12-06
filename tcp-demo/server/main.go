package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const ListenAddr = "localhost:10086"

func main() {
	fmt.Println("Server starting...")

	l, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		log.Printf("failed to start listener: %s", err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("failed to get connection: %s", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	bs, err := io.ReadAll(conn)

	if err != nil {
		log.Printf("failed to get message from %s: %s", conn.RemoteAddr(), err)
		return
	}
	fmt.Printf("get message from %s: %s", conn.RemoteAddr(), string(bs))
}
