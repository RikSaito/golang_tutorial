package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	min := l[0]
	for _, v := range l {
		if min > v {
			min = v
		}
	}

	fmt.Println(min)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _, v := range m {
		sum += v
	}

	fmt.Println(sum)
}
