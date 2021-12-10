package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(numbers []int) int {
	increased := 0
	prev := 0
	for _, depth := range numbers {
		if depth > prev && prev != 0 {
			increased++
		}

		prev = depth
	}

	return increased
}

func part2(numbers []int) int {
	increased := 0
	prev := 0
	for i := 0; i < len(numbers)-2; i++ {
		sum := numbers[i] + numbers[i+1] + numbers[i+2]

		if sum > prev && prev != 0 {
			increased++
		}

		prev = sum
	}

	return increased
}

func main() {
	input, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	numbers := make([]int, 0, 2000)
	for fileScanner.Scan() {
		number, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}
