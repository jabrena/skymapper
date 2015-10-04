package main

import (
	"github.com/blackjack/webcam"
	"fmt"
	"reflect"
	"os"
)


func main() {
	pic := webcam.GetImg("/dev/video0")
	fmt.Println(reflect.TypeOf(pic))
	fo, err := os.Create("image")
	if err != nil {
		panic(err)
	}

	if _, err := fo.Write(pic); err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
}