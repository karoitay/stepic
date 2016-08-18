package antibiotics

// SpectralConvolution returns the spectral convolution of a spectrum.
func SpectralConvolution(spectrum []int) []int {
	var convolution []int
	for i := 0; i < len(spectrum)-1; i++ {
		for j := i + 1; j < len(spectrum); j++ {
			val := spectrum[j] - spectrum[i]
			if val > 0 {
				convolution = append(convolution, val)
			}
		}
	}
	return convolution
}
