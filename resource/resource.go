package resource

import (
	"github.com/bill-rich/screepsgo-stub/common"
	"github.com/bill-rich/screepsgo-stub/room"
)

// Resource is a dropped piece of resource.
type Resource struct {
	common.Object
	Effects      []common.Effect `js:"effects"`
	Room         room.Room       `js:"room"`
	Amount       int             `js:"amount"`
	ID           string          `js:"id"`
	ResourceType string          `js:"resourceType"`
}
