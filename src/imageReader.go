/*
package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"fmt"
	"reflect"
)

type Pixel struct {
    color color.Color
}

type Blob struct {
	color color.Color 
    x1 int
    x2 int
    y1 int
    y2 int
}

func getPixelsFromImage(imagePath string) [800][530]color.Color {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(os.Stderr, "%s: %v\n", imagePath, err)
	}

	b := img.Bounds()

	//Defining a Fixed array
	var imgSet [800][530]color.Color

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			oldPixel := img.At(x, y)
			_, g, _, a := oldPixel.RGBA()
			pixel := color.RGBA{uint8(g), uint8(g), uint8(g), uint8(a)}
			imgSet[x][y] = pixel
		}
	}

	return imgSet
}


func getBlobs(imageData [800][530]color.Color, colorPattern color.Color ) Blob {

	//https://golang.org/pkg/image/color/

	var counter int
	counter = 1

	for y := 0; y < 530; y++ {
		for x := 0; x < 800; x++ {
			pixel := imageData[x][y]

			if (reflect.DeepEqual(pixel, colorPattern)){
				fmt.Println(counter, x, y, "FOUND")
				counter++
			}
			

		}
	}

	a := Blob{color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},1,1,1,1}
	return a
}

func main() {

	imageFramePath := "./psmove.jpg"
	imageData := getPixelsFromImage(imageFramePath)
	//x, y := getImageDimension(imageData)
	//x := 800
	//y := 530

	//Cyan
	colorPattern := color.RGBA{uint8(143), uint8(143), uint8(143), uint8(255)}
	getBlobs(imageData, colorPattern)

}

*/