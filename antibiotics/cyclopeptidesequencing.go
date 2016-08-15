package antibiotics

import (
	"fmt"
	"reflect"
)

type peptideCandidate struct {
	peptide []int
	mass    int
}

var inverseMassTable = inverse(integerMassTable)

// CyclopeptideSequencing returns the possible peptides with the given
// cyclo-spectrm.
func CyclopeptideSequencing(specturm []int) [][]int {
	parentMass := specturm[len(specturm)-1]
	spec := map[int]int{}
	for _, v := range specturm {
		spec[v] = spec[v] + 1
	}

	matchingPeptides := [][]int{}
	candidates := []peptideCandidate{peptideCandidate{[]int{}, 0}}
	candLen := 0
	for len(candidates) != 0 {
		var expandedCandidates []peptideCandidate
		for _, candidate := range candidates {
			if len(candidate.peptide) != candLen {
				fmt.Println("len is expected to be", candLen, "but is", len(candidate.peptide))
				panic("")
			}
			for mass := range inverseMassTable {
				cpy := make([]int, len(candidate.peptide)+1)
				copy(cpy, candidate.peptide)
				cpy[len(candidate.peptide)] = mass
				newCandidate := peptideCandidate{cpy, candidate.mass + mass}
				if newCandidate.mass == parentMass {
					if reflect.DeepEqual(Cyclospectrum(toStringPeptide(newCandidate.peptide)), specturm) {
						matchingPeptides = append(matchingPeptides, newCandidate.peptide)
					}
				} else if newCandidate.mass < parentMass && consistent(newCandidate.peptide, spec) {
					expandedCandidates = append(expandedCandidates, newCandidate)
				}
			}
		}
		candLen = candLen + 1
		candidates = expandedCandidates
	}
	return matchingPeptides
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
