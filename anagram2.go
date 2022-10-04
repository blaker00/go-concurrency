package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Anagram struct {
	index1, index2 string
}

func (a *Anagram) solve() bool {
	if len(a.index1) == len(a.index2) {
		data := strings.Split(a.index1, "")
		data2 := strings.Split(a.index2, "")
		sort.Strings(data)
		sort.Strings(data2)
		return reflect.DeepEqual(data, data2)
	}
	return false
}
func main() {
	s := "anagram"
	t := "angaram"
	f := Anagram{index1: s, index2: t}
	fmt.Println(f.solve())
}
