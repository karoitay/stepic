package part1

type profileFunc func([]string) map[byte][]float64

func ProfileMostProbable(text string, k int, profile map[byte][]float64) string {
	res := ""
	max := -1.0
	for i := 0; i <= len(text)-k; i++ {
		prob := 1.0
		for j := i; j < i+k; j++ {
			prob *= profile[text[j]][j-i]
		}
		if prob > max {
			max = prob
			res = text[i : i+k]
		}
	}
	return res
}

func GreedyMotifSearch(dna []string, k int) []string {
	return greedyMotifSearch(dna, k, motifsProfile)
}

func PseudoCountGreedyMotifSearch(dna []string, k int) []string {
	return greedyMotifSearch(dna, k, motifsPseudoCountProfile)
}

func greedyMotifSearch(dna []string, k int, profile profileFunc) []string {
	var bestMotifs []string
	// The lower the score the better, so we initialize it with
	// a value greater than the max possible score.
	bestScore := k * len(dna)
	for i := 0; i <= len(dna[0])-k; i++ {
		motifs := []string{dna[0][i : i+k]}
		for j := 1; j < len(dna); j++ {
			nextMotif := ProfileMostProbable(dna[j], k, profile(motifs))
			motifs = append(motifs, nextMotif)
		}
		score := MotifsScore(motifs)
		if score < bestScore {
			bestMotifs = motifs
			bestScore = score
		}
	}
	return bestMotifs
}
