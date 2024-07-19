package main

// // #include <stdlib.h>
// import "C"

import (
	"image/jpeg"
	"os"

	"github.com/slainless/digides-ogimaker/pkg/draw"
	"github.com/slainless/digides-ogimaker/pkg/wasm"
)

func main() {
	param, err := wasm.ReadParameters()
	if err != nil {
		wasm.Exit(err)
		return
	}

	img, err := draw.Draw(param)
	if err != nil {
		wasm.Exit(err)
		return
	}

	err = jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: 100})
	if err != nil {
		wasm.Exit(err)
		return
	}
}
