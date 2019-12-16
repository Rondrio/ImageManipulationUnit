package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"image/color"
)

func (list *ImageUnit.ImageList) Invert(flags Flags.Flags, selection *ImageUnit.Selection) error {
	if !flags.CheckIfFlagsAreSet("alias") {
		return Flags.ErrUnsetFlags
	}
	alias := flags.Flag["alias"]
	image := list.GetImageByAlias(alias)
	if err := image.InvertColor(selection); err != nil {
		return err
	}
	return nil
}

func (image *ImageUnit.Image) InvertColor(selection *ImageUnit.Selection) error {
	paint := func(width, height int, img ImageUnit.SetColor) {
		oldR, oldG, oldB, oldA := image.Image.At(width, height).RGBA()
		c := color.Color(color.RGBA64{
			R: getInvertedValue(uint16(oldR)),
			G: getInvertedValue(uint16(oldG)),
			B: getInvertedValue(uint16(oldB)),
			A: uint16(oldA),
		})
		ImageUnit.Set(width, height, c)
	}
	return image.IterateOverPixels(paint, selection)
}

func getInvertedValue(value uint16) uint16 {
	return ImageUnit.max16bit - value
}
