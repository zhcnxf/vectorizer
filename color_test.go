package main

import (
	"image/color"
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	c0, c1 := color.RGBA64{0xffff, 0xffff, 0xffff, 0xffff}, color.RGBA64{0xfffe, 0xfffe, 0xfffe, 0xffff}
	md := ManhattanDistance(c0, c1)
	if md != 3 {
		t.Errorf("Expecting %v, got %v", 3, md)
	}
}
