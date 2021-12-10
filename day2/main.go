package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(commands []string) int {
	horizontal := 0
	depth := 0

	for _, command := range commands {
		parts := strings.Split(command, " ")
		if len(parts) != 2 {
			log.Fatalf("wrong command size: %d, %s", len(parts), command)
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		default:
			log.Fatalf("weird command: %s", parts[0])
		}
	}

	return depth * horizontal
}

func part2(commands []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, command := range commands {
		parts := strings.Split(command, " ")
		if len(parts) != 2 {
			log.Fatalf("wrong command size: %d, %s", len(parts), command)
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			horizontal += value
			depth += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		default:
			log.Fatalf("weird command: %s", parts[0])
		}
	}

	return depth * horizontal
}

func main() {
	input, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	commands := make([]string, 0, 1000)
	for fileScanner.Scan() {
		commands = append(commands, fileScanner.Text())
	}

	fmt.Println(part1(commands))
	fmt.Println(part2(commands))
}
