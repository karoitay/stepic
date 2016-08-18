package antibiotics

import (
	"strconv"
	"strings"
)

// Peptide represents a peptide as a slice of masses.
type Peptide struct {
	Masses []int
}

// NewPeptide creates a new Peptide from the given masses slice.
func NewPeptide(peptide []int) Peptide {
	return Peptide{peptide}
}

// NewPeptideFromString uses the integer mass table in order to create a Peptide
// from a peptide string.
func NewPeptideFromString(peptide string) Peptide {
	p := make([]int, len(peptide))
	for i := 0; i < len(peptide); i++ {
		p[i] = proteinogenicMassesTable[peptide[i]]
	}
	return Peptide{p}
}

// Expand creates a new Peptide with an additional mass over the given Peptide.
func (p Peptide) Expand(mass int) Peptide {
	cpy := make([]int, len(p.Masses)+1)
	copy(cpy, p.Masses)
	cpy[len(p.Masses)] = mass
	return Peptide{cpy}
}

// ToString returns a string of masses separated by "-".
func (p Peptide) ToString() string {
	massesStrings := make([]string, len(p.Masses))
	for i := 0; i < len(p.Masses); i++ {
		massesStrings[i] = strconv.Itoa(p.Masses[i])
	}
	return strings.Join(massesStrings, "-")
}

// IsCyclicEqualTo returns true if another cyclic peptide is equal to the
// the receiver.
func (p Peptide) IsCyclicEqualTo(other Peptide) bool {
	l := len(p.Masses)
	if l != len(other.Masses) {
		return false
	}
	for i := 0; i < l; i++ {
		eq := true
		for j := 0; j < l; j++ {
			if p.Masses[j] != other.Masses[(i+j)%l] {
				eq = false
			}
		}
		if eq {
			return true
		}
	}
	return false
}
