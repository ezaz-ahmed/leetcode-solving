package main

import "fmt"

func main() {
		fmt.Println(removeElement([]int{3,2,2,3}, 3))
		fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}

func removeElement(nums []int, val int) int {
	count := 0
	for idx, el := range nums{
		if el != val {
			nums[count] = nums[idx]
			count++
		}
	}
  return count
}