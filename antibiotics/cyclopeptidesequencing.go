package antibiotics

import (
	"reflect"
	"sort"
)

type peptideCandidate struct {
	peptide []int
	mass    int
	score   int
}

func newCandidate(peptide []int, mass int) peptideCandidate {
	return newScoredCandidate(peptide, mass, 0)
}

func newScoredCandidate(peptide []int, mass, score int) peptideCandidate {
	return peptideCandidate{peptide, mass, score}
}

var inverseMassTable = inverse(integerMassTable)

// CyclopeptideSequencing returns the possible peptides with the given
// cyclo-spectrm.
func CyclopeptideSequencing(spectrum []int) [][]int {
	parentMass := spectrum[len(spectrum)-1]
	spec := spectrumToMap(spectrum)

	matchingPeptides := [][]int{}
	candidates := []peptideCandidate{newCandidate([]int{}, 0)}
	for len(candidates) != 0 {
		var expandedCandidates []peptideCandidate
		for _, candidate := range candidates {
			for mass := range inverseMassTable {
				cpy := make([]int, len(candidate.peptide)+1)
				copy(cpy, candidate.peptide)
				cpy[len(candidate.peptide)] = mass
				newCandidate := newCandidate(cpy, candidate.mass+mass)
				if newCandidate.mass == parentMass {
					if reflect.DeepEqual(Cyclospectrum(toStringPeptide(newCandidate.peptide)), spectrum) {
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

// LeaderboardCyclopeptideSequencing returns the highest scoring peptide.
// In every iteration only the n higest scoring (keeping ties)
// candidates will be kept then the one with the highest score will be returned.
func LeaderboardCyclopeptideSequencing(spectrum []int, n int) []int {
	parentMass := spectrum[len(spectrum)-1]
	spec := spectrumToMap(spectrum)

	leaderPeptide := newScoredCandidate([]int{}, 0, 0)
	leaderboard := []peptideCandidate{leaderPeptide}
	for len(leaderboard) > 0 {
		var expandedLeaderboard []peptideCandidate
		for _, candidate := range leaderboard {
			for mass := range inverseMassTable {
				if candidate.mass+mass <= parentMass {
					cpy := make([]int, len(candidate.peptide)+1)
					copy(cpy, candidate.peptide)
					cpy[len(candidate.peptide)] = mass
					score := LinearScore(toStringPeptide(cpy), spec)
					cand := newScoredCandidate(cpy, candidate.mass+mass, score)
					expandedLeaderboard = append(expandedLeaderboard, cand)
					if cand.mass == parentMass {
						if cand.score > leaderPeptide.score {
							leaderPeptide = cand
						}
					}
				}
			}
		}
		leaderboard = trim(expandedLeaderboard, n)
	}
	return leaderPeptide.peptide
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

func consistent(peptide []int, spectrum map[int]int) bool {
	lSpec := LinearSpectrum(toStringPeptide(peptide))
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

func toStringPeptide(peptide []int) string {
	bytes := make([]byte, len(peptide))
	for i := 0; i < len(peptide); i++ {
		bytes[i] = inverseMassTable[peptide[i]]
	}
	return string(bytes)
}

func inverse(table map[byte]int) map[int]byte {
	inverse := map[int]byte{}
	for k, v := range table {
		inverse[v] = k
	}
	return inverse
}
