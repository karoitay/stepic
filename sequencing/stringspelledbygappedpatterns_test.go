package sequencing

import "testing"

func TestStringSpelledByGappedPatterns(t *testing.T) {
	type testInput struct {
		k         int
		d         int
		readPairs []string
	}
	input := testInput{4, 2, []string{"GACC|GCGC", "ACCG|CGCC", "CCGA|GCCG", "CGAG|CCGG", "GAGC|CGGA"}}

	text := StringSpelledByGappedPatterns(input.readPairs, input.k, input.d)

	expected := "GACCGAGCGCCGGA"
	if text != expected {
		t.Error("For", input, "expected", expected, "got", text)
	}
}
