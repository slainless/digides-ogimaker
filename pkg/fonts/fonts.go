package fonts

import (
	_ "embed"

	"github.com/goki/freetype/truetype"
)

//go:embed Outfit-Regular.ttf
var OutfitRegular []byte

var (
	OutfitRegularFont, _ = truetype.Parse(OutfitRegular)
)
