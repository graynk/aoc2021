package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DeterminantDice int

func (dd *DeterminantDice) Roll() int {
	*dd++
	if *dd == 101 {
		*dd = 1
	}
	return int(*dd)
}

func bustAMove(dice *DeterminantDice, position *int) int {
	for i := 0; i < 3; i++ {
		*position += dice.Roll()
	}
	*position %= 10
	if *position == 0 {
		*position = 10
	}
	return *position
}

func play(first, second int) int {
	var dice DeterminantDice = 0
	firstScore := 0
	secondScore := 0
	iterations := 0

	for currentScore, currentPosition := &secondScore, &second; *currentScore < 1000; iterations += 3 {
		if currentPosition == &first {
			currentPosition = &second
			currentScore = &secondScore
		} else {
			currentPosition = &first
			currentScore = &firstScore
		}
		*currentScore += bustAMove(&dice, currentPosition)
	}

	loser := firstScore
	if secondScore < loser {
		loser = secondScore
	}

	return loser * iterations
}

func parseInput(filename string) (int, int) {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Scan()
	firstStr := fileScanner.Text()
	fileScanner.Scan()
	secondStr := fileScanner.Text()
	fileScanner.Scan()
	first, err := strconv.Atoi(firstStr[len("Player 1 starting position: "):])
	if err != nil {
		log.Fatal(err)
	}
	second, err := strconv.Atoi(secondStr[len("Player 2 starting position: "):])
	if err != nil {
		log.Fatal(err)
	}
	return first, second
}

func main() {
	first, second := parseInput("./day21/input.txt")
	fmt.Println(play(first, second))
}
