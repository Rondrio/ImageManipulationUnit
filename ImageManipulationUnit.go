package main

import (
	"ImageManipulationUnit/CommandParser"
	"ImageManipulationUnit/Functions"
	"ImageManipulationUnit/ImageUnit"
	"os"
)

func main() {
	var ImageList ImageUnit.ImageList
	var SelectionList ImageUnit.Selection
	var FunctionList Functions.FunctionList

	ImageList.LoadedImages = make([]ImageUnit.Image, 0)
	FunctionList.List = make([]Functions.Function, 0)

	commands := ImageUnit.CommandList(GetCommandList())

	CommandParser.ScanInput(&ImageList, &SelectionList, &FunctionList, &commands, os.Stdin)
}

func GetCommandList() []ImageUnit.Command {
	return []ImageUnit.Command{
		ImageUnit.LoadCommand{Keyword: "load"},
		ImageUnit.ExportCommand{Keyword: "export"},
		ImageUnit.UnloadCommand{Keyword: "unload"},
		ImageUnit.InvertCommand{Keyword: "invert"},
		ImageUnit.GrayscaleCommand{Keyword: "grayscale"},
		ImageUnit.OverlayCommand{Keyword: "overlay"},
		ImageUnit.MergeCommand{Keyword: "merge"},
		ImageUnit.AddColorCommand{Keyword: "addcolor"},
		ImageUnit.MirrorCommand{Keyword: "mirror"},
		ImageUnit.SetColorCommand{Keyword: "setcolor"},
		ImageUnit.SelectCommand{Keyword: "select"},
		ImageUnit.UnselectCommand{Keyword: "unselect"},
	}
}
