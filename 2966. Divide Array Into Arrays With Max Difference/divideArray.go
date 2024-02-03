package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	results := [][]int{}

	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		} else {
			results = append(results, nums[i:i+3])
		}
	}

	return results
}

func main() {
	// divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2)
	fmt.Println(divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray([]int{1, 3, 3, 2, 7, 3}, 3))
}
