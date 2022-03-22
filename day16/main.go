package main

import (
	"fmt"
)

func parseInput(filename string) Packet {
	packetStream := NewBitStream(filename, 5300)

	defer packetStream.Close()

	return parsePacket(&packetStream)
}

func main() {
	packet := parseInput("./day16/input.txt")
	fmt.Println(packet.VersionSum())
	fmt.Println(packet.Value())
}
