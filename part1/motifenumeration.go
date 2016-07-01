package part1

func MotifEnumeration(dna []string, k, d int) []string {
	m := map[string]bool{}
	seq := dna[0]
	for i := 0; i < len(seq)-k; i++ {
		for n := range Neighbours(seq[i:i+k], d) {
			m[n] = false
		}
	}
	for _, seq := range dna[1:] {
		for i := 0; i <= len(seq)-k; i++ {
			for n := range Neighbours(seq[i:i+k], d) {
				if _, ok := m[n]; ok {
					m[n] = true
				}
			}
		}
		for k, v := range m {
			if v {
				m[k] = false
			} else {
				delete(m, k)
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
