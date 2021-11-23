package main

import (
	"log"
	"time"

	"github.com/sshaman1101/actions/foo"
)

var version = "HEAD"

func main() {
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

	log.Printf("starting app version=%s\n", version)

	i := 1
	for {
		select {
		case <-t.C:
			// make one foo each 100msec
			f, _ := foo.New(i)
			log.Printf("got a foo instance: %v\n", f)
			i++
		}
	}
}
