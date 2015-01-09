// topten reads a text from standard input and outputs the
// frequency of each words in the text.
//
// Output format is count word so it can be easily compared with
// output from the uniq -c command
//
// Copyright 2015 Morris Bernstein
package topten

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
)

const (
	punctuationPattern = "[^a-zA-Z]+"
)

var (
	// Ensure that a syntax error in the regular expression will cause
	// a panic at program intialization.  Thread-safety of regexps is
	// unknown, so to be conservative, each instance of TopTen
	// constructs its own regexp.
	punct = regexp.MustCompile(punctuationPattern)
)

type wordsByCount struct {
	words  []string
	counts map[string]int
}

func New() *wordsByCount {
	return &wordsByCount{counts: make(map[string]int)}
}

func (w *wordsByCount) Len() int           { return len(w.words) }
func (w *wordsByCount) Swap(i, j int)      { w.words[i], w.words[j] = w.words[j], w.words[i] }
func (w *wordsByCount) Less(i, j int) bool { return w.counts[w.words[i]] < w.counts[w.words[j]] }

// TopTen: read tex from standard input & output to standard output.
func TopTen(in io.Reader, out io.Writer) {
	buf := bufio.NewReader(in)

	punct, _ := regexp.Compile(punctuationPattern)

	counter := New()

	for line, _, err := buf.ReadLine(); err == nil; line, _, err = buf.ReadLine() {
		words := punct.Split(string(line), -1)
		for _, w := range words {
			if w != "" {
				counter.counts[w]++
			}
		}
	}

	for w, _ := range counter.counts {
		counter.words = append(counter.words, w)
	}
	sort.Sort(counter)

	n := len(counter.words)
	m := 10
	if n < 10 {
		m = n
	}
	for i := m; i > 0; i-- {
		fmt.Printf("%d\t%s\n", i, counter.words[n-i])
	}
}
