package main

import (
	"os/exec"
	"runtime"
)

func play(f string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	cmd := exec.Command("aplay", "../sound/"+f+".wav")
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
