package ImageUnit

import (
	"errors"
	"image"
	"image/color"
)

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
