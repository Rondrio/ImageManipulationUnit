package CommandParser

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/Functions"
	"ImageManipulationUnit/ImageUnit"
	"bufio"
	"io"
	"log"
	"strings"
)

type Header struct{
	ImageList *ImageUnit.ImageList
	Selection *ImageUnit.Selection
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
		if err := header.ImageList.LoadImage(flags); err != nil {
			log.Println(err)
		}
	case "export":
		if err := header.ImageList.ExportImage(flags); err != nil {
			log.Println(err)
		}
	case "grayscale":
		if err := header.ImageList.Grayscale(flags, header.Selection); err != nil {
			log.Println(err)
		}
	case "add":
		if err := header.ImageList.AddColor(flags, header.Selection); err != nil {
			log.Println(err)
		}
	case "invert":
		if err := header.ImageList.Invert(flags, header.Selection); err != nil {
			log.Println(err)
		}
	case "set":
		if err := header.ImageList.SetColor(flags, header.Selection); err != nil {
			log.Println(err)
		}
	case "mirror":
		if err := header.ImageList.MirrorImage(flags); err != nil {
			log.Println(err)
		}
	case "select":
		if err := header.Selection.Select(flags); err != nil {
			log.Println(err)
		}
	case "merge":
		if err := header.ImageList.Merge(flags); err != nil {
			log.Println(err)
		}
	case "overlay":
		if err := header.ImageList.Overlay(flags); err != nil {
			log.Println(err)
		}
	case "unload":
		if err := header.ImageList.Unload(flags); err != nil {
			log.Println(err)
		}
	case "unselect":
		header.Selection.Points = make([]ImageUnit.Point, 0)
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
		}
	}
}

func recovery() {
	if err := recover(); err != nil {
		log.Println("error with flags", err)
	}
}
