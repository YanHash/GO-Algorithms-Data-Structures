package main

import (
	datastruc "github.com/YanHash/GO-Algorithms-Data-Structures/data-structures"
)

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
