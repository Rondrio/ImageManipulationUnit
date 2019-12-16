package utils

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image"
	"image/color"
	"os"
)

const (
	Max16bit = 65535
)

type IImageList interface{
	GetImageByAlias(string)IImage
	Unload(Flags.Flags)(IImageList,error)
	LoadImage(Flags.Flags)error
	ExportImage(Flags.Flags)error
}
type ISelection interface{
	Select(Flags.Flags)(ISelection,error)
	CheckIfSelected(Point)bool
}
type IImage interface{
	IterateOverPixels(func(width, height int, img SetColor),*Selection) error
	EncodeImage(*os.File) error
}


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

func (list ImageList) GetImageByAlias(alias string) IImage {
	for _, image := range list.LoadedImages {
		if image.Alias == alias {
			return image
		}
	}
	return nil
}

func (imgStruct Image) IterateOverPixels(paint func(width, height int, img SetColor), selection *Selection) error {
	if img, ok := imgStruct.Image.(SetColor); ok {
		for height := 0; height < imgStruct.Image.Bounds().Max.Y; height++ {
			for width := 0; width < imgStruct.Image.Bounds().Max.X; width++ {
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
		return errors.New("imgStruct unchangeable")
	}
}

func (list ImageList) Unload(flags Flags.Flags) (IImageList,error) {
	if set := flags.CheckIfFlagsAreSet("alias"); !set {
		return nil,errors.New("wanted flags unset")
	}
	alias := flags.Flag["alias"]

	for index := range list.LoadedImages {
		if list.LoadedImages[index].Alias == alias {
			list.LoadedImages = append(list.LoadedImages[:index], list.LoadedImages[index+1:]...)
		}
	}
	return list,nil
}
