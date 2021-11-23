package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type ExportCommand struct {
	Keyword string
}

func (cmd ExportCommand) GetKeyword() string {
	return cmd.Keyword
}

func (cmd ExportCommand) Execute(list *ImageList, flags Flags.Flags, selection *Selection) error {
	if !flags.CheckIfFlagsAreSet("output", "alias") {
		return Flags.ErrUnsetFlags
	}
	output := flags.Flag["output"]
	alias := flags.Flag["alias"]

	image := list.GetImageByAlias(alias)

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = image.EncodeImage(file); err != nil {
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
