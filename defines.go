package main

import (
	"ImageManipulationUnit/ImageUnit/colorModifiers"
	"ImageManipulationUnit/ImageUnit/utils"
)

type ImageList struct {
	utils.IImageList
	colorModifiers.ColorModifiersList
	utils.ImageList
}

type Selection struct{
	utils.ISelection
	utils.Selection
}

