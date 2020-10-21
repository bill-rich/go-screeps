package roomposition

import "github.com/gopherjs/gopherjs/js"

// RoomPosition is an object representing the specified position in the room.
type RoomPosition struct {
	*js.Object
	X        int    `js:"x"`
	Y        int    `js:"y"`
	RoomName string `js:"roomName"`
}
