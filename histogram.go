package main

import "sort"

type Histogram map[string]int

func (h Histogram) Add(s string) {
	h[s] += 1
}

func (h Histogram) Total() int {
	total := 0
	for _, v := range h {
		total += v
	}
	return total
}

func (h Histogram) Top(n int) []Pair {
	return h.toList()[:n]
}

func (h Histogram) toList() []Pair {
	pl := make(PairList, len(h))
	i := 0
	for k, v := range h {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	string
	int
}

type PairList []Pair

func (pl PairList) Len() int           { return len(pl) }
func (pl PairList) Less(i, j int) bool { return pl[i].int < pl[j].int }
func (pl PairList) Swap(i, j int)      { pl[i], pl[j] = pl[j], pl[i] }
