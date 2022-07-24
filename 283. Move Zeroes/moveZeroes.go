package main

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0})
}

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n], nums[i] = nums[i], nums[n]
			n++
		}
	}
}
