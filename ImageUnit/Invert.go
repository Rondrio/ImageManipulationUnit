package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/color"
)

func (list *ImageList) Invert(flags Flags.Flags) error {
	var alias string
	if flags.CheckIfFlagsAreSet("alias") {
		alias = flags.Flag["alias"]
	} else {
		return errors.New("unset flags")
	}
	image := list.GetImageByAlias(alias)
	if err := image.InvertColor();err != nil{
		return err
	}
	return nil
}

func (image *Image) InvertColor() error{
	paint := func(width, height int, img SetColor) {
		oldR, oldG, oldB, oldA := image.Image.At(width, height).RGBA()
		c := color.Color(color.RGBA64{
			R: getInvertedValue(uint16(oldR)),
			G: getInvertedValue(uint16(oldG)),
			B: getInvertedValue(uint16(oldB)),
			A: uint16(oldA),
		})
		img.Set(width, height, c)
	}
	return image.IterateOverPixels(paint)
}
func getInvertedValue(value uint16) uint16 {
	return 65535 - value
}
