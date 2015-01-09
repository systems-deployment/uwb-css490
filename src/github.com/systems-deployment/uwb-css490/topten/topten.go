// topten reads a text from standard input and outputs the
// frequency of each words in the text.
//
// Output format is count word so it can be easily compared with
// output from the uniq -c command
//
// Copyright 2015 Morris Bernstein
package main

import (
	"fmt"
	"github.com/systems-deployment/uwb-css490/lib/topten"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		topten.TopTen(os.Stdin, os.Stdout)
		return
	}

	for _, name := range os.Args[1:] {
		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "can't open %s: %s\n", name, err)
			continue
		}
		fmt.Printf("\nTop Ten Words for %s\n", name)
		topten.TopTen(f, os.Stdout)
	}

}
