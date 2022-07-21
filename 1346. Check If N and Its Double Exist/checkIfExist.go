package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkIfExist([]int {10, 2, 5, 3})) // true
	fmt.Println(checkIfExist([]int {7, 1, 14, 11})) // true
	fmt.Println(checkIfExist([]int {3, 1, 7, 11})) // false
	fmt.Println(checkIfExist([]int {-10,12,-20,-8,15})) // true
}

func checkIfExist(arr []int) bool {

	check := make(map[int]bool, len(arr))
    
	for i := 0 ; i < len(arr) ; i++ {
			if _, exists := check[arr[i]*2] ; exists {
					return true
			}
			if _, exists := check[arr[i]/2] ; exists && arr[i]%2==0{
					return true
			}
			check[arr[i]] = true
	}
	
	return false
}