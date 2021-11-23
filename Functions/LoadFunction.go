package Functions

import (
	"io/ioutil"
	"os"
	"strings"
)

func (List *FunctionList) CheckForFunctionFile(keyWord string) (*Function, error) {
	var function *Function
	FunctionDirInfo, err := ioutil.ReadDir("Functions")
	if err != nil {
		return nil, err
	}
	for _, info := range FunctionDirInfo {
		if strings.TrimSuffix(info.Name(), ".rondrio") == keyWord {
			currentPath, err := os.Getwd()
			if err != nil {
				return nil, err
			}
			file, err := os.Open(currentPath + "/Functions/" + info.Name())
			if err != nil {
				return nil, err
			}
			function, err = ParseFunctionFile(file, keyWord)
			if err != nil {
				return nil, err
			}
		}
	}

	List.List = append(List.List, *function)
	return function, nil
}
