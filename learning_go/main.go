package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	v := []float64{1.1, 5.2, 12.5, 9.2, 10.02}
	average(v)
}

// FOR LOOPS
func fill_arr(max int) {
	var arr [10]int
	for i := range max {
		arr[i] = i
	}
	fmt.Println(arr)
}

func for_loop(max int) {
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
}

func goto_loop(max int) {
	i := 0
Printt:
	if i < max {
		fmt.Println(i)
		i++
		goto Printt
	}
}

// FIZZBUZZ
func fizzbuzz() {
	for i := range 100 {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// STRINGS
func multicarac(c string, l int) {
	chars := c
	i := 0
	for i < l {
		fmt.Println(chars)
		chars += c
		i++
	}
}

func carac_counter(chars string) {
	m := make(map[byte]int)
	sum := 0
	for i := range len(chars) {
		if _, exist := m[chars[i]]; exist {
			m[chars[i]] += 1
		} else {
			m[chars[i]] = 1
		}
		if m[chars[i]] != ' ' {
			sum++
		}
	}
	fmt.Println("distribution : ", m)
	fmt.Println("total caracters : ", sum)
	fmt.Println("total bytes : ", utf8.RuneCountInString(chars))
}

func reverse(chars string) {
	i := len(chars) - 1
	str := ""

	for i >= 0 {
		str += string(chars[i])
		i--
	}
	chars = str
	fmt.Println(chars)
}

// Average
func average(values []float64) {
	sum := float64(0)
	for _, v := range values {
		sum += v
	}
	av := sum / float64(len(values))
	fmt.Println(av)
}
