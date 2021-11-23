package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
)

type UnloadCommand struct {
	Keyword string
}

func (cmd UnloadCommand) GetKeyword() string {
	return cmd.Keyword
}

func (cmd UnloadCommand) Execute(list *ImageList, flags Flags.Flags, selection *Selection) error {
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
