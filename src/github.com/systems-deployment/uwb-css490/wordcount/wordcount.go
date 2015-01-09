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
	punct  = regexp.MustCompile("[^a-zA-Z]+")
	counts = make(map[string]int)
)

type WordsByCount []string

func (w WordsByCount) Len() int           { return len(w) }
func (w WordsByCount) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w WordsByCount) Less(i, j int) bool { return counts[w[i]] < counts[w[j]] }

func main() {

	buf := bufio.NewReader(os.Stdin)
	for line, _, err := buf.ReadLine(); err == nil; line, _, err = buf.ReadLine() {
		words := punct.Split(string(line), -1)
		for _, w := range words {
			if w != "" {
				counts[w]++
			}
		}
	}

	words := make([]string, 0)
	for w, _ := range counts {
		words = append(words, w)
	}
	sort.Sort(WordsByCount(words))

	for _, w := range words {
		fmt.Printf("%d\t%s\n", counts[w], w)
	}
}
