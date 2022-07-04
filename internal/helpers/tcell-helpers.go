package helpers

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func DrawString(s tcell.Screen, str string, x, y int) {
	for i, e := range str {
		style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
		s.SetContent(x + i, y, e, nil, style)
	}
}

func (c *Container) DrawUI(s tcell.Screen) {
	DrawString(s, "     -PRiSM-", 0, 0)
	DrawString(s, "-----------------", 0, 1)
	DrawString(s, "Arrows: operators", 0, 2)
	DrawString(s, "F1:     options", 0, 3)
	DrawString(s, "F2:     pause", 0, 4)
	DrawString(s, "F3:     colour", 0, 5)
	DrawString(s, "Esc:    quit", 0, 6)
	DrawString(s, "-----------------", 0, 7)
	
	c.DrawOperators(s)
}

func (c *Container) DrawOperators(s tcell.Screen) {
	DrawString(s, "   -Operators-", 0, 8)
	DrawString(s, fmt.Sprintf("  t %v z %v y %v x", c.operators[0], c.operators[1], c.operators[2]), 0, 9)
	DrawString(s, "                 ", 0, 10)
	DrawString(s, "^", (c.opIndex * 4) + 4, 10)
}

