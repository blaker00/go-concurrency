package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
)

func threesum(input []int) [][]int {
	//sanity check
	if len(input) < 3 {
		fmt.Println("invalid input")
		os.Exit(1)
	}
	//sort
	data := input
	sort.Ints(data)

	//3sum
	left := 0
	right := len(data) - 1
	var final [][]int
	for left < right {
		if data[left]+data[left+1]+data[right] < 0 {
			left++
		} else if data[left]+data[left+1]+data[right] > 0 {
			right--
		} else {
			if reflect.DeepEqual(final, [][]int{{data[left], data[left+1], data[right]}}) == false {
				final = append(final, []int{data[left], data[left+1], data[right]})
			}
			left++
		}

	}
	return final
}

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threesum(input))

}
