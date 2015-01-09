// Copyright 2015 Morris Bernstein
package wordcounter

import (
	"sort"
)

type wordCounter struct {
	words  []string
	counts map[string]int
}

func New() *wordCounter {
	return &wordCounter{counts: make(map[string]int)}
}

func (w *wordCounter) Len() int           { return len(w.words) }
func (w *wordCounter) Swap(i, j int)      { w.words[i], w.words[j] = w.words[j], w.words[i] }
func (w *wordCounter) Less(i, j int) bool { return w.counts[w.words[i]] < w.counts[w.words[j]] }

func (this *wordCounter) Incr(word string) {
	this.counts[word]++
}

func (this *wordCounter) Sort() {
	this.words = make([]string, 0)
	for word, _ := range this.counts {
		this.words = append(this.words, word)
	}
	sort.Sort(this)
}

func (counter *wordCounter) Get(n int) (string, int) {
	word := counter.words[n]
	return word, counter.counts[word]
}
