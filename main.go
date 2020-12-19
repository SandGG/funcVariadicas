package main

import "fmt"

func main() {
	calcArea(0, 0)
	calcArea(10, 10, "Square")
	calcArea(15, 20, "Triangle", "Rectangle", "Trapezium")
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
