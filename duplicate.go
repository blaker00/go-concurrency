package main

import "fmt"

type duplicate struct {
	data []string
}

func (d *duplicate) run() bool {
	data := make(map[string]int)
	for _, v := range d.data {
		data[v] += 1
	}

	for k := range data {
		if data[k] > 1 {
			return true
		}
	}
	return false

}

func main() {
	st := []string{"c", "d", "b"}
	d := duplicate{data: st}
	fmt.Println(d.run())
}
