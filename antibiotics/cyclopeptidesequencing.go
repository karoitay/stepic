package antibiotics

import (
	"reflect"
	"sort"
)

type peptideCandidate struct {
	peptide Peptide
	mass    int
	score   int
}

func newCandidate(peptide Peptide, mass int) peptideCandidate {
	return newScoredCandidate(peptide, mass, 0)
}

func newScoredCandidate(peptide Peptide, mass, score int) peptideCandidate {
	return peptideCandidate{peptide, mass, score}
}

// CyclopeptideSequencing returns the possible peptides with the given
// cyclo-spectrm.
func CyclopeptideSequencing(spectrum, availableMasses []int) []Peptide {
	parentMass := spectrum[len(spectrum)-1]
	spec := spectrumToMap(spectrum)

	matchingPeptides := []Peptide{}
	candidates := []peptideCandidate{newCandidate(NewPeptideFromString(""), 0)}
	for len(candidates) != 0 {
		var expandedCandidates []peptideCandidate
		for _, candidate := range candidates {
			for _, mass := range availableMasses {
				newCandidate := newCandidate(candidate.peptide.Expand(mass), candidate.mass+mass)
				if newCandidate.mass == parentMass {
					if reflect.DeepEqual(Cyclospectrum(newCandidate.peptide), spectrum) {
						matchingPeptides = append(matchingPeptides, newCandidate.peptide)
					}
				} else if newCandidate.mass < parentMass && consistent(newCandidate.peptide, spec) {
					expandedCandidates = append(expandedCandidates, newCandidate)
				}
			}
		}
		candidates = expandedCandidates
	}
	return matchingPeptides
}

// LeaderboardCyclopeptideSequencing returns the highest scoring peptides.
// In every iteration only the n higest scoring (keeping ties)
// candidates will be kept then the one with the highest score will be returned.
func LeaderboardCyclopeptideSequencing(spectrum []int, n int, availableMasses []int) []Peptide {
	parentMass := spectrum[len(spectrum)-1]
	spec := spectrumToMap(spectrum)

	leaderPeptides := []peptideCandidate{newScoredCandidate(NewPeptideFromString(""), 0, 0)}
	leaderboard := []peptideCandidate{leaderPeptides[0]}
	for len(leaderboard) > 0 {
		var expandedLeaderboard []peptideCandidate
		for _, candidate := range leaderboard {
			for _, mass := range availableMasses {
				if candidate.mass+mass <= parentMass {
					expanded := candidate.peptide.Expand(mass)
					score := LinearScore(expanded, spec)
					cand := newScoredCandidate(expanded, candidate.mass+mass, score)
					expandedLeaderboard = append(expandedLeaderboard, cand)
					if cand.mass == parentMass {
						cand.score = CyclicScore(cand.peptide, spec)
						if cand.score > leaderPeptides[0].score {
							leaderPeptides = []peptideCandidate{cand}
						} else if cand.score == leaderPeptides[0].score {
							leaderPeptides = append(leaderPeptides, cand)
						}
					}
				}
			}
		}
		leaderboard = trim(expandedLeaderboard, n)
	}
	res := make([]Peptide, len(leaderPeptides))
	for i := 0; i < len(leaderPeptides); i++ {
		res[i] = leaderPeptides[i].peptide
	}
	return res
}

// ConvolutionCyclopeptideSequencing first finds the most m frequent masses
// in the spectrum's spectral convolution and then runs
// LeaderboardCyclopeptideSequencing with these masses only keeping n leaders.
func ConvolutionCyclopeptideSequencing(spectrum []int, m, n int) []Peptide {
	massCount := map[int]int{}
	for _, mass := range SpectralConvolution(spectrum) {
		if mass >= 57 && mass <= 200 {
			massCount[mass] = massCount[mass] + 1
		}
	}
	counts := make([]int, len(massCount))
	i := 0
	for _, v := range massCount {
		counts[i] = v
		i++
	}
	sort.Ints(counts)
	minCount := counts[0]
	if len(counts) > m {
		minCount = counts[len(counts)-m]
	}
	masses := []int{}
	for k, v := range massCount {
		if v >= minCount {
			masses = append(masses, k)
		}
	}
	return LeaderboardCyclopeptideSequencing(spectrum, n, masses)
}

func trim(leaderboard []peptideCandidate, numberToKeep int) []peptideCandidate {
	if numberToKeep >= len(leaderboard) {
		return leaderboard
	}
	scores := make([]int, len(leaderboard))
	for i := 0; i < len(leaderboard); i++ {
		scores[i] = leaderboard[i].score
	}
	sort.Ints(scores)
	minScoreToKeep := scores[len(scores)-numberToKeep]

	var res []peptideCandidate
	for _, cand := range leaderboard {
		if cand.score >= minScoreToKeep {
			res = append(res, cand)
		}
	}
	return res
}

func spectrumToMap(spectrum []int) map[int]int {
	spec := map[int]int{}
	for _, v := range spectrum {
		spec[v] = spec[v] + 1
	}
	return spec
}

func consistent(peptide Peptide, spectrum map[int]int) bool {
	lSpec := LinearSpectrum(peptide)
	counts := map[int]int{}
	for _, v := range lSpec {
		counts[v] = counts[v] + 1
	}
	for k, v := range counts {
		if v > spectrum[k] {
			return false
		}
	}
	return true
}
