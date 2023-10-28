package main

func lengthOfLongestSubstring(s string) int {
	var max int
	bs := []byte(s)

	sl := make([]byte, 0, len(s))
	m := make(map[byte]int)

	for i := 0; i < len(bs); i++ {
		if _, ok := m[bs[i]]; !ok {
			sl = append(sl, bs[i])
			m[bs[i]] = len(sl) - 1
		} else {
			if max < len(sl) {
				max = len(sl)
			}
			sl = sl[m[bs[i]]+1:]
			sl = append(sl, bs[i])
			m = make(map[byte]int)
			for j := 0; j < len(sl); j++ {
				m[sl[j]] = j
			}
		}
	}

	if max < len(sl) {
		max = len(sl)
	}

	return max
}

func main() {
	// lengthOfLongestSubstring("abcabcbb")
	lengthOfLongestSubstring("dvdf")
}
