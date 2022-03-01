package main

import "fmt"

func main() {

	n := 5

	ans := make([]int, n+1)

	// It was my solution

	// for i := 0; i <= n; i++ {

	// 	count := 0
	// 	pick := i

	// 	for pick >= 1 {
	// 		if pick%2 == 1 {
	// 			count++
	// 		}
	// 		pick = pick / 2
	// 	}

	// 	ans[i] = count
	// }

	fmt.Println(ans)

	// This is solution of others.. -> Better approach

	ans[0] = 0

	for i := 1; i <= n; i++ {
		ans[i] = ans[i/2] + i%2
	}

	fmt.Println(ans)

}
