package antibiotics

import (
	"reflect"
	"sort"
	"testing"
)

func TestSpectralConvolution(t *testing.T) {
	spectrum := []int{0, 137, 186, 323}

	convolution := SpectralConvolution(spectrum)

	expectedConvolution := []int{49, 137, 137, 186, 186, 323}
	sort.Ints(convolution)

	if !reflect.DeepEqual(convolution, expectedConvolution) {
		t.Error("for", spectrum, "expected", expectedConvolution, "got", convolution)
	}
}
