package main

import "fmt"

func main () {
	fmt.Println(removeDuplicates([]int {1,1,2}))
	fmt.Println(removeDuplicates([]int {0,0,1,1,1,2,2,3,3,4}))
}

func removeDuplicates(nums []int) int {

	left:= 1
	right:= 1

	for right < len(nums) {
		if (nums[right] == nums[right - 1]) {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}

	return left
}