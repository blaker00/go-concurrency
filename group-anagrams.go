package main

import (
	"fmt"
	"sort"
	"strings"
)

type Ganagram struct {
	strs []string
}

func (a *Ganagram) solve() map[string][]string {
	m := make(map[string][]string)
	for _, val := range a.strs {
		ss := strings.Split(val, "")
		sort.Strings(ss)
		str := strings.Join(ss, "")
		m[str] = append(m[str], val)

	}
	return m
}
func main() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	s := Ganagram{strs: data}
	fmt.Println(s.solve())

}
