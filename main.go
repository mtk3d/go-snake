package main

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	rectEdge = 20
	width    = 800
	height   = 600
)

var (
	backgroundColor = colornames.Dimgrey
	foodColor       = colornames.Crimson
	snakeColor      = colornames.Ivory
	randSource      = rand.NewSource(time.Now().UnixNano())
	random          = rand.New(randSource)
	win             pixelgl.Window
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Snake. Golang is Awesome!",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	snake := NewSnake()
	food := NewFood()

	tick := time.Tick(80 * time.Millisecond)

	for !win.Closed() {
		if win.Pressed(pixelgl.KeyLeft) {
			snake.Left()
		}
		if win.Pressed(pixelgl.KeyRight) {
			snake.Right()
		}
		if win.Pressed(pixelgl.KeyDown) {
			snake.Down()
		}
		if win.Pressed(pixelgl.KeyUp) {
			snake.Up()
		}
		if win.Pressed(pixelgl.KeyR) {
			snake = NewSnake()
		}

		select {
		case <-tick:
			win.Clear(backgroundColor)
			snake.Move()
			if food.HasBeenEaten(snake) {
				snake.Ate()
				food = NewFood()
			}
			food.Draw(win)
			snake.Draw(win)
			win.Update()
		}
	}
}

func main() {
	pixelgl.Run(run)
}
