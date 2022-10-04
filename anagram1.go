package main

import (
	"fmt"
	"reflect"
)

type Anagram struct {
	index1, index2 string
}

func (a *Anagram) solve() bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)

	for _, v := range a.index1 {
		map1[string(v)] += 1
	}

	for _, v := range a.index2 {
		map2[string(v)] += 1
	}

	return reflect.DeepEqual(map1, map2)

}
func main() {
	s := "anagram"
	t := "nagaram"
	f := Anagram{index1: s, index2: t}
	fmt.Println(f.solve())
}
