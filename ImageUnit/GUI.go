package ImageUnit

import (
	//"fmt"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"image"
	"log"
)
var Win screen.Window

func DrawGUI(img *Image) {
	var err error
	driver.Main(func(s screen.Screen) {
		Win, err = s.NewWindow(&screen.NewWindowOptions{
			Title: "MemeMaker3000",
			Width: img.Image.Bounds().Max.X,
			Height: img.Image.Bounds().Max.Y,
		})

		if err != nil {
			log.Fatal(err)
		}
		defer Win.Release()

		imgText, err := s.NewTexture(image.Point{X: img.Image.Bounds().Max.X, Y: img.Image.Bounds().Max.Y})
		if err != nil {
			log.Fatal(err)
		}
		defer imgText.Release()
		c, err := s.NewBuffer(image.Point{X: img.Image.Bounds().Max.X, Y: img.Image.Bounds().Max.Y})
		imgText.Upload(image.Point{}, c, img.Image.Bounds())
		for {
			e := Win.NextEvent()

			switch e := e.(type) {

			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}

			case paint.Event:
				if e.External{
					break
				}

				for height := 0; height < img.Image.Bounds().Max.Y; height++ {
					for width := 0; width < img.Image.Bounds().Max.X; width++ {
						imgText.Fill(image.Rect(width-1, height-1, width+1, height+1), img.Image.At(width, height), screen.Src)
					}
				}

				Win.Copy(image.Point{0, 0}, imgText, imgText.Bounds(), screen.Src, nil)
				Win.Publish()
				Win.SendFirst(paint.Event{})


			case error:
				log.Print(e)
			}
		}
	})
}
