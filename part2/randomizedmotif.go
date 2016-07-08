package part2

import (
	"math/rand"
	"time"
)

const timesToRun = 1000

func RandomizedMotifSearch(dna []string, k int) []string {
	return randomizedMotifSearch(dna, k, time.Now().UnixNano())
}

func randomizedMotifSearch(dna []string, k int, seed int64) []string {
	rand.Seed(seed)
	bestScore := len(dna)*len(dna[0]) + 1
	var bestMotifs []string
	for i := 0; i < timesToRun; i++ {
		motifs := singleRandomizedMotifSearch(dna, k)
		newScore := MotifsScore(motifs)
		if newScore < bestScore {
			bestMotifs = motifs
			bestScore = newScore
		}
	}
	return bestMotifs
}

func singleRandomizedMotifSearch(dna []string, k int) []string {
	motifs := randomMotifs(dna, k)
	bestMotifs := motifs
	bestScore := MotifsScore(motifs)
	for {
		profile := motifsPseudoCountProfile(bestMotifs)
		for i := 0; i < len(dna); i++ {
			motifs[i] = ProfileMostProbable(dna[i], k, profile)
		}
		newScore := MotifsScore(motifs)
		if newScore < bestScore {
			bestMotifs = motifs
			bestScore = newScore
		} else {
			return bestMotifs
		}
	}
	return bestMotifs
}

func randomMotifs(dna []string, k int) []string {
	motifs := make([]string, len(dna), len(dna))
	for i := 0; i < len(dna); i++ {
		j := rand.Intn(len(dna[i]) - k + 1)
		motifs[i] = dna[i][j : j+k]
	}
	return motifs
}
