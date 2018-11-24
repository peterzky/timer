package main

import (
	"flag"
	"os"
	"strconv"
)

var client bool

func init() {
	flag.BoolVar(&client, "c", false, "timer client report time")
}

func main() {
	flag.Parse()
	if client {
		SocketClient()
	} else {
		t, _ := strconv.Atoi(os.Args[1])
		p := make(chan string)
		c := make(chan string)
		q := make(chan bool)
		done := make(chan bool)
		go SocketServer(3333, c, q)
		go timer(60*t, c, p, q, done)
		go audio(p)
		p <- "start"
		<-done
	}

}
