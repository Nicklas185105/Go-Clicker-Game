package frontend

import (
	"clicker-game/engine"

	"github.com/dtylman/gowd"
)

type app struct {
	*gowd.Element

	score      *gowd.Element
	cps        *gowd.Element
	clicks     *gowd.Element
	clickValue *gowd.Element

	em gowd.ElementsMap

	homeBody *gowd.Element
	gameBody *gowd.Element

	game *engine.Game
}

func NewApplication() *app {
	// Set up the app
	a := &app{}
	a.game = engine.NewGame()
	a.score = gowd.NewStyledText("Score: 0", gowd.BoldText)
	a.Element = gowd.NewElement("section")
	a.SetClass("container-fluid")
	a.em = gowd.NewElementMap()

	a.pagesSetup()

	a.elementsSetup()

	a.eventsSetup()

	return a
}

func (a *app) pagesSetup() {
	a.homeBody = gowd.NewElement("div")
	a.homeBody.SetClass("row")
	a.homeBody.AddHTML(homeBody, a.em)

	a.gameBody = gowd.NewElement("div")
	a.gameBody.SetClass("row")
	a.gameBody.AddHTML(gameBody, a.em)

	// Add the pages to the app
	a.AddElement(a.homeBody)
	a.AddElement(a.gameBody)

	// Hide the pages that is not supposed to be shown
	a.gameBody.Hide()
}

func (a *app) elementsSetup() {
	a.score = a.em["score"]
	a.cps = a.em["cps"]
	a.clicks = a.em["clicks"]
	a.clickValue = a.em["clickValue"]
}

func (a *app) eventsSetup() {
	a.em["click"].OnEvent(gowd.OnClick, a.click)
	a.em["start"].OnEvent(gowd.OnClick, a.start)
}
