package mineral

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/room"
)

// Mineral is a mineral deposit.
type Mineral struct {
	common.Object
	Effects       []common.Effect `js:"effects"`
	Room          room.Room       `js:"room"`
	Density       int             `js:"density"`
	MineralAmount int             `js:"mineralAmount"`
	MineralType   string          `js:"mineralType"`
	ID            string          `js:"id"`
	TicksToDecay  int             `js:"ticksToDecay"`
}
