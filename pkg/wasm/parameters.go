package wasm

import (
	"bytes"
	"image"
	"io"
	"os"
	"unsafe"

	"github.com/goki/freetype/truetype"
	"github.com/slainless/digides-ogimaker/pkg/fonts"
)

type Parameters struct {
	title            string
	subtitle         string
	icon             image.Image
	background       image.Image
	fontFaceTitle    *truetype.Font
	fontFaceSubtitle *truetype.Font
	Quality          uint8
}

func (p *Parameters) Title() string                    { return p.title }
func (p *Parameters) Subtitle() string                 { return p.subtitle }
func (p *Parameters) Icon() image.Image                { return p.icon }
func (p *Parameters) Background() image.Image          { return p.background }
func (p *Parameters) FontFaceTitle() *truetype.Font    { return p.fontFaceTitle }
func (p *Parameters) FontFaceSubtitle() *truetype.Font { return p.fontFaceSubtitle }

func ReadParameters() (*Parameters, error) {
	stdin := os.Stdin

	sizeOffset := 1
	// allocate u32 size for each parameter (6 parameter) and misc data
	// for now, should only contains quality + 6 parameters size
	initialRead := make([]uint8, 6*4+sizeOffset)

	// read parameter's size first
	n, err := stdin.Read(initialRead)
	if err != nil {
		return nil, err
	}

	if n != len(initialRead) {
		return nil, io.EOF
	}

	parameters := make([][]uint8, 6)
	for i := 0; i < 6; i++ {
		// read the size as uint32
		size := *(*uint32)(unsafe.Pointer(&initialRead[4*i+sizeOffset]))
		parameters[i] = make([]uint8, size)
		stdin.Read(parameters[i])
	}

	_parameters := &Parameters{
		title:    string(parameters[0]),
		subtitle: string(parameters[1]),
		Quality:  initialRead[0],
	}

	_parameters.icon, _, err = image.Decode(bytes.NewReader(parameters[2]))
	if err != nil {
		return nil, err
	}

	_parameters.background, _, err = image.Decode(bytes.NewReader(parameters[3]))
	if err != nil {
		return nil, err
	}

	_parameters.fontFaceTitle = fonts.OutfitRegularFont
	_parameters.fontFaceSubtitle = fonts.OutfitRegularFont

	return _parameters, nil
}
