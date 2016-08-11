package sequencing

// DeBruijnFromText returns the DeBruijn graph built from the given text.
func DeBruijnFromText(k int, text string) *Graph {
	m := map[Node][]Node{}
	for i := 0; i <= len(text)-k; i++ {
		first := text[i : i+k-1]
		second := text[i+1 : i+k]
		m[Read(first)] = append(m[Read(first)], Read(second))
	}
	g := Graph(m)
	return &g
}

// DeBruijnFromKmers returns the DeBruijn graph built from the given kmers.
func DeBruijnFromKmers(kmers []string) *Graph {
	m := map[Node][]Node{}
	for _, kmer := range kmers {
		e := Read(kmer).ToEdge()
		m[e.From] = append(m[e.From], e.To)
	}
	g := Graph(m)
	return &g
}
