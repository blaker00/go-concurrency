package main

import "fmt"

func moveZeros(nums []int) []int {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
	}
	return nums
}

func main() {
	input := []int{1, 2, 0, 3, 5, 2, 0, 1}
	fmt.Println(moveZeros(input))

}
