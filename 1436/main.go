package main

func destCity(paths [][]string) string {
	sp := make(map[string]struct{}, len(paths))

	for i := 0; i < len(paths); i++ {
		sp[paths[i][0]] = struct{}{}
	}

	for i := 0; i < len(paths); i++ {
		if _, ok := sp[paths[i][len(paths[i])-1]]; !ok {
			return paths[i][len(paths[i])-1]
		}
	}

	return ""
}
