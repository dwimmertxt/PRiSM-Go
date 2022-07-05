package prism

import (
	"fmt"
	"prism-go/internal/helpers"
	"prism-go/internal/parser"
	"github.com/hisamafahri/coco"
	"github.com/gdamore/tcell/v2"
)

func Render(s tcell.Screen, ops [3]string, colours [3]int, termWH [2]int, n, t int) {
	// border
	s.Clear()
	d := n * 2
	for z := 0; z < n; z++ {
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if ((z != (n - 1)) && (x == 0 || y == 0)) || z == (n - 1) {
					thueMorseN := thueMorse(n, t, z, y, x, ops[0], ops[1], ops[2])
					style := tcell.StyleDefault.Foreground(getColour(colours, thueMorseN, n, t)).Background(tcell.ColorReset)
					s.SetContent(
						helpers.Centre(termWH[0], d)+x+z, helpers.Centre(termWH[1], d)+y+z,
						rune('█'), nil, style)
				}
			}
		}
	}
}

func thueMorse(n, t, z, y, x int, o1, o2, o3 string) int {
	var parsed int
	t = helpers.NumberToBase(t, n)
	z = helpers.NumberToBase(z, n)
	y = helpers.NumberToBase(y, n)
	x = helpers.NumberToBase(x, n)
	parsed = parser.Calculate(fmt.Sprintf("%v%v%v%v%v%v%v", t, o1, z, o2, y, o3, x))
	return helpers.Modulo(parsed, n)
}

func getColour(colours [3]int, thueMorseN, n, t int) tcell.Color {
	var hue, sat, light float64
	var colour [3]float64
	for i, e := range colours {
		switch i {
		case 0:
			switch e {
			case 0:
				hue = float64(0)
			case 1:
				hue = float64((360/n)*helpers.Modulo(t, n))
			case 2:
				hue = float64((360/n)*helpers.Modulo(thueMorseN+t, n))
			case 3:
				hue = float64((360/n)*thueMorseN + t)
			}
		case 1:
			switch e {
			case 0:
				sat = float64(0)
			case 1:
				sat = float64(100)
			case 2:
				sat = float64((100/n)*thueMorseN)
			case 3:
				sat = float64((100/n)*thueMorseN + t)
			}
		case 2:
			switch e {
			case 0:
				light = float64(0)
			case 1:
				light = float64(50)
			case 2:
				light = float64((100/n)*thueMorseN)
			case 3:
				light = float64((100/n)*thueMorseN + t)
			}
		}
	}
	colour = coco.Hsl2Rgb(hue, sat, light)
	return tcell.NewRGBColor(int32(colour[0]), int32(colour[1]), int32(colour[2]))
}
