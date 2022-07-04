package helpers

import (
	"sync"
)

type InterfaceState struct {
	mu      		sync.Mutex
	opIndex 		int
	opState 		[3]int
	uiState 		int
	colourIndex		int
	colourState		[3]int	 
}

func (is *InterfaceState) InitInterfaceState() {
	is.opIndex = 0
	is.colourIndex = 0
	is.uiState = 0
	is.colourState[0] = 0
	is.colourState[1] = 0
	is.colourState[2] = 3 
	is.opState[0] = 0
	is.opState[1] = 1
	is.opState[2] = 1
}

func (is *InterfaceState) SelectState(key, state string) {
	is.mu.Lock()
	defer is.mu.Unlock()
	switch key {
	case "left":
		if state == "opIndex" {
			is.opIndex = Modulo(is.opIndex-1, 3)
		} 
		if state == "colourIndex" {
			is.colourIndex = Modulo(is.colourIndex-1, 3)
		}	
	case "right":
		if state == "opIndex" {
			is.opIndex = Modulo(is.opIndex+1, 3)
		} 
		if state == "colourIndex" {
			is.colourIndex = Modulo(is.colourIndex+1, 3)
		}	
	}
}

func (is *InterfaceState) CycleState(key, state string) {
	is.mu.Lock()
	defer is.mu.Unlock()
	switch key {
	case "up":
		if state == "opState" {
			is.opState[is.opIndex] = Modulo((is.opState[is.opIndex] + 1), 3)
		}
		if state == "colourState" {
			is.colourState[is.colourIndex] = Modulo((is.colourState[is.colourIndex] + 1), 4)
		}		
	case "down":
		if state == "opState" {
			is.opState[is.opIndex] = Modulo((is.opState[is.opIndex] - 1), 3)
		}
		if state == "colourState" {
			is.colourState[is.colourIndex] = Modulo((is.colourState[is.colourIndex] - 1), 4)
		}		
	}
}

func (is *InterfaceState) CycleUIState() {
	is.mu.Lock()
	defer is.mu.Unlock()
	is.uiState = Modulo(is.uiState + 1, 2)
}

func (is *InterfaceState) OperatorStringFromIndex(index int) string {
	switch is.opState[index] {
	case 0:
		return "+"
	case 1:
		return "-"
	case 2:
		return "*"
	default:
		return ""
	}
}

func (is *InterfaceState) GetUIState() int {
	is.mu.Lock()
	defer is.mu.Unlock()
	return is.uiState
}

func (is *InterfaceState) GetOps() [3]string {
	is.mu.Lock()
	defer is.mu.Unlock()
	var ops [3]string
	for i := 0; i < 3; i++ {
		ops[i] = is.OperatorStringFromIndex(is.opState[i])
	}
	return ops
}

func (is *InterfaceState) GetColours() [3]int {
	is.mu.Lock()
	defer is.mu.Unlock()
	return is.colourState
}


