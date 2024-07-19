package draw

import (
	"github.com/disintegration/gift"
)

func NewResizeToFill(w int, h int) *gift.GIFT {
	return gift.New(gift.ResizeToFill(w, h, gift.LanczosResampling, gift.CenterAnchor))
}

func NewResizeToFit(w int, h int) *gift.GIFT {
	return gift.New(gift.ResizeToFit(w, h, gift.LanczosResampling))
}
