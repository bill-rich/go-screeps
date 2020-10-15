package nuke

import (
	"github.com/bill-rich/go-screeps/object"
)

// Nuke is a nuke landing position.
type Nuke struct {
	object.Object
	LaunchRoomName string `js:"launchRoomName"`
	TimeToLand     int    `js:"timeToLand"`
}
