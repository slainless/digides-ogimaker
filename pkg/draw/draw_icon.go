package draw

import (
	"image"
	"image/draw"

	"github.com/disintegration/gift"
)

func drawIcon(canvas draw.Image, icon image.Image, bound image.Rectangle) {
	filter := NewResizeToFit(bound.Dx(), bound.Dy())
	filter.DrawAt(canvas, icon, bound.Min, gift.OverOperator)
}
