package main

import "fmt"

func main() {
	fmt.Println("hi")
}

// Collected Solution

func deleteAndEarn(nums []int) int {
	freq := make([]int, 10001)
	for _, v := range nums {
		freq[v]++
	}
	dp1, dp2 := 0, freq[1]
	dp := dp2
	for i := 2; i < 10001; i++ {
		dp = max(dp1+freq[i]*i, dp2)
		dp1 = dp2
		dp2 = dp
	}
	return dp
}

func max(in ...int) int {
	max := 0
	for _, v := range in {
		if v > max {
			max = v
		}
	}
	return max
}
