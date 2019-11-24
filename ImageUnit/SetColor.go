package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/color"
	"strconv"
)

func (list *ImageList) SetColor(flags Flags.Flags, selection *Selection) error {

	if set := flags.CheckIfFlagsAreSet("alias", "red", "blue", "green", "alpha"); !set {
		return errors.New("unset flags")
	}
	alias := flags.Flag["alias"]

	red, err := strconv.Atoi(flags.Flag["red"])
	if err != nil {
		return err
	}

	green, err := strconv.Atoi(flags.Flag["green"])
	if err != nil {
		return err
	}

	blue, err := strconv.Atoi(flags.Flag["blue"])
	if err != nil {
		return err
	}

	alpha, err := strconv.Atoi(flags.Flag["alpha"])
	if err != nil {
		return err
	}
	if red == 0 && green == 0 && blue == 0 && alpha == 0 {
		return Flags.ErrUnsetFlags
	}

	image := list.GetImageByAlias(alias)

	if err := image.SetColor(uint32(red), uint32(green), uint32(blue), uint32(alpha), selection); err != nil {
		return err
	}
	return nil

}
func (image *Image) SetColor(r, g, b, a uint32, selection *Selection) error {
	paint := func(width, height int, img SetColor) {
		c := color.Color(color.RGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: uint16(a),
		})
		img.Set(width, height, c)
	}
	return image.IterateOverPixels(paint, selection)
}
