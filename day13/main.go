package main

import (
	"bufio"
	"fmt"
	"github.com/graynk/advent_of_code"
	"log"
	"os"
	"strconv"
	"strings"
)

type paperField []paperRow
type paperRow []bool

func newPaperField(x, y int) paperField {
	field := make(paperField, y)

	for row := range field {
		field[row] = make([]bool, x)
	}

	return field
}

func parseInput(filename string) (paperField, []string) {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	input := make([][]int, 0, 735)
	maxX := 0
	maxY := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}
		coordinateStrings := strings.Split(line, ",")
		if len(coordinateStrings) != 2 {
			log.Fatalf("wrong input %s", line)
		}
		x, err := strconv.Atoi(coordinateStrings[0])
		if err != nil {
			log.Fatal(err)
		}
		if x > maxX {
			maxX = x
		}
		y, err := strconv.Atoi(coordinateStrings[1])
		if err != nil {
			log.Fatal(err)
		}
		if y > maxY {
			maxY = y
		}
		input = append(input, []int{x, y})
	}
	field := newPaperField(maxX+1, maxY+1)

	for _, coordinates := range input {
		field[coordinates[1]][coordinates[0]] = true
	}

	instructions := make([]string, 0, 1)

	for fileScanner.Scan() {
		instructions = append(instructions, fileScanner.Text())
	}

	return field, instructions
}

func (pf paperField) countDots() int {
	sum := 0
	for _, row := range pf {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}
	return sum
}

func (pf paperField) foldInstruction(instruction string) paperField {
	axisIndex := strings.Split(instruction, "=")
	if len(axisIndex) != 2 {
		log.Fatalf("wrong instruction %s", instruction)
	}
	index, err := strconv.Atoi(axisIndex[1])
	if err != nil {
		log.Fatal(err)
	}
	var newPf paperField
	if axisIndex[0] == "fold along y" {
		newPf = newPaperField(len(pf[0]), index)
		for rowIndex, row := range pf {
			if rowIndex == index {
				continue
			}
			for colIndex, cell := range row {
				if cell {
					newIndex := index - advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(index-rowIndex)
					newPf[newIndex][colIndex] = true
				}
			}
		}
	} else {
		newPf = newPaperField(index, len(pf))
		for rowIndex, row := range pf {
			for colIndex, cell := range row {
				if colIndex == index {
					continue
				}
				if cell {
					newIndex := index - advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(index-colIndex)
					newPf[rowIndex][newIndex] = true
				}
			}
		}
	}

	return newPf
}

func (pf paperField) String() string {
	builder := strings.Builder{}
	for _, row := range pf {
		for _, cell := range row {
			if cell {
				builder.WriteRune('#')
			} else {
				builder.WriteRune('.')
			}
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func main() {
	input, instructions := parseInput("./day13/input.txt")
	newPf := input.foldInstruction(instructions[0])
	fmt.Println(newPf.countDots())
	for i := 1; i < len(instructions); i++ {
		newPf = newPf.foldInstruction(instructions[i])
	}
	fmt.Println(newPf)
}
