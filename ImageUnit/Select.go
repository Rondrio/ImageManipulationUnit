package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"strconv"
	"strings"
)


func (selection *Selection) Select(flags Flags.Flags)error{

	for key,value := range flags.Flag{
		if !strings.Contains(key,"point"){
			return errors.New("no point given")
		}
		coords := strings.Split(value,"|")
		xCoord,err := strconv.Atoi(coords[0])
		if err != nil{
			return err
		}
		yCoord,err := strconv.Atoi(coords[1])
		if err != nil{
			return err
		}
		selection.Points = append(selection.Points,Point{X: xCoord,Y:yCoord})
	}



	return nil
}

func (selection *Selection)CheckIfSelected(Point Point){
	//TODO algorithmus erdenken
}