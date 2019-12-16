package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"fmt"
	"image"
	"image/color"
)

func (list *ImageUnit.ImageList) Overlay(flags Flags.Flags) error {
	if !flags.CheckIfFlagsAreSet("alias1", "alias2") {
		return Flags.ErrUnsetFlags
	}
	alias2 := flags.Flag["alias2"]
	alias1 := flags.Flag["alias1"]
	image1 := list.GetImageByAlias(alias1)
	image2 := list.GetImageByAlias(alias2)

	result, err := image1.OverlayImages(image2)
	if err != nil {
		return err
	}
	list.LoadedImages = append(list.LoadedImages, ImageUnit.Image{
		Id:    0,
		Alias: "overlaid",
		Path:  "",
		Image: result,
	})

	fmt.Println("overlaid image has alias 'overlaid'")
	return nil
}

func (image1 *ImageUnit.Image) OverlayImages(image2 *ImageUnit.Image) (image.Image, error) {
	maxX := image1.Image.Bounds().Max.X
	if image2.Image.Bounds().Max.X > image1.Image.Bounds().Max.X {
		maxX = image2.Image.Bounds().Max.X
	}
	maxY := image1.Image.Bounds().Max.Y
	if image2.Image.Bounds().Max.Y > image1.Image.Bounds().Max.Y {
		maxX = image2.Image.Bounds().Max.Y
	}
	rect := image.Rect(0, 0, maxX, maxY)
	result := image.NewRGBA64(rect)

	paint := func(height, width int, img ImageUnit.SetColor) {
		color1 := image1.Image.At(width, height)
		color2 := image2.Image.At(width, height)

		result.Set(width, height, mixColors(color1, color2))
	}

	err := image1.IterateOverPixels(paint, nil)

	return result, err
}

func mixColors(color1, color2 color.Color) color.Color {
	r1, g1, b1, a1 := color1.RGBA()
	r2, g2, b2, a2 := color2.RGBA()

	return color.RGBA64{
		R: uint16(r1 + (r1 - r2)),
		G: uint16(g1 + (g1 - g2)),
		B: uint16(b1 + (b1 - b2)),
		A: uint16(a1 + (a1 - a2)),
	}

}
