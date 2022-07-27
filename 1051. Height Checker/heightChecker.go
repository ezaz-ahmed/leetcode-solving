package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3})) // 3
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))    // 5
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))    // 0
}

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)

	count := 0

	for in, el := range sorted {
		if el != heights[in] {
			count++
		}
	}

	return count
}
