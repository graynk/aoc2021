package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type BoardCell struct {
	value  int
	marked bool
}
type Coordinates struct {
	row    int
	column int
}
type Board struct {
	cells   [][]BoardCell
	mapping map[int]Coordinates
}

func newBoard() Board {
	board := Board{}
	board.cells = make([][]BoardCell, 5)
	for i := 0; i < len(board.cells); i++ {
		board.cells[i] = make([]BoardCell, 5)
	}
	board.mapping = make(map[int]Coordinates)
	return board
}

func (b Board) rowWins(row int) bool {
	for column := 0; column < len(b.cells[0]); column++ {
		if !b.cells[row][column].marked {
			return false
		}
	}

	return true
}

func (b Board) columnWins(column int) bool {
	for row := range b.cells {
		if !b.cells[row][column].marked {
			return false
		}
	}

	return true
}

func (b Board) unmarkedSum() int {
	sum := 0
	for row := range b.cells {
		for _, cell := range b.cells[row] {
			if cell.marked {
				continue
			}
			sum += cell.value
		}
	}

	return sum
}

func parseInput(filename string) ([]int, []Board) {
	input, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fileScanner.Scan()
	numbersLine := strings.Split(fileScanner.Text(), ",")
	numbers := make([]int, 0, len(numbersLine))
	for _, numberStr := range numbersLine {
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	boards := make([]Board, 0, 100)
	for fileScanner.Scan() {
		board := newBoard()
		for row := 0; row < 5; row++ {
			if !fileScanner.Scan() {
				log.Fatal("input ended too early")
			}
			values := strings.ReplaceAll(strings.TrimSpace(fileScanner.Text()), "  ", " ")
			rowColumns := strings.Split(values, " ")
			for column, columnValue := range rowColumns {
				value, err := strconv.Atoi(columnValue)
				if err != nil {
					log.Fatal(err)
				}
				board.mapping[value] = Coordinates{row, column}
				board.cells[row][column] = BoardCell{
					value:  value,
					marked: false,
				}
			}
		}
		boards = append(boards, board)

	}

	return numbers, boards
}

func part1(numbers []int, boards []Board) int {
	for _, number := range numbers {
		for _, board := range boards {
			if coordinates, ok := board.mapping[number]; ok {
				board.cells[coordinates.row][coordinates.column].marked = true
				if board.rowWins(coordinates.row) || board.columnWins(coordinates.column) {
					return number * board.unmarkedSum()
				}
			}
		}
	}
	// should not happen
	return 0
}

func part2(numbers []int, boards []Board) int {
	won := make(map[int]bool)
	for _, number := range numbers {
		for boardIndex, board := range boards {
			if coordinates, ok := board.mapping[number]; ok {
				x := coordinates.row
				y := coordinates.column
				board.cells[x][y].marked = true
				if board.rowWins(x) || board.columnWins(y) {
					if len(won) == len(boards)-1 && !won[boardIndex] {
						return number * board.unmarkedSum()
					}
					if _, ok := won[boardIndex]; !ok {
						won[boardIndex] = true
					}
				}
			}
		}
	}
	// still should not happen
	return 0
}

func main() {
	start := time.Now()
	numbers, boards := parseInput("./day4/input.txt")
	res1 := part1(numbers, boards)
	res2 := part2(numbers, boards)
	fmt.Println(time.Now().Sub(start))
	fmt.Println(res1)
	fmt.Println(res2)
}
