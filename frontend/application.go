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
	a := &app{}
	a.game = engine.NewGame()
	a.score = gowd.NewStyledText("Score: 0", gowd.BoldText)
	a.Element = gowd.NewElement("section")
	a.SetClass("container-fluid")
	a.em = gowd.NewElementMap()

	a.homeBody = gowd.NewElement("div")
	a.homeBody.SetClass("row")
	a.homeBody.AddHTML(homeBody, a.em)

	a.gameBody = gowd.NewElement("div")
	a.gameBody.SetClass("row")
	a.gameBody.AddHTML(gameBody, a.em)

	a.AddElement(a.homeBody)
	a.AddElement(a.gameBody)

	a.gameBody.Hide()

	a.score = a.em["score"]
	a.em["click"].OnEvent(gowd.OnClick, a.click)
	a.em["start"].OnEvent(gowd.OnClick, a.start)

	return a
}
