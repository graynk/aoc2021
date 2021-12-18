package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type snailFishNumber struct {
	left   *snailFishNumber
	right  *snailFishNumber
	parent *snailFishNumber
	value  int
}

func (s *snailFishNumber) add(other *snailFishNumber) *snailFishNumber {
	sum := &snailFishNumber{
		left:  s,
		right: other,
		value: 0,
	}
	sum.left.parent = sum
	sum.right.parent = sum
	changes := 0
	for {
		changes = sum.reduceExplode(sum, 0)
		if changes != 0 {
			continue
		}
		changes = sum.reduceSplit(sum)
		if changes == 0 {
			break
		}
	}
	return sum
}

//func (s *snailFishNumber) reduce(head *snailFishNumber, depth int) {
//	if depth == 5 {
//		s.parent.explode()
//		head.reduce(head, 0)
//		return
//	}
//	if s.parent == nil && depth != 0 {
//		return
//	}
//	if s.left == nil && s.right == nil {
//		if s.value > 9 {
//			s.split()
//			head.reduce(head, 0)
//		}
//		return
//	}
//	if s.left == nil {
//		return
//	}
//	s.left.reduce(head, depth+1)
//	if s.right == nil {
//		return
//	}
//	s.right.reduce(head, depth+1)
//}

func (s *snailFishNumber) reduceExplode(head *snailFishNumber, depth int) int {
	if depth == 5 {
		s.parent.explode()
		return 1
	}
	if s.parent == nil && depth != 0 {
		return 0
	}
	if s.left == nil {
		return 0
	}
	changes := s.left.reduceExplode(head, depth+1)
	if s.right == nil || changes != 0 {
		return changes
	}
	return s.right.reduceExplode(head, depth+1)
}

func (s *snailFishNumber) reduceSplit(head *snailFishNumber) int {
	//if s.parent == nil && depth != 0 {
	//	return
	//}
	if s.left == nil && s.right == nil {
		if s.value > 9 {
			s.split()
			return 1
		}
		return 0
	}
	if s.left == nil {
		return 0
	}
	changes := s.left.reduceSplit(head)
	if s.right == nil || changes != 0 {
		return changes
	}

	return s.right.reduceSplit(head)
}

func (s *snailFishNumber) split() {
	s.left = &snailFishNumber{
		left:   nil,
		right:  nil,
		parent: s,
		value:  int(math.Floor(float64(s.value) / 2)),
	}
	s.right = &snailFishNumber{
		left:   nil,
		right:  nil,
		parent: s,
		value:  int(math.Ceil(float64(s.value) / 2)),
	}
	s.value = 0
}

func (s *snailFishNumber) explode() {
	left := s.findAnyLeftRegular()
	if left != nil {
		left.value += s.left.value
	}
	right := s.findAnyRightRegular()
	if right != nil {
		right.value += s.right.value
	}
	s.left.parent = nil
	s.left = nil
	s.right.parent = nil
	s.right = nil
	s.value = 0
}

func (s *snailFishNumber) findAnyLeftRegular() *snailFishNumber {
	prev := s
	for number := s.parent; number != nil; {
		if number.right == prev {
			prev = number
			number = number.left
			continue
		}
		if number.left == prev {
			if s.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.right
	}

	return nil
}

func (s *snailFishNumber) findAnyRightRegular() *snailFishNumber {
	prev := s
	for number := s.parent; number != nil; {
		if number.left == prev {
			prev = number
			number = number.right
			continue
		}
		if number.right == prev {
			if s.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.left
	}

	return nil
}

func (s *snailFishNumber) String() string {
	builder := strings.Builder{}
	if s.left != nil {
		builder.WriteString("[")
		builder.WriteString(s.left.String())
		builder.WriteString(",")
	}
	if s.right != nil {
		builder.WriteString(s.right.String())
		builder.WriteString("]")
	}
	if s.left == nil && s.right == nil {
		builder.WriteString(strconv.Itoa(s.value))
	}

	return builder.String()
}

func sumNumbers(numbers []*snailFishNumber) *snailFishNumber {
	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = sum.add(numbers[i])
	}
	return sum
}

func (s *snailFishNumber) magnitude() int {
	sum := 0
	if s.left != nil {
		sum += s.left.magnitude() * 3
	}
	if s.right != nil {
		sum += s.right.magnitude() * 2
	}
	if s.left == nil && s.right == nil {
		return s.value
	}

	return sum
}

func parseInput(filename string) []*snailFishNumber {
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	numbers := make([]*snailFishNumber, 0, 1)
	for fileScanner.Scan() {
		number := &snailFishNumber{}
		start := number
		line := fileScanner.Text()
		for _, symbol := range line {
			switch symbol {
			case '[':
				number.left = &snailFishNumber{parent: number}
				number.right = &snailFishNumber{parent: number}
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
	numbers := parseInput("./day18/input.txt")
	sum := sumNumbers(numbers)
	fmt.Println(sum.magnitude())
}
