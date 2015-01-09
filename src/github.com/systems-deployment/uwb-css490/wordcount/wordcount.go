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
	// Map from word to count
	counts := make(map[string]int)

	// For each line in standard input, split the line by punctuation
	// (any non-letter character) and for each word, increment its
	// count.
	//
	// Note how the code takes advantage of the default zero-value for
	// the map element.
	//
	// Note the idiomatic use of underscore to discard unused results
	// of the Readline function and range operator.
	buf := bufio.NewReader(os.Stdin)
	for line, _, err := buf.ReadLine(); err == nil; line, _, err = buf.ReadLine() {
		words := punct.Split(string(line), -1)
		for _, w := range words {
			counts[w]++
		}
	}

	// Stepping through a map: order of elements is indeterminate.
	// Run this program several times with the same input and note
	// different output.
	for w, c := range counts {
		fmt.Printf("%d\t%s\n", c, w)
	}
}
