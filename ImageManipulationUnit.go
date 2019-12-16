package main

import (
	"ImageManipulationUnit/CommandParser"
	"ImageManipulationUnit/Functions"
	"ImageManipulationUnit/ImageUnit/utils"
	"os"
)

func main() {
	var ImageList CommandParser.ImageList
	var SelectionList CommandParser.Selection
	var FunctionList Functions.FunctionList

	ImageList.L.LoadedImages = make([]utils.Image, 0)
	FunctionList.List = make([]Functions.Function, 0)

	var header CommandParser.Header
	header.Selection = SelectionList
	header.ImageList = ImageList
	header.Functions = &FunctionList

	CommandParser.ScanInput(header, os.Stdin)
}
