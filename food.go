package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Food struct {
	rect pixel.Rect
}

func NewFood() *Food {
	x := float64(random.Intn(width/rectEdge) * rectEdge)
	y := float64(random.Intn(height/rectEdge) * rectEdge)
	rect := pixel.R(x, y, x+rectEdge, y+rectEdge)
	return &Food{rect}
}

func (f *Food) HasBeenEaten(s *Snake) bool {
	return f.rect.Center().Eq(s.GetHead().Center())
}

func (f *Food) Draw(win pixel.Target) {
	imd := imdraw.New(nil)
	imd.Color = foodColor
	for _, v := range f.rect.Vertices() {
		imd.Push(v)
	}
	imd.Polygon(0)
	imd.Draw(win)
}
