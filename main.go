package main

import (
	//"fmt"
	"log"
	"os"
	"prism-go/internal/helpers"
	"prism-go/internal/prism"
	"time"

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

	ps := new(helpers.PrismState)
	width, height := s.Size()
	ps.InitPrismState(width, height)

	is := new(helpers.InterfaceState)
	is.InitInterfaceState()

	go func(s tcell.Screen, ps *helpers.PrismState, is *helpers.InterfaceState) {
		t := 0
		for {
			pause, drawUI, colour, termWH, n := ps.ReadPrismState()
			timeNow := time.Now()
			if !pause {
				prism.Render(s, colour, is.GetOps(), termWH, n, t)
				t++
			} else {
				prism.Render(s, colour, is.GetOps(), termWH, n, t)
			}
			if drawUI {
				ps.DrawUI(s, is)
			}
			s.Show()
			helpers.Sleep(timeNow, 24)
		}
	}(s, ps, is)

	for {
		s.Show()
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			width, height := s.Size()
			ps.UpdateTermDimensions(width, height)
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				quit()
			case tcell.KeyCtrlC:
				quit()
			case tcell.KeyF1:
				ps.UpdateUI()
			case tcell.KeyF2:
				ps.UpdatePause()
			case tcell.KeyF3:
				ps.UpdateColour()
			case tcell.KeyLeft:
				is.SelectOperator("left")
			case tcell.KeyRight:
				is.SelectOperator("right")
			case tcell.KeyUp:
				is.CycleOperator("up")
			case tcell.KeyDown:
				is.CycleOperator("down")
			case tcell.KeyEnter:
			}

		}
	}
}
