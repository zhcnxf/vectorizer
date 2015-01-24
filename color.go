package main

import (
	"image/color"
	"math"
	"sort"
)

func ManhattanDistance(c0, c1 color.Color) uint64 {
	var (
		r0, g0, b0, a0 = c0.RGBA()
		r1, g1, b1, a1 = c1.RGBA()
	)
	return uint64(math.Abs(float64(r0)-float64(r1)) + math.Abs(float64(g0)-float64(g1)) + math.Abs(float64(b0)-float64(b1)) + math.Abs(float64(a0)-float64(a1)))
}

type colorSorter []color.Color

func (cs colorSorter) Len() int {
	return len(cs)
}

func (cs colorSorter) Less(i, j int) bool {
	ir, ig, ib, ia := cs[i].RGBA()
	jr, jg, jb, ja := cs[j].RGBA()
	return ir < jr || ig < jg || ib < jb || ia < ja
}

func (cs colorSorter) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

func SortColors(colors []color.Color) {
	sort.Sort(colorSorter(colors))
}
