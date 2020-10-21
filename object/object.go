package object

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/roomposition"
	"github.com/gopherjs/gopherjs/js"
)

// Object is anything that has a 'pos'.
type Object struct {
	*js.Object
	Pos     roomposition.RoomPosition `js:"pos"`
	Effects []common.Effect           `js:"effects"`
	ID      string                    `js:"id"`
}
