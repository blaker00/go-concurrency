package main

import "fmt"

func owner() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func consumer(ch <-chan int) {
	data := ch
	for v := range data {
		fmt.Println(v)
	}
	fmt.Println("done!")

}

func main() {
	ow := owner()
	consumer(ow)
}
