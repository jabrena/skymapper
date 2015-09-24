
package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"fmt"
)

/*
type ImageSet interface {
	Set(x, y int, c color.Color)
}
*/
type PixelSet struct {
    x int 
    y int 
    pixel color.Color
}

type Pixels struct {
    pixs []PixelSet
}

func main() {
	file, err := os.Open("./psmove.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", "./psmove.jpg", err)
	}

	b := img.Bounds()

	//imgSet := Pixels{}
	//var imgSet [5]PixelSet
	for y := b.Min.Y; y < b.Max.Y; y++ {
		fmt.Println(y);
		for x := b.Min.X; x < b.Max.X; x++ {

			oldPixel := img.At(x, y)
			_, g, _, a := oldPixel.RGBA()
			fmt.Println(g, b, a)
			//pixel := color.RGBA{uint8(g), uint8(g), uint8(g), uint8(a)}
			//imgSet := PixelSet{x,y, pixel}
			
		}
	}
	//fmt.Println(pixel)
}