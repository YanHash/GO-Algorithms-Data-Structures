package algorithms

import "fmt"

func Bubble(input []int) []int {
	l := len(input)
	swaped := true
	fmt.Println(input)
	for swaped {
		swaped = false
		for i := 0; i < l-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swaped = true
				fmt.Println(input)
			}
		}
	}
	return input
}
