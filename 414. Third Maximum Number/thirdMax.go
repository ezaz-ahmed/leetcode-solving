package main

import "fmt"

func main() {
	fmt.Println(thirdMax([]int{3, 2, 1}))    // 1
	fmt.Println(thirdMax([]int{1, 2}))       // 2
	fmt.Println(thirdMax([]int{2, 2, 3, 1})) // 1
}

// What type of problem it is?
// It's an array based probelm

func thirdMax(nums []int) int {
	var max int
	var second_max int
	var third_max int

	for _, num := range nums {
		if num == max || num == second_max || num == third_max {
			continue
		}

		if max == null || max > num {
			max = num
		} else if second_max == num || second_max > num {
			second_max = num
		} else if third_max == num || third_max > num {
			third_max = num
		}
	}

	return 0
}
