package utils

import (
    "fmt"
    "os"
    "image"
    "image/color"
    "image/draw"
    "image/png"
)

const imageWidth = 100
const imageHeight = 100
const pixelByte = 3

func CreateDigitalImage(imagePath string){

    //Create an image object with some dimensions
    img := image.NewRGBA(image.Rect(0,0,imageWidth,imageHeight))

    //Define colors used in the digital examples
    white := color.RGBA{uint8(255), uint8(255) , uint8(255), uint8(255)}
    red := color.RGBA{uint8(255), uint8(0) , uint8(0), uint8(255)}

    //Add a white background.
    draw.Draw(img, image.Rect(0,  0, imageWidth, imageHeight), &image.Uniform{white}, image.ZP, draw.Src)

    //Add some red blobs.
    draw.Draw(img, image.Rect(50, 50, 55, 55), &image.Uniform{red}, image.ZP, draw.Src)

    toimg, err := os.Create("digitalImage.png")
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer toimg.Close()

    png.Encode(toimg, img)

}

func GetPixelsFromImage(imagePath string) []byte {
    file, err := os.Open(imagePath)
    if err != nil {
        panic("File not found");
    }
    defer file.Close()

    img, err := png.Decode(file)
    if err != nil {
        panic("Problem with PNG file");
    }

    b := img.Bounds()

    //Defining a Fixed array
    imgSet2 := make([]byte, (b.Max.Y - b.Min.Y) * (b.Max.X - b.Min.X) * 3)
    var i int = 0
    for y := b.Min.Y; y < b.Max.Y; y++ {
        for x := b.Min.X; x < b.Max.X; x++ {
            oldPixel := img.At(x, y)
            r, g, b, _ := oldPixel.RGBA()
            imgSet2[i] = byte(r)
            i++
            imgSet2[i] = byte(g)
            i++
            imgSet2[i] = byte(b)
            i++
        }
    }

    return imgSet2
}