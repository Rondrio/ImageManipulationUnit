package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func (list *ImageList) LoadImage(flags Flags.Flags) error {
	var image Image

	if !flags.CheckIfFlagsAreSet("alias", "path") {
		return Flags.ErrUnsetFlags
	}
	alias := flags.Flag["alias"]
	path := flags.Flag["path"]

	aliasCheck := list.GetImageByAlias(alias)
	if aliasCheck != nil {
		return errors.New("alias already taken")
	}

	image.Alias = alias
	image.Path = path

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	image.Image, err = DecodeFile(file)
	if err != nil {
		return err
	}
	ImageTunnel = &image
	go DrawGUI(ImageTunnel)
	list.LoadedImages = append(list.LoadedImages, image)
	//go StartGUI(list.GetImageByAlias(alias))
	return nil
}

func DecodeFile(file *os.File) (image.Image, error) {
	switch strings.Split(file.Name(), ".")[1] {
	case "png":
		return png.Decode(file)
	case "jpg":
		return jpeg.Decode(file)
	default:
		return nil, errors.New("file type not supported")
	}
}
