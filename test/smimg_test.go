package smimg

import (
	"../src/smimg"
    "fmt"
    "io/ioutil"
	"testing"
)

func TestGetBoxes(t *testing.T) {
    data, err := ioutil.ReadFile("./test01.data")
    if (err != nil) {
        panic("File not found");
    }
    boxes := smimg.GetBoxes(data, 640)
    fmt.Println(boxes)
}

func BenchmarkGetBoxes(b *testing.B) {
	//640 * 400
	data, err := ioutil.ReadFile("./test01.data")
	if (err != nil) {
		panic("File not found");
	}
    boxes := smimg.GetBoxes(data, 640)
    for i := 0; i < b.N; i++ {
        boxes = smimg.GetBoxes(data, 640)
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