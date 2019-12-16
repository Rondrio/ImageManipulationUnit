package colorModifiers

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit/utils"
	"image/color"
	"strconv"
)

type imageList struct{
	utils.IImageList
	ColorModifiersList
	images []Image
}
type Image struct{
	utils.IImage
	ColorModImage
	utils.Image
}
func init(){

}

func (list *imageList) AddColor(flags Flags.Flags, selection *utils.Selection) error {
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

	if err := image.ChangeColor(uint32(red), uint32(green), uint32(blue), uint32(alpha), selection); err != nil {
		return err
	}
	return nil
}

func (imgStruct *Image) ChangeColor(r, g, b, a uint32, selection *utils.Selection) error {

	paint := func(width, height int, img utils.SetColor) {
		oldR, oldG, oldB, oldA := imgStruct.Image.Image.At(width, height).RGBA()
		c := color.Color(color.RGBA64{
			R: uint16(oldR - r),
			G: uint16(oldG - g),
			B: uint16(oldB - b),
			A: uint16(oldA - a),
		})
		img.Set(width, height, c)
	}
	return imgStruct.Image.IterateOverPixels(paint, selection)
}
