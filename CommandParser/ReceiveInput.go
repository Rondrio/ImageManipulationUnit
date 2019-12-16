package CommandParser

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/Functions"
	"ImageManipulationUnit/ImageUnit/colorModifiers"
	"ImageManipulationUnit/ImageUnit/utils"
	"bufio"
	"io"
	"log"
	"strings"
)
type ImageList struct{
	IL utils.IImageList
	CML colorModifiers.ColorModifiersList
	L utils.ImageList
}

type Selection struct{
	IS utils.ISelection
	S utils.Selection
}



type Header struct{
	ImageList ImageList
	Selection Selection
	Functions *Functions.FunctionList
}

func ScanInput(header Header, input io.Reader) {

	reader := bufio.NewReader(input)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		cmd = strings.Replace(cmd, "\r\n", "", -1)
		ParseCommand(cmd, header)
	}
}

func ParseCommand(cmd string, header Header) {
	var flags Flags.Flags
	var err error

	defer recovery()
	flags.Flag = make(map[string]string)
	words := strings.Fields(cmd)
	for _, word := range words {
		if strings.HasPrefix(word, "-") {
			parts := strings.Split(word, "=")
			flags.Flag[strings.Replace(parts[0], "-", "", 1)] = parts[1]
		}
	}

	switch keyWord := strings.ToLower(words[0]); keyWord {
	case "load":
		if err := header.ImageList.L.LoadImage(flags); err != nil {
			log.Println(err)
		}
	case "export":
		if err := header.ImageList.L.ExportImage(flags); err != nil {
			log.Println(err)
		}
	case "grayscale":
		if err := header.ImageList.CML.Grayscale(flags, &header.Selection.S); err != nil {
			log.Println(err)
		}
	case "add":
		if err := header.ImageList.CML.AddColor(flags, &header.Selection.S); err != nil {
			log.Println(err)
		}
	case "invert":
		if err := header.ImageList.CML.Invert(flags, &header.Selection.S); err != nil {
			log.Println(err)
		}
	case "set":
		if err := header.ImageList.CML.SetColor(flags, &header.Selection.S); err != nil {
			log.Println(err)
		}
/*	case "mirror":
		if err := header.ImageList.MirrorImage(flags); err != nil {
			log.Println(err)
		}*/
	case "select":
		header.Selection.IS,err = header.Selection.IS.Select(flags)
		if err != nil {
			log.Println(err)
		}
/*	case "merge":
		if err := header.ImageList.Merge(flags); err != nil {
			log.Println(err)
		}
	case "overlay":
		if err := header.ImageList.Overlay(flags); err != nil {
			log.Println(err)
		}
	case "unload":
		header.ImageList,err = header.ImageList.Unload(flags)
		if err != nil {
			log.Println(err)
		}
	case "unselect":
		header.Selection.Points = make([]utils.Point, 0)
	default:
		function, err := header.Functions.GetFunctionByKeyWord(keyWord)
		if err != nil {
			log.Println(err)
		}
		commands, err := function.ExecuteFunction(flags)
		if err != nil {
			log.Println(err)
		}
		for _, command := range commands {
			ParseCommand(command, header)
		}*/
	}
}

func recovery() {
	if err := recover(); err != nil {
		log.Println("error with flags", err)
	}
}
