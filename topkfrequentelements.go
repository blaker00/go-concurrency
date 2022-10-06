package main

import (
	"fmt"
	"sort"
)

type kfrequent struct {
	input  []int
	output int
}

func (k *kfrequent) solve() {
	data := make(map[int]int)
	for _, v := range k.input {
		data[v] += 1
	}

	var keys []int

	for _, v := range data {
		keys = append(keys, v)
	}

	fmt.Println(data)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]

	})
	fmt.Println(keys[:k.output])

}

func main() {
	i := []int{1, 1, 2, 2, 3, 1}
	k := 2
	kf := kfrequent{input: i, output: k}
	kf.solve()

}
