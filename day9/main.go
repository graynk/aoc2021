package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Heatmap [][]int

func (h Heatmap) findRisk() int {
	sum := 0
	for row := range h {
		for col, cell := range h[row] {
			if row > 0 && h[row-1][col] <= cell {
				continue
			}
			if row < len(h)-1 && h[row+1][col] <= cell {
				continue
			}
			if col > 0 && h[row][col-1] <= cell {
				continue
			}
			if col < len(h[row])-1 && h[row][col+1] <= cell {
				continue
			}
			sum += cell + 1
		}
	}
	return sum
}

func (h *Heatmap) basinExplorer(row, col int) int {
	heatmap := *h
	cell := heatmap[row][col]
	if cell == -1 || cell == 9 {
		return 0
	}
	sum := 1
	heatmap[row][col] = -1
	if row > 0 {
		sum += h.basinExplorer(row-1, col)
	}
	if row < len(heatmap)-1 {
		sum += h.basinExplorer(row+1, col)
	}
	if col > 0 {
		sum += h.basinExplorer(row, col-1)
	}
	if col < len(heatmap[row])-1 {
		sum += h.basinExplorer(row, col+1)
	}
	return sum
}

func (h *Heatmap) basinCounter() int {
	heatmap := *h
	basins := make([]int, 0, 3)
	for row := range heatmap {
		for col := range heatmap[row] {
			basinSize := h.basinExplorer(row, col)
			basins = append(basins, basinSize)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func parseInput(filename string) Heatmap {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	input := make([][]int, 0, 100)
	for fileScanner.Scan() {
		cells := strings.Split(fileScanner.Text(), "")
		row := make([]int, 0, 100)
		for _, cell := range cells {
			value, err := strconv.Atoi(cell)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, value)
		}
		input = append(input, row)
	}
	return input
}

func main() {
	input := parseInput("./day9/input.txt")
	fmt.Println(input.findRisk())
	fmt.Println(input.basinCounter())
}
