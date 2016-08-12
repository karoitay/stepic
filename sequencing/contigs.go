package sequencing

import "strings"

// GetContigs returns the contigs from a kmers reads.
func GetContigs(kmers []string) []string {
	g := DeBruijnFromKmers(kmers)
	paths := MaximalNonBranchingPaths(g)
	contigs := make([]string, len(paths))
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		s := make([]string, len(path))
		s[0] = path[0].ToString()
		for j := 1; j < len(path); j++ {
			s[j] = path[j].ToString()
			s[j] = s[j][len(s[j])-1:]
		}
		contigs[i] = strings.Join(s, "")
	}
	return contigs
}
