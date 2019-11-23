package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (selection *Selection) Select(flags Flags.Flags) error {

	for key, value := range flags.Flag {
		if !strings.Contains(key, "point") {
			return errors.New("no point given")
		}
		coords := strings.Split(value, "|")
		xCoord, err := strconv.Atoi(coords[0])
		if err != nil {
			return err
		}
		yCoord, err := strconv.Atoi(coords[1])
		if err != nil {
			return err
		}
		selection.Points = append(selection.Points, Point{X: xCoord, Y: yCoord})
	}
	return nil
}

func (selection *Selection) CheckIfSelected(Point Point) bool {
	poly := append(selection.Points,selection.Points[0])
	num := len(selection.Points)
	j := num-1
	c := false
	for i:=0;i<=num;i++{
		if (poly[i].Y > Point.Y) != (poly[j].Y > Point.Y) &&
			(Point.X < poly[i].X + (poly[j].X - poly[i].X) *(Point.Y - poly[i].Y)/(poly[j].Y - poly[i].Y)){
			c = !c
		}
		j=i
	}
	return c
}