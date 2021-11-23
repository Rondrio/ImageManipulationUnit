package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"image/color"
	"strconv"
)

type SetColorCommand struct {
	Keyword string
}

func (cmd SetColorCommand) GetKeyword() string {
	return cmd.Keyword
}

func (cmd SetColorCommand) Execute(list *ImageList, flags Flags.Flags, selection *Selection) error {
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
