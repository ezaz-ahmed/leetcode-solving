package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3,4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 7)) // [1,2]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1,-1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1,-1]
}

// I wasn't able to solve it...
func searchRange(nums []int, target int) []int {
	output := []int{-1, -1}
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l] == target && output[0] == -1 {
			output[0] = l
		}

		if nums[r] == target && output[1] == -1 {
			output[1] = r
		}

		if output[0] != -1 && output[1] != -1 {
			return output
		} else if output[0] == -1 && output[1] != -1 {
			l++
		} else if output[0] != -1 && output[1] == -1 {
			r--
		} else {
			l++
			r--
		}
	}

	return output
}
