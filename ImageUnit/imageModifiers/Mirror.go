package imageModifiers

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"ImageManipulationUnit/ImageUnit/utils"
	"errors"
	"image/color"
	"strconv"
)

func (list *ImageUnit.ImageList) MirrorImage(flags Flags.Flags) error {
	var alias string
	horizontal, vertical := false, false
	var err error
	if flags.CheckIfFlagsAreSet("alias") {
		alias = flags.Flag["alias"]
	}
	if flags.CheckIfFlagsAreSet("horizontal") {
		horizontal, err = strconv.ParseBool(flags.Flag["horizontal"])
		if err != nil {
			return err
		}
	}
	if flags.CheckIfFlagsAreSet("vertical") {
		vertical, err = strconv.ParseBool(flags.Flag["vertical"])
		if err != nil {
			return err
		}
	}
	if (!horizontal && !vertical) || alias == "" {
		return Flags.ErrUnsetFlags
	}
	image := list.GetImageByAlias(alias)
	if err := image.Mirror(horizontal, vertical); err != nil {
		return err
	}
	return nil
}

func (imgStruct *ImageUnit.Image) Mirror(horizontal, vertical bool) error {
	buffer := make([][]color.Color, imgStruct.Image.Bounds().Max.Y+1)
	for row := range buffer {
		buffer[row] = make([]color.Color, imgStruct.Image.Bounds().Max.X+1)
	}

	if horizontal {
		buffer = MirrorHorizontal(imgStruct, buffer)
	}
	if vertical {
		buffer = MirrorVertical(imgStruct, buffer)
	}

	if img, ok := imgStruct.Image.(utils.SetColor); ok {
		for height := 0; height < len(buffer); height++ {
			for width := 0; width < len(buffer[height]); width++ {
				ImageUnit.Set(width, height, buffer[height][width])
			}
		}
		return nil
	} else {
		return errors.New("imgStruct unchangeable")
	}
}
func MirrorHorizontal(image *ImageUnit.Image, buffer [][]color.Color) [][]color.Color {
	for height := 0; height < image.Image.Bounds().Max.Y; height++ {
		for width := 0; width < image.Image.Bounds().Max.X; width++ {
			buffer[height][image.Image.Bounds().Max.X-1-width] = image.Image.At(width, height)
		}
	}
	return buffer
}
func MirrorVertical(image *ImageUnit.Image, buffer [][]color.Color) [][]color.Color {
	for height := 0; height < image.Image.Bounds().Max.Y; height++ {
		for width := 0; width < image.Image.Bounds().Max.X; width++ {
			buffer[image.Image.Bounds().Max.Y-1-height][width] = image.Image.At(width, height)
		}
	}
	return buffer
}
