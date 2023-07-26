package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	soja := flag.Int("sleep", 5, "probing delay in seconds")
	flag.Parse()

	verifyString := []byte("Monitors: 1")
	for {
		xrandrOutput, err := exec.Command("xrandr", "--listmonitors").Output()
		if err != nil {
			log.Fatal(err)
		}
		if bytes.Contains(xrandrOutput, verifyString) {
			fmt.Println("HDMI port connected")
		} else {
			fmt.Println("! HDMI port not connected !")
			// issue shutdown
		}
		time.Sleep(time.Second * time.Duration(*soja)) // strange behavior in Pi for 5 and 10 seconds
	}
}
