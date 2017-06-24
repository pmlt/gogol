package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const NumRows = 50
const RowHeightPx = 20
const NumColumns = 50
const ColWidthPx = 20
const GenerationTick = 16

func main() {
	var running bool
	var changed bool
	var dragging bool
	var dragbtn uint8
	var event sdl.Event
	var state GameState
	state.Started = false

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("GoGOL - Go Game Of Life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		NumColumns*ColWidthPx, NumRows*RowHeightPx, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	draw(state.Current, renderer)

	dragging = false
	running = true
	for running {
		changed = false
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == 32 {
					state.Started = !state.Started
				}
			case *sdl.MouseMotionEvent:
				if !state.Started {
					row := (t.Y / RowHeightPx)
					col := (t.X / ColWidthPx)
					if dragging {
						changed = state.Current[row][col] == (dragbtn != 1)
						state.Current[row][col] = dragbtn == 1
					}
				}
			case *sdl.MouseButtonEvent:
				if !state.Started {
					if dragging {
						if t.Button == dragbtn && t.State == 0 {
							// dragend
							dragging = false
						}
						// We don't listen to any other button while dragging
					} else {
						if t.State == 1 {
							// Button is pressed, start drag
							dragging = true
							dragbtn = t.Button
							row := (t.Y / RowHeightPx)
							col := (t.X / ColWidthPx)
							changed = state.Current[row][col] == (dragbtn != 1)
							state.Current[row][col] = dragbtn == 1
						}
					}
				}
			}
		}
		if state.Started {
			state = step(state)
			changed = true
		}
		if changed {
			draw(state.Current, renderer)
		}
		sdl.Delay(GenerationTick)
	}

	sdl.Quit()
	os.Exit(0)
}

func draw(state [NumRows][NumColumns]bool, renderer *sdl.Renderer) {
	for i, row := range state {
		for j, cell := range row {
			rect := sdl.Rect{X: int32(j * ColWidthPx), Y: int32(i * RowHeightPx), W: int32(ColWidthPx), H: int32(RowHeightPx)}
			if cell {
				renderer.SetDrawColor(0, 0, 0, 255)
			} else {
				renderer.SetDrawColor(255, 255, 255, 255)
			}
			renderer.FillRect(&rect)
			renderer.SetDrawColor(196, 196, 196, 255)
			renderer.DrawRect(&rect)
		}
	}
	renderer.Present()
}
