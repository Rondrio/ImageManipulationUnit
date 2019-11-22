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
	if err := image.ChangeToGrayscale();err != nil{
		return err
	}
	return nil
}

func (image *Image) ChangeToGrayscale()error {
	paint := func(width, height int, img SetColor) {
		img.Set(width, height, color.Gray16Model.Convert(image.Image.At(width, height)))
	}
	return image.IterateOverPixels(paint)
}
