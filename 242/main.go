package main

func isAnagram(s string, t string) bool {
	m := make(map[rune]int, len(s))
	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		if _, ok := m[v]; ok {
			if m[v] == 0 {
				return false
			}
			m[v]--
		} else {
			return false
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
