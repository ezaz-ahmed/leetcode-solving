package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, val := range nums {
		if m[val] {
			return true
		}
		m[val] = true
	}

	return false
}
