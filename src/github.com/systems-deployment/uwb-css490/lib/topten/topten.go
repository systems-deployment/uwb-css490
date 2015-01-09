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
	"github.com/systems-deployment/uwb-css490/lib/wordcounter"
	"io"
	"regexp"
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

// TopTen: read tex from standard input & output to standard output.
func TopTen(in io.Reader, out io.Writer) {
	buf := bufio.NewReader(in)

	punct, _ := regexp.Compile(punctuationPattern)

	counter := wordcounter.New()

	for line, _, err := buf.ReadLine(); err == nil; line, _, err = buf.ReadLine() {
		words := punct.Split(string(line), -1)
		for _, w := range words {
			if w != "" {
				counter.Counts[w]++
			}
		}
	}

	for w, _ := range counter.Counts {
		counter.Words = append(counter.Words, w)
	}
	counter.Sort()

	n := len(counter.Words)
	m := 10
	if n < 10 {
		m = n
	}
	for i := m; i > 0; i-- {
		fmt.Printf("%d\t%s\n", i, counter.Words[n-i])
	}
}
