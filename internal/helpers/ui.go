package helpers

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func (is *InterfaceState) DrawUI(s tcell.Screen, termWH [2]int, n, fps int) {
	var yCount int 
	if is.GetShowInfo() {
		yCount = 23
	} else {
		yCount = 1
	}
	y := &yCount
	x := 3
	is.drawKeyInfo(s, y, x)
	is.drawOperators(s, y, x)
	is.drawColours(s, y, x)
	is.drawVars(s, termWH, y, x, n, fps)
}

func drawUIString(s tcell.Screen, style tcell.Style, str string, y *int, x int) {
	for i, e := range str {
		s.SetContent(x+i, *y, e, nil, style)
	}
	*y = *y + 1
}

func insertToString(baseStr, toInsert string, index int) string {
	returnStr := baseStr[:index] + toInsert + baseStr[index:]
	return returnStr
}

func (is *InterfaceState) drawKeyInfo(s tcell.Screen, y *int, x int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	drawUIString(s, style, "    -Options-    ", y, x)
	drawUIString(s, style, "-----------------", y, x)
	drawUIString(s, style, "Arrows: operators", y, x)
	drawUIString(s, style, "Tab:    section  ", y, x)
	drawUIString(s, style, "F1:     options  ", y, x)
	drawUIString(s, style, "F2:     info     ", y, x)
	drawUIString(s, style, "F3:     pause    ", y, x)
	drawUIString(s, style, "Esc:    quit     ", y, x)
	drawUIString(s, style, "-----------------", y, x)
}

func (is *InterfaceState) drawOperators(s tcell.Screen, y *int, x int) {
	is.mu.Lock()
	defer is.mu.Unlock()
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	
	if is.uiState == 0 {
		drawUIString(s, style, ">> -Operators- <<", y, x)
	} else {
		drawUIString(s, style, "   -Operators-   ", y, x)
	}
	drawUIString(s, style, fmt.Sprintf(
		"  t %v z %v y %v x  ", 
		is.OperatorStringFromIndex(0), 
		is.OperatorStringFromIndex(1), 
		is.OperatorStringFromIndex(2)), 
		y, x)
	drawUIString(s, style, insertToString("                 ", "^", (is.opIndex*4)+4), y, x)
	drawUIString(s, style, "-----------------", y, x)
}

func (is *InterfaceState) drawColours(s tcell.Screen, y *int, x int) {
	is.mu.Lock()
	defer is.mu.Unlock()
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)

	if is.uiState == 1 {
		drawUIString(s, style, ">>  -Colours-  <<", y, x)
	} else {
		drawUIString(s, style, "    -Colours-    ", y, x)
	}
	drawUIString(s, style, " Hue   Sat  Light", y, x)
	
	var strState, strHue, strSat, strLight string
	for i, e := range is.colourState {
		switch e {
		case 0:
			strState = " 0 "
		case 1:
			strState = "100"
		case 2:
			strState = "TMn"
		}
		switch i {
		case 0:
			strHue = strState
		case 1:
			strSat = strState
		case 2:
			strLight = strState
		}
	}
	drawUIString(s, style, fmt.Sprintf(" %v   %v   %v ", strHue, strSat, strLight), y, x)
	drawUIString(s, style, insertToString("                 ", "^", (is.colourIndex*6)+2), y, x)
	drawUIString(s, style, "-----------------", y, x)
}

func (is *InterfaceState) drawVars(s tcell.Screen, termWH [2]int, y *int, x, n, fps  int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	if is.uiState == 2 {
		drawUIString(s, style, ">> -Variables- <<", y, x)
	} else {
		drawUIString(s, style, "   -Variables-   ", y, x)
	}

	drawUIString(s, style, fmt.Sprintf("Term W: %v", termWH[0]), y, x)
	drawUIString(s, style, fmt.Sprintf("Term W: %v", termWH[1]), y, x)
	drawUIString(s, style, fmt.Sprintf("Base N: %v", n), y, x)
	drawUIString(s, style, fmt.Sprintf("tgtFPS: %v", fps), y, x)

}

func (is *InterfaceState) DrawInfo(s tcell.Screen, termWH [2]int) {
	yCount := 2
	y := &yCount
	x := 2
	style := tcell.StyleDefault.Foreground(tcell.ColorReset).Background(tcell.ColorReset)
	
	drawUIString(s, style, "            -- PRiSM --            ", y, x)
	drawUIString(s, style, "|---------------------------------|", y, x)
	drawUIString(s, style, "|   A visual exploration of the   |", y, x)
	drawUIString(s, style, "|                                 |", y, x)
	drawUIString(s, style, "|   --- Thue-Morse sequence ---   |", y, x)
	drawUIString(s, style, "|                                 |", y, x)
	drawUIString(s, style, "| generalised to any base number, |", y, x)
	drawUIString(s, style, "| extended into four dimensions:  |", y, x)
	drawUIString(s, style, "| three spatial and one temporal. |", y, x)
	drawUIString(s, style, "|                                 |", y, x)
	drawUIString(s, style, "| \"Mathematics are the result of  |", y, x)
	drawUIString(s, style, "| mysterious powers which no one  |", y, x)
	drawUIString(s, style, "|    understands, and which the   |", y, x)
	drawUIString(s, style, "|    unconscious recognition of   |", y, x)
	drawUIString(s, style, "|  beauty must play an important  |", y, x)
	drawUIString(s, style, "|   part. Out of an infinity of   |", y, x)
	drawUIString(s, style, "| designs a mathematician chooses |", y, x)
	drawUIString(s, style, "|  one pattern for beauty's sake  |", y, x)
	drawUIString(s, style, "|   and pulls it down to earth.\"  |", y, x)
	drawUIString(s, style, "|       -- Marston Morse --       |", y, x)
	drawUIString(s, style, "|---------------------------------|", y, x)
	
	
	if is.GetShowUI() {
		drawUIString(s, style, "   -Details-   ", y, 21)
		drawUIString(s, style, "---------------", y, 21)
		drawUIString(s, style, "Art:    dwimmer", y, 21)
		*y = *y + 1
		drawUIString(s, style, "Font:   Square ", y, 21)
		*y = *y + 1
		drawUIString(s, style, "Author: Wouter ", y, 21)
		drawUIString(s, style, "van Oortmerssen", y, 21)
	} else {
		drawUIString(s, style, "                     dwimmer 2022  ", y, x)
	}
}
