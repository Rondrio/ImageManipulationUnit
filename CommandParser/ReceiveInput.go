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

func ScanInput(list *ImageUnit.ImageList, selection *ImageUnit.Selection, functions *Functions.FunctionList, commands *ImageUnit.CommandList, input io.Reader) {

	reader := bufio.NewReader(input)
	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		cmd = strings.Replace(cmd, "\r\n", "", -1)
		ParseCommand(cmd, list, selection, functions, commands)
	}
}

func ParseCommand(cmd string, list *ImageUnit.ImageList, selection *ImageUnit.Selection, functions *Functions.FunctionList, commands *ImageUnit.CommandList) {
	var flags Flags.Flags
	defer recovery()
	flags.Flag = make(map[string]string)
	words := strings.Fields(cmd)
	for _, word := range words {
		if strings.HasPrefix(word, "--") {
			parts := strings.Split(word, "=")
			flags.Flag[strings.Replace(parts[0], "--", "", 1)] = parts[1]
		}
	}

	for _, command := range *commands {
		if strings.ToLower(words[0]) == command.GetKeyword() {
			command.Execute(list, flags, selection)
			return
		}
	}

	switch keyword := strings.ToLower(words[0]); keyword {
	default:
		function, err := functions.GetFunctionByKeyWord(keyword)
		if err != nil {
			log.Println(err)
		}
		functionCommands, err := function.ExecuteFunction(list, selection, functions, flags)
		if err != nil {
			log.Println(err)
		}
		for _, command := range functionCommands {
			ParseCommand(command, list, selection, functions, commands)
		}
	}

}

func recovery() {
	if err := recover(); err != nil {
		log.Println("error with flags", err)
	}
}
