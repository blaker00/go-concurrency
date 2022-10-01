package main

import (
	"fmt"
	"os"
	"unicode"
)

type CamelCase struct {
	data string
}

func (c *CamelCase) deploy() {

	var tr string
	var fl string
	for _, v := range c.data {
		if unicode.IsUpper(v) == true {
			tr += string(v)
		} else {
			fl += string(v)
		}
	}
	fmt.Println(tr)
	fmt.Println(fl)

}

func main() {
	fmt.Println("Please input your query.")
	var input string
	_, err := fmt.Scanln(&input) 
	if err != nil {
		os.Exit(1)
	}
	c := CamelCase{data: input}
	c.deploy()
}
