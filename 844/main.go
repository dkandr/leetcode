package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	if s == t {
		return true
	}

	sp, tp := len(s)-1, len(t)-1 // текущая позиция в строке
	si, ti := 0, 0               // количеество удаленных символов

	for {
		if sp >= 0 && s[sp] == '#' {
			si++
			sp--
			continue
		}
		if tp >= 0 && t[tp] == '#' {
			ti++
			tp--
			continue
		}

		if si > 0 && sp >= 0 {
			si--
			sp--
			continue
		}
		if ti > 0 && tp >= 0 {
			ti--
			tp--
			continue
		}

		if sp != -1 && tp != -1 {
			if s[sp] != t[tp] {
				return false
			}
		}

		if sp == -1 || tp == -1 {
			switch {
			case sp == -1 && tp == -1:
				return true
			case sp == -1 && tp-ti < 0:
				return true
			case tp == -1 && sp-si < 0:
				return true
			default:
				return false
			}
		}

		sp--
		tp--
	}
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("bxj##tw", "bxj###tw"))
	fmt.Println(backspaceCompare("a#c", "b"))
}
