package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	indexMap := make(map[int]int)

	for currIndex, currentNumber := range nums {

		if value, ok := indexMap[target-currentNumber]; ok {
			fmt.Println("[", value, currIndex, "]")
		} else {
			indexMap[currentNumber] = currIndex
		}

	}
	fmt.Println("[]")

}
