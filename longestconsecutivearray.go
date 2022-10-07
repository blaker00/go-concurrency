package main

import (
	"fmt"
	"sort"
)

type Consecutive struct {
	input []int
}

func (c *Consecutive) solve() int {
	data := c.input
	var results int
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	for i := 0; i < len(data); i++ {
		if i < len(data)-1 {
			if data[i+1]-data[i] == 1 {
				results = data[i+1]
			}
		}
	}

	return results

}

func main() {
	input := []int{100, 4, 5, 200, 1, 3, 2}
	c := Consecutive{input}
	fmt.Println(c.solve())

}
