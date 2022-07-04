package helpers

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func DrawString(s tcell.Screen, str string, x, y int) {
	for i, e := range str {
		style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
		s.SetContent(x+i, y, e, nil, style)
	}
}

func (ps *PrismState) DrawUI(s tcell.Screen, is *InterfaceState) {
	DrawString(s, "     -PRiSM-", 0, 0)
	DrawString(s, "-----------------", 0, 1)
	DrawString(s, "Arrows: operators", 0, 2)
	DrawString(s, "F1:     options", 0, 3)
	DrawString(s, "F2:     pause", 0, 4)
	DrawString(s, "F3:     colour", 0, 5)
	DrawString(s, "Esc:    quit", 0, 6)
	DrawString(s, "-----------------", 0, 7)

	ps.DrawOperators(s, is)
}

func (ps *PrismState) DrawOperators(s tcell.Screen, is *InterfaceState) {
	DrawString(s, "   -Operators-", 0, 8)
	DrawString(s, fmt.Sprintf("  t %v z %v y %v x", is.OperatorStringFromIndex(0), is.OperatorStringFromIndex(1), is.OperatorStringFromIndex(2)), 0, 9)
	DrawString(s, "                 ", 0, 10)
	DrawString(s, "^", (is.opIndex*4)+4, 10)
}
