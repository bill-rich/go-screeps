package room

import (
	"github.com/bill-rich/screepsgo-stub/common"
	"github.com/gopherjs/gopherjs/js"
)

type Room struct {
	*js.Object
	Name string `js:"name"`
}

func (r Room) find(args ...interface{}) *js.Object {
	return r.Call("find", args...)
}

type Source struct {
	common.Object
	Effect              []common.Effect `js:"effects"`
	Room                Room            `js:"room"`
	Energy              int             `js:"energy"`
	EnergyCapacity      int             `js:"energyCapacity"`
	ID                  string          `js:"id"`
	TicksToRegeneration int             `js:"ticksToRegeneration"`
}
