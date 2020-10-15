package object

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/gopherjs/gopherjs/js"
)

// Object is anything that has a 'pos'.
type Object struct {
	*js.Object
	Pos     common.RoomPosition `js:"pos"`
	Effects []common.Effect     `js:"effects"`
	ID      string              `js:"id"`
}
