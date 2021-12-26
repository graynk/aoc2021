package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func sumNumbers(numbers []*SnailFishNumber) *SnailFishNumber {
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = sum.Add(numbers[i])
	}
	return sum
}

func maxSumMagnitude(numbers []*SnailFishNumber) int {
	maxFound := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			sum := numbers[i].Clone().Add(numbers[j].Clone()).Magnitude()
			if sum > maxFound {
				maxFound = sum
			}
		}
	}
	return maxFound
}

func parseInput(filename string) []*SnailFishNumber {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	numbers := make([]*SnailFishNumber, 0, 1)
	for fileScanner.Scan() {
		number := &SnailFishNumber{}
		start := number
		line := fileScanner.Text()
		for _, symbol := range line {
			switch symbol {
			case '[':
				number.left = &SnailFishNumber{parent: number}
				number.right = &SnailFishNumber{parent: number}
				number = number.left
			case ',':
				number = number.parent.right
			case ']':
				number = number.parent
			default:
				parsed, err := strconv.Atoi(string(symbol))
				if err != nil {
					log.Fatalf("wrong input %s", line)
				}
				number.value = parsed
			}
		}
		numbers = append(numbers, start)
	}

	return numbers
}

func main() {
	start := time.Now()
	numbersPart1 := parseInput("./day18/input.txt")
	numbersPart2 := cloneSlice(numbersPart1)
	sum := sumNumbers(numbersPart1)
	magnitude := sum.Magnitude()
	fmt.Printf("%d in %d ms\n", magnitude, time.Now().Sub(start).Milliseconds())
	start = time.Now()
	maxMagnitude := maxSumMagnitude(numbersPart2)
	fmt.Printf("%d in %d ms\n", maxMagnitude, time.Now().Sub(start).Milliseconds())
}
