package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type polymerInserter struct {
	rules   map[string]string
	counter map[string]int64
}

func newPolymerInserter() polymerInserter {
	return polymerInserter{
		rules:   make(map[string]string),
		counter: make(map[string]int64),
	}
}

func parseInput(filename string) polymerInserter {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)

	fileScanner.Scan()
	template := fileScanner.Text()
	fileScanner.Scan()

	pi := newPolymerInserter()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " -> ")
		if len(split) != 2 {
			log.Fatalf("wrong input %s", line)
		}
		pi.rules[split[0]] = string(split[1][0])
	}

	for i := 0; i < len(template)-1; i++ {
		a := string(template[i])
		b := string(template[i+1])
		pi.counter[a+b]++
	}

	return pi
}

func (pi *polymerInserter) insertPolymers() {
	newCounter := make(map[string]int64)

	for key, count := range pi.counter {
		insert := pi.rules[key]
		newCounter[string(key[0])+insert] += count
		newCounter[insert+string(key[1])] += count
	}

	pi.counter = newCounter
}

func (pi *polymerInserter) commonCounter() int64 {
	var leastCommon int64 = math.MaxInt64
	var mostCommon int64 = 0

	common := make(map[rune]int64)
	for key, count := range pi.counter {
		common[rune(key[1])] += count
	}

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
	pi := parseInput("./day14/input.txt")

	for i := 0; i < 10; i++ {
		pi.insertPolymers()
	}

	fmt.Println(pi.commonCounter())

	for i := 10; i < 40; i++ {
		pi.insertPolymers()
	}

	fmt.Println(pi.commonCounter())

}
