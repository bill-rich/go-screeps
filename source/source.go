package source

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/object"
	"github.com/bill-rich/go-screeps/room"
)

// Source is an energy source object.
type Source struct {
	object.Object
	Effect              []common.Effect `js:"effects"`
	Room                room.Room       `js:"room"`
	Energy              int             `js:"energy"`
	EnergyCapacity      int             `js:"energyCapacity"`
	ID                  string          `js:"id"`
	TicksToRegeneration int             `js:"ticksToRegeneration"`
}
