package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const NumRows = 50
const RowHeightPx = 20
const NumColumns = 50
const ColWidthPx = 20
const HelpHeightPx = 100
const GenerationTick = 16

func main() {
	var running bool
	var changed bool
	var dragging bool
	var dragbtn uint8
	var event sdl.Event
	var ui UIState
	var state GameState

	ui = CreateUI()
	defer ui.Free()

	ui.draw(state)

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
					changed = true
				}
			case *sdl.MouseMotionEvent:
				if !state.Started {
					row := (t.Y / RowHeightPx)
					col := (t.X / ColWidthPx)
					if dragging && row >= 0 && row < NumRows && col >= 0 && col < NumColumns {
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
							if row >= 0 && row < NumRows && col >= 0 && col < NumColumns {
								changed = state.Current[row][col] == (dragbtn != 1)
								state.Current[row][col] = dragbtn == 1
							}
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
			ui.draw(state)
		}
		sdl.Delay(GenerationTick)
	}

	os.Exit(0)
}
