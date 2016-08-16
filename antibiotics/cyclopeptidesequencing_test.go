package antibiotics

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestCyclopeptideSequencing(t *testing.T) {
	spectrum := []int{0, 113, 128, 186, 241, 299, 314, 427}

	var output []string
	for _, seq := range CyclopeptideSequencing(spectrum) {
		output = append(output, strings.Join(toStringSlice(seq), "-"))
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

	peptide := LeaderboardCyclopeptideSequencing(spectrum, n)
	peptideScore := LinearScore(toStringPeptide(peptide), spectrumToMap(spectrum))

	expectedPeptide := []int{113, 147, 71, 129}
	expectedPeptideScore := LinearScore(toStringPeptide(expectedPeptide), spectrumToMap(spectrum))

	if peptideScore < expectedPeptideScore {
		t.Error("for", n, ",", spectrum, "expected", expectedPeptide, "got", peptide)
	}
}

func toStringSlice(ints []int) []string {
	strings := make([]string, len(ints))
	for i := 0; i < len(ints); i++ {
		strings[i] = strconv.Itoa(ints[i])
	}
	return strings
}
