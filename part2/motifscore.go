package part2

import "math"

func motifsCount(motifs []string) map[byte][]int {
	l := len(motifs[0])
	m := map[byte][]int{
		'A': make([]int, l, l),
		'C': make([]int, l, l),
		'G': make([]int, l, l),
		'T': make([]int, l, l),
	}
	for _, motif := range motifs {
		for i := 0; i < len(motif); i++ {
			m[motif[i]][i] += 1
		}
	}
	return m
}

func motifsProfile(motifs []string) map[byte][]float64 {
	m := map[byte][]float64{}
	l := float64(len(motifs))
	for k, v := range motifsCount(motifs) {
		m[k] = make([]float64, len(v), len(v))
		for i := 0; i < len(v); i++ {
			m[k][i] = float64(v[i]) / l
		}
	}
	return m
}

func motifsPseudoCountProfile(motifs []string) map[byte][]float64 {
	m := map[byte][]float64{}
	l := float64(len(motifs)) * 2
	for k, v := range motifsCount(motifs) {
		m[k] = make([]float64, len(v), len(v))
		for i := 0; i < len(v); i++ {
			m[k][i] = float64(v[i]+1) / l
		}
	}
	return m
}

func MotifsEntropy(motifs []string) float64 {
	entropy := 0.0
	for _, v := range motifsProfile(motifs) {
		for _, i := range v {
			if i != 0 {
				entropy -= i * math.Log2(i)
			}
		}
	}
	return entropy
}

func MotifsScore(motifs []string) int {
	m := motifsCount(motifs)
	l := len(motifs)
	score := 0
	for i := 0; i < len(m['A']); i++ {
		max := m['A'][i]
		if m['C'][i] > max {
			max = m['C'][i]
		}
		if m['G'][i] > max {
			max = m['G'][i]
		}
		if m['T'][i] > max {
			max = m['T'][i]
		}
		score += l - max
	}
	return score
}
