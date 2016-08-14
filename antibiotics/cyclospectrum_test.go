package antibiotics

import (
	"reflect"
	"testing"
)

func TestCyclospectrum(t *testing.T) {
	peptide := "LEQN"

	actualSpectrum := Cyclospectrum(peptide)

	expectedSpectrum := []int{0, 113, 114, 128, 129, 227, 242, 242, 257, 355, 356, 370, 371, 484}
	if !reflect.DeepEqual(actualSpectrum, expectedSpectrum) {
		t.Error("for", peptide, "expected", expectedSpectrum, "got", actualSpectrum)
	}
}
