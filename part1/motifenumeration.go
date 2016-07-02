package part1

func MotifEnumeration(dna []string, k, d int) []string {
	m := map[string]bool{}
	seq := dna[0]
	for i := 0; i < len(seq)-k; i++ {
		for n := range Neighbours(seq[i:i+k], d) {
			m[n] = false
		}
	}
	for pattern := range m {
		for _, seq := range dna[1:] {
			found := false
			for i := 0; i <= len(seq)-k; i++ {
				if HammingDistance(pattern, seq[i:i+k]) <= d {
					found = true
				}
			}
			if !found {
				delete(m, pattern)
				break
			}
		}
	}
	res := make([]string, len(m), len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}
	return res
}
