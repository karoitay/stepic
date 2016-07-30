package sequencing

// StringComposition returns all the strings of length k that are substrings
// of text.
// The returned slice is ordered lexicographically.
func StringComposition(k int, text string) []string {
	l := len(text) - k + 1
	res := make([]string, l, l)
	for i := 0; i < l; i++ {
		res[i] = text[i : i+k]
	}
	return res
}
