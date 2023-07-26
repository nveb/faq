package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	check := []byte("Monitors: 0")
	for {
		out, err := exec.Command("xrandr", "--listmonitors").Output()
		if err != nil {
			log.Fatal(err)
		}
		if bytes.Contains(out, check) {
			fmt.Println("HDMI port connected")
		} else {
			fmt.Println("HDMI port not connected !")
		}
		time.Sleep(10000)
	}
}
