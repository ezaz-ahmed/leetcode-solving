package main

import "fmt"

func main() {
	num1, num2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	m, n := 3, 3

	newArr := append(num1[0:m], num2[0:n]...)

	fmt.Println(newArr)
}
