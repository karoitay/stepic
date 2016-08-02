package sequencing

// DeBruijn returns the DeBruijn graph built of text.
func DeBruijn(k int, text string) map[string][]string {
	m := map[string][]string{}
	for i := 0; i <= len(text)-k; i++ {
		first := text[i : i+k-1]
		second := text[i+1 : i+k]
		m[first] = append(m[first], second)
	}
	return m
}
