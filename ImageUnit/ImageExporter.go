package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func (list *ImageList) ExportImage(flags Flags.Flags) error {
	var output, alias string
	if flags.CheckIfFlagsAreSet("output", "alias") {
		alias = flags.Flag["alias"]
		output = flags.Flag["output"]
	} else {
		return errors.New("unset flags")
	}

	image := list.GetImageByAlias(alias)

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	err = image.EncodeImage(file)
	if err != nil {
		return err
	}
	return nil
}

func (image *Image) EncodeImage(file *os.File) error {
	switch strings.Split(file.Name(), ".")[1] {
	case "png":
		return png.Encode(file, image.Image)
	case "jpg":
		return jpeg.Encode(file, image.Image, nil)
	default:
		return errors.New("file type not supported")
	}
}
