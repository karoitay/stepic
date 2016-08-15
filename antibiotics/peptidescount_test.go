package antibiotics

import "testing"

func TestPeptideCountForMass(t *testing.T) {
	testCases := map[int]int{
		1024: 14712706211,
		1307: 34544458837656,
		1414: 679487374524349,
	}

	for mass, expected := range testCases {
		c := PeptideCountForMass(mass)
		if c != expected {
			t.Error("for", mass, "expected", expected, "got", c)
		}
	}
}
