
package main

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"fmt"
)

type Pixel struct {
    x int 
    y int 
    color color.Color
}

func getPixelsFromImage(imagePath string) [800]Pixel {
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
	const limit = 800
	var imgSet [limit]Pixel

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			oldPixel := img.At(x, y)
			_, g, _, a := oldPixel.RGBA()
			pixel := color.RGBA{uint8(g), uint8(g), uint8(g), uint8(a)}
			imgSet[x] = Pixel{x:x,y:y,color:pixel}
		}
	}

	return imgSet
}

func main() {

	pixels := getPixelsFromImage("./psmove.jpg")

	fmt.Println(pixels)
}