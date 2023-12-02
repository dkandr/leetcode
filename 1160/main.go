package main

func countCharacters(words []string, chars string) int {
	mChars := getWordMap(chars)

	var res int
	var f bool
	for i := 0; i < len(words); i++ {
		mWord := getWordMap(words[i])

		for k, v := range mWord {
			if v2, ok := mChars[k]; ok {
				if v > v2 {
					f = false
					break
				}
				f = true
			} else {
				f = false
				break
			}
		}

		if f {
			res += len(words[i])
			f = false
		}
	}

	return res
}

func getWordMap(word string) map[rune]int {
	m := make(map[rune]int, len(word))
	for _, v := range word {
		m[v]++
	}

	return m
}
