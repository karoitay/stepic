package sequencing

// DeBruijnFromText returns the DeBruijn graph built from the given text.
func DeBruijnFromText(k int, text string) map[string][]string {
	m := map[string][]string{}
	for i := 0; i <= len(text)-k; i++ {
		first := text[i : i+k-1]
		second := text[i+1 : i+k]
		m[first] = append(m[first], second)
	}
	return m
}

// DeBruijnFromKmers returns the DeBruijn graph built from the given kmers.
func DeBruijnFromKmers(kmers []string) map[string][]string {
	m := map[string][]string{}
	for _, kmer := range kmers {
		from := kmer[:len(kmer)-1]
		to := kmer[1:]
		m[from] = append(m[from], to)
	}
	return m
}
