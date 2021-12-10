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

type BoardCell int
type Board [][]BoardCell
type Coordinates struct {
	row    int
	column int
}

func newBoard(size int) Board {
	board := Board{}
	board = make([][]BoardCell, size)
	for i := 0; i < len(board); i++ {
		board[i] = make([]BoardCell, size)
	}
	return board
}

func (b *Board) updateVentsVertical(from, to Coordinates) {
	xMovement := from.row == to.row
	yMovement := from.column == to.column

	fromMin := Coordinates{}
	toMax := Coordinates{}
	if to.row < from.row || to.column < from.column {
		fromMin = to
		toMax = from
	} else {
		fromMin = from
		toMax = to
	}
	if to.column < from.column {
		fromMin.column = to.column
		toMax.column = from.column
	} else {
		fromMin.column = from.column
		toMax.column = to.column
	}

	if xMovement {
		for column := fromMin.column; column <= toMax.column; column++ {
			(*b)[from.row][column]++
		}
		return
	} else if yMovement {
		for row := fromMin.row; row <= toMax.row; row++ {
			(*b)[row][from.column]++
		}
		return
	}
}

func (b *Board) updateVentsFull(from, to Coordinates) {
	xMovement := from.row == to.row
	yMovement := from.column == to.column

	fromMin := Coordinates{}
	toMax := Coordinates{}
	if to.row < from.row || to.column < from.column {
		fromMin = to
		toMax = from
	} else {
		fromMin = from
		toMax = to
	}

	if xMovement {
		for column := fromMin.column; column <= toMax.column; column++ {
			(*b)[from.row][column]++
		}
		return
	} else if yMovement {
		for row := fromMin.row; row <= toMax.row; row++ {
			(*b)[row][from.column]++
		}
		return
	}

	rowDownTo := from.row > to.row
	columnDownTo := from.column > to.column
	row := from.row
	column := from.column
	for i := 0; i <= int(math.Abs(float64(toMax.row-fromMin.row))); i++ {
		(*b)[row][column]++
		if rowDownTo {
			row--
		} else {
			row++
		}
		if columnDownTo {
			column--
		} else {
			column++
		}
	}
}

func (b *Board) String() string {
	builder := strings.Builder{}
	for _, row := range *b {
		for _, cell := range row {
			if cell == 0 {
				builder.WriteString(".")
				continue
			}
			builder.WriteString(strconv.Itoa(int(cell)))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func parseCoordinates(coords string) Coordinates {
	split := strings.Split(coords, ",")
	if len(split) != 2 {
		log.Fatalf("wrong input: %s", coords)
	}
	x, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	// eh
	return Coordinates{
		row:    y,
		column: x,
	}
}

func parseInput(filename string, size int, useDiagonals bool) Board {
	input, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	board := newBoard(size)
	fileScanner := bufio.NewScanner(input)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fromTo := strings.Split(line, " -> ")
		if len(fromTo) != 2 {
			log.Fatalf("wrong input: %s", line)
		}
		from := parseCoordinates(fromTo[0])
		to := parseCoordinates(fromTo[1])
		if !useDiagonals {
			board.updateVentsVertical(from, to)
		} else {
			board.updateVentsFull(from, to)
		}
	}

	return board
}

func (b Board) countDangerousSpots() int {
	count := 0
	for _, row := range b {
		for _, cell := range row {
			if cell >= 2 {
				count++
			}
		}
	}
	return count
}

func part1(board Board) int {
	return board.countDangerousSpots()
}

func part2(board Board) int {
	return board.countDangerousSpots()
}

func main() {
	startTime := time.Now()
	board := parseInput("./day5/input.txt", 1000, false)
	fmt.Println(part1(board))
	log.Println("Took", time.Now().Sub(startTime))
	startTime = time.Now()
	board = parseInput("./day5/input.txt", 1000, true)
	fmt.Println(part2(board))
	log.Println("Took", time.Now().Sub(startTime))
}
