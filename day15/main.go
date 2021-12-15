package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type riskTraveler struct {
	start  *riskCell
	target *riskCell
}

type riskCell struct {
	options []*riskCell
	chosen  bool
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

func (rt *riskTraveler) traverse() int {
	cell := rt.start
	cell.minPath = 0
	options := []*riskCell{cell.options[0]}
	cell = cell.options[1]

	for len(options) != 0 {
		if cell.minPath == math.MaxInt {
			//sort.Slice(cell.options, func(i, j int) bool {
			//	return cell.options[i].minPath < cell.options[j].minPath
			//})

			options = append(options, cell.options...)
		}

		for _, neighbour := range cell.options {
			if neighbour.minPath != math.MaxInt && neighbour.minPath+cell.risk < cell.minPath {
				cell.minPath = neighbour.minPath + cell.risk
			}
		}
		cell = options[0]
		options = options[1:]
	}

	doubletake := 0
	for cell := rt.target; cell != rt.start; {
		sort.Slice(cell.options, func(i, j int) bool {
			return cell.options[i].minPath < cell.options[j].minPath
		})
		cell.chosen = true
		cell = cell.options[0]
		doubletake += cell.risk
	}

	return doubletake
}

func main() {
	rt := parseInput("./day15/input.txt", 100, 100)
	fmt.Println(rt.traverse())
}
