package main

type PacketType int

const (
	Sum PacketType = iota
	Product
	Min
	Max
	Literal
	Greater
	Less
	Equal
)

type PacketHeader struct {
	version uint
	typeID  PacketType
}

type Packet struct {
	header     PacketHeader
	value      uint
	subpackets []Packet
}

type Bit bool

func bitsToNumber(bits []Bit) uint {
	length := len(bits) - 1
	result := uint(0)
	for i, bit := range bits {
		if !bit {
			continue
		}
		result |= 1 << (length - i)
	}
	return result
}

func parseHeader(version, typeID []Bit) PacketHeader {
	header := PacketHeader{
		version: bitsToNumber(version),
		typeID:  PacketType(bitsToNumber(typeID)),
	}

	return header
}

func parseLiteral(stream *BitStream) uint {
	value := uint(0)
	for stillReading := Bit(true); stillReading; {
		portion := stream.ReadBits(5)
		stillReading = portion[0]
		value = value << 4
		value |= bitsToNumber(portion[1:])
	}

	return value
}

func parseSubpackets(stream *BitStream) []Packet {
	lengthTypeID := stream.ReadBits(1)[0]
	var subpackets []Packet
	if lengthTypeID {
		packetCount := int(bitsToNumber(stream.ReadBits(11)))
		subpackets = make([]Packet, 0, packetCount)
		for i := 0; i < packetCount; i++ {
			subpackets = append(subpackets, parsePacket(stream))
		}
	} else {
		bitsToRead := int(bitsToNumber(stream.ReadBits(15)))
		subpackets = make([]Packet, 0, 1)

		start := stream.CurrentIndex()

		for currentIndex := start; currentIndex < start+bitsToRead; currentIndex = stream.CurrentIndex() {
			subpackets = append(subpackets, parsePacket(stream))
		}
	}

	return subpackets
}

func parsePacket(stream *BitStream) Packet {
	packet := Packet{}

	packet.header = parseHeader(stream.ReadBits(3), stream.ReadBits(3))

	switch packet.header.typeID {
	case Literal:
		packet.value = parseLiteral(stream)
	default:
		packet.subpackets = parseSubpackets(stream)
	}

	return packet
}

func (p Packet) VersionSum() uint {
	sum := p.header.version

	for _, packet := range p.subpackets {
		sum += packet.VersionSum()
	}

	return sum
}

func (p Packet) sum() uint {
	var value uint

	for _, packet := range p.subpackets {
		value += packet.Value()
	}

	return value
}

func (p Packet) product() uint {
	var value uint = 1

	for _, packet := range p.subpackets {
		value *= packet.Value()
	}

	return value
}

func (p Packet) min() uint {
	value := p.subpackets[0].Value()

	for _, packet := range p.subpackets[1:] {
		currentValue := packet.Value()
		if currentValue < value {
			value = currentValue
		}
	}

	return value
}

func (p Packet) max() uint {
	value := p.subpackets[0].Value()

	for _, packet := range p.subpackets[1:] {
		currentValue := packet.Value()
		if currentValue > value {
			value = currentValue
		}
	}

	return value
}

func (p Packet) greater() uint {
	if p.subpackets[0].Value() > p.subpackets[1].Value() {
		return 1
	}
	return 0
}

func (p Packet) less() uint {
	if p.subpackets[0].Value() < p.subpackets[1].Value() {
		return 1
	}
	return 0
}

func (p Packet) equal() uint {
	if p.subpackets[0].Value() == p.subpackets[1].Value() {
		return 1
	}
	return 0
}

func (p Packet) Value() uint {
	switch p.header.typeID {
	case Sum:
		p.value = p.sum()
	case Product:
		p.value = p.product()
	case Min:
		p.value = p.min()
	case Max:
		p.value = p.max()
	case Greater:
		p.value = p.greater()
	case Less:
		p.value = p.less()
	case Equal:
		p.value = p.equal()
	}

	return p.value
}
