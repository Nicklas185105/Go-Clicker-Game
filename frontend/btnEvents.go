package frontend

import (
	"github.com/dtylman/gowd"
)

func (a *app) click(sender *gowd.Element, event *gowd.EventElement) {
	a.game.Click()
	a.score.SetText(a.game.GetScore().String())
}

func (a *app) start(sender *gowd.Element, event *gowd.EventElement) {
	a.homeBody.Hide()
	a.gameBody.Show()
}
