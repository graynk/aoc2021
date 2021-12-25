package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CucumberField [][]rune

func parseInput(filename string) CucumberField {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	field := make(CucumberField, 0, 1)
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		row := []rune(fileScanner.Text())
		field = append(field, row)
	}
	return field
}

func (cf CucumberField) GetNextColumn(column int) int {
	if column == len(cf[0])-1 {
		return 0
	} else {
		return column + 1
	}
}

func (cf CucumberField) GetNextRow(row int) int {
	if row == len(cf)-1 {
		return 0
	} else {
		return row + 1
	}
}

func (cf *CucumberField) Step() int {
	moved := 0
	field := *cf
	for row := 0; row < len(field); row++ {
		first := field[row][0]
		for column := 0; column < len(field[row]); column++ {
			if field[row][column] != '>' {
				continue
			}
			nextColumn := field.GetNextColumn(column)
			if field[row][nextColumn] == '.' {
				if nextColumn == 0 && first != '.' {
					continue
				}
				field[row][column] = '.'
				field[row][nextColumn] = '>'
				moved++
				column++
			}
		}
	}
	for column := 0; column < len(field[0]); column++ {
		first := field[0][column]
		for row := 0; row < len(field); row++ {
			if field[row][column] != 'v' {
				continue
			}
			nextRow := field.GetNextRow(row)
			if field[nextRow][column] == '.' {
				if nextRow == 0 && first != '.' {
					continue
				}
				field[row][column] = '.'
				field[nextRow][column] = 'v'
				moved++
				row++
			}
		}
	}
	//fmt.Println(cf.String())
	return moved
}

func (cf CucumberField) String() string {
	builder := strings.Builder{}
	for row := 0; row < len(cf); row++ {
		for column := 0; column < len(cf[row]); column++ {
			builder.WriteRune(cf[row][column])
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

func (cf *CucumberField) StepTillWeStop() int {
	i := 1
	for moved := cf.Step(); moved != 0; moved = cf.Step() {
		i++
	}
	return i
}

func main() {
	field := parseInput("./day25/input.txt")
	fmt.Println(field.StepTillWeStop())
}
