package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	LCS := make([][]int, len(text1)+1)
	LCS[0] = make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		LCS[i] = make([]int, len(text2)+1)
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				LCS[i][j] = LCS[i-1][j-1] + 1
			} else {
				LCS[i][j] = max(LCS[i-1][j], LCS[i][j-1])
			}
		}
	}
	return LCS[len(text1)][len(text2)]

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
}
