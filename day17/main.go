package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) (minX, maxX, minY, maxY int) {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)

	fileScanner.Scan()
	line := fileScanner.Text()
	arguments := strings.Split(line, ",")
	if len(arguments) != 2 {
		log.Fatalf("wrong input %s", line)
	}
	x := strings.Split(arguments[0], "..")
	if len(x) != 2 {
		log.Fatalf("wrong input %s", arguments[0])
	}
	y := strings.Split(arguments[1], "..")
	if len(y) != 2 {
		log.Fatalf("wrong input %s", arguments[1])
	}
	minX, err = strconv.Atoi(x[0][len("target area: x="):])
	if err != nil {
		log.Fatalf("wrong input %s", x[0])
	}
	maxX, err = strconv.Atoi(x[1])
	if err != nil {
		log.Fatalf("wrong input %s", x[1])
	}
	minY, err = strconv.Atoi(y[0][len(" y="):])
	if err != nil {
		log.Fatalf("wrong input %s", y[0])
	}
	maxY, err = strconv.Atoi(y[1])
	if err != nil {
		log.Fatalf("wrong input %s", y[1])
	}
	return minX, maxX, minY, maxY
}

func findMaxY(minX, maxX, minY, maxY int) int {
	totalMaxReachedY := 0

	for startXVelocity := maxX; startXVelocity > 0; startXVelocity-- {
		for startYVelocity := 0; startYVelocity < -1*minY; startYVelocity++ {
			xVelocity := startXVelocity
			yVelocity := startYVelocity
			x := 0
			y := 0
			maxReachedY := 0
			landedInZone := false
			for x < maxX {
				x += xVelocity
				y += yVelocity
				if x >= minX && y >= minY && y <= maxY {
					landedInZone = true
					break
				}
				if xVelocity == 0 && y < minY {
					break
				}
				if xVelocity > 0 {
					xVelocity--
				}
				yVelocity--
				if y > maxReachedY {
					maxReachedY = y
				}
			}
			if landedInZone && maxReachedY > totalMaxReachedY {
				totalMaxReachedY = maxReachedY
				fmt.Printf("current max is %d\n", totalMaxReachedY)
			}
		}
	}

	return totalMaxReachedY
}

func main() {
	minX, maxX, minY, maxY := parseInput("./day17/input.txt")
	fmt.Println(findMaxY(minX, maxX, minY, maxY))
}
