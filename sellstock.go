package main

import "fmt"

func buyandsell(perf []int) {
	buy := perf[0]
	sell := perf[1]
	for i := 0; i < len(perf)-1; i++ {
		//Sell
		if perf[i] > perf[i+1] && perf[i] > sell {
			sell = perf[i]
		}
		//buy
		if perf[i] < perf[i+1] && perf[i] < buy {
			buy = perf[i]
			sell = 0
		}
	}

	fmt.Println(sell - buy)
}
func main() {
	i := []int{7, 1, 5, 3, 6, 4}
	buyandsell(i)

}
