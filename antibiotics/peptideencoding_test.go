package antibiotics

import (
	"reflect"
	"testing"
)

func TestPeptideEncoding(t *testing.T) {
	dna := "ATGGCCATGGCCCCCAGAACTGAGATCAATAGTACCCGTATTAACGGGTGA"
	peptide := "MA"

	actual := PeptideEncoding(dna, peptide)

	expected := []string{"ATGGCC", "GGCCAT", "ATGGCC"}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("for", dna, "expected", expected, "got", actual)
	}
}
