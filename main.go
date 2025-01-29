package main

import (
	//datastruc "github.com/YanHash/GO-Algorithms-Data-Structures/data-structures"
	"fmt"

	algorithms "github.com/YanHash/GO-Algorithms-Data-Structures/algorithms"
)

func main() {
	input := []int{9, 1, 2, 4, 6, 89, 52, 3, 5, 78, 54526, 69, 8, 54}
	l := algorithms.Bubble(input)

	fmt.Println(l)
}
