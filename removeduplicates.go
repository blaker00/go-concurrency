package main

import "fmt"

func removeDuplicates(nums []int) {
	pos := 1

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1]  {
			nums[pos] = nums[i+1]
			pos++

		}
	}
	fmt.Println(pos, nums[:pos])
}

func main() {
	nums := []int{1, 1, 2, 3, 4, 4}
	removeDuplicates(nums)
}
