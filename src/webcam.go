package main

import (
	"fmt"
	"os/exec"
	"log"
)


func main() {
	cmd := exec.Command("/usr/bin/streamer", "-c", "/dev/video0", "-o", "frame.jpeg", "-s", "1600x1200")
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Frame stored")
}