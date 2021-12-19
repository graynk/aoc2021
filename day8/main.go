package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type SegmentedDisplay [][]Patterns
type Patterns []string

type DisplayDigitMapping []DisplayDigit

func parseInput(filename string) SegmentedDisplay {
	data := make(SegmentedDisplay, 0, 200)
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)

	for fileScanner.Scan() {
		row := fileScanner.Text()
		rowSplit := strings.Split(row, " | ")
		if len(rowSplit) != 2 {
			log.Fatalf("wrong input %s\n", row)
		}
		input := strings.Split(rowSplit[0], " ")
		output := strings.Split(rowSplit[1], " ")
		data = append(data, []Patterns{input, output})
	}

	return data
}

func (sd SegmentedDisplay) countEasyNumbers() int {
	sum := 0
	for _, row := range sd {
		for _, output := range row[1] {
			switch len(output) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	return sum
}

func (sd SegmentedDisplay) countSum() int {
	sum := 0
	for _, row := range sd {
		mapping := deductMapping(row[0])
		number := 0
		for i, digit := range row[1] {
			rank := int(math.Pow(10, float64(3-i)))
			number += mapping.MapPattern(digit) * rank
		}
		sum += number
	}
	return sum
}

func main() {
	data := parseInput("./day8/input.txt")
	fmt.Println(data.countEasyNumbers())
	fmt.Println(data.countSum())
}
