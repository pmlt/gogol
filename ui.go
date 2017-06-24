package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const HelpText string = "Welcome to Conway's Game of Life!"
const FontSize = 24

var TextColor = sdl.Color{R: 0, G: 0, B: 0, A: 0}

type UIState struct {
	Window      *sdl.Window
	Renderer    *sdl.Renderer
	Font        *ttf.Font
	HelpRect    sdl.Rect
	HelpTexture *sdl.Texture
}

func CreateUI() UIState {
	var ui UIState
	var err error

	sdl.Init(sdl.INIT_EVERYTHING)

	ui.Window, err = sdl.CreateWindow("GoGOL - Go Game Of Life", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		NumColumns*ColWidthPx, NumRows*RowHeightPx+HelpHeightPx, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	ui.Renderer, err = sdl.CreateRenderer(ui.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	if ui.Font, err = ttf.OpenFont("whitrabt.ttf", FontSize); err != nil {
		panic(err)
	}
	return ui
}

func (ui UIState) Free() {
	ui.HelpTexture.Destroy()
	ui.Font.Close()
	ui.Renderer.Destroy()
	ui.Window.Destroy()
	sdl.Quit()
}

func (ui UIState) draw(state GameState) {
	for i, row := range state.Current {
		for j, cell := range row {
			rect := sdl.Rect{X: int32(j * ColWidthPx), Y: int32(i * RowHeightPx), W: int32(ColWidthPx), H: int32(RowHeightPx)}
			if cell {
				ui.Renderer.SetDrawColor(0, 0, 0, 255)
			} else {
				ui.Renderer.SetDrawColor(255, 255, 255, 255)
			}
			ui.Renderer.FillRect(&rect)
			ui.Renderer.SetDrawColor(196, 196, 196, 255)
			ui.Renderer.DrawRect(&rect)
		}
	}
	helpRect := sdl.Rect{W: NumColumns * ColWidthPx, H: HelpHeightPx, X: 0, Y: NumRows * RowHeightPx}
	ui.Renderer.SetDrawColor(255, 255, 255, 255)
	ui.Renderer.FillRect(&helpRect)
	x, y := ui.DrawLineOfText(0, helpRect.Y+5, "Welcome to Conway's Game of Life!")
	if state.Started {
		ui.DrawLineOfText(x, y+20, "Game started... press Spacebar to interrupt.")
	} else {
		x, y = ui.DrawLineOfText(x, y+20, "Draw initial state using the mouse,")
		ui.DrawLineOfText(x, y, "then start the game with Spacebar!")
	}
	ui.Renderer.Present()
}

func (ui UIState) DrawLineOfText(x int32, y int32, text string) (int32, int32) {
	var solid *sdl.Surface
	var texture *sdl.Texture
	var err error
	if solid, err = ui.Font.RenderUTF8_Solid(text, TextColor); err != nil {
		panic(err)
	}
	defer solid.Free()
	if texture, err = ui.Renderer.CreateTextureFromSurface(solid); err != nil {
		panic(err)
	}
	defer texture.Destroy()
	textRect := sdl.Rect{W: solid.W, H: solid.H, X: x, Y: y}
	ui.Renderer.Copy(texture, nil, &textRect)
	return x, y + int32(ui.Font.Height())
}
