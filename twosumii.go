package main

import (
	"fmt"
)

type Twosum2 struct {
	numbers []int
	target  int
}

func (t *Twosum2) solve() []int {
	left, right := 0, len(t.numbers)-1
	var data []int
	for {
		if t.numbers[left]+t.numbers[right] < t.target {
			left++
		} else if t.numbers[right]+t.numbers[left] > t.target {
			right--
		} else {
			data = append(data, left, right)
		}

		return data
	}
}
func main() {
	input := []int{-1, 0}
	tgt := -1
	t := Twosum2{numbers: input, target: tgt}
	fmt.Println(t.solve())

}
