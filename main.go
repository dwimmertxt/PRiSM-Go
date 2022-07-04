package main

import (
	//"fmt"
	"log"
	"os"
	"time"
	"prism-go/internal/helpers"
	"prism-go/internal/prism"
	"github.com/gdamore/tcell/v2"
	
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()
	
	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	c := new(helpers.Container)
	width, height := s.Size()
	c.InitContainer(width, height)
	
	go func(s tcell.Screen, c *helpers.Container) {
		t := 0
		for {
			pause, drawUI, colour, termWH, n, operators := c.ReadContainer() 
			timeNow := time.Now()
			if pause == false {
				prism.Render(s, colour, operators, termWH, n, t)
				t++
			} else {
				prism.Render(s, colour, operators, termWH, n, t)
			}
			if drawUI == true {
				c.DrawUI(s)
			}
			s.Show()
			helpers.Sleep(timeNow, 24)
			}
	}(s, c)	

	for {
		s.Show()
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			width, height := s.Size()
			c.UpdateTermDimensions(width, height)
		case *tcell.EventKey: 
			switch ev.Key() {
			case tcell.KeyEscape:
				quit()
			case tcell.KeyCtrlC:
				quit()
			case tcell.KeyF1:
				c.UpdateUI()
			case tcell.KeyF2:
				c.UpdatePause()
			case tcell.KeyF3:
				c.UpdateColour()
			case tcell.KeyLeft:
				c.SelectOperator("left")	
			case tcell.KeyRight:
				c.SelectOperator("right")
			case tcell.KeyUp:
				c.CycleOperator("up")
			case tcell.KeyDown:
				c.CycleOperator("down")
			case tcell.KeyEnter:					
			}

		}
	}		
}
