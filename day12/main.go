package main

import (
	"bufio"
	"fmt"
	"github.com/graynk/advent_of_code"
	"log"
	"os"
	"strings"
)

func parseInput(filename string) map[string]*Cave {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	caveMap := make(map[string]*Cave)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, "-")
		if len(split) != 2 {
			log.Fatalf("wrong input %s", line)
		}
		for _, c := range split {
			cave, ok := caveMap[c]
			if !ok {
				cave = &Cave{options: make([]*Cave, 0, 1), name: c}
				caveMap[c] = cave
			}
			if c == "start" {
				cave.start = true
			} else if c == "end" {
				cave.end = true
			} else if advent_of_code.IsUpper(c) {
				cave.big = true
			}
		}
		caveMap[split[0]].options = append(caveMap[split[0]].options, caveMap[split[1]])
		if !caveMap[split[0]].start && !caveMap[split[1]].end {
			caveMap[split[1]].options = append(caveMap[split[1]].options, caveMap[split[0]])
		}
	}

	return caveMap
}

func main() {
	caveMap := parseInput("./day12/input.txt")
	paths := make([]Path, 0, 1)
	paths = caveMap["start"].Explore(make(Path, 0, 1), paths)
	fmt.Println(len(paths))
}
