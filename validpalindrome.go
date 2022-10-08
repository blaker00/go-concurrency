package main

import (
	"fmt"
	"strings"
)

type Twopointer struct {
	input string
}

func (T *Twopointer) solve() bool {
	left := 0
	right := len(T.input) - 1

	for left < right {
		if T.input[left] != T.input[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	i := "A man, a plan, a canal: Panama"
	i = strings.ToLower(i)
	i = strings.ReplaceAll(i, " ", "")
	var final string
	for _, v := range i {
		if strings.Contains(alpha, string(v)) {
			final += string(v)
		}
	}

	t := Twopointer{input: final}
	fmt.Println(t.solve())
}
