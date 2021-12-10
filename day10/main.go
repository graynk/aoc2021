package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type runeStack struct {
	stack []rune
}

func newRuneStack(capacity int) runeStack {
	return runeStack{stack: make([]rune, 0, capacity)}
}

func (rs *runeStack) push(r rune) {
	rs.stack = append(rs.stack, r)
}

func (rs *runeStack) pop() rune {
	last := len(rs.stack) - 1
	if last == -1 {
		return -1
	}
	r := rs.stack[last]
	rs.stack = rs.stack[:last]
	return r
}

func (rs *runeStack) peek() rune {
	last := len(rs.stack) - 1
	if last == -1 {
		return -1
	}
	return rs.stack[last]
}

func findCorrupted(lines []string) int {
	sum := 0
	openers := make(map[rune]rune)
	openers[')'] = '('
	openers[']'] = '['
	openers['}'] = '{'
	openers['>'] = '<'
	values := make(map[rune]int)
	values[')'] = 3
	values[']'] = 57
	values['}'] = 1197
	values['>'] = 25137

	for _, row := range lines {
		stack := newRuneStack(len(row))
	loop:
		for _, symbol := range row {
			switch symbol {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack.push(symbol)
			case ')':
				fallthrough
			case ']':
				fallthrough
			case '}':
				fallthrough
			case '>':
				if stack.peek() != openers[symbol] {
					sum += values[symbol]
					break loop
				}
				stack.pop()
			}

		}

	}
	return sum
}

func completeIncompletes(lines []string) int {
	openers := make(map[rune]rune)
	openers[')'] = '('
	openers[']'] = '['
	openers['}'] = '{'
	openers['>'] = '<'
	values := make(map[rune]int)
	values['('] = 1
	values['['] = 2
	values['{'] = 3
	values['<'] = 4

	scores := make([]int, 0, 1)

	for _, row := range lines {
		stack := newRuneStack(len(row))
		isCorrupted := false
	loop:
		for _, symbol := range row {
			switch symbol {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				fallthrough
			case '<':
				stack.push(symbol)
			case ')':
				fallthrough
			case ']':
				fallthrough
			case '}':
				fallthrough
			case '>':
				if stack.peek() != openers[symbol] {
					isCorrupted = true
					break loop
				}
				stack.pop()
			}
		}
		if isCorrupted {
			continue
		}
		sum := 0
		for symbol := stack.pop(); symbol != -1; symbol = stack.pop() {
			sum = (sum * 5) + values[symbol]
		}
		scores = append(scores, sum)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func parseInput(filename string) []string {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	input := make([]string, 0, 106)
	for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
	}
	return input
}

func main() {
	input := parseInput("./day10/input.txt")
	fmt.Println(findCorrupted(input))
	fmt.Println(completeIncompletes(input))
}
