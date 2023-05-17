package frontend

import "github.com/dtylman/gowd"

const homeBody = `
<div id="homeBody">
	<div class="row">
		<div class="col-md-12">
			<div class="jumbotron">
				<h1>Clicker Game</h1>
			</div>
		</div>
	</div>
	<div class="row">
		<div class="col-md-4 col-md-offset-4">
			<button class="btn btn-primary btn-block" id="start">Start</button>
		</div>
	</div>
</div>
`

func (a *app) start(sender *gowd.Element, event *gowd.EventElement) {
	a.homeBody.Hide()
	a.gameBody.Show()
}
