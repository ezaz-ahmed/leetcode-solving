package main

import "fmt"

func main() {
	// nums1, nums2 := []int{1, 2, 2, 1}, []int{2, 2}
	// nums1, nums2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
	nums1, nums2 := []int{1, 2, 2, 1}, []int{2}

	output := []int{}

	for _, firstNumber := range nums1 {
		for secIndex, secNumber := range nums2 {
			if firstNumber == secNumber {
				output = append(output, secNumber)
				nums2 = append(nums2[:secIndex], nums2[secIndex+1:]...)
				break
			}
		}
	}

	fmt.Println(output)
}
