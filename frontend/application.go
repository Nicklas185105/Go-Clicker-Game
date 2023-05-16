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

	body string

	game *engine.Game
}

func NewApplication() *app {
	a := &app{}
	a.game = engine.NewGame()
	a.score = gowd.NewStyledText("Score: 0", gowd.BoldText)
	a.Element = gowd.NewElement("section")
	a.SetClass("container-fluid")
	em := gowd.NewElementMap()

	a.body = gameBody

	a.AddHTML(a.body, em)

	a.score = em["score"]
	em["click"].OnEvent(gowd.OnClick, a.click)

	return a
}
