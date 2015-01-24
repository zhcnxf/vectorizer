package main

import (
	"image"
)

type Point image.Point

func (p *Point) Top() Point {
	return Point{p.X, p.Y - 1}
}

func (p *Point) Right() Point {
	return Point{p.X + 1, p.Y}
}

func (p *Point) Bottom() Point {
	return Point{p.X, p.Y + 1}
}

func (p *Point) Left() Point {
	return Point{p.X - 1, p.Y}
}
