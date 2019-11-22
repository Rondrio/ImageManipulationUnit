package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/color"
	"strconv"
)

func (list *ImageList) AddColor(flags Flags.Flags) error {
	var alias string
	var red, green, blue, alpha int
	var err error
	if flags.CheckIfFlagsAreSet("alias") {
		alias = flags.Flag["alias"]
		red, err = strconv.Atoi(flags.Flag["red"])
		if err != nil {
			return err
		}

		green, err = strconv.Atoi(flags.Flag["green"])
		if err != nil {
			return err
		}
		blue, err = strconv.Atoi(flags.Flag["blue"])
		if err != nil {
			return err
		}

		alpha, err = strconv.Atoi(flags.Flag["alpha"])
		if err != nil {
			return err
		}
	} else {
		return errors.New("unset flags")
	}
	if red == 0 && green == 0 && blue == 0 && alpha == 0 {
		return errors.New("no rgb values set")
	}
	image := list.GetImageByAlias(alias)
	if err := image.ChangeColor(uint32(red), uint32(green), uint32(blue), uint32(alpha)); err != nil {
		return err
	}
	return nil
}

func (image *Image) ChangeColor(r, g, b, a uint32) error {
	paint := func(width, height int, img SetColor) {
		oldR, oldG, oldB, oldA := image.Image.At(width, height).RGBA()
		c := color.Color(color.RGBA64{
			R: uint16(oldR - r),
			G: uint16(oldG - g),
			B: uint16(oldB - b),
			A: uint16(oldA - a),
		})
		img.Set(width, height, c)
	}
	return image.IterateOverPixels(paint)
}
