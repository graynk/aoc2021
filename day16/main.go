package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parseInput(filename string) Packet {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileReader := bufio.NewReader(inputFile)

	packetStream := NewBitStream(5300)
	for inputByte, err := fileReader.ReadByte(); err == nil && inputByte != '\n'; inputByte, err = fileReader.ReadByte() {
		packetStream.WriteByte(inputByte)
	}

	return parsePacket(&packetStream)
}

func main() {
	packet := parseInput("./day16/input.txt")
	fmt.Println(packet.VersionSum())
	fmt.Println(packet.Value())
}
