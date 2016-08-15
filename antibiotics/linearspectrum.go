package antibiotics

import "sort"

// LinearSpectrum computes the linear spectrum of a peptide.
func LinearSpectrum(peptide string) []int {
	prefixMass := make([]int, len(peptide)+1)
	prefixMass[0] = 0
	for i := 0; i < len(peptide); i++ {
		prefixMass[i+1] = prefixMass[i] + integerMassTable[peptide[i]]
	}

	spectrum := []int{0}
	for i := 0; i < len(prefixMass)-1; i++ {
		for j := i + 1; j < len(prefixMass); j++ {
			spectrum = append(spectrum, prefixMass[j]-prefixMass[i])
		}
	}

	sort.Ints(spectrum)
	return spectrum
}
