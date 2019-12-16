package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"fmt"
	"image"
)

func (list *ImageList) Merge(flags Flags.Flags) error {
	if !flags.CheckIfFlagsAreSet("alias1", "alias2") {
		return Flags.ErrUnsetFlags
	}
	alias2 := flags.Flag["alias2"]
	alias1 := flags.Flag["alias1"]
	image1 := list.GetImageByAlias(alias1)
	image2 := list.GetImageByAlias(alias2)

	list.LoadedImages = append(list.LoadedImages, Image{
		Id:    0,
		Alias: "merged",
		Path:  "",
		Image: image1.MergeImages(image2),
	})
	fmt.Println("merged image has alias 'merged'")
	return nil
}

func (image1 *Image) MergeImages(image2 *Image) image.Image {
	rect := image.Rect(0, 0, image1.Image.Bounds().Max.X+image2.Image.Bounds().Max.X, image1.Image.Bounds().Max.Y+image2.Image.Bounds().Max.Y)
	result := image.NewRGBA64(rect)
	firstImage := true

	for height := 0; height < rect.Bounds().Max.Y; height++ {
		for width := 0; width < rect.Bounds().Max.X; width++ {

			if firstImage {
				result.Set(width, height, image1.Image.At(width, height))
			} else {
				result.Set(width, height, image2.Image.At(width, height))
			}
			firstImage = !firstImage
		}
		firstImage = !firstImage
	}
	return result
}
