package main

import "fmt"

func main() {
	s, t := "abc", "ahbgdc"

	sL, sR := 0, len(s)-1
	tL, tR := 0, len(t)-1

	for tL <= tR && sL <= sR {
		if s[sL] == t[tL] {
			sL++
		}
		if s[sR] == t[tR] {
			sR--
		}

		tL++
		tR--
	}

	fmt.Println(sL > sR)
}
