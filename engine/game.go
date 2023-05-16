package engine

import (
	"github.com/shopspring/decimal"
)

type Game struct {
	running         bool
	score           decimal.Decimal
	clicks          int
	clickValue      decimal.Decimal
	clicksPerSecond decimal.Decimal
}

func NewGame() *Game {
	return &Game{
		running:         true,
		score:           decimal.NewFromInt(0),
		clicks:          0,
		clickValue:      decimal.NewFromInt(1),
		clicksPerSecond: decimal.NewFromInt(0),
	}
}

func (g *Game) IsRunning() bool {
	return g.running
}

func (g *Game) GetScore() decimal.Decimal {
	return g.score
}

func (g *Game) SetScore(score decimal.Decimal) {
	g.score = score
}

func (g *Game) AddScore(score decimal.Decimal) {
	g.score = g.score.Add(score)
}

func (g *Game) PrintScore() {
	println(g.score.String())
}

func (g *Game) Quit() {
	g.running = false
}

func (g *Game) Click() {
	g.AddScore(g.clickValue)
	g.clicks++
}

func (g *Game) SetClicksPerSecond(clicksPerSecond decimal.Decimal) {
	g.clicksPerSecond = clicksPerSecond
}

func (g *Game) GetClicksPerSecond() decimal.Decimal {
	return g.clicksPerSecond
}

func (g *Game) ProcessInput(input string) {
	switch input {
	case "quit":
		g.Quit()
	case "score":
		g.PrintScore()
	case "click":
		g.Click()
		g.PrintScore()
	case "buy":
		if g.score.GreaterThanOrEqual(decimal.NewFromInt(5)) {
			g.SetClicksPerSecond(g.GetClicksPerSecond().Add(decimal.NewFromFloat(0.1)))
			g.AddScore(decimal.NewFromInt(-5))
		}
	case "cps":
		println(g.GetClicksPerSecond().String())
	default:
		println("Unknown command")
	}
}

func (g *Game) Tick() {
	print("tick: ")
	g.AddScore(g.GetClicksPerSecond())
	g.PrintScore()
}
