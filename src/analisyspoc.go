package main

import (
	"io/ioutil"
	"fmt"
)

type Point struct { x, y int }
type Rectangle struct { min, max Point }

func main() {
	const HEIGTH, WIDTH = 400, 640
	const TOTAL = HEIGTH * WIDTH
	rects := make([]Rectangle, 0, 4)
	data, err := ioutil.ReadFile("./test01.data")
	if (err != nil) {
		panic("File not found");
	}
	var r, g, b byte
	var y, x int
	added := false
	for pixel := 0; pixel < TOTAL; pixel++ {
		i := pixel * 3
		r, g, b = data[i], data[i + 1], data[i + 2]
		y, x = pixel / WIDTH, pixel % WIDTH;
		if r == 255 && g == 0 && b == 0 {
			added = false
			size := len(rects)
			for index := 0; index < size; index++ {
				rect := &rects[index] 
				if rect.max.y == y && rect.max.x + 1 == x {
					rect.max.x++
					added = true
				} else if rect.max.y + 1 >= y && rect.min.x -1 <= x && rect.max.x + 1 >= x {
					rect.max.y = y
					added = true
				}
			}
			if !added {
				rects = append(rects, Rectangle{Point{x, y}, Point{x, y}})
			}	
		}
	}
	fmt.Println(rects)
}