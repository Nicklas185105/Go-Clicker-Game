package frontend

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
