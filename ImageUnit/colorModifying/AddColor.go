package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"image/color"
	"strconv"
)

func (list *ImageUnit.ImageList) AddColor(flags Flags.Flags, selection *ImageUnit.Selection) error {
	var alias string
	var red, green, blue, alpha int
	var err error
	if flags.CheckIfFlagsAreSet("alias") {
		alias = flags.Flag["alias"]
	}
	if flags.CheckIfFlagsAreSet("red") {
		red, err = strconv.Atoi(flags.Flag["red"])
		if err != nil {
			return err
		}
	}
	if flags.CheckIfFlagsAreSet("green") {
		green, err = strconv.Atoi(flags.Flag["green"])
		if err != nil {
			return err
		}
	}
	if flags.CheckIfFlagsAreSet("blue") {
		blue, err = strconv.Atoi(flags.Flag["blue"])
		if err != nil {
			return err
		}
	}
	if flags.CheckIfFlagsAreSet("alpha") {
		alpha, err = strconv.Atoi(flags.Flag["alpha"])
		if err != nil {
			return err
		}
	}
	if (red == 0 && green == 0 && blue == 0 && alpha == 0) || alias == "" {
		return Flags.ErrUnsetFlags
	}
	image := list.GetImageByAlias(alias)
	if err := image.ChangeColor(uint32(red), uint32(green), uint32(blue), uint32(alpha), selection); err != nil {
		return err
	}
	return nil
}

func (image *ImageUnit.Image) ChangeColor(r, g, b, a uint32, selection *ImageUnit.Selection) error {
	paint := func(width, height int, img ImageUnit.SetColor) {
		oldR, oldG, oldB, oldA := image.Image.At(width, height).RGBA()
		c := color.Color(color.RGBA64{
			R: uint16(oldR - r),
			G: uint16(oldG - g),
			B: uint16(oldB - b),
			A: uint16(oldA - a),
		})
		ImageUnit.Set(width, height, c)
	}
	return image.IterateOverPixels(paint, selection)
}
