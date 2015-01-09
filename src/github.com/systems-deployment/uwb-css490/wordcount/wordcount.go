// wordcount reads a text from standard input and outputs the
// frequency of each words in the text.
//
// Output format is count word so it can be easily compared with
// output from the uniq -c command
//
// Copyright 2015 Morris Bernstein
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var (
	punct = regexp.MustCompile("[^a-zA-Z]+")
)

func main() {
	counts := make(map[string]int)
	buf := bufio.NewReader(os.Stdin)
	for line, _, err := buf.ReadLine(); err == nil; line, _, err = buf.ReadLine() {
		words := punct.Split(string(line), -1)
		for _, w := range words {
			counts[w]++
		}
	}

	for w, c := range counts {
		fmt.Printf("%d\t%s\n", c, w)
	}
}
