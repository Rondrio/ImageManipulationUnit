package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/color"
)

func (list *ImageList) Grayscale(flags Flags.Flags) error {
	var alias string
	if flags.CheckIfFlagsAreSet("alias") {
		alias = flags.Flag["alias"]
	} else {
		return errors.New("unset flags")
	}
	image := list.GetImageByAlias(alias)
	image.ChangeToGrayscale()
	return nil
}

func (image *Image) ChangeToGrayscale() {
	paint := func(width, height int, img SetColor) {
		img.Set(width, height, color.Gray16Model.Convert(image.Image.At(width, height)))
	}
	image.IterateOverPixels(paint)
}
