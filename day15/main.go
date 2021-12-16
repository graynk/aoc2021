package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type riskTraveler struct {
	start  *riskCell
	target *riskCell
}

type riskCell struct {
	options []*riskCell
	visited bool
	risk    int
	minPath int
}

func parseInput(filename string, width, height int) riskTraveler {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)

	field := make([][]riskCell, height)
	for row := 0; row < height; row++ {
		field[row] = make([]riskCell, width)
	}
	for row := 0; row < height; row++ {
		fileScanner.Scan()
		line := fileScanner.Text()
		split := strings.Split(line, "")

		for col, value := range split {
			cellValue, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			field[row][col].risk = cellValue
			field[row][col].minPath = math.MaxInt
			options := make([]*riskCell, 0, 4)
			if row > 0 {
				options = append(options, &field[row-1][col])
			}
			if col > 0 {
				options = append(options, &field[row][col-1])
			}
			if row < height-1 {
				options = append(options, &field[row+1][col])
			}
			if col < width-1 {
				options = append(options, &field[row][col+1])
			}
			field[row][col].options = options
		}
	}

	traveller := riskTraveler{
		start:  &field[0][0],
		target: &field[height-1][width-1],
	}

	return traveller
}

func parseInputPart2(filename string, width, height int) riskTraveler {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)

	field := make([][]riskCell, height*5)
	for row := 0; row < height*5; row++ {
		field[row] = make([]riskCell, width*5)
	}
	for row := 0; row < height*5; row++ {
		fileScanner.Scan()
		line := fileScanner.Text()
		split := strings.Split(line, "")

		for col := 0; col < width*5; col++ {
			if row < height && col < width {
				cellValue, err := strconv.Atoi(split[col])
				if err != nil {
					log.Fatal(err)
				}
				field[row][col].risk = cellValue
			} else {
				rowIndex := row
				colIndex := col
				if rowIndex >= height {
					rowIndex -= height
				} else if colIndex >= width {
					colIndex -= width
				}
				field[row][col].risk = field[rowIndex][colIndex].risk%9 + 1
			}
			field[row][col].minPath = math.MaxInt
			options := make([]*riskCell, 0, 4)
			if row > 0 {
				options = append(options, &field[row-1][col])
			}
			if col > 0 {
				options = append(options, &field[row][col-1])
			}
			if row < height*5-1 {
				options = append(options, &field[row+1][col])
			}
			if col < width*5-1 {
				options = append(options, &field[row][col+1])
			}
			field[row][col].options = options
		}
	}

	traveller := riskTraveler{
		start:  &field[0][0],
		target: &field[height*5-1][width*5-1],
	}

	return traveller
}

func (rt *riskTraveler) traverse() int {
	cell := rt.start
	cell.minPath = 0
	cell.visited = true
	options := []*riskCell{cell.options[0]}
	cell = cell.options[1]

	for len(options) != 0 {
		minPathUpdated := false
		for _, neighbour := range cell.options {
			sumRisk := neighbour.minPath + cell.risk
			if neighbour.visited && sumRisk < cell.minPath {
				cell.minPath = sumRisk
				minPathUpdated = true
			}
		}
		if !cell.visited || minPathUpdated {
			options = append(options, cell.options...)
		}

		cell.visited = true

		cell = options[0]
		options = options[1:]
	}

	return rt.target.minPath
}

func main() {
	start := time.Now()
	rt := parseInput("./day15/input.txt", 100, 100)
	result := rt.traverse()
	fmt.Printf("%v seconds, %d\n", time.Now().Sub(start).Seconds(), result)
	start = time.Now()
	rt2 := parseInputPart2("./day15/input.txt", 100, 100)
	result = rt2.traverse()
	fmt.Printf("%v seconds, %d\n", time.Now().Sub(start).Seconds(), result)
}
