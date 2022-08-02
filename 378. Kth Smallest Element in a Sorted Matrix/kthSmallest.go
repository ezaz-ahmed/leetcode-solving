package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8)) // 13
	fmt.Println(kthSmallest([][]int{{-5}}, 1))                                  // -5
	fmt.Println(kthSmallest([][]int{{1, 2}, {1, 3}}, 2))                        // 1
}

func kthSmallest(matrix [][]int, k int) int {

	arr := make([]int, 0, len(matrix)*len(matrix[0]))

	for i := range matrix {
		for _, v := range matrix[i] {
			arr = append(arr, v)
		}
	}

	sort.Ints(arr)

	return arr[k-1]
}
