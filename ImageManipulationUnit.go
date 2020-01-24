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

	CommandParser.ScanInput(&ImageList, &SelectionList, &FunctionList, os.Stdin)
}
