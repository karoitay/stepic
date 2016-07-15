package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

type Fasta struct {
	ID, Sequence string
}

type FastaParser struct {
	file   string
	f      *os.File
	out    chan Fasta
	closed bool
	err    error
}

func NewFastaParser(file string) *FastaParser {
	parser := FastaParser{}
	parser.file = file
	parser.closed = true
	return &parser
}

func (p *FastaParser) Close() {
	if !p.closed {
		p.closed = true
		p.f.Close()
		p.err = nil
		close(p.out)
	}
}

func (p *FastaParser) Err() error {
	return p.err
}

func (p *FastaParser) Parse() (chan Fasta, error) {
	if !p.closed {
		return nil, errors.New("FastaParser: Parse should not be called on an opened parser")
	}
	f, err := os.Open(p.file)
	if err != nil {
		return nil, err
	}
	p.closed = false
	p.f = f
	p.out = make(chan Fasta)

	r := bufio.NewReader(f)
	c, err := r.ReadByte()
	if err == nil && c != '>' {
		err = errors.New("FastaParser: invalid format, first byte is expected to be '>'")
	}
	if err != nil {
		p.Close()
		return nil, err
	}

	go p.parse(r)
	return p.out, nil
}

func (p *FastaParser) parse(r *bufio.Reader) {
	for {
		bs, err := r.ReadBytes('>')
		if err == nil || err == io.EOF {
			lines := strings.Split(string(bs), "\n")
			for i := 0; i < len(lines); i++ {
				lines[i] = strings.Trim(lines[i], " \n\r>")
			}
			p.out <- Fasta{
				ID:       lines[0],
				Sequence: strings.Join(lines[1:], ""),
			}
		}
		if err != nil {
			p.Close()
			p.err = err
			break
		}
	}
}
