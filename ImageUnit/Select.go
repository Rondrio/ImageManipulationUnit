package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SelectCommand struct {
	Keyword string
}

func (cmd SelectCommand) GetKeyword() string {
	return cmd.Keyword
}

func (cmd SelectCommand) Execute(list *ImageList, flags Flags.Flags, selection *Selection) error {
	var numPoints int
	for key := range flags.Flag {
		if strings.Contains(key, "point") {
			numPoints++
		}
	}
	selection.Points = make([]Point, numPoints)
	for key, value := range flags.Flag {
		index := strings.TrimPrefix(key, "point")
		i, err := strconv.Atoi(index)
		if err != nil {
			return errors.New("point not indexed")
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
		selection.Points[i-1] = Point{xCoord, yCoord}
	}

	fmt.Println("Currently selected :", selection.Points)
	return nil
}

func (selection *Selection) CheckIfSelected(Point Point) bool {
	poly := append(selection.Points, selection.Points[0])
	c := false
	for i := 0; i < len(selection.Points); i++ {
		j := i + 1

		if (poly[i].Y > Point.Y) != (poly[j].Y > Point.Y) && (Point.X < poly[i].X+(poly[j].X-poly[i].X)*(Point.Y-poly[i].Y)/(poly[j].Y-poly[i].Y)) {
			c = !c
		}
	}
	return c
}
