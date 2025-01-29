package datastruc

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Length int
	Last   *Node
}

func (s *Stack) Push(data int) int {
	current_Last := s.Last
	n := Node{
		Value: data,
		Next:  current_Last,
	}

	s.Last = &n
	s.Length++

	return s.Last.Value
}

func (s *Stack) Pop() (int, error) {
	var output int

	if s.Length > 1 {
		output = s.Last.Value
		s.Last = s.Last.Next
		s.Length--
		return output, nil
	}
	if s.Length == 1 {
		output = s.Last.Value
		s.Last = nil
		s.Length--
		return output, nil
	}
	return -1, errors.New("empty stack")
}

func (s Stack) ToString() {
	output := "Stack from the Last to the First: \n"
	current := s.Last
	for s.Length > 0 {
		output += fmt.Sprint(current.Value) + "\n"
		current = current.Next
		s.Length--
	}
	fmt.Print(output)
}

/*
func main() {
	n0 := datastruc.Node{
		Value: 10,
		Next:  nil,
	}

	n1 := datastruc.Node{
		Value: 20,
		Next:  &n0,
	}
	n2 := datastruc.Node{
		Value: 30,
		Next:  &n1,
	}
	n3 := datastruc.Node{
		Value: 40,
		Next:  &n2,
	}
	n4 := datastruc.Node{
		Value: 50,
		Next:  &n3,
	}
	n5 := datastruc.Node{
		Value: 60,
		Next:  &n4,
	}

	s := datastruc.Stack{
		Length: 6,
		Last:   &n5,
	}

	s.ToString()
	s.Pop()
	s.ToString()

	s.Push(70)
	s.ToString()
}
*/
