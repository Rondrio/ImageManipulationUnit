package Functions

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"ImageManipulationUnit/ImageUnit"
	"strings"
)

type Function struct {
	KeyWord  string
	Args     []string
	Commands []string
}
type FunctionList struct {
	List []Function
}

func (List *FunctionList) GetFunctionByKeyWord(keyWord string) (*Function, error) {
	for _, function := range List.List {
		if function.KeyWord == keyWord {
			return &function, nil
		}
	}
	function, err := List.CheckForFunctionFile(keyWord)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func (function *Function) ExecuteFunction(list *ImageUnit.ImageList, selection *ImageUnit.Selection, functions *FunctionList, flags Flags.Flags) ([]string, error) {

	functionFlags := make(map[string]string)

	for i := 0; i < len(function.Args); i++ {
		if flagExist := flags.CheckIfFlagsAreSet(function.Args[i]); !flagExist {
			return nil, Flags.ErrUnsetFlags
		}
		functionFlags[function.Args[i]] = flags.Flag[function.Args[i]]
	}
	commands := make([]string, 0)
	for _, line := range function.Commands {
		line = parseCommandsWithVars(line, functionFlags)
		commands = append(commands, line)
	}

	return commands, nil
}

func parseCommandsWithVars(command string, args map[string]string) string {
	var result string
	words := strings.Fields(command)
	for _, word := range words {
		if !strings.HasPrefix(word, "--") {
			result += word + " "
			continue
		}
		parts := strings.Split(word, "=")
		if strings.HasPrefix(parts[1], "$") {
			word = strings.Replace(word, parts[1], args[strings.TrimPrefix(parts[1], "$")], 1)
		}
		result += word + " "
	}
	return result
}
