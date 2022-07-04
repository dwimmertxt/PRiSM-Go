package helpers

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func DrawString(s tcell.Screen, str string, x, y int, style tcell.Style) {
	for i, e := range str {
		s.SetContent(x+i, y, e, nil, style)
	}
}

func (is *InterfaceState) DrawUI(s tcell.Screen) {
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	
	DrawString(s, "     -PRiSM-", 0, 1, style)
	DrawString(s, "-----------------", 0, 3, style)
	DrawString(s, "Arrows: operators", 0, 5, style)
	DrawString(s, "F1:     options", 0, 6, style)
	DrawString(s, "F2:     pause", 0, 7, style)
	DrawString(s, "Esc:    quit", 0, 8, style)
	DrawString(s, "-----------------", 0, 10, style)

	is.DrawOperators(s)
	is.DrawColours(s)
}

func (is *InterfaceState) DrawOperators(s tcell.Screen) {
	is.mu.Lock()
	defer is.mu.Unlock()
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	
	if is.uiState == 0 {
		DrawString(s, ">> -Operators- <<", 0, 12, style)
	} else {
		DrawString(s, "   -Operators-", 0, 12, style)
	}
	DrawString(s, fmt.Sprintf(
		"  t %v z %v y %v x", 
		is.OperatorStringFromIndex(0), 
		is.OperatorStringFromIndex(1), 
		is.OperatorStringFromIndex(2)), 
		0, 14, style)
	DrawString(s, "^", (is.opIndex*4)+4, 15, style)
	DrawString(s, "-----------------", 0, 17, style)
}

func (is *InterfaceState) DrawColours(s tcell.Screen) {
	is.mu.Lock()
	defer is.mu.Unlock()
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)

	if is.uiState == 1 {
		DrawString(s, ">>  -Colours-  <<", 0, 19, style)
	} else {
		DrawString(s, "    -Colours-    ", 0, 19, style)
	}
	DrawString(s, " Hue   Sat  Light", 0, 21, style)
	for i, e := range is.colourState {
		if e == 0 {
			DrawString(s, "0", (i*6)+2, 23, style)
		}
		if e == 1 {
			DrawString(s, "100", (i*6)+1, 23, style)
		} 
		if e == 2 {
			DrawString(s, "TMn", (i*6)+1, 23, style)
		}
		if e == 3 {
			DrawString(s, "+t", (i*6)+2, 23, style)
		}
	}
	DrawString(s, "^", (is.colourIndex*6)+2, 24, style)	
}