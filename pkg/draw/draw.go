package draw

import (
	"fmt"
	"image"

	"github.com/disintegration/gift"
	"github.com/fogleman/gg"
)

// 0.262145631 s/op
// using parameters: [github.com/slainless/digides-ogimage/pkg/ogimage_test.LoadParameters]
func Draw(param Parameters) (*image.RGBA, error) {
	const (
		canvasWidth  = 1200
		canvasHeight = 630

		canvasMarginInline = 95

		elementRatio    = 2.805
		elementMargin   = 65
		iconToStringGap = 60
	)

	fmt.Println("Creating canvas")
	img := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasHeight))
	canvas := gg.NewContextForRGBA(img)

	fmt.Println("Drawing background")
	// draw the background
	filter := NewResizeToFill(canvasWidth, canvasHeight)
	filter.DrawAt(img, param.Background(), image.Pt(0, 0), gift.CopyOperator)

	// draw the backdrop
	elementWidth := canvasWidth - (canvasMarginInline * 2)
	elementHeight := float64(elementWidth) / elementRatio
	canvasMarginBlock := (canvasHeight - elementHeight) / 2

	elementBound := image.Rect(
		canvasMarginInline,
		int(canvasMarginBlock),
		canvasMarginInline+elementWidth,
		int(canvasMarginBlock+elementHeight),
	)
	drawBackdrop(canvas, elementBound)

	fmt.Println("Drawing icon")
	// draw the icon
	iconSize := elementHeight - (elementMargin * 2)
	iconBound := image.Rect(
		canvasMarginInline+elementMargin,
		int(canvasMarginBlock+elementMargin),
		int(canvasMarginInline+elementMargin+iconSize),
		int(canvasMarginBlock+elementMargin+iconSize),
	)
	drawIcon(img, param.Icon(), iconBound)

	fmt.Println("Drawing string")
	// draw the string
	stringWidth := canvasWidth - ((canvasMarginInline + elementMargin) * 2) - iconSize - iconToStringGap
	stringBound := image.Rect(
		int(canvasMarginInline+elementMargin+iconSize+iconToStringGap),
		int(canvasMarginBlock+elementMargin),
		int(canvasMarginInline+elementMargin+iconSize+iconToStringGap+stringWidth),
		int(canvasMarginBlock+elementMargin+iconSize),
	)
	err := drawStrings(img, param, stringBound)
	fmt.Println("Drawing result", err)

	if err != nil {
		return nil, err
	}

	return img, nil
}
