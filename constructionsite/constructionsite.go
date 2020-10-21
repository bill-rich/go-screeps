package constructionsite

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/room"
	"github.com/bill-rich/go-screeps/roomposition"
	"github.com/gopherjs/gopherjs/js"
)

// ConstructionSite is a site of a structure which is currently under
// construction.
type ConstructionSite struct {
	*js.Object
	Effects       []common.Effect           `js:"effects"`
	Pos           roomposition.RoomPosition `js:"pos"`
	Room          room.Room                 `js:"room"`
	ID            string                    `js:"name"`
	My            bool                      `js:"my"`
	Owner         common.Owner              `js:"owner"`
	Progress      int                       `js:"progress"`
	ProgressTotal int                       `js:"progressTotal"`
	StructureType string                    `js:"structureType"`
}

// Remove the construction site.
func (c ConstructionSite) Remove() error {
	return common.ErrT(c.Call("remove").Int())
}
