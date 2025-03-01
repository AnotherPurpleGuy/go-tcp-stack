package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func main() {
	// Open a raw socket for IPv4 TCP packets.
	// Note: This requires root privileges or NET_RAW capability.
	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		log.Fatalf("Error opening raw socket: %v", err)
	}
	defer conn.Close()

	buffer := make([]byte, 65535)
	fmt.Println("Listening for TCP packets on raw socket...")

	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Printf("Error reading from socket: %v", err)
			continue
		}
		fmt.Printf("Received %d bytes from %v:\n%s\n", n, addr, hex.Dump(buffer[:n]))
	}
}
