package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(1))
}

func isPowerOfFour(num int) bool {
	if num == 0 {
		return false
	}
	for num%4 == 0 {
		num /= 4
	}
	return num == 1
}
