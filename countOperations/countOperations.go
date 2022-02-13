package main

import "fmt"

func main() {

	num1 := 10
	num2 := 10
	operation := 0

	for num1 != 0 && num2 != 0 {
		if num1 > num2 {

			num1 = num1 - num2

			operation++
		} else if num2 > num1 {
			num2 = num2 - num1

			operation++
		} else {
			num1 = 0
			num2 = 0
			operation++
			break
		}
	}

	fmt.Println(operation)

}
