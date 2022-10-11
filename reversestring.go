package main

import "fmt"

type Reverse struct {
	input []string
}

func (r *Reverse) solve() []string {
	fmt.Println(r.input)
	l := 0
	re := len(r.input) - 1

	for l < re {
		r.input[l], r.input[re] = r.input[re], r.input[l]
		l++
		re--
	}
	return r.input

}

func main() {
	i := []string{"h", "e", "l", "l", "o"}
	r := Reverse{input: i}
	fmt.Println(r.solve())

}
