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
	var err error
	var alias, path string

	if flags.CheckIfFlagsAreSet("alias", "path") {
		alias = flags.Flag["alias"]
		path = flags.Flag["path"]
	} else {
		return errors.New("unset flags")
	}

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

	list.LoadedImages = append(list.LoadedImages, image)
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
