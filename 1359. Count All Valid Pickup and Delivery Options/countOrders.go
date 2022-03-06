package main

import "fmt"

func main() {
	fmt.Println(countOrders(1))
	fmt.Println(countOrders(2))
	fmt.Println(countOrders(3))
	fmt.Println(countOrders(4))
}

// Ans Collected

func countOrders(n int) int {
	ans, m := 1, 1000000007
	for i := 2; i <= n; i++ {
		ans *= (2*i - 1) * i
		ans %= m
	}
	return ans
}
