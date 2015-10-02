package main

import "fmt"
import "github.com/lucasb-eyer/go-colorful"
import "image"
import "image/draw"
import "image/png"
import "os"

func main() {
    blocks := 10
    blockw := 40
    img := image.NewRGBA(image.Rect(0,0,blocks*blockw,530))

    c1, _ := colorful.Hex("#ffffff")
    c2, _ := colorful.Hex("#ffffff")
    c3, _ := colorful.Hex("#000000")

    for i := 0 ; i < blocks ; i++ {
        draw.Draw(img, image.Rect(i*blockw,  0,(i+1)*blockw, 40), &image.Uniform{c3}, image.ZP, draw.Src)
        draw.Draw(img, image.Rect(i*blockw, 40,(i+1)*blockw, 80), &image.Uniform{c3}, image.ZP, draw.Src)
        draw.Draw(img, image.Rect(i*blockw, 80,(i+1)*blockw,120), &image.Uniform{c1.BlendRgb(c2, float64(i)/float64(blocks-1))}, image.ZP, draw.Src)
        draw.Draw(img, image.Rect(i*blockw,120,(i+1)*blockw,160), &image.Uniform{c1.BlendLab(c2, float64(i)/float64(blocks-1))}, image.ZP, draw.Src)
        draw.Draw(img, image.Rect(i*blockw,160,(i+1)*blockw,200), &image.Uniform{c1.BlendHcl(c2, float64(i)/float64(blocks-1))}, image.ZP, draw.Src)
    }

    toimg, err := os.Create("testingImage.png")
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer toimg.Close()

    png.Encode(toimg, img)
}