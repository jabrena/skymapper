package smimg

import (
    "../src/smimg"
    "fmt"
    "io/ioutil"
    "testing"
    "./utils"
)

const imagePath = "digitalImage.png"

func TestCreateDigitalImage(t *testing.T) {
    utils.CreateDigitalImage(imagePath)
    fmt.Println("Image created")
}

func TestCreateDigitalInput(t *testing.T) {
    utils.CreateDigitalImage(imagePath)
    fmt.Println("Image created")
    data := utils.GetPixelsFromImage(imagePath)
    //utils.GetPixelsFromImage(imagePath) 
    
    // TODO: Avoid this problem: 
    // cannot use data (type [30000]byte) as type []byte in argument to smimg.GetBoxes
    pattern := smimg.Color{uint8(255), uint8(0) , uint8(0)}
    boxes := smimg.GetBoxes(data, 100, pattern)
    if len(boxes) != 1 {
        fmt.Println(boxes)
        t.Errorf("CreateDigitalInput Got: %d , Want: %d", len(boxes), 1)
    }
}

func BenchmarkCreateDigitalImage(b *testing.B) {
    utils.CreateDigitalImage(imagePath)
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
    pattern := smimg.Color{uint8(255), uint8(0) , uint8(0)}
    boxes := smimg.GetBoxes(data, 640, pattern)
    if len(boxes) != 4 {
        fmt.Println(boxes)
        t.Errorf("Got: %d , Want: %d", len(boxes), 4)
    }  
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
    pattern := smimg.Color{uint8(255), uint8(0) , uint8(0)}
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