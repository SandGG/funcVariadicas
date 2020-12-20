package main

import "fmt"

var figure = make([]string, 0)
var otherFigure = make([]string, 0)

func main() {
	calcArea(0, 0)
	calcArea(10, 10, "Square")
	calcArea(15, 20, "Triangle", "Rectangle", "Trapezium")

	figure = append(figure, "Square", "Triangle", "Rectangle", "Trapezium")
	otherFigure = append(otherFigure, figure...)
	otherFigure = append(otherFigure, "circle", "hexagon")

	calcArea(0, 0, otherFigure...)
	fmt.Println(otherFigure)

}

func calcArea(width, height int, figure ...string) {
	var area int = width * height
	if len(figure) <= 0 {
		fmt.Println("It isn't a figure")
	} else {
		for _, n := range figure {
			fmt.Println(n, area)
		}
	}

}
