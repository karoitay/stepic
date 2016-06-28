package part1

func HammingDistance(text1, text2 string) int {
	h := 0
	for i := 0; i < len(text1); i++ {
		if text1[i] != text2[i] {
			h += 1
		}
	}
	return h
}
