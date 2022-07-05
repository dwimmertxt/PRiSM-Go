package helpers

import (
	"math"
	"sync"
)

type PrismState struct {
	mu     	sync.Mutex
	pause 	bool
	termWH 	[2]int
	fps		int
	n      	int
}

func (ps *PrismState) InitPrismState(width, height int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.termWH[0] = width
	ps.termWH[1] = height
	ps.fps = 24
	ps.n = int(math.Round(float64(MinInt(width, height)-1) / float64(3)))
	ps.pause = false
}

func (ps *PrismState) ReadPrismState() (bool, [2]int, int, int) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.pause, ps.termWH, ps.fps, ps.n
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

func (ps *PrismState) GetTermWH() [2]int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.termWH
}



