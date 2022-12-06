package main

import (
	"fmt"
	"log"
	"net"
)

const DailAddr = "localhost:10086"

func main() {
	fmt.Println("Client starting...")

	conn, err := net.Dial("tcp", DailAddr)
	if err != nil {
		log.Printf("failed to dail [%s]: %s", DailAddr, err)
		return
	}
	defer conn.Close()
	_, err = fmt.Fprintf(conn, "Hello TCP\n")
	if err != nil {
		log.Panicf("failed to send msg to %s: %s", DailAddr, err)
		return
	}
}
