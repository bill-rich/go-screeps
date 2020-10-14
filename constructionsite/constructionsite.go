package constructionsite

import (
	"github.com/bill-rich/screepsgo-stub/common"
	"github.com/bill-rich/screepsgo-stub/room"
	"github.com/gopherjs/gopherjs/js"
)

// ConstructionSite is a site of a structure which is currently under
// construction.
type ConstructionSite struct {
	*js.Object
	Effects       []common.Effect     `js:"effects"`
	Pos           common.RoomPosition `js:"pos"`
	Room          room.Room           `js:"room"`
	ID            string              `js:"name"`
	My            bool                `js:"my"`
	Owner         common.Owner        `js:"owner"`
	Progress      int                 `js:"progress"`
	ProgressTotal int                 `js:"progressTotal"`
	StructureType string              `js:"structureType"`
}

// Remove the construction site.
func (c ConstructionSite) Remove() error {
	return common.ErrT(c.Call("remove").Int())
}
