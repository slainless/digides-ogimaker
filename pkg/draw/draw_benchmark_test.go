package draw_test

import (
	_ "embed"
	"image"
	"os"
	"testing"

	"github.com/slainless/digides-ogimaker/pkg/draw"
	"github.com/slainless/digides-ogimaker/pkg/fonts"
	"github.com/slainless/digides-ogimaker/pkg/wasm"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func LoadParameters() *wasm.Parameters {
	params := &wasm.Parameters{}
	params.
		SetTitle("This is a title").
		SetSubtitle("This is a subtitle")

	background, err := os.Open("../../assets/35ade0a022c7b566dbffdc934f4cb174.png")
	panicIf(err)

	icon, err := os.Open("../../assets/300_barru.png")
	panicIf(err)

	imageBackground, _, _ := image.Decode(background)
	imageIcon, _, _ := image.Decode(icon)

	params.
		SetBackground(imageBackground).
		SetIcon(imageIcon).
		SetFontFaceTitle(fonts.OutfitRegularFont).
		SetFontFaceSubtitle(fonts.OutfitRegularFont)

	return params
}

var parameters = LoadParameters()

func BenchmarkDirectDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = draw.Draw(parameters)
	}
}
