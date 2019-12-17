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

func ScanInput(list *ImageUnit.ImageList, selection *ImageUnit.Selection, functions *Functions.FunctionList, input io.Reader) {

	reader := bufio.NewReader(input)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		cmd = strings.Replace(cmd, "\r\n", "", -1)
		ParseCommand(cmd, list, selection, functions)
	}
}

func ParseCommand(cmd string, list *ImageUnit.ImageList, selection *ImageUnit.Selection, functions *Functions.FunctionList) {
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
		if err := list.LoadImage(flags); err != nil {
			log.Println(err)
		}
	case "export":
		if err := list.ExportImage(flags); err != nil {
			log.Println(err)
		}
	case "grayscale":
		if err := list.Grayscale(flags, selection); err != nil {
			log.Println(err)
		}
	case "add":
		if err := list.AddColor(flags, selection); err != nil {
			log.Println(err)
		}
	case "invert":
		if err := list.Invert(flags, selection); err != nil {
			log.Println(err)
		}
	case "set":
		if err := list.SetColor(flags, selection); err != nil {
			log.Println(err)
		}
	case "mirror":
		if err := list.MirrorImage(flags); err != nil {
			log.Println(err)
		}
	case "select":
		if err := selection.Select(flags); err != nil {
			log.Println(err)
		}
	case "merge":
		if err := list.Merge(flags); err != nil {
			log.Println(err)
		}
	case "overlay":
		if err := list.Overlay(flags); err != nil {
			log.Println(err)
		}
	case "unload":
		if err := list.Unload(flags); err != nil {
			log.Println(err)
		}
	case "unselect":
		selection.Points = make([]ImageUnit.Point, 0)
	default:
		function, err := functions.GetFunctionByKeyWord(keyWord)
		if err != nil {
			log.Println(err)
		}
		commands, err := function.ExecuteFunction(list, selection, functions, flags)
		if err != nil {
			log.Println(err)
		}
		for _, command := range commands {
			ParseCommand(command, list, selection, functions)
		}
	}
}


func recovery() {
	if err := recover(); err != nil {
		log.Println("error with flags", err)
	}
}
