package main

import (
	"image/color"
	"pixart/apptype"
	"pixart/pxcanvas"
	"pixart/swatch"
	"pixart/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	pixartApp := app.New()
	pixartWindow := pixartApp.NewWindow("pixart")
	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}
	pixartCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}
	pixartCanvas := pxcanvas.NewPxCanvas(&state, pixartCanvasConfig)
	appInit := ui.AppInit{
		PixartCanvas: pixartCanvas,
		PixartWindow: pixartWindow,
		State:        &state,
		Swatches:     make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixartWindow.ShowAndRun()
}
