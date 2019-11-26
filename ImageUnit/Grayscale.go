package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"image/color"
)

func (list *ImageList) Grayscale(flags Flags.Flags, selection *Selection) error {
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

func (image *Image) ChangeToGrayscale(selection *Selection) error {
	paint := func(width, height int, img SetColor) {
		img.Set(width, height, color.Gray16Model.Convert(image.Image.At(width, height)))
	}
	return image.IterateOverPixels(paint, selection)
}
