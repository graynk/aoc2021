package main

import (
	"reflect"
	"testing"
)

func Test_parseLiteralPacket(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		want           Packet
		wantVersionSum uint
	}{
		{
			"0",
			"./testinput.txt",
			Packet{
				header: PacketHeader{
					version: 6,
					typeID:  4,
				},
				value: 2021,
			},
			6,
		},
		{
			"1",
			"./testinput2.txt",
			Packet{
				header: PacketHeader{
					version: 1,
					typeID:  6,
				},
				subpackets: []Packet{
					{
						header: PacketHeader{version: 6, typeID: 4},
						value:  10,
					},
					{
						header: PacketHeader{version: 2, typeID: 4},
						value:  20,
					},
				},
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.input); !reflect.DeepEqual(got, tt.want) && got.VersionSum() != tt.wantVersionSum {
				t.Errorf("fourBitHexToBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
