package antibiotics

import (
	"strings"

	"github.com/karoitay/stepic/part1"
)

// PeptideEncoding returns all substrings s in dna so that either s or its
// reverse complement encodes peptide.
func PeptideEncoding(dna, peptide string) []string {
	var res []string
	l := len(peptide) * 3
	for i := 0; i <= len(dna)-l; i++ {
		d := dna[i : i+l]
		rd := part1.ReverseComplement(d)
		rna := strings.Replace(d, "T", "U", -1)
		rrna := strings.Replace(rd, "T", "U", -1)
		if RNAToProtein(rna) == peptide || RNAToProtein(rrna) == peptide {
			res = append(res, d)
		}
	}
	return res
}
