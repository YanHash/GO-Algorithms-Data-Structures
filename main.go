package main

import datastruc "github.com/YanHash/GO-Algorithms-Data-Structures/data-structures"

func main() {
	head := &datastruc.Linked_list{Data: 1, Link: nil}
	head.AddToTail(2)
	head.AddToTail(3)
	head.Print()
}