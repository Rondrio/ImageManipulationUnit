package colorModifiers

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit/utils"
	"image/color"
	"strconv"
)

func (list imageList) SetColor(flags Flags.Flags, selection *utils.Selection) error {
	var alias string
	var red, green, blue, alpha int
	var err error
	var image Image

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
	image.IImage = list.GetImageByAlias(alias)

	if err := image.SetColor(uint32(red), uint32(green), uint32(blue), uint32(alpha), selection); err != nil {
		return err
	}
	return nil

}
func (imgStruct Image) SetColor(r, g, b, a uint32, selection *utils.Selection) error {
	paint := func(width, height int, img utils.SetColor) {
		c := color.Color(color.RGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: uint16(a),
		})
		img.Set(width, height, c)
	}
	return imgStruct.Image.IterateOverPixels(paint, selection)
}
