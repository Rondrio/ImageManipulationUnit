package CommandParser

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ScanCommandLine(list *ImageUnit.ImageList, selection *ImageUnit.Selection) {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		cmd = strings.Replace(cmd, "\r\n", "", -1)
		ParseCommand(cmd, list, selection)
	}
}

func ParseCommand(cmd string, list *ImageUnit.ImageList, selection *ImageUnit.Selection) {
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

	switch strings.ToLower(words[0]) {
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
	default:
		fmt.Println("Command not recognized")
	}

}

func recovery() {
	if err := recover(); err != nil {
		log.Println("error with flags",err)
	}
}
