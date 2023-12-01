package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var w1, w2, p1, p2 int
	for {
		if p1 == len(word1[w1]) {
			p1 = 0
			w1++
		}
		if p2 == len(word2[w2]) {
			p2 = 0
			w2++
		}

		if w1 == len(word1) || w2 == len(word2) {
			if w1 == len(word1) && w2 == len(word2) {
				return true
			}

			return false
		}

		if word1[w1][p1] != word2[w2][p2] {
			return false
		}

		p1++
		p2++
	}
}
