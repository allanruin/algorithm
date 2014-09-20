package graph

import (
// "fmt"
)

type Stack struct {
	values []int
	// size   int
}

func (s *Stack) Push(val int) {
	s.values = append(s.values, val)
	// size = size + 1
}

func (s *Stack) Pop() int {
	val := -1
	if len(s.values) != 0 {

		if len(s.values) == 1 {
			val = s.values[0]
			s.values = s.values[0:0]
		} else {
			val = s.values[len(s.values)-1]
			s.values = s.values[:len(s.values)-1]
		}
		// fmt.Println("len:", len(s.values))
		// size = size - 1
	}
	return val
}

func (s *Stack) GetValues() []int {
	return s.values
}

func (s *Stack) Size() int {
	return len(s.values)
}
