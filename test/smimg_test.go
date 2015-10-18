package smimg

import (
    "../src/smimg"
    "fmt"
    "io/ioutil"
    "testing"
    "os"
    "image"
    "image/color"
    "image/draw"
    "image/png"
)

const imageWidth = 100
const imageHeight = 100
const pixelByte = 3
const imagePath = "digitalImage.png"

func createDigitalImage(imagePath string){

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

func getPixelsFromImage(imagePath string) [imageWidth * imageHeight * pixelByte]byte {
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
    var imgSet2 [imageWidth * imageHeight * pixelByte] byte
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

func TestCreateDigitalImage(t *testing.T) {
    createDigitalImage(imagePath)
    fmt.Println("Image created")
}

func TestCreateDigitalInput(t *testing.T) {
    createDigitalImage(imagePath)
    fmt.Println("Image created")
    //data := getPixelsFromImage(imagePath) 
    getPixelsFromImage(imagePath) 
    
    // TODO: Avoid this problem: 
    // cannot use data (type [30000]byte) as type []byte in argument to smimg.GetBoxes
    //pattern := smimg.ColorPixel{uint8(255), uint8(0) , uint8(0)}
    //boxes := smimg.GetBoxes(data, 640, pattern)
    //fmt.Println(boxes)
}

func BenchmarkCreateDigitalImage(b *testing.B) {
    createDigitalImage(imagePath)
    fmt.Println("Image created")
}

// This method uses a image exported from GIMP 
// to get the number of Blobs
// TODO: What is t *testing.T?
func TestGetBoxes(t *testing.T) {
    data, err := ioutil.ReadFile("./test01.data")
    if (err != nil) {
        panic("File not found");
    }
    pattern := smimg.ColorPixel{uint8(255), uint8(0) , uint8(0)}
    boxes := smimg.GetBoxes(data, 640, pattern)
    fmt.Println(boxes)
}

func TestOpenDataFile(t *testing.T) {
    //640 * 400
    data, err := ioutil.ReadFile("./test01.data")
    if (err != nil) {
        panic("File not found");
    }
    total := len(data)
    fmt.Println(total)
    fmt.Println(640 * 400 * 3)
}

func BenchmarkGetBoxes(b *testing.B) {
    //640 * 400
    data, err := ioutil.ReadFile("./test01.data")
    if (err != nil) {
        panic("File not found");
    }
    pattern := smimg.ColorPixel{uint8(255), uint8(0) , uint8(0)}
    boxes := smimg.GetBoxes(data, 640, pattern)
    for i := 0; i < b.N; i++ {
        boxes = smimg.GetBoxes(data, 640, pattern)
    }
    fmt.Println(boxes)
}

func BenchmarkDivMod(b *testing.B) {
    x, y, w, z := 200, 34, 0, 0
    for i := 0; i < b.N; i++ {
        w, z = x / y, x % y
    }
    fmt.Println(w, z)
}

func BenchmarkEmpty(b *testing.B) {
    for i := 0; i < b.N; i++ {
    }
}