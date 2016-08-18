package antibiotics

import (
	"reflect"
	"sort"
	"testing"
)

func TestCyclopeptideSequencing(t *testing.T) {
	spectrum := []int{0, 113, 128, 186, 241, 299, 314, 427}

	var output []string
	for _, peptide := range CyclopeptideSequencing(spectrum, proteinogenicMasses) {
		output = append(output, peptide.ToString())
	}

	sort.Strings(output)
	expectedOutput := []string{"113-128-186", "113-186-128", "128-113-186", "128-186-113", "186-113-128", "186-128-113"}

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Error("for", spectrum, "expected", expectedOutput, "got", output)
	}
}

func TestLeaderboardCyclopeptideSequencing(t *testing.T) {
	n := 10
	spectrum := []int{0, 71, 113, 129, 147, 200, 218, 260, 313, 331, 347, 389, 460}

	peptide := LeaderboardCyclopeptideSequencing(spectrum, n, proteinogenicMasses)[0]
	peptideScore := CyclicScore(peptide, spectrumToMap(spectrum))

	expectedPeptide := NewPeptide([]int{113, 147, 71, 129})
	expectedPeptideScore := CyclicScore(expectedPeptide, spectrumToMap(spectrum))

	if peptideScore < expectedPeptideScore {
		t.Error("for", n, ",", spectrum,
			"expected", expectedPeptide.ToString(), "with score", expectedPeptideScore,
			"got", peptide.ToString(), "with score", peptideScore)
	}
}

func TestConvolutionCyclopeptideSequencing(t *testing.T) {
	m, n := 20, 60
	spectrum := []int{57, 57, 71, 99, 129, 137, 170, 186, 194, 208, 228, 265, 285, 299, 307, 323, 356, 364, 394, 422, 493}

	peptides := ConvolutionCyclopeptideSequencing(spectrum, m, n)

	expectedPeptide := NewPeptide([]int{99, 71, 137, 57, 72, 57})

	for _, peptide := range peptides {
		if peptide.IsCyclicEqualTo(expectedPeptide) {
			return
		}
	}

	t.Error("for", m, n, spectrum, "failed to find", expectedPeptide, "in", peptides)
}
