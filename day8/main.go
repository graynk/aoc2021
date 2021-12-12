package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type SegmentedDisplay [][][]string

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
		data = append(data, [][]string{input, output})
	}

	return data
}

func (s SegmentedDisplay) countEasyNumbers() int {
	sum := 0
	for _, row := range s {
		for _, output := range row[1] {
			switch len(output) {
			case 2, 4, 3, 7:
				sum++
			}
		}
	}
	return sum
}

func main() {
	data := parseInput("./day8/input.txt")
	fmt.Println(data.countEasyNumbers())
}
