package colorModifiers

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit/utils"
	"image/color"
)

func (list *imageList) Invert(flags Flags.Flags, selection *utils.Selection) error {
	var image Image
	if !flags.CheckIfFlagsAreSet("alias") {
		return Flags.ErrUnsetFlags
	}
	alias := flags.Flag["alias"]
	image.IImage = list.GetImageByAlias(alias)
	if err := image.InvertColor(selection); err != nil {
		return err
	}
	return nil
}

func (imgStruct *Image) InvertColor(selection *utils.Selection) error {
	paint := func(width, height int, img utils.SetColor) {
		oldR, oldG, oldB, oldA := imgStruct.Image.Image.At(width, height).RGBA()
		c := color.Color(color.RGBA64{
			R: getInvertedValue(uint16(oldR)),
			G: getInvertedValue(uint16(oldG)),
			B: getInvertedValue(uint16(oldB)),
			A: uint16(oldA),
		})
		img.Set(width, height, c)
	}
	return imgStruct.Image.IterateOverPixels(paint, selection)
}

func getInvertedValue(value uint16) uint16 {
	return utils.Max16bit - value
}
