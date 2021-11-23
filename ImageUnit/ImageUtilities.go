package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image"
	"image/color"
)

const (
	max16bit = 65535
)

var ImageTunnel *Image

type Command interface {
	GetKeyword() string
	Execute(*ImageList, Flags.Flags, *Selection) error
}

type CommandList []Command

type Image struct {
	Id    int64
	Alias string
	Path  string
	Image image.Image
}
type ImageList struct {
	LoadedImages []Image
}
type Selection struct {
	Points []Point
}
type Point struct {
	X int
	Y int
}
type SetColor interface {
	Set(x, y int, c color.Color)
}

func (list *ImageList) GetImageByAlias(alias string) *Image {
	for _, image := range list.LoadedImages {
		if image.Alias == alias {
			return &image
		}
	}
	return nil
}

func (image *Image) IterateOverPixels(paint func(width, height int, img SetColor), selection *Selection) error {
	if img, ok := image.Image.(SetColor); ok {
		for height := 0; height < image.Image.Bounds().Max.Y; height++ {
			for width := 0; width < image.Image.Bounds().Max.X; width++ {
				if len(selection.Points) > 2 {
					if selected := selection.CheckIfSelected(Point{width, height}); selected {
						paint(width, height, img)
					}
				} else {
					paint(width, height, img)
				}
			}
		}
		return nil
	} else {
		return errors.New("image unchangeable")
	}
}

func (list *ImageList) Unload(flags Flags.Flags) error {
	if set := flags.CheckIfFlagsAreSet("alias"); !set {
		return errors.New("wanted flags unset")
	}
	alias := flags.Flag["alias"]

	for index := range list.LoadedImages {
		if list.LoadedImages[index].Alias == alias {
			list.LoadedImages = append(list.LoadedImages[:index], list.LoadedImages[index+1:]...)
		}
	}
	return nil
}
