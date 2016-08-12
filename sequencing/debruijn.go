package sequencing

// DeBruijnFromText returns the DeBruijn graph built from the given text.
func DeBruijnFromText(k int, text string) *Graph {
	m := map[Node][]Node{}
	for i := 0; i <= len(text)-k; i++ {
		first := text[i : i+k-1]
		second := text[i+1 : i+k]
		m[KmerRead(first)] = append(m[KmerRead(first)], KmerRead(second))
	}
	g := Graph(m)
	return &g
}

// DeBruijnFromKmers returns the DeBruijn graph built from the given kmers.
func DeBruijnFromKmers(kmers []string) *Graph {
	return DeBruijnFromReads(kmers, KmerReadParser{})
}

// DeBruijnFromReads return the DeBruijin graph for some reads.
func DeBruijnFromReads(reads []string, p ReadParser) *Graph {
	m := map[Node][]Node{}
	for _, read := range reads {
		e := p.ParseRead(read).ToEdge()
		m[e.From] = append(m[e.From], e.To)
	}
	g := Graph(m)
	return &g
}
