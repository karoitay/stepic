package antibiotics

import "testing"

func TestRNAToProteing(t *testing.T) {
	rna := "AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA"
	expectedProtein := "MAMAPRTEINSTRING"

	protein := RNAToProtein(rna)
	if protein != expectedProtein {
		t.Error("for", rna, "expected", expectedProtein, "got", protein)
	}
}
