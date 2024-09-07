package editor

import (
	"fmt"
	"log"

	"github.com/RowMur/gql/lexer"
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

	g.SetManagerFunc(mainLayout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	shouldQuery := false
	if err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		shouldQuery = true
		return quit(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	mainView, err := g.View("main")
	if err != nil {
		log.Panicln(err)
	}

	_, err = g.View("lexer")
	if err != nil {
		log.Panicln(err)
	}

	content := ""
	if shouldQuery {
		content = mainView.Buffer()
	}

	return &content
}

func mainLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	mv, err := g.SetView("main", 0, 0, 2*maxX/3-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		mv.Title = " query "
		mv.Editable = true
		mv.Wrap = true

		if _, err := g.SetCurrentView("main"); err != nil {
			log.Panicln(err)
		}

		g.SetViewOnTop("main")
	}

	lv, err := g.SetView("lexer", 2*maxX/3, 0, maxX, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		lv.Title = " lexer output "
		lv.Editable = false
		lv.Wrap = true
		lv.Autoscroll = true
	}

	tokens, err := lexer.Tokenize([]byte(mv.Buffer()))
	if err != nil {
		return err
	}
	lv.Clear()
	for _, token := range tokens {
		fmt.Fprintf(lv, "%s - %s\n", token.Name, token.Value)
	}

	return nil
}
