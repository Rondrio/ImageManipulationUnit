package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/color"
)

func (list *ImageList) MirrorImage(flags Flags.Flags) error {
	if set := flags.CheckIfFlagsAreSet("alias"); !set {
		return Flags.ErrUnsetFlags
	}
	alias := flags.Flag["alias"]
	image := list.GetImageByAlias(alias)
	if err := image.Mirror(); err != nil {
		return err
	}
	return nil
}

func (image *Image) Mirror() error {
	buffer := make([][]color.Color, image.Image.Bounds().Max.Y+1)
	for row := range buffer {
		buffer[row] = make([]color.Color, image.Image.Bounds().Max.X+1)
	}

	for height := 0; height < image.Image.Bounds().Max.Y; height++ {
		for width := 0; width < image.Image.Bounds().Max.X; width++ {
			buffer[height][image.Image.Bounds().Max.X-1-width] = image.Image.At(width, height)
		}
	}

	if img, ok := image.Image.(SetColor); ok {
		for height := 0; height < len(buffer); height++ {
			for width := 0; width < len(buffer[height]); width++ {
				img.Set(width, height, buffer[height][width])
			}
		}
		return nil
	} else {
		return errors.New("image unchangeable")
	}
}
