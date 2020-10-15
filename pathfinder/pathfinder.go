package pathfinder

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/gopherjs/gopherjs/js"
)

// PathFinder contains powerful methods for pathfinding in the game world.
type PathFinder struct {
	*js.Object
}

// Path is an object that contains path result information.
type Path struct {
	*js.Object
	Path       []common.RoomPosition `js:"path"`
	Ops        int                   `js:"ops"`
	Cost       int                   `js:"cost"`
	Incomplete bool                  `js:"incomplete"`
}

// Search find an optimal path between origin and goal.
func (p PathFinder) Search(origin common.Object, goal common.Object, opts string) Path {
	return p.Call("search", origin, goal, opts).Interface().(Path)
}
