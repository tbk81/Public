package main

import (
	"fmt"
	"github.com/Pitrified/go-turtle"
	"image/color"
)

func circle(td *turtle.TurtleDraw) {
	// move somewhere else
	td.PenUp()
	td.SetPos(450, 300)

	// draw a circle with increasing brightness
	td.PenDown()
	td.SetHeading(turtle.North)
	td.SetSize(5)
	for i := 0; i < 360; i++ {
		val := uint8(float64(i) * 255 / 360)
		td.SetColor(color.RGBA{val, val / 2, 0, 255})
		td.Right(1)
		td.Forward(3)
	}
}

func main() {
	//w := turtle.NewWorld(900, 600)
	w := turtle.NewWorldWithColor(900, 900, turtle.White)

	//img := image.NewRGBA(image.Rect(0, 0, 900, 600))
	//draw.Draw(img, img.Bounds(), &image.Uniform{turtle.Cyan}, image.Point{0, 0}, draw.Src)
	//wi := turtle.NewWorldWithImage(img)

	// create a turtle attached to w
	td := turtle.NewTurtleDraw(w)

	// position/orientation
	td.SetPos(50, 50)
	td.SetHeading(turtle.East)

	// line style
	td.SetColor(turtle.Blue)
	td.SetSize(4)

	// start drawing
	td.PenDown()

	// same interface as Turtle
	td.Forward(200)
	td.Left(90)
	td.Forward(200)

	circle(td)

	err := w.SaveImage("world.png")
	if err != nil {
		fmt.Println("Could not save the image: ", err)
	}

	w.Close()

	// this is an error: the turtle tries to send the line
	// to the world input channel that has been closed
	// td.Forward(50)
}
