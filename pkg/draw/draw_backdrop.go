package draw

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

func drawBackdrop(canvas *gg.Context, bound image.Rectangle) {
	canvas.SetColor(color.NRGBA{0, 0, 0, 89})
	canvas.DrawRectangle(
		float64(bound.Min.X),
		float64(bound.Min.Y),
		float64(bound.Dx()),
		float64(bound.Dy()),
	)
	canvas.Fill()
}
