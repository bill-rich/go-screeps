package resource

import (
	"github.com/bill-rich/go-screeps/object"
	"github.com/bill-rich/go-screeps/room"
)

// Resource is a dropped piece of resource.
type Resource struct {
	object.Object
	Room         room.Room `js:"room"`
	Amount       int       `js:"amount"`
	ResourceType string    `js:"resourceType"`
}
