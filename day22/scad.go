package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

type cuboid struct {
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

func fromString(instruction string) (cuboid, string) {
	split := strings.Split(instruction, " ")
	coordinates := strings.Split(split[1], ",")
	xRange := strings.Split(coordinates[0][len("x="):], "..")
	yRange := strings.Split(coordinates[1][len("y="):], "..")
	zRange := strings.Split(coordinates[2][len("z="):], "..")

	c := cuboid{}

	c.minX, _ = strconv.Atoi(xRange[0])
	c.maxX, _ = strconv.Atoi(xRange[1])
	c.minY, _ = strconv.Atoi(yRange[0])
	c.maxY, _ = strconv.Atoi(yRange[1])
	c.minZ, _ = strconv.Atoi(zRange[0])
	c.maxZ, _ = strconv.Atoi(zRange[1])

	return c, split[0]
}

func (c cuboid) String() string {
	return fmt.Sprintf("translate([%d, %d, %d])\ncube([%d, %d, %d]);\n",
		c.minX, c.minY, c.minZ, abs(c.maxX-c.minX), abs(c.maxY-c.minY), abs(c.maxZ-c.minZ))
}

func generateScad(filename string) string {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)
	builder := strings.Builder{}
	prev := "start"
	for fileScanner.Scan() {
		c, instruction := fromString(fileScanner.Text())
		if instruction == "on" {
			if prev != "on" {
				if prev == "off" {
					builder.WriteString("}\n")
				}
				prevResult := builder.String()
				builder.Reset()
				builder.WriteString("union(){\n")
				builder.WriteString(prevResult)
			}
			builder.WriteString(c.String())
			prev = "on"
		} else {
			if prev != "off" {
				builder.WriteString("}\n")
				prevResult := builder.String()
				builder.Reset()
				builder.WriteString("difference(){\n")
				builder.WriteString(prevResult)
			}
			builder.WriteString(c.String())
			prev = "off"
		}
	}
	builder.WriteString("}")
	return builder.String()
}
