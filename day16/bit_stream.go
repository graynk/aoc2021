package main

import (
	"io"
	"log"
	"os"
)

type BitStream struct {
	bits   []Bit
	buffer []byte
	index  int
	file   *os.File
}

func NewBitStream(filename string, capacity int) BitStream {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return BitStream{
		bits:   make([]Bit, 0, capacity),
		buffer: make([]byte, 20),
		file:   inputFile,
	}
}

func (bs *BitStream) Close() {
	bs.file.Close()
}

func (bs *BitStream) writeByte(hex byte) {
	if hex >= '0' && hex <= '9' {
		hex -= '0'
	} else if hex >= 'A' && hex <= 'F' {
		hex -= 'A' - 10
	}

	bs.bits = append(bs.bits, (hex>>3)&1 == 1, (hex>>2)&1 == 1, (hex>>1)&1 == 1, hex&1 == 1)
}

func (bs *BitStream) ReadBits(n int) []Bit {
	read, err := bs.file.Read(bs.buffer[:n])
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	for i := 0; i < read; i++ {
		bs.writeByte(bs.buffer[i])
	}

	slice := bs.bits[bs.index : bs.index+n]
	bs.index += n
	return slice
}

func (bs *BitStream) CurrentIndex() int {
	return bs.index
}
