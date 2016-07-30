package sequencing

import "testing"

func TestGenomePath(t *testing.T) {
	path := []string{"ACCGA", "CCGAA", "CGAAG", "GAAGC", "AAGCT"}
	genome := GenomePath(path)

	expectedGenome := "ACCGAAGCT"
	if genome != expectedGenome {
		t.Error("For", path, "expected", expectedGenome, "got", genome)
	}
}
