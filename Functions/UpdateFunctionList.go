package Functions

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (List *FunctionList) UpdateListCycle(){
	ticker := time.NewTicker(5*time.Minute)
	for {
		<-ticker.C
		if err := List.CheckForNewFunction();err != nil{
			log.Println(err)
		}
	}
}

func (List *FunctionList) CheckForNewFunction()error{
	file,err := os.OpenFile("Functions",os.O_CREATE|os.O_RDWR,777)
	if err != nil{
		return err
	}
	if isDir,err := checkIfDir(file);!isDir || err!= nil{
		if err != nil{
			return err
		}
		return errors.New("functions file is not a dir")
	}
	filesInFunctions,err := file.Readdir(-1)
	if err := List.ParseFilesInDir(filesInFunctions);err != nil{
		return err
	}
	return nil
}

func checkIfDir(file *os.File)(bool,error){
	fileInfo,err := file.Stat()
	if err != nil{
		return false,err
	}
	if !fileInfo.IsDir(){
		return false,errors.New("functions is no directory")
	}
	return true,nil
}

func (List *FunctionList)ParseFilesInDir(files []os.FileInfo)error{
outerLoop:
	for _,fileInfo := range files{
		_,name := filepath.Split(fileInfo.Name())
		if strings.Contains(name,"rondrio") {
			continue outerLoop
		}
		if result := List.GetFunctionByKeyWord(strings.TrimSuffix(name,".rondrio"));result != nil{
			continue outerLoop
		}
		function := Function{
			KeyWord:  strings.TrimSuffix(name,".rondrio"),
		}
		List.List = append(List.List,function)
	}
	return nil
}