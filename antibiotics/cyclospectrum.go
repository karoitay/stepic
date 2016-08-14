package antibiotics

import "sort"

func peptideMassFunc() func(string) int {
	cache := map[string]int{}
	return func(peptide string) int {
		mass := 0
		if m, ok := cache[peptide[1:]]; ok {
			mass = integerMassTable[peptide[0]] + m
		} else {
			for i := 0; i < len(peptide); i++ {
				mass = mass + integerMassTable[peptide[i]]
			}
		}
		cache[peptide] = mass
		return mass
	}
}

// Cyclospectrum returns the theoretical spectrum of a peptide.
func Cyclospectrum(peptide string) []int {
	spectrum := []int{0}
	peptideMass := peptideMassFunc()
	doubledPeptide := peptide + peptide
	for i := 1; i < len(peptide); i++ {
		for j := 0; j < len(peptide); j++ {
			sub := doubledPeptide[j : j+i]
			m := peptideMass(sub)
			spectrum = append(spectrum, m)
		}
	}
	spectrum = append(spectrum, peptideMass(peptide))
	sort.Ints(spectrum)
	return spectrum
}
