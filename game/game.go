package game

import (
	"github.com/bill-rich/screepsgo-stub/creep"
	"github.com/bill-rich/screepsgo-stub/room"
	"github.com/gopherjs/gopherjs/js"
)

// Game object.
type Game struct {
	*js.Object
	Rooms  map[string]room.Room   `js:"rooms"`
	Creeps map[string]creep.Creep `js:"creeps"`
}
