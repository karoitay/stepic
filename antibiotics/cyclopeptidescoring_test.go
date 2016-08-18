package antibiotics

import "testing"

func TestCyclicScore(t *testing.T) {
	spectrum := map[int]int{0: 1, 99: 1, 113: 1, 114: 1, 128: 1, 227: 1, 257: 1, 299: 1, 355: 1, 356: 1, 370: 1, 371: 1, 484: 1}
	peptide := NewPeptideFromString("NQEL")

	score := CyclicScore(peptide, spectrum)

	expectedScore := 11
	if score != expectedScore {
		t.Error("for", peptide, ",", spectrum, "expected", expectedScore, "got", score)
	}
}

func TestLinearScore(t *testing.T) {
	spectrum := map[int]int{0: 1, 99: 1, 113: 1, 114: 1, 128: 1, 227: 1, 257: 1, 299: 1, 355: 1, 356: 1, 370: 1, 371: 1, 484: 1}
	peptide := NewPeptideFromString("NQEL")

	score := LinearScore(peptide, spectrum)

	expectedScore := 8
	if score != expectedScore {
		t.Error("for", peptide, ",", spectrum, "expected", expectedScore, "got", score)
	}
}
