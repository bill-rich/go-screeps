package pathfinder

import (
	"github.com/bill-rich/go-screeps/object"
	"github.com/bill-rich/go-screeps/roomposition"
	"github.com/gopherjs/gopherjs/js"
)

// PathFinder contains powerful methods for pathfinding in the game world.
type PathFinder struct {
	*js.Object
}

// Path is an object that contains path result information.
type Path struct {
	*js.Object
	Path       []roomposition.RoomPosition `js:"path"`
	Ops        int                         `js:"ops"`
	Cost       int                         `js:"cost"`
	Incomplete bool                        `js:"incomplete"`
}

// Search find an optimal path between origin and goal.
func (p PathFinder) Search(origin object.Object, goal object.Object, opts string) Path {
	return p.Call("search", origin, goal, opts).Interface().(Path)
}
