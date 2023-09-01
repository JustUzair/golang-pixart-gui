package ui

import (
	"pixart/apptype"
	"pixart/pxcanvas"
	"pixart/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixartCanvas *pxcanvas.PxCanvas
	PixartWindow fyne.Window
	State        *apptype.State
	Swatches     []*swatch.Swatch
}
