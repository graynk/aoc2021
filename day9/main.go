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

func findRisk(heatmap [][]int) int {
	sum := 0
	for row := range heatmap {
		for col, cell := range heatmap[row] {
			if row > 0 && heatmap[row-1][col] <= cell {
				continue
			}
			if row < len(heatmap)-1 && heatmap[row+1][col] <= cell {
				continue
			}
			if col > 0 && heatmap[row][col-1] <= cell {
				continue
			}
			if col < len(heatmap[row])-1 && heatmap[row][col+1] <= cell {
				continue
			}
			sum += cell + 1
		}
	}
	return sum
}

func basinExplorer(heatmap [][]int, row, col int) int {
	cell := heatmap[row][col]
	if cell == -1 || cell == 9 {
		return 0
	}
	sum := 1
	heatmap[row][col] = -1
	if row > 0 {
		sum += basinExplorer(heatmap, row-1, col)
	}
	if row < len(heatmap)-1 {
		sum += basinExplorer(heatmap, row+1, col)
	}
	if col > 0 {
		sum += basinExplorer(heatmap, row, col-1)
	}
	if col < len(heatmap[row])-1 {
		sum += basinExplorer(heatmap, row, col+1)
	}
	return sum
}

func basinCounter(heatmap [][]int) int {
	basins := make([]int, 0, 3)
	for row := range heatmap {
		for col := range heatmap[row] {
			basinSize := basinExplorer(heatmap, row, col)
			basins = append(basins, basinSize)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func parseInput(filename string) [][]int {
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
	fmt.Println(findRisk(input))
	fmt.Println(basinCounter(input))
}
