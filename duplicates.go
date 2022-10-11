package main

import (
	"fmt"
	"os"
	"sort"
)

func duplicates(input []int) {
	if len(input) < 2 {
		fmt.Println("invalid length of numbers")
		os.Exit(1)
	}

	data := input
	sort.Ints(data)

	var final []int
	for i := 1; i < len(data); i++ {
		if data[i] != data[i-1] {
			final = append(final, data[i-1])
		} else if len(data)-1 == i && data[i-1] == data[i] {
			final = append(final, data[i])
		}
	}
	fmt.Println(final)
}

func main() {

	data := []int{6, 6, 1, 1, 1, 2, 3, 4, 5, 5, 1}
	duplicates(data)

}
