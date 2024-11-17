package main

import (
	//"encoding/json" // use it for saving data
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, _ := gocui.NewGui(gocui.OutputNormal)
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, handl_exit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()

	if v, err := g.SetView("main", maxX/2-15, maxY/2, maxX/2+15, maxY/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Wrap = true
		v.Title = "Welcome to lazy-money"
		if _, err := g.SetCurrentView("main"); err != nil {
			return err
		}
		g.SetViewOnTop("main")
	}

	if v, err := g.SetView("money_bottle", maxX/2-30, maxY/2-15, maxX/2+30, maxY/2+15); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = true
		fmt.Fprintln(v, "here will be your money in bottle")
		g.SetViewOnBottom("money_bottle")
	}

	return nil
}

func handl_exit(gui *gocui.Gui, view *gocui.View) error {
	return gocui.ErrQuit
}
