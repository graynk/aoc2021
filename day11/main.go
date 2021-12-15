package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type octopusField []octopusRow
type octopusRow []octopusCell
type octopusCell struct {
	value   int
	flashed bool
}

func parseInput(filename string) octopusField {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	field := make(octopusField, 0, 10)
	for fileScanner.Scan() {
		row := make(octopusRow, 10)
		field = append(field, row)
		line := fileScanner.Text()
		octopiStrings := strings.Split(line, "")
		if len(octopiStrings) != 10 {
			log.Fatalf("wrong input %s", line)
		}
		for i, octopus := range octopiStrings {
			value, err := strconv.Atoi(octopus)
			if err != nil {
				log.Fatal(err)
			}
			row[i].value = value
		}
	}

	return field
}

func (of octopusField) step() int {
	sum := 0
	for row := range of {
		for col := range of[row] {
			of[row][col].value++
		}
	}
	for row := range of {
		for col := range of[row] {
			of.flash(row, col, false)
		}
	}
	for row := range of {
		for col := range of[row] {
			if of[row][col].flashed {
				sum++
			}
			if of[row][col].value > 9 {
				of[row][col].value = 0
				of[row][col].flashed = false
			}
		}
	}
	return sum
}

func (of octopusField) flash(row, col int, recursive bool) {
	if recursive {
		of[row][col].value++
	}
	if of[row][col].flashed || of[row][col].value <= 9 {
		return
	}
	of[row][col].flashed = true
	if row > 0 {
		of.flash(row-1, col, true)
	}
	if row < len(of)-1 {
		of.flash(row+1, col, true)
	}
	if col > 0 {
		of.flash(row, col-1, true)
	}
	if col < len(of[row])-1 {
		of.flash(row, col+1, true)
	}
	if row > 0 && col > 0 {
		of.flash(row-1, col-1, true)
	}
	if row < len(of)-1 && col < len(of[row])-1 {
		of.flash(row+1, col+1, true)
	}
	if row > 0 && col < len(of[row])-1 {
		of.flash(row-1, col+1, true)
	}
	if row < len(of)-1 && col > 0 {
		of.flash(row+1, col-1, true)
	}
}

func (of octopusField) iWouldStepOneHundredSteps() int {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += of.step()
	}
	return sum
}

func (of octopusField) andIWouldStepOneHundredMore() int {
	for i := 0; ; i++ {
		flashed := of.step()
		if flashed == 100 {
			//fmt.Println(of)
			return i + 1
		}
	}
}

func (of octopusField) String() string {
	builder := strings.Builder{}
	for _, row := range of {
		for _, cell := range row {
			builder.WriteString(strconv.Itoa(cell.value))
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func main() {
	field := parseInput("./day11/input.txt")
	fmt.Println(field.iWouldStepOneHundredSteps())
	fmt.Println(field.andIWouldStepOneHundredMore() + 100)
}
