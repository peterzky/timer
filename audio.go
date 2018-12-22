package main

import (
	"os"
	"os/exec"
	"path"
)

func play(f string) {
	filename, _ := os.Executable()
	cmd := exec.Command("aplay", path.Dir(filename)+"/../sound/"+f+".wav")
	cmd.Run()
}

func audio(c chan string) {
	for {
		switch <-c {
		case "start":
			play("start")
		case "end":
			play("end")
		case "5":
			play("alarm")
		}
	}
}
