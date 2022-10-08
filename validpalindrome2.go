package main

import "fmt"

type Palindromeii struct {
	input string
}

func (p *Palindromeii) solve() bool {
	left := 0
	right := len(p.input) - 1
	for left < right {
		if string(p.input[left]) != string(p.input[right]) {
			if string(p.input[left+1]) == string(p.input[right]) {
			} 
			if string(p.input[left]) == string(p.input[right-1]) {
			} else {
				return false
			}

		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	i := "aba"
	p := Palindromeii{input: i}
	fmt.Println(p.solve())
}
