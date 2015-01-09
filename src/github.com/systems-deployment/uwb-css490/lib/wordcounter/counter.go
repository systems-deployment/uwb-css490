// Copyright 2015 Morris Bernstein
package wordcounter

import (
	"sort"
)

type wordCounter struct {
	Words  []string
	Counts map[string]int
}

func New() *wordCounter {
	return &wordCounter{Counts: make(map[string]int)}
}

func (w *wordCounter) Len() int           { return len(w.Words) }
func (w *wordCounter) Swap(i, j int)      { w.Words[i], w.Words[j] = w.Words[j], w.Words[i] }
func (w *wordCounter) Less(i, j int) bool { return w.Counts[w.Words[i]] < w.Counts[w.Words[j]] }

func (w *wordCounter) Sort() {
	sort.Sort(w)
}
