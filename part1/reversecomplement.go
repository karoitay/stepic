package part1

var rev = map[byte]byte{
	'A': 'T',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}

// ReverseComplement returns the reverse complement of text.
func ReverseComplement(text string) string {
	l := len(text)
	var res []byte
	for i := l - 1; i >= 0; i-- {
		res = append(res, rev[text[i]])
	}
	return string(res)
}
