package antibiotics

import "strings"

// RNAToProtein uses the codon table to translate an RNA string into a protein.
func RNAToProtein(rna string) string {
	l := len(rna) / 3
	res := make([]string, l)
	for i := 0; i < l; i++ {
		codon := rna[3*i : 3*i+3]
		res[i] = codonTable[codon]
	}
	return strings.Join(res, "")
}
