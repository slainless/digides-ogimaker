package draw

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"math"

	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
	"github.com/goki/freetype/truetype"
)

func drawStrings(canvas draw.Image, param Parameters, bound image.Rectangle) (err error) {
	const (
		canvasRatio = 2.5

		rootEm float64 = 16

		titleFontSize    = rootEm * 6
		subtitleFontSize = rootEm * 3
	)

	var (
		titleFontFace = truetype.NewFace(param.FontFaceTitle(), &truetype.Options{
			Size: titleFontSize,
		})
		subtitleFontFace = truetype.NewFace(param.FontFaceSubtitle(), &truetype.Options{
			Size: subtitleFontSize,
		})

		titleAscent     = float64(titleFontFace.Metrics().Ascent.Round())
		subtitleXHeight = float64(subtitleFontFace.Metrics().XHeight.Round())
		subtitleHeight  = float64(subtitleFontFace.Metrics().Height.Round())

		canvasHeight = titleAscent + (subtitleHeight * 2)
		canvasWidth  = canvasHeight * canvasRatio
	)

	drawCanvas := gg.NewContext(int(math.Round(canvasWidth)), int(math.Round(canvasHeight)))

	drawCanvas.SetFontFace(titleFontFace)
	titleWidth, _ := drawCanvas.MeasureString(param.Title())
	if titleWidth >= canvasWidth {
		return errors.New("title surpasses maximum text length limit, consider using another smaller font-face or reduce string length")
	}

	drawCanvas.SetFontFace(subtitleFontFace)
	_, _subtitleHeight := drawCanvas.MeasureMultilineString(param.Subtitle(), 0)
	if _subtitleHeight > subtitleHeight*2 {
		return errors.New("subtitle surpasses maximum text length/height limit, consider using another smaller font-face or reduce string length")
	}

	drawCanvas.SetColor(color.NRGBA{255, 255, 255, 255})
	drawCanvas.SetFontFace(titleFontFace)
	drawCanvas.DrawString(param.Title(), 0, titleAscent)

	drawCanvas.SetColor(color.NRGBA{255, 255, 255, 255})
	drawCanvas.SetFontFace(subtitleFontFace)
	drawCanvas.DrawStringWrapped(param.Subtitle(), 0, titleAscent, 0, 0, float64(drawCanvas.Width()), 0, gg.AlignLeft)

	result := drawCanvas.Image()
	resultHeight := titleAscent + _subtitleHeight + subtitleXHeight

	filter := gift.New()
	filter.Add(gift.CropToSize(
		drawCanvas.Width(),
		int(resultHeight),
		gift.TopLeftAnchor,
	))
	filter.Add(gift.Resize(bound.Dx(), 0, gift.LanczosResampling))
	point := image.Pt(
		bound.Min.X,
		bound.Min.Y+int((float64(bound.Dy())-resultHeight)/2),
	)
	filter.DrawAt(canvas, result, point, gift.OverOperator)

	return nil
}
