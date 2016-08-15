package antibiotics

// PeptideCountForMass returns the number of peptides that has a given mass.
func PeptideCountForMass(mass int) int {
	cache := map[int]int{}
	return cachedPeptideCountForMass(cache, mass)
}

func cachedPeptideCountForMass(cache map[int]int, m int) int {
	if m < 0 {
		return 0
	} else if m == 0 {
		return 1
	} else if c, ok := cache[m]; ok {
		return c
	}
	c := 0
	for p, mass := range integerMassTable {
		if p != 'L' && p != 'Q' {
			c = c + cachedPeptideCountForMass(cache, m-mass)
		}
	}
	cache[m] = c
	return c
}
