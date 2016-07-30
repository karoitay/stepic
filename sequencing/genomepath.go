package sequencing

import "bytes"

// GenomePath constructs the genome string out of its constructing kmers.
func GenomePath(kmers []string) string {
	var buffer bytes.Buffer
	buffer.WriteString(kmers[0])
	for i := 1; i < len(kmers); i++ {
		s := kmers[i]
		buffer.WriteByte(s[len(s)-1])
	}
	return buffer.String()
}
