package main

import (
	"encoding/json"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	game := NewGame()

	game.State.Dice.Roll()

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(game)

	// if err := ebiten.RunGame(game); err != nil {
	// 	log.Fatal(err)
	// }
}
