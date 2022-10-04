package main

import "fmt"

func twosum(list []int, num int) {

	data := make(map[int]int)

	for idx, val := range list {
		v, found := data[num-val]
		if found {
			fmt.Println(v, idx)
		}
		data[val] = idx
	}

}

func main() {
	l := []int{1, 2, 4, 5, 5}
	n := 5
	twosum(l, n)

}
