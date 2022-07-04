package helpers

import (
	"log"
	"time"
)

func Sleep(timeNow time.Time, framerate int) {
	nap := time.Duration(1000000000/framerate) - time.Since(timeNow)
	if nap > 0 {
		time.Sleep(nap * time.Nanosecond)
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

