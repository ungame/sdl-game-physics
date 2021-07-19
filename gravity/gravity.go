package gravity

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"sdl-game-physics/sdlutil"
	"time"
)

const (
	WINDOW_WIDTH  = 500
	WINDOW_HEIGHT = 500
)

func Run() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	sdlutil.HandleError(err)
	defer sdl.Quit()

	var g float64 = 1 //sdl.STANDARD_GRAVITY

	window, err := sdl.CreateWindow(
		fmt.Sprintf("Gravity %f", g),
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		sdl.WINDOW_ALLOW_HIGHDPI)
	sdlutil.HandleError(err)
	defer sdlutil.HandleDestroy(window)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	sdlutil.HandleError(err)
	defer sdlutil.HandleDestroy(renderer)

	ball := sdl.Rect{X: WINDOW_WIDTH/2 - 50, Y: -2000, W: 100, H: 100}
	var velocity int32 = 0

	start := time.Now()
	var isGround bool

GameLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.GetType() {
			case sdl.QUIT:
				break GameLoop
			}
		}

		err = renderer.SetDrawColor(255, 255, 255, 255)
		sdlutil.HandleError(err)
		err = renderer.Clear()
		sdlutil.HandleError(err)

		velocity += int32(g)
		ball.Y += velocity

		if ball.Y+ball.H > WINDOW_HEIGHT {
			ball.Y = WINDOW_HEIGHT - ball.H
			if !isGround {
				fmt.Printf("Elapsed=%v, Velocity=%d\n", time.Since(start).Seconds(), velocity)
			}
			velocity = 0
			isGround = true
		}

		err = renderer.SetDrawColor(255, 0, 0, 255)
		sdlutil.HandleError(err)
		err = renderer.FillRect(&ball)
		sdlutil.HandleError(err)

		renderer.Present()
	}

}
