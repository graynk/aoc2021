package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Image [][]bool

func (im Image) Fold(row, col int, infinity rune) int {
	var builder = strings.Builder{}
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i < 0 || i >= len(im) || j < 0 || j >= len(im[0]) {
				if infinity == '#' {
					builder.WriteRune('1')
				} else {
					builder.WriteRune('0')
				}
				continue
			}
			if im[i][j] {
				builder.WriteRune('1')
			} else {
				builder.WriteRune('0')
			}
		}
	}
	value, err := strconv.ParseInt(builder.String(), 2, 32)
	if err != nil {
		log.Fatalln(err)
	}
	return int(value)
}

func (im Image) Enhance(enhancement string, infinity rune) Image {
	newIm := make(Image, 0, 1)
	for i := -1; i < len(im)+1; i++ {
		row := make([]bool, 0, 1)
		for j := -1; j < len(im[0])+1; j++ {
			value := im.Fold(i, j, infinity)
			enhanced := rune(enhancement[value])
			if enhanced == '.' {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}
		newIm = append(newIm, row)
	}

	return newIm
}

func (im Image) CountLit() int {
	sum := 0
	for _, row := range im {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}
	return sum
}

func (im Image) String() string {
	var builder = strings.Builder{}
	for _, row := range im {
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

func FromFile(filename string) (string, Image) {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Scan()
	enhancement := fileScanner.Text()
	fileScanner.Scan()
	image := make(Image, 0, 1)
	for fileScanner.Scan() {
		row := make([]bool, 0, 1)
		line := fileScanner.Text()

		for _, symbol := range line {
			if symbol == '.' {
				row = append(row, false)
			} else {
				row = append(row, true)
			}
		}

		image = append(image, row)
	}

	return enhancement, image
}
