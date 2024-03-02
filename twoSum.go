package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for key, val := range nums {
		e := target - val

		if _, ok := m[e]; ok {
			return []int{key, m[e]}
		} else {
			m[val] = key
		}
	}

	return []int{-1, -1}
}
