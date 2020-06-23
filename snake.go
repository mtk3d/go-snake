package main

import (
	"errors"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Snake struct {
	segments  []pixel.Rect
	direction pixel.Vec
	ate       int
}

func NewSnake() *Snake {
	x, y := pixel.R(0, 0, width, height).Center().XY()
	segments := make([]pixel.Rect, 3)
	for i, _ := range segments {
		segments[i] = pixel.R(x, y-float64(i*rectEdge), x+rectEdge, y-float64((i+1)*rectEdge))
	}

	return &Snake{segments, pixel.V(0, 1), 0}
}

func (s *Snake) Up() {
	s.changeDirection(pixel.V(0, 1))
}

func (s *Snake) Down() {
	s.changeDirection(pixel.V(0, -1))
}

func (s *Snake) Right() {
	s.changeDirection(pixel.V(1, 0))
}

func (s *Snake) Left() {
	s.changeDirection(pixel.V(-1, 0))
}

func (s *Snake) changeDirection(direction pixel.Vec) {
	x, y := direction.XY()
	if !(s.direction.Eq(pixel.V(-x, -y))) {
		s.direction = direction
	}
}

func (s *Snake) nextMove() pixel.Vec {
	angle := s.direction.Angle()
	return pixel.V(rectEdge*math.Cos(angle), rectEdge*math.Sin(angle))
}

func (s *Snake) Move() error {
	newHead := s.segments[0].Moved(s.nextMove())
	x, y := newHead.Center().XY()

	if y > height {
		newHead = newHead.Moved(pixel.V(0, -height))
	} else if y < 0 {
		newHead = newHead.Moved(pixel.V(0, height))
	} else if x < 0 {
		newHead = newHead.Moved(pixel.V(width, 0))
	} else if x > width {
		newHead = newHead.Moved(pixel.V(-width, 0))
	}

	if s.selfCollided() {
		return errors.New("Game Over")
	}

	lastIndex := len(s.segments) - 1 + s.ate
	s.ate = 0
	s.segments = append([]pixel.Rect{newHead}, s.segments[:lastIndex]...) // prepend, and remove last
	return nil
}

func (s *Snake) selfCollided() bool {
	for _, segment := range s.segments[1:] {
		if segment.Center().Eq(s.GetHead().Center()) {
			return true
		}
	}
	return false
}

func (s *Snake) GetHead() pixel.Rect {
	return s.segments[0]
}

func (s *Snake) Ate() {
	s.ate = 1
}

func (s *Snake) Draw(win pixel.Target) {
	for _, seg := range s.segments {
		imd := imdraw.New(nil)
		imd.Color = snakeColor
		for _, v := range seg.Vertices() {
			imd.Push(v)
		}
		imd.Polygon(0)
		imd.Draw(win)
	}
}
