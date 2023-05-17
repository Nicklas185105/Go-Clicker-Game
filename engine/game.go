package engine

import (
	"github.com/shopspring/decimal"
)

type Game struct {
	score           decimal.Decimal
	clicks          int
	clickValue      decimal.Decimal
	clicksPerSecond decimal.Decimal
}

func NewGame() *Game {
	return &Game{
		score:           decimal.NewFromInt(0),
		clicks:          0,
		clickValue:      decimal.NewFromInt(1),
		clicksPerSecond: decimal.NewFromInt(0),
	}
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

func (g *Game) Tick() {
	print("tick: ")
	g.AddScore(g.GetClicksPerSecond())
	g.PrintScore()
}
