package part1

func Skew(text string) []int {
	c := 0
	res := make([]int, len(text)+1)
	res[0] = 0
	for i := 0; i < len(text); i++ {
		if text[i] == 'C' {
			c -= 1
		} else if text[i] == 'G' {
			c += 1
		}
		res[i+1] = c
	}
	return res
}

func MinimumSkew(text string) []int {
	min := 0
	c := 0
	res := []int{0}
	for i := 0; i < len(text); i++ {
		if text[i] == 'C' {
			c -= 1
		} else if text[i] == 'G' {
			c += 1
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
