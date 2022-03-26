package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {
	index := -1
	mid := len(nums) / 2

	if nums[mid] == target {
		index = mid
	}

	return index
}
