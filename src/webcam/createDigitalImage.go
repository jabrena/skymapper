package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"
import "image"
import "image/color"
import "image/draw"
import "image/png"
import "os"
import "log"
import "reflect"

type Blob struct {
    color color.Color 
    x1 int
    x2 int
    y1 int
    y2 int
}

const imageWidth = 100
const imageHeight = 100
const imagePath = "digitalImage.png"

func createDigitalImage(imagePath string){

    //Create an image object with some dimensions
    img := image.NewRGBA(image.Rect(0,0,imageWidth,imageHeight))

    //Define colors used in the digital examples
    white, _ := colorful.Hex("#ffffff")
    red, _ := colorful.Hex("#ff0000")

    //Add a white background.
    draw.Draw(img, image.Rect(0,  0, imageWidth, imageHeight), &image.Uniform{white}, image.ZP, draw.Src)

    //Add some red blobs.
    draw.Draw(img, image.Rect(50, 50, 55, 55), &image.Uniform{red}, image.ZP, draw.Src)
    //draw.Draw(img, image.Rect(550, 50, 600, 100), &image.Uniform{red}, image.ZP, draw.Src)
    //draw.Draw(img, image.Rect(50, 300, 100, 350), &image.Uniform{red}, image.ZP, draw.Src)
    //draw.Draw(img, image.Rect(300, 200, 400, 300), &image.Uniform{red}, image.ZP, draw.Src)

    toimg, err := os.Create("digitalImage.png")
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer toimg.Close()

    png.Encode(toimg, img)

}

func getPixelsFromImage(imagePath string) [imageWidth][imageHeight]color.Color {
    file, err := os.Open(imagePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    img, err := png.Decode(file)
    if err != nil {
        log.Fatal(os.Stderr, "%s: %v\n", imagePath, err)
    }

    b := img.Bounds()

    //Defining a Fixed array
    var imgSet [imageWidth][imageHeight]color.Color

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

func getBlobs(imageData [imageWidth][imageHeight]color.Color, colorPattern color.Color ) Blob {

    //https://golang.org/pkg/image/color/

    var counter int
    counter = 1

    for y := 0; y < imageHeight; y++ {
        for x := 0; x < imageHeight; x++ {
            pixel := imageData[x][y]

            //fmt.Println(pixel,colorPattern);

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

    createDigitalImage(imagePath)
    imageData := getPixelsFromImage(imagePath)

    //Patter: RED
    colorPattern := color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)}
    getBlobs(imageData, colorPattern)
}