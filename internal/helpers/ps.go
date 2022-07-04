package helpers

import (
	"math"
	"sync"
)

type PrismState struct {
	mu     	sync.Mutex
	pause 	bool
	drawUI 	bool
	termWH 	[2]int
	n      	int
}

func (ps *PrismState) InitPrismState(width, height int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.termWH[0] = width
	ps.termWH[1] = height
	ps.n = int(math.Round(float64(MinInt(width, height)-1) / float64(3)))
	ps.pause = false
	ps.drawUI = true
}

func (ps *PrismState) ReadPrismState() (bool, bool, [2]int, int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.pause, ps.drawUI, ps.termWH, ps.n
}

func (ps *PrismState) SetTermWH(width, height int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()	
	ps.termWH[0] = width
	ps.termWH[1] = height
}

func (ps *PrismState) SetN() {
	ps.mu.Lock()
	defer ps.mu.Unlock()	
	ps.n = int(math.Round(float64(MinInt(ps.termWH[0], ps.termWH[1])-1) / float64(3)))
}

func (ps *PrismState) SetPause() {
	ps.mu.Lock()
	defer ps.mu.Unlock()	
	if ps.pause == false {
		ps.pause = true
	} else {
		ps.pause = false
	}
}

func (ps *PrismState) SetDrawUI() {
	ps.mu.Lock()
	defer ps.mu.Unlock()	
	if ps.drawUI == false {
		ps.drawUI = true
	} else {
		ps.drawUI = false
	}
}

