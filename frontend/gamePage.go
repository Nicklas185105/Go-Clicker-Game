package frontend

import "github.com/dtylman/gowd"

const gameBody = `
<div id="gameBody">
	<div class="row">
		<div class="col-md-12">
			<div class="well">
				<h2>Score: <span id="score">0</span></h2>
				<h3>Clicks per second: <span id="cps">0</span></h3>
				<h3>Clicks: <span id="clicks">0</span></h3>
				<h3>Click value: <span id="clickValue">1</span></h3>
			</div>
		</div>
	</div>
	<button class="btn btn-primary" id="click">Click</button>
	<button class="btn btn-primary" id="buy">Buy</button>
</div>
`

func (a *app) click(sender *gowd.Element, event *gowd.EventElement) {
	a.game.Click()
	a.score.SetText(a.game.GetScore().String())
}
