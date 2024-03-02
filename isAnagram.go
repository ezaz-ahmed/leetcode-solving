package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[string]int)
	n := make(map[string]int)

	for _, val := range s {
		if _, found := m[string(val)]; found {
			m[string(val)] += 1
		} else {
			m[string(val)] = 1
		}
	}

	for _, val := range t {
		if _, found := n[string(val)]; found {
			n[string(val)] += 1
		} else {
			n[string(val)] = 1
		}
	}

	if len(m) != len(n) {
		return false
	} else {
		for key := range m {
			if m[key] != n[key] {
				return false
			}
		}
	}

	return true
}
