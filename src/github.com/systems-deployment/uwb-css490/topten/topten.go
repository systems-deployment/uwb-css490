// topten reads a text from standard input and outputs the
// frequency of each words in the text.
//
// Output format is count word so it can be easily compared with
// output from the uniq -c command
//
// Copyright 2015 Morris Bernstein
package main

import (
	"github.com/systems-deployment/uwb-css490/lib/topten"
)

func main() {
	topten.TopTen()
}
