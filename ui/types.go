package ui

import (
	"fyne.io/fyne/v2"
	"github.com/simpleittools/pixl/apptype"
	"github.com/simpleittools/pixl/pxcanvas"
	"github.com/simpleittools/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
