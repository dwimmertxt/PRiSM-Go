package helpers

import (
	"math"
	"sync"
)

type PrismState struct {
	mu     sync.Mutex
	pause  bool
	drawUI bool
	colour bool
	termWH [2]int
	n      int
}

type InterfaceState struct {
	mu        sync.Mutex
	opIndex   int
	opState   [3]int
	ctrlState bool
}

func (is *InterfaceState) InitInterfaceState() {
	is.opIndex = 0
	for i := 0; i < 3; i++ {
		is.opState[i] = 0
	}
}

func (ps *PrismState) InitPrismState(width, height int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.termWH[0] = width
	ps.termWH[1] = height
	ps.n = int(math.Round(float64(minInt(width, height)-1) / float64(3)))
	ps.pause = false
	ps.drawUI = true
	ps.colour = false
}

func (ps *PrismState) ReadPrismState() (bool, bool, bool, [2]int, int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.pause, ps.drawUI, ps.colour, ps.termWH, ps.n
}

func (ps *InterfaceState) CycleOperator(key string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	switch key {
	case "up":
		ps.opState[ps.opIndex] = Modulo((ps.opState[ps.opIndex] + 1), 3)
	case "down":
		ps.opState[ps.opIndex] = Modulo((ps.opState[ps.opIndex] - 1), 3)
	}
	/*
		switch ps.opState[ps.opIndex] {
		case 0:
			ps.operators[ps.opIndex] = "+"
		case 1:
			ps.operators[ps.opIndex] = "-"
		case 2:
			ps.operators[ps.opIndex] = "*"
		}
	*/
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

func (is *InterfaceState) GetOps() [3]string {
	var ops [3]string
	for i := 0; i < 3; i++ {
		ops[i] = is.OperatorStringFromIndex(is.opState[i])
	}
	return ops
}

func (is *InterfaceState) SelectOperator(key string) {
	is.mu.Lock()
	defer is.mu.Unlock()
	switch key {
	case "left":
		is.opIndex = Modulo(is.opIndex-1, 3)
	case "right":
		is.opIndex = Modulo(is.opIndex+1, 3)
	}
}

func (ps *PrismState) UpdateTermDimensions(width, height int) {
	ps.termWH[0] = width
	ps.termWH[1] = height
	ps.n = int(math.Round(float64(minInt(width, height)-1) / float64(3)))
}

func (ps *PrismState) UpdatePause() {
	if ps.pause == false {
		ps.pause = true
	} else {
		ps.pause = false
	}
}

func (ps *PrismState) UpdateUI() {
	if ps.drawUI == false {
		ps.drawUI = true
	} else {
		ps.drawUI = false
	}
}

func (ps *PrismState) UpdateColour() {
	if ps.colour == false {
		ps.colour = true
	} else {
		ps.colour = false
	}
}
