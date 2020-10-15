package common

// For more complete and up to date documentation,
// see: https://docs.screeps.com/api.

import (
	"fmt"

	"github.com/bill-rich/go-screeps/constants"
	"github.com/gopherjs/gopherjs/js"
)

// AttackTarget is and object that may be attacked.
type AttackTarget interface {
	NotifyWhenAttacked(bool) Enabled
}

// Enabled represents if an option is enabled.
type Enabled struct {
	*js.Object
	Enabled bool `js:"enabled"`
}

// Effect represents an effect on an object.
type Effect struct {
	*js.Object
	Effect         int `js:"effect"`
	Level          int `js:"level"`
	TicksRemaining int `js:"ticksRemaining"`
}

// RoomPosition is an object representing the specified position in the room.
type RoomPosition struct {
	*js.Object
	X        int    `js:"x"`
	Y        int    `js:"y"`
	RoomName string `js:"roomName"`
}

// Owner is an object with the creep’s owner info.
type Owner struct {
	*js.Object
	Username string `js:"username"`
}

// Storage is an attribute of objects that contain storage.
type Storage struct {
	Store Store `js:"store"`
}

// Store is an object that can contain resources in its cargo.
type Store struct {
	*js.Object
}

// PathOp is a step in a path.
type PathOp struct {
	X         int `js:"x"`
	Y         int `js:"y"`
	Dx        int `js:"dx"`
	Dy        int `js:"dy"`
	Direction int `js:"direction"`
}

// GetCapacity returns capacity of this store for the specified resource. For a
// general purpose store, it returns total capacity if resource is undefined.
func (s Store) GetCapacity(resource string) error {
	return ErrT(s.Call("getCapacity", resource).Int())
}

// GetUsedCapacity returns the capacity used by the specified resource. For a
// general purpose store, it returns total used capacity if resource is
// undefined.
func (s Store) GetUsedCapacity(resource string) error {
	return ErrT(s.Call("getUsedCapacity", resource).Int())
}

// GetFreeCapacity returns free capacity for the store. For a limited store, it
// returns the capacity available for the specified resource if resource is
// defined and valid for this store.
func (s Store) GetFreeCapacity(resource string) error {
	return ErrT(s.Call("getFreeCapacity", resource).Int())
}

// ErrT translates Screeps error codes into useful strings.
func ErrT(errCode int) error {
	if errCode == 0 {
		return nil
	}
	for key, value := range constants.ErrorCodes {
		if value == errCode {
			return fmt.Errorf(key)
		}
	}
	return fmt.Errorf("UNKNOWN_ERROR: %d", errCode)
}
