package helpers

import (
	"log"
	//"sync"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/hisamafahri/coco"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Modulo(a, b int) int {
	return ((a % b) + b) % b
}

func Centre(a int, b int) int {
	centre := (a / 2) - (b / 2) 
	return centre
}

func Colour(key string, index, n, t int) tcell.Color {
	var colour [3]float64
	switch key {
	case "hue-fun":
		hue := float64((360 / n) * index + t)
		sat := float64((100 / n) * index + t)
		light := float64((100 / n) * index + t)
		//colour = coco.Hsl2Rgb(hue, 100, 50)
		colour = coco.Hsl2Rgb(hue, sat, light)
	case "hue":
		hue := float64((360 / n) * index)
		sat := float64((100 / n) * index)
		light := float64((100 / n) * index)
		//colour = coco.Hsl2Rgb(hue, 100, 50)
		colour = coco.Hsl2Rgb(hue, sat, light)
	case "saturation":
		sat := float64((100 / n) * index)
		colour = coco.Hsl2Rgb(0, sat, 50)
	case "lightness":
		light := float64((100 / n) * index)
		colour = coco.Hsl2Rgb(0, 0, light)
	default:
	}
	r := int32(colour[0])
	g := int32(colour[1])
	b := int32(colour[2])
	return tcell.NewRGBColor(r, g, b)
}

func Sleep(timeNow time.Time, framerate int) {
	nap := time.Duration(1000000000 / framerate) - time.Since(timeNow)
	if nap > 0 {
		time.Sleep(nap * time.Nanosecond)
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func ParseRune(rune tcell.KeyRune, )