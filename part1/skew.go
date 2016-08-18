package part1

// Skew returns the skew of text.
func Skew(text string) []int {
	c := 0
	res := make([]int, len(text)+1)
	res[0] = 0
	for i := 0; i < len(text); i++ {
		if text[i] == 'C' {
			c--
		} else if text[i] == 'G' {
			c++
		}
		res[i+1] = c
	}
	return res
}

// MinimumSkew return the indices of all elements with the minimum skew value.
func MinimumSkew(text string) []int {
	min := 0
	c := 0
	res := []int{0}
	for i := 0; i < len(text); i++ {
		if text[i] == 'C' {
			c--
		} else if text[i] == 'G' {
			c++
		}
		if c < min {
			min = c
			res = nil
		}
		if c <= min {
			res = append(res, i+1)
		}
	}
	return res
}
