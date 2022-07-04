package prism

import (
	"fmt"
	"prism-go/internal/helpers"
	"prism-go/internal/parser"
	"github.com/gdamore/tcell/v2"
	
)
/*
func Calculate(operators [3]string, t, n int) (map[int]map[int]int, int) {
	d := (n * 2) - 1
	p := make(map[int]map[int]int)
	for i := 0; i < d; i++ {
		p[i] = make(map[int]int)
		for j := 0; j < d; j++ {
			p[i][j] = 32
		}
	}
	o1 := operators[0]
	o2 := operators[1]
	o3 := operators[2]
	for z := 0; z < n; z++ {
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if p[y + z][x + z] >> 5 == 1 {
					//p[y + z][x + z] = parser.Calculate(fmt.Sprintf("%v%v%v%v%v%v%v", t, o1, z, o2, y, o3, x)) % n
					p[y + z][x + z] = helpers.Modulo(parser.Calculate(fmt.Sprintf("%v%v%v%v%v%v%v", t, o1, z, o2, y, o3, x)), n)

				}
			}
		}
	}
	return p, d
}

func Render(s tcell.Screen, c *helpers.Container, pause, drawUI, colour bool, termWH [2]int, p map[int]map[int]int, n, d int) {
	// border
	if pause == false {
		var c string
		if colour == true {
			c = "hue"
		} 
		if colour == false {
			c = "lightness"
		}
		s.Clear()
		for y := 0; y < d; y++ {
			for x := 0; x < d; x++ {
				if p[y][x] != 32 {
					style := tcell.StyleDefault.Foreground(helpers.Colour(c, p[y][x], n)).Background(tcell.ColorReset)  
					s.SetContent(
						helpers.Centre(termWH[0], d) + x, helpers.Centre(termWH[1], d) + y, 
						rune('█'), nil, style)
				}
			}
		}
	}
	if drawUI == true {
		c.DrawUI(s)
	}
	s.Show()
}*/



func Render(s tcell.Screen, colour bool, ops [3]string, termWH [2]int, n, t int) {
	// border
	var c string
	if colour == true {
		c = "hue-fun"
	} else {
		c = "lightness"
	}
	s.Clear()
	d := n * 2
	for z := 0; z < n; z++ {
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if z != (n - 1) {
					if x == 0 || y == 0 {
						thueMorseN := ThueMorse(x, y, z, t, n, ops[0], ops[1], ops[2])
						style := tcell.StyleDefault.Foreground(helpers.Colour(c, thueMorseN, n, t)).Background(tcell.ColorReset)  
						s.SetContent(
							helpers.Centre(termWH[0], d) + x + z, helpers.Centre(termWH[1], d) + y + z, 
							rune('█'), nil, style)
					}
				} else {
					thueMorseN := ThueMorse(x, y, z, t, n, ops[0], ops[1], ops[2])
					style := tcell.StyleDefault.Foreground(helpers.Colour(c, thueMorseN, n, t)).Background(tcell.ColorReset)  
					s.SetContent(
						helpers.Centre(termWH[0], d) + x + z, helpers.Centre(termWH[1], d) + y + z, 
						rune('█'), nil, style)
				}
			}
		}		
	}
}

func ThueMorse(t, z, y, x, n int, o1, o2, o3 string) int {
	//var toSum []int
	var parsed int
	/*
	for _, i := range []int{t, z, y, x} {
		toSum = append(toSum, numberToBase(i, b))
	}*/
	t = NumberToBase(t, n)
	z = NumberToBase(z, n)
	y = NumberToBase(y, n)
	x = NumberToBase(x, n)
	parsed = parser.Calculate(fmt.Sprintf("%v%v%v%v%v%v%v", t, o1, z, o2, y, o3, x))
	return helpers.Modulo(parsed, n)
}

func NumberToBase(n, b int) int {
	if n == 0 {
		return 0
	}
	var digits []int
	var condition bool = true
	for condition {
		digits = append(digits, helpers.Modulo(n, b))
		n = n / b
		if n == 0 {
			condition = false
		}
	}
	return Sum(digits)
}

func Sum(array []int) int {  
	 result := 0  
	 for _, v := range array {  
	  result += v  
	 }  
	 return result  
} 