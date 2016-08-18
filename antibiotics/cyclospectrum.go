package antibiotics

import "sort"

// Cyclospectrum returns the theoretical spectrum of a peptide.
func Cyclospectrum(peptide Peptide) []int {
	prefixMass := make([]int, len(peptide.Masses)+1)
	prefixMass[0] = 0
	for i := 0; i < len(peptide.Masses); i++ {
		prefixMass[i+1] = prefixMass[i] + peptide.Masses[i]
	}
	peptideMass := prefixMass[len(prefixMass)-1]

	spectrum := []int{0}
	for i := 0; i < len(prefixMass)-1; i++ {
		for j := i + 1; j < len(prefixMass); j++ {
			mass := prefixMass[j] - prefixMass[i]
			spectrum = append(spectrum, mass)
			if i > 0 && j < len(prefixMass)-1 {
				spectrum = append(spectrum, peptideMass-mass)
			}
		}
	}

	sort.Ints(spectrum)
	return spectrum
}
