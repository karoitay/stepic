package antibiotics

// CyclicScore returns the number of matches between peptide's cyclic spectrum
// and a given spectrum.
func CyclicScore(peptide string, spectrum map[int]int) int {
	return score(peptide, spectrum, Cyclospectrum)
}

// LinearScore returns the number of matches between peptide's linear spectrum
// and a given spectrum.
func LinearScore(peptide string, spectrum map[int]int) int {
	return score(peptide, spectrum, LinearSpectrum)
}

func score(peptide string, spectrum map[int]int, spectrumFunc func(string) []int) int {
	actualSpectrum := map[int]int{}
	for _, m := range spectrumFunc(peptide) {
		actualSpectrum[m] = actualSpectrum[m] + 1
	}
	c := 0
	for k, v := range actualSpectrum {
		if spectrum[k] < v {
			c = c + spectrum[k]
		} else {
			c = c + v
		}
	}
	return c
}
