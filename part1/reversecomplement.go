package part1

var rev = map[byte]byte{
	'A': 'T',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}

func ReverseComplement(text string) string {
	l := len(text)
	res := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		res = append(res, rev[text[i]])
	}
	return string(res)
}
