package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/karoitay/stepic/part1"
)

var inFile = flag.String("in", "in.txt", "Input file name")

func main() {
	flag.Parse()
	file, err := os.Open(*inFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	text := nextLine(r)
	k, l, t := 9, 500, 3

	res := part1.FindClumps(text, k, l, t)
	fmt.Println(len(res))
}

func nextLine(r *bufio.Reader) string {
	isPrefix := true
	var res []string
	for isPrefix {
		bytes, prefix, err := r.ReadLine()
		if err != nil {
			panic(err)
		}
		res = append(res, string(bytes))
		isPrefix = prefix
	}
	return strings.Join(res, "")
}
