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
	"sort"
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

	words := make([]string, 0)
	for w, _ := range counts {
		words = append(words, w)
	}
	sort.Strings(words)

	for _, w := range words {
		fmt.Printf("%d\t%s\n", counts[w], w)
	}
}
