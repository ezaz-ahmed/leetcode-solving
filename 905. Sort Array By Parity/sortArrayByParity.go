package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
}

func sortArrayByParity(nums []int) []int {
	n := 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}

	return nums
}
