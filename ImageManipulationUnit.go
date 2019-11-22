package main

import (
	"ImageManipulationUnit/CommandParser"
	"ImageManipulationUnit/ImageUnit"
	"time"
)

func main() {
	var ImageList ImageUnit.ImageList
	var SelectionList ImageUnit.Selection
	ImageList.LoadedImages = make([]ImageUnit.Image, 0)

	go CommandParser.ScanCommandLine(&ImageList,&SelectionList)
	for {
		time.Sleep(1 * time.Second)
	}
}
