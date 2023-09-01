package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/lusingander/colorpicker"
)

func SetupColorPicker(app *AppInit) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(c color.Color) {
		app.Swatches[app.State.SwatchSelected].SetColor(c)
		app.State.BrushColor = c

	})

	return container.NewVBox(picker)
}
