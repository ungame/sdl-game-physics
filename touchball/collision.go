package touchball

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"sdl-game-physics/sdlutil"
)

const (
	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 720

	GRAVITY = 1
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
		sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_RESIZABLE)
	sdlutil.HandleError(err)
	defer sdlutil.HandleDestroy(window)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	sdlutil.HandleError(err)
	defer sdlutil.HandleDestroy(renderer)

	obj := sdl.Rect{X: 200, Y: 100, W: 50, H: 50}
	velocity := sdl.Point{X: 0, Y: -10}
	touch := sdl.Point{}
	var touched bool

GameLoop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.GetType() {
			case sdl.QUIT:
				break GameLoop

			case sdl.MOUSEBUTTONDOWN:
				touched = true
				mouse := event.(*sdl.MouseButtonEvent)
				touch.X = mouse.X
				touch.Y = mouse.Y
				fmt.Printf("TOUCH: %+v\n", touch)

			case sdl.MOUSEBUTTONUP:
				touched = false
				touch.X = 0
				touch.Y = 0
			}
		}

		err = renderer.SetDrawColor(255, 255, 255, 255)
		sdlutil.HandleError(err)
		err = renderer.Clear()
		sdlutil.HandleError(err)

		update(&obj, &velocity, &touch, touched)
		draw(renderer, obj)

		renderer.Present()
	}

}

func update(obj *sdl.Rect, velocity *sdl.Point, touch *sdl.Point, touched bool) {
	obj.X += velocity.X
	obj.Y += velocity.Y

	if touched {

		ab := sdl.Point{X: touch.X - obj.X, Y: touch.Y - obj.Y}
		NormalizeVector(&ab)
		ab.X *= 3
		ab.Y *= 3
		velocity.X += ab.X
		velocity.Y += ab.Y
	}

	if obj.X < 0 {
		obj.X = 0
		velocity.X *= -1
	}

	if obj.X+obj.W >= WINDOW_WIDTH {
		obj.X = WINDOW_WIDTH - obj.W
		velocity.X *= -1
	}

	if obj.Y+obj.H >= WINDOW_HEIGHT {
		obj.Y = WINDOW_HEIGHT - obj.H
		velocity.Y *= -1
	}

	if obj.Y < 0 {
		obj.Y = 0
		velocity.Y *= -1
	}

	velocity.Y += GRAVITY
}

func draw(renderer *sdl.Renderer, obj sdl.Rect) {
	err := renderer.SetDrawColor(255, 0, 0, 255)
	sdlutil.HandleError(err)
	err = renderer.FillRect(&obj)
	sdlutil.HandleError(err)
}

func NormalizeVector(v *sdl.Point) {
	x := float64(v.X) / VectorLength(*v)
	y := float64(v.Y) / VectorLength(*v)
	fmt.Printf("Normalized: X=%f, Y=%f\n", x, y)
	v.X = int32(math.Round(x))
	v.Y = int32(math.Round(y))
}

func VectorLength(v sdl.Point) float64 {
	x := v.X * v.X
	y := v.Y * v.Y
	return math.Sqrt(float64(x) + float64(y))
}
