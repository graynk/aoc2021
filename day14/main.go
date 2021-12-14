package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type polymerInserter map[string]string

func parseInput(filename string) (string, polymerInserter) {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)

	fileScanner.Scan()
	initial := fileScanner.Text()
	fileScanner.Scan()

	instructions := make(polymerInserter)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " -> ")
		if len(split) != 2 {
			log.Fatalf("wrong input %s", line)
		}
		instructions[split[0]] = split[1]
	}

	return initial, instructions
}

func (pi polymerInserter) insertPolymer(template string) string {
	builder := strings.Builder{}

	prev := string(template[0])
	builder.WriteString(prev)
	for i := 0; i < len(template)-1; i++ {
		next := string(template[i+1])
		builder.WriteString(pi[prev+next])
		builder.WriteString(next)
		prev = next
	}

	return builder.String()
}

func commonCounter(template string) int64 {
	common := make(map[rune]int64)
	for _, value := range template {
		common[value]++
	}
	var leastCommon int64 = math.MaxInt64
	var mostCommon int64 = 0

	for _, value := range common {
		if value < leastCommon {
			leastCommon = value
		}
		if value > mostCommon {
			mostCommon = value
		}
	}

	return mostCommon - leastCommon
}

func main() {
	input, pi := parseInput("./day14/input.txt")

	for i := 0; i < 10; i++ {
		input = pi.insertPolymer(input)
	}

	fmt.Println(commonCounter(input))

	for i := 10; i < 40; i++ {
		input = pi.insertPolymer(input)
	}

}
