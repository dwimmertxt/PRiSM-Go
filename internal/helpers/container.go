package helpers

import (
	"math"
	"sync"
)

type PrismState struct {
	mu			sync.Mutex
	pause		bool
	drawUI		bool 
	colour 		bool
	termWH		[2]int
	n			int
	operators 	[3]string
}

type InterfaceState struct {
	opIndex 	int
	opState 	[3]int
	ctrlState	bool
}

func (c *Container) InitContainer(width, height int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.termWH[0] = width
	c.termWH[1] = height
	c.n = int(math.Round(float64(minInt(width, height) - 1) / float64(3)))
	c.pause = false
	c.drawUI = true
	c.colour = false
	c.opIndex = 0
	for i := 0; i < 3; i++ {
    	c.opState[i] = 0
    	c.operators[i] = "-"
	}
}

func (c *Container) ReadContainer() (bool, bool, bool, [2]int, int, [3]string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.pause, c.drawUI, c.colour, c.termWH, c.n, c.opIndex, c.opState, c.operators
}

func (c *Container) CycleOperator(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	switch key {
	case "up":
		c.opState[c.opIndex] = Modulo((c.opState[c.opIndex] + 1), 3)
	case "down":
		c.opState[c.opIndex] = Modulo((c.opState[c.opIndex] - 1), 3)
	}
	switch c.opState[c.opIndex] {
	case 0:
		c.operators[c.opIndex] = "+"
	case 1:
		c.operators[c.opIndex] = "-"
	case 2:
		c.operators[c.opIndex] = "*"
	}
}

func (c *Container) SelectOperator(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	switch key {
	case "left":
		c.opIndex = Modulo(c.opIndex - 1, 3)
	case "right":
		c.opIndex = Modulo(c.opIndex + 1, 3)
	}
}

func (c *Container) UpdateTermDimensions(width, height int) {
	c.termWH[0] = width
	c.termWH[1] = height
	c.n = int(math.Round(float64(minInt(width, height) - 1) / float64(3)))
}

func (c *Container) UpdatePause() {
	if c.pause == false {
		c.pause = true
	} else {
		c.pause = false
	}
}

func (c *Container) UpdateUI() {
	if c.drawUI == false {
		c.drawUI = true
	} else {
		c.drawUI = false
	}
}

func (c *Container) UpdateColour() {
	if c.colour == false {
		c.colour = true
	} else {
		c.colour = false
	}
}



