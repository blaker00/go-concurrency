package main

import (
	"fmt"
	"sort"
	"strings"
)

type Ganagram struct {
	strs []string
}

func (a *Ganagram) solve() {
	m := make(map[string][]string)
	for _, val := range a.strs {
		ss := strings.Split(val, "")
		sort.Strings(ss)
		str := strings.Join(ss, "")
		m[str] = append(m[str], val)

	}
	fmt.Println(m)
}
func main() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	s := Ganagram{strs: data}
	s.solve()

}
