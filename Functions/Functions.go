package Functions

import (
	"ImageManipulationUnit/CommandParser"
	"ImageManipulationUnit/ImageUnit"
	"io"
)

type Function struct{
	KeyWord string
	Args string
	Commands []string

}
type FunctionList struct{
	List []Function
}

func (List *FunctionList) GetFunctionByKeyWord(keyWord string)*Function{
 	for _,function := range List.List{
 		if function.KeyWord == keyWord{
 			return &function
		}
	}
	return nil
}

func (function *Function)ExecuteFunction(list *ImageUnit.ImageList, selection *ImageUnit.Selection, functions *FunctionList)error{
	var reader io.Reader

	for _,line := range function.Commands{
		if n,err := reader.Read([]byte(line));err != nil || n <len(line){
			return err
		}
		CommandParser.ScanInput(list,selection,functions,reader)
	}
}
