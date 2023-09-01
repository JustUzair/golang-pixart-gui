package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

type Brushable interface {
	SetColor(c color.Color, x int, y int)
	MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int)
}

func (state *State) setFilePath(path string) {
	state.FilePath = path
}
