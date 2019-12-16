package colorModifiers

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit/utils"
)

type ColorModifiersList interface{
	AddColor(Flags.Flags,*utils.Selection)error
	Grayscale(Flags.Flags, *utils.Selection) error
	Invert(Flags.Flags, *utils.Selection) error
	SetColor(Flags.Flags,*utils.Selection) error
}

type ColorModImage interface{
	ChangeColor(uint32,uint32,uint32,uint32,*utils.Selection) error
	ChangeToGrayscale(*utils.ISelection) error
	InvertColor(*utils.ISelection) error
	SetColor(uint32,uint32,uint32,uint32,*utils.Selection) error
}