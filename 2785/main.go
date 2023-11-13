package main

import "sort"

func sortVowels(s string) string {
	m := make(map[byte]struct{}, 10)
	m['A'] = struct{}{}
	m['E'] = struct{}{}
	m['I'] = struct{}{}
	m['O'] = struct{}{}
	m['U'] = struct{}{}
	m['a'] = struct{}{}
	m['e'] = struct{}{}
	m['i'] = struct{}{}
	m['o'] = struct{}{}
	m['u'] = struct{}{}

	sl := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			sl = append(sl, s[i])
		}
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[j] > sl[i]
	})

	out := make([]byte, 0, len(s))

	pos := 0
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			out = append(out, sl[pos])
			pos++
		} else {
			out = append(out, s[i])
		}
	}

	return string(out)
}
