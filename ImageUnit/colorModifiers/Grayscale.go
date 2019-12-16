package colorModifiers

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit/utils"
	"image/color"
)

func (list imageList) Grayscale(flags Flags.Flags, selection *utils.Selection) error {
	var image Image
	if !flags.CheckIfFlagsAreSet("alias") {
		return Flags.ErrUnsetFlags
	}
	alias := flags.Flag["alias"]
	image.IImage = list.GetImageByAlias(alias)
	if err := image.ChangeToGrayscale(selection); err != nil {
		return err
	}
	return nil
}

func (imgStruct Image) ChangeToGrayscale(selection *utils.Selection) error {
	paint := func(width, height int, img utils.SetColor) {
		img.Set(width, height, color.Gray16Model.Convert(imgStruct.Image.Image.At(width, height)))
	}
	return imgStruct.Image.IterateOverPixels(paint, selection)
}
