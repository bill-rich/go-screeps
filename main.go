package main

import (
	"github.com/bill-rich/screepsgo-stub/game"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("addTwo", addTwo)
}

func addTwo() interface{} {
	game := game.Game{Object: js.Global.Get("game")}
	creep := game.Creeps["creep1"]
	enemy := game.Creeps["creep2"]
	return creep.Attack(enemy)
}
