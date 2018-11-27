package main

import (
	"fmt"
	"strconv"
	"time"
)

func digit(i int) string {
	switch {
	case i < 10:
		return fmt.Sprintf("<fc=#87d37c>%s</fc>", "0"+strconv.Itoa(i))
	default:
		return fmt.Sprintf("<fc=#87d37c>%s</fc>", strconv.Itoa(i))
	}
}

func formatTime(i int) string {
	min := digit(i / 60)
	sec := digit(i % 60)
	return fmt.Sprintf("<fc=#d64541><fn=1></fn></fc>  [%s:%s]\n", min, sec)
}

func timer(t int, c, p chan string, q, done chan bool) {
	tick := time.Tick(time.Second)
Loop:
	for {
		if t <= -1 {
			p <- "end"
			done <- true
			break Loop
		}
		if t == 30 {
			select {
			case p <- "30":
			default:
			}

		}
		select {
		case <-q:
			c <- formatTime(t)
		case <-tick:
			t--
		default:
		}
		time.Sleep(10 * time.Millisecond)
	}

}
