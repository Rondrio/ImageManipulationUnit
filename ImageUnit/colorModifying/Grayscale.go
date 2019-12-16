package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"image/color"
)

func (list *ImageUnit.ImageList) Grayscale(flags Flags.Flags, selection *ImageUnit.Selection) error {
	if !flags.CheckIfFlagsAreSet("alias") {
		return Flags.ErrUnsetFlags
	}
	alias := flags.Flag["alias"]
	image := list.GetImageByAlias(alias)
	if err := image.ChangeToGrayscale(selection); err != nil {
		return err
	}
	return nil
}

func (image *ImageUnit.Image) ChangeToGrayscale(selection *ImageUnit.Selection) error {
	paint := func(width, height int, img ImageUnit.SetColor) {
		ImageUnit.Set(width, height, color.Gray16Model.Convert(image.Image.At(width, height)))
	}
	return image.IterateOverPixels(paint, selection)
}
