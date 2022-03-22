package main

type BitStream struct {
	bits  []Bit
	index int
}

func NewBitStream(capacity int) BitStream {
	return BitStream{bits: make([]Bit, 0, capacity)}
}

func (bs *BitStream) WriteByte(hex byte) {
	if hex >= '0' && hex <= '9' {
		hex -= '0'
	} else if hex >= 'A' && hex <= 'F' {
		hex -= 'A' - 10
	}

	bs.bits = append(bs.bits, (hex>>3)&1 == 1, (hex>>2)&1 == 1, (hex>>1)&1 == 1, hex&1 == 1)
}

func (bs *BitStream) ReadBits(n int) []Bit {
	slice := bs.bits[bs.index : bs.index+n]
	bs.index += n
	return slice
}

func (bs *BitStream) CurrentIndex() int {
	return bs.index
}
