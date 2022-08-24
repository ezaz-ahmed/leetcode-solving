package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfThree(81)) // truez
	fmt.Println(isPowerOfThree(27)) // true
	fmt.Println(isPowerOfThree(30)) // true
	fmt.Println(isPowerOfThree(45)) // false
	fmt.Println(isPowerOfThree(0))  // false
	fmt.Println(isPowerOfThree(9))  // true
}

// func isPowerOfThree(n int) bool {

// 	fmt.Println(int(math.Log10(float64(n)) / math.Log10(3)))

// 	return true
// }

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	x := int(math.Log(float64(n)) / math.Log(float64(3)))
	return n == int(math.Pow(3, float64(x)))
}
