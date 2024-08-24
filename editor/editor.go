package editor

import (
	"log"

	"github.com/jroimartin/gocui"
)

type Editor struct {
}

func NewEditor() *Editor {
	return &Editor{}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (e *Editor) Run() *string {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	mainView, err := g.View("main")
	if err != nil {
		log.Panicln(err)
	}
	content := mainView.Buffer()

	return &content
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("main", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " query "
		v.Editable = true
		v.Wrap = true

		if _, err := g.SetCurrentView("main"); err != nil {
			log.Panicln(err)
		}

		g.SetViewOnTop("main")
	}

	return nil
}
