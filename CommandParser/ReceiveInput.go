package CommandParser

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"bufio"
	"log"
	"os"
	"strings"
)

func ScanCommandLine(list *ImageUnit.ImageList) {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		cmd = strings.Replace(cmd, "\r\n", "", -1)
		ParseCommand(cmd, list)
	}
}

func ParseCommand(cmd string, list *ImageUnit.ImageList) {
	var flags Flags.Flags
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
		err := list.LoadImage(flags)
		if err != nil {
			log.Println(err)
		}
	case "export":
		err := list.ExportImage(flags)
		if err != nil {
			log.Println(err)
		}
	case "grayscale":
		err := list.Grayscale(flags)
		if err != nil {
			log.Println(err)
		}
	case "add":
		err := list.AddColor(flags)
		if err != nil {
			log.Println(err)
		}
	case "invert":
		err := list.Invert(flags)
		if err != nil {
			log.Println(err)
		}
	case "set":
		err := list.SetColor(flags)
		if err != nil {
			log.Println(err)
		}
	}

}
