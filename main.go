package main

import (
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const NumRows = 50
const RowHeightPx = 20
const NumColumns = 50
const ColWidthPx = 20
const (
	ColorAlive uint32 = 0xffff0000
	ColorDead  uint32 = 0x00000000
)

func main() {
	var running bool
	var changed bool
	var dragging bool
	var dragbtn uint8
	var event sdl.Event
	var state [NumRows][NumColumns]bool

	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		NumColumns*ColWidthPx, NumRows*RowHeightPx, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	dragging = false
	running = true
	for running {
		changed = false
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				row := (t.Y / RowHeightPx)
				col := (t.X / ColWidthPx)
				if dragging {
					changed = state[row][col] == (dragbtn != 1)
					state[row][col] = dragbtn == 1
				}
			case *sdl.MouseButtonEvent:
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
						changed = state[row][col] == (dragbtn != 1)
						state[row][col] = dragbtn == 1
					}
				}
			}
		}
		if changed {
			draw(state, window)
		}
		sdl.Delay(16)
	}

	sdl.Quit()
	os.Exit(0)
}

func draw(state [NumRows][NumColumns]bool, window *sdl.Window) {
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	for i, row := range state {
		for j, cell := range row {
			rect := sdl.Rect{X: int32(j * ColWidthPx), Y: int32(i * RowHeightPx), W: int32(ColWidthPx), H: int32(RowHeightPx)}
			color := ColorDead
			if cell {
				color = ColorAlive
			}
			surface.FillRect(&rect, color)
		}
	}
	window.UpdateSurface()
}
