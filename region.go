package main

import (
	//	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Region struct {
	Color  color.Color
	Points []Point
	W, H   int
}

func (r *Region) ColorModel() color.Model {
	return color.AlphaModel
}

func (r *Region) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.W, r.H)
}

func (r *Region) At(x, y int) color.Color {
	for i, p := range r.Points {
		if p.X == x && p.Y == y && isEdge(i, r.Points) {
			return color.Alpha{0xff}
		}
	}
	return color.Alpha{0}
}

func isEdge(i int, points []Point) bool {
	p, count := points[i], 0
	for _, pt := range points {
		if p.Left() == pt || p.Right() == pt || p.Top() == pt || p.Bottom() == pt {
			count++
		}
	}
	return count < 4
}

type imape struct {
	width, height int
	pixels        []color.Color
}

func (imp *imape) at(x, y int) color.Color {
	return imp.pixels[(y*imp.width + x)]
}

func newimape(img image.Image) *imape {
	bounds := img.Bounds()
	pixels := make([]color.Color, 0, bounds.Dx()*bounds.Dy())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixels = append(pixels, img.At(x, y))
		}
	}
	return &imape{bounds.Dx(), bounds.Dy(), pixels}
}

func Print(img image.Image) {
	var (
		imp = newimape(img)
		//		regions = make([]Region, 0, 16)
	)
	//	for y := 0; y < imp.height; y++ {
	//		for x := 0; x < imp.width; x++ {
	//			color := imp.at(x, y)
	//			if color != nil {
	//				r := new(Region)
	//				imp.etch(r, &Point{x, y})
	//				regions = append(regions, *r)
	//			}
	//		}
	//	}

	m := image.NewRGBA(img.Bounds())
	r := new(Region)
	imp.etch(r, &Point{0, 0})
	draw.DrawMask(m, m.Bounds(), &image.Uniform{color.RGBA{0xcc, 0x40, 0x00, 0xff}}, image.ZP, r, image.ZP, draw.Over)
	png.Encode(os.Stdout, m)
}

func (imp imape) etch(r *Region, p *Point) {
	if r.Color == nil {
		r.Color = imp.at(p.X, p.Y)
		r.Points = []Point{*p}
		r.W = imp.width
		r.H = imp.height
	} else {
		r.Points = append(r.Points, *p)
	}
	imp.pixels[p.Y*imp.width+p.X] = nil
	for _, neighbor := range []Point{p.Left(), p.Top(), p.Right(), p.Bottom()} {
		if neighbor.X < 0 || neighbor.X >= imp.width || neighbor.Y < 0 || neighbor.Y >= imp.height {
			continue
		}
		color := imp.at(neighbor.X, neighbor.Y)
		if color != nil && ManhattanDistance(color, r.Color) <= 16 {
			imp.etch(r, &neighbor)
		}
	}
}
