package main

import (
	"math"
	"strconv"
	"strings"
)

type SnailFishNumber struct {
	left   *SnailFishNumber
	right  *SnailFishNumber
	parent *SnailFishNumber
	value  int
}

func (s *SnailFishNumber) Add(other *SnailFishNumber) *SnailFishNumber {
	sum := &SnailFishNumber{
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

func (s *SnailFishNumber) reduceExplode(head *SnailFishNumber, depth int) int {
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

func (s *SnailFishNumber) reduceSplit(head *SnailFishNumber) int {
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

func (s *SnailFishNumber) split() {
	s.left = &SnailFishNumber{
		left:   nil,
		right:  nil,
		parent: s,
		value:  int(math.Floor(float64(s.value) / 2)),
	}
	s.right = &SnailFishNumber{
		left:   nil,
		right:  nil,
		parent: s,
		value:  int(math.Ceil(float64(s.value) / 2)),
	}
	s.value = 0
}

func (s *SnailFishNumber) explode() {
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

func (s *SnailFishNumber) findAnyLeftRegular() *SnailFishNumber {
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

func (s *SnailFishNumber) findAnyRightRegular() *SnailFishNumber {
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

func (s *SnailFishNumber) String() string {
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

func (s *SnailFishNumber) Magnitude() int {
	sum := 0
	if s.left != nil {
		sum += s.left.Magnitude() * 3
	}
	if s.right != nil {
		sum += s.right.Magnitude() * 2
	}
	if s.left == nil && s.right == nil {
		return s.value
	}

	return sum
}

func (s *SnailFishNumber) Clone() *SnailFishNumber {
	clone := &SnailFishNumber{}
	if s.left != nil {
		clone.left = s.left.Clone()
		clone.left.parent = clone
	}
	if s.right != nil {
		clone.right = s.right.Clone()
		clone.right.parent = clone
	}
	clone.value = s.value

	return clone
}
