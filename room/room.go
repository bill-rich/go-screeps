package room

import (
	"github.com/gopherjs/gopherjs/js"
)

// Room is an object representing the room in which your units and structures are in.
type Room struct {
	*js.Object
	Name string `js:"name"`
}

func (r Room) Find(args ...interface{}) *js.Object {
	return r.Call("find", args...)
}
