package Functions

import (
	"io/ioutil"
	"os"
	"strings"
)

func ParseFunctionFile(file *os.File, keyWord string) (*Function, error) {
	function := &Function{
		KeyWord:  keyWord,
		Args:     make([]string, 0),
		Commands: make([]string, 0),
	}
	DataType := "Args"
	fileText, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileText), "\n")
	for _, line := range lines {
		if line == "Args :" {
			DataType = "Args"
			continue
		} else if line == "Methods :" {
			DataType = "Methods"
			continue
		}

		switch DataType {
		case "Args":
			function.Args = append(function.Args, strings.TrimPrefix(line, "--"))

		case "Methods":
			function.Commands = append(function.Commands, line)
		}
	}
	return function, nil
}
