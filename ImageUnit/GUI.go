package ImageUnit

import (
	"ImageManipulationUnit/CommandParser/Flags"
	"log"
	"strconv"

	"golang.org/x/exp/shiny/unit"
	"golang.org/x/exp/shiny/widget/theme"

	"golang.org/x/exp/shiny/widget/node"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
)

func stretch(n node.Node, alongWeight int) node.Node {
	return widget.WithLayoutData(n, widget.FlowLayoutData{
		AlongWeight:  alongWeight,
		ExpandAlong:  true,
		ShrinkAlong:  true,
		ExpandAcross: true,
		ShrinkAcross: true,
	})
}

func StartGUI(img *Image) {
	go driver.Main(func(s screen.Screen) {
		w := UpdateImage(img)
		if err := widget.RunWindow(s, w, &widget.RunWindowOptions{
			NewWindowOptions: screen.NewWindowOptions{
				Height: img.Image.Bounds().Max.Y + 35,
				Width:  img.Image.Bounds().Max.X,
				Title:  "MemeMaker3000",
			},
		}); err != nil {
			log.Panic(err)
		}
	})
}

func ShowImage(flags Flags.Flags, list *ImageList) {
	var alias string
	if set := flags.CheckIfFlagsAreSet("alias"); set {
		alias = flags.Flag["alias"]
	}
	UpdateImage(list.GetImageByAlias(alias))
}

func UpdateImage(img *Image) *widget.Flow {
	header := widget.NewUniform(theme.Light,
		widget.NewPadder(widget.AxisVertical, unit.Ems(0.5),
			widget.NewFlow(widget.AxisHorizontal,
				widget.NewLabel("Alias : "+img.Alias),
				stretch(widget.NewSpace(), 1),
				widget.NewLabel("Size : "+strconv.Itoa(img.Image.Bounds().Max.X)+" x "+strconv.Itoa(img.Image.Bounds().Max.Y)),
			),
		),
	)
	body := widget.NewImage(img.Image, img.Image.Bounds())

	return widget.NewFlow(widget.AxisVertical,
		stretch(widget.NewSheet(header), 1),
		stretch(widget.NewSheet(body), 2),
	)
}
