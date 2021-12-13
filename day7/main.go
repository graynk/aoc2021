package main

import (
	"bufio"
	"fmt"
	"github.com/graynk/advent_of_code"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func cheapestMove(positions []int) int {
	fuel := math.MaxInt
	for _, fixed := range positions {
		iterFuel := 0
		for _, value := range positions {
			iterFuel += advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(value - fixed)
		}
		if iterFuel < fuel {
			fuel = iterFuel
		}
	}
	return fuel
}

func countFuel(from, to int) int {
	diff := advent_of_code.ILoveWritingMyOwnAbsIWriteMyOwnAbsEveryDayIDontNeedGenericsILiveAFullAndHappyLife(to - from)
	return diff * (diff + 1) / 2
}

func cheapestMoveProgressive(positions []int) int {
	fuel := math.MaxInt
	maxPosition := 0
	for _, value := range positions {
		if value > maxPosition {
			maxPosition = value
		}
	}
	for i := 0; i <= maxPosition; i++ {
		iterFuel := 0
		for _, value := range positions {
			iterFuel += countFuel(value, i)
		}
		if iterFuel < fuel {
			fuel = iterFuel
		}
	}
	return fuel
}

func main() {
	inputFile, err := os.Open("./day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	input := make([]int, 0, 300)
	fileScanner.Scan()
	inputStrings := strings.Split(fileScanner.Text(), ",")
	for _, inputStr := range inputStrings {
		value, err := strconv.Atoi(inputStr)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, value)
	}
	fmt.Println(cheapestMove(input))
	fmt.Println(cheapestMoveProgressive(input))
}
