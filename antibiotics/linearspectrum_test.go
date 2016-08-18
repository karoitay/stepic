package antibiotics

import (
	"reflect"
	"testing"
)

func TestLinearSpectrum(t *testing.T) {
	peptide := NewPeptideFromString("NQEL")

	spectrum := LinearSpectrum(peptide)

	expectedSpectrum := []int{0, 113, 114, 128, 129, 242, 242, 257, 370, 371, 484}
	if !reflect.DeepEqual(spectrum, expectedSpectrum) {
		t.Error("for", peptide, "expected", expectedSpectrum, "got", spectrum)
	}
}
