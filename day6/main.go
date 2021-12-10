package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countFishesBlunt(fishes []int, days int) int {
	for i := 0; i < days; i++ {
		newIteration := make([]int, 0, 1)
		for index, fish := range fishes {
			if fish != 0 {
				fishes[index]--
				continue
			}
			fishes[index] = 6
			newIteration = append(newIteration, 8)
		}
		//fmt.Printf("%d\n", len(newIteration))
		fishes = append(fishes, newIteration...)
	}
	return len(fishes)
}

func countFishesNoMemory(fishes []int, days int) uint64 {
	var increaseIn [9]uint64
	for _, fish := range fishes {
		increaseIn[fish]++
	}
	count := uint64(len(fishes))
	for i := 0; i < days; i++ {
		newborns := increaseIn[0]
		count += newborns
		for j := 0; j < 8; j++ {
			increaseIn[j] = increaseIn[j+1]
		}
		increaseIn[8] = newborns
		increaseIn[6] += newborns
	}

	return count
}

func main() {
	input, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	fish := make([]int, 0, 300)
	fileScanner.Scan()
	fishStrings := strings.Split(fileScanner.Text(), ",")
	for _, fishStr := range fishStrings {
		value, err := strconv.ParseUint(fishStr, 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		fish = append(fish, int(value))
	}
	fmt.Println(countFishesNoMemory(fish, 80))
	fmt.Println(countFishesNoMemory(fish, 256))
}
