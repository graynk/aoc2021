package main

import (
	"bufio"
	"fmt"
	"github.com/graynk/advent_of_code"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type reactor [][][]bool

func reboot(filename string, size int) reactor {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	r := make(reactor, size+1)
	for x := 0; x < size+1; x++ {
		xes := make([][]bool, size+1)
		for y := 0; y < size+1; y++ {
			xes[y] = make([]bool, size+1)
		}
		r[x] = xes
	}

	fmt.Println("it's done!")

	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		instruction := fileScanner.Text()
		split := strings.Split(instruction, " ")
		coordinates := strings.Split(split[1], ",")
		xRange := strings.Split(coordinates[0][len("x="):], "..")
		yRange := strings.Split(coordinates[1][len("y="):], "..")
		zRange := strings.Split(coordinates[2][len("z="):], "..")

		half := size / 2
		minX, _ := strconv.Atoi(xRange[0])
		maxX, _ := strconv.Atoi(xRange[1])
		minY, _ := strconv.Atoi(yRange[0])
		maxY, _ := strconv.Atoi(yRange[1])
		minZ, _ := strconv.Atoi(zRange[0])
		maxZ, _ := strconv.Atoi(zRange[1])
		if advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(minX) > half ||
			advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(minY) > half ||
			advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(minZ) > half ||
			advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(maxX) > half ||
			advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(maxY) > half ||
			advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(maxZ) > half {
			continue
		}
		minX += half
		minY += half
		minZ += half
		maxX += half
		maxY += half
		maxZ += half
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				for z := minZ; z <= maxZ; z++ {
					//fmt.Printf("turning on %d,%d,%d\n", x, y, z)
					if split[0] == "on" {
						r[x][y][z] = true
					} else {
						r[x][y][z] = false
					}
				}
			}
		}
	}
	return r
}

func (r reactor) countOn() int {
	count := 0
	for x := 0; x < len(r); x++ {
		for y := 0; y < len(r[0]); y++ {
			for z := 0; z < len(r[0][0]); z++ {
				if r[x][y][z] {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	r := reboot("./day22/input.txt", 100)
	fmt.Println(r.countOn())
	ioutil.WriteFile("./day22/small.scad", []byte(generateScad("./day22/testscad.txt")), 0644)
	ioutil.WriteFile("./day22/big.scad", []byte(generateScad("./day22/input.txt")), 0644)
}
