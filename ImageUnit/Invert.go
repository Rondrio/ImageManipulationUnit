package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"image/color"
)

type InvertCommand struct {
	Keyword string
}

func (cmd InvertCommand) GetKeyword() string {
	return cmd.Keyword
}

func (cmd InvertCommand) Execute(list *ImageList, flags Flags.Flags, selection *Selection) error {
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

func (image *Image) InvertColor(selection *Selection) error {
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
	return image.IterateOverPixels(paint, selection)
}

func getInvertedValue(value uint16) uint16 {
	return max16bit - value
}
