package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
	"prism-go/internal/helpers"
	"prism-go/internal/prism"
	"time"
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
			pause, drawUI, termWH, n := ps.ReadPrismState()
			timeNow := time.Now()
			if !pause {
				prism.Render(s, is.GetOps(), is.GetColours(), termWH, n, t)
				t++
			} else {
				prism.Render(s, is.GetOps(), is.GetColours(), termWH, n, t)
			}
			if drawUI {
				is.DrawUI(s)
			}
			s.Show()
			helpers.Sleep(timeNow, 60)
		}
	}(s, ps, is)
	for {
		s.Show()
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
			width, height := s.Size()
			ps.SetTermWH(width, height)
			ps.SetN()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				quit()
			case tcell.KeyCtrlC:
				quit()
			case tcell.KeyF1:
				ps.SetDrawUI()
			case tcell.KeyF2:
				ps.SetPause()
				case tcell.KeyLeft:
					switch is.GetUIState() {
					case 0:
						is.SelectState("left", "opIndex")
					case 1:
						is.SelectState("left", "colourIndex")
					}
				case tcell.KeyRight:
					switch is.GetUIState() {
					case 0:
						is.SelectState("right", "opIndex")
					case 1:
						is.SelectState("right", "colourIndex")
					}
				case tcell.KeyUp:
					switch is.GetUIState() {
					case 0:
						is.CycleState("up", "opState")
					case 1:
						is.CycleState("up", "colourState")
					}
				case tcell.KeyDown:
					switch is.GetUIState() {
					case 0:
						is.CycleState("up", "opState")
					case 1:
						is.CycleState("down", "colourState")
					}
				case tcell.KeyTab:
					is.CycleUIState()
			}
		}
	}
}
