package sequencing

// OverlapGraph returns the overlap graph in the form of adjacency list.
func OverlapGraph(kmers []string) map[string][]string {
	graph := map[string][]string{}
	for i := 0; i < len(kmers); i++ {
		for j := 0; j < len(kmers); j++ {
			if i == j {
				continue
			}
			if shouldConnect(kmers[i], kmers[j]) {
				graph[kmers[i]] = append(graph[kmers[i]], kmers[j])
			}
		}
	}
	return graph
}

func shouldConnect(kmer1, kmer2 string) bool {
	return kmer1[1:] == kmer2[:len(kmer2)-1]
}
