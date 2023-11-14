package main

func countPalindromicSubsequence(s string) int {
	if len(s) < 3 {
		return 0
	}

	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]] = i
	}

	var count int
	// для уникальных символов
	um := make(map[byte]struct{}, len(s))

	for i := 0; i < len(s)-2; i++ {
		if v, ok := m[s[i]]; ok {
			// если расстояние меньше 2
			if v-i < 2 {
				delete(m, s[i])
				continue
			}

			// ищу количество различных символов
			clear(um)

			for j := i + 1; j < v; j++ {
				if _, ok := um[s[j]]; !ok {
					um[s[j]] = struct{}{}
				}
			}

			count += len(um)
			delete(m, s[i])
		}
	}

	return count
}
