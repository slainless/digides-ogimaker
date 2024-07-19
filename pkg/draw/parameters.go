package draw

import (
	"image"

	"github.com/goki/freetype/truetype"
)

type Parameters interface {
	Title() string
	Subtitle() string

	Icon() image.Image
	Background() image.Image

	FontFaceTitle() *truetype.Font
	FontFaceSubtitle() *truetype.Font
}
