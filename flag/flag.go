package flag

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/room"
	"github.com/bill-rich/go-screeps/roomposition"
)

// Flag can be used to mark particular spots in a room.
type Flag struct {
	common.Object
	Effects        []common.Effect `js:"effects"`
	Room           room.Room       `js:"room"`
	Color          int             `js:"color"`
	Memory         interface{}     `js:"memory"`
	Name           string          `js:"name"`
	SecondaryColor int             `js:"secondaryColor"`
}

// Remove the flag.
func (f Flag) Remove() error {
	return common.ErrT(f.Call("remove").Int())
}

// SetColor sets a new color of the flag.
func (f Flag) SetColor(color, secondary int) error {
	if secondary == 0 {
		secondary = color
	}
	return common.ErrT(f.Call("setColor", color, secondary).Int())
}

// SetPosition sets a new position of the flag.
func (f Flag) SetPosition(pos roomposition.RoomPosition) error {
	return common.ErrT(f.Call("setPosition", pos).Int())
}
