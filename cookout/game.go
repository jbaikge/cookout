package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	PersonalityExtrovert = 0
	PersonalityIntrovert = 1
)

type Battery struct {
	Charge int `json:"charge"`
}

type Player struct {
	Name        string  `json:"name"`
	Avatar      int     `json:"avatar"`
	Personality int     `json:"personality"`
	Battery     Battery `json:"battery"`
	Space       int     `json:"space"`
}

type Space struct {
	Avatar   int  `json:"avatar"`
	LoseTurn bool `json:"loseTurn"`
}

type Board struct {
	Spaces []Space `json:"spaces"`
}

type Condition struct {
	Values  []int  `json:"values"`
	Label   string `json:"label"`
	Outcome string `json:"outcomes"`
	Charges []int  `json:"charges"`
}

type Card struct {
	Title      string      `json:"title"`
	Body       string      `json:"body"`
	Function   string      `json:"function,omitempty"`
	Conditions []Condition `json:"conditions"`
}

type Deck struct {
	Cards []Card `json:"cards"`
}

type Dice struct {
	Value int `json:"value"`
}

type State struct {
	Players []Player `json:"players"`
	Dice    Dice     `json:"dice"`
}

type Game struct {
	Board Board
	Deck  Deck
	State State
}

func (d *Dice) Roll() {
	d.Value = rand.Intn(5) + 1
}

func NewGame() (g *Game) {
	return &Game{
		Board: Board{
			Spaces: make([]Space, 0, 36),
		},
		Deck: Deck{
			Cards: make([]Card, 0, 50),
		},
		State: State{
			Players: make([]Player, 0, 4),
		},
	}
}

// Update is called every tick
func (g *Game) Update() (err error) {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
