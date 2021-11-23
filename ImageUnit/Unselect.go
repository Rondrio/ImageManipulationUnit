package ImageUnit

import "ImageManipulationUnit/CommandParser/Flags"

type UnselectCommand struct {
	Keyword string
}

func (cmd UnselectCommand) GetKeyword() string {
	return cmd.Keyword
}

func (cmd UnselectCommand) Execute(list *ImageList, flags Flags.Flags, selection *Selection) error {
	selection.Points = make([]Point, 0)
	return nil
}
