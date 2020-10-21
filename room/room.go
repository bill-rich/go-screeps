package room

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/roomposition"
	"github.com/gopherjs/gopherjs/js"
)

// Room is an object representing the room in which your units and structures are in.
type Room struct {
	*js.Object
	Controller      StructureInterface `js:"structureController"`
	EnergyAvailable int                `js:"energyAvailable"`
	Memory          interface{}        `js:"memory"`
	Name            string             `js:"name"`
	Storage         StructureInterface `js:"storage"`
	Terminal        StructureInterface `js:"terminal"`
	Visual          Visual             `js:"visual"`
}

// Visual provides a way to show various visual debug info in game rooms.
type Visual struct {
	*js.Object
	RoomName string `js:"roomName"`
}

// Event represents some game action.
type Event struct {
	*js.Object
	Event    string      `js:"event"`
	ObjectID string      `js:"objectId"`
	Data     interface{} `js:"data"`
}

// Terrain is an object which provides fast access to room terrain data.
type Terrain struct {
	*js.Object
}

// StructureInterface is a bit of a hack to get around an import cycle.
type StructureInterface interface {
	Destroy() error
	IsActive() bool
	NotifyWhenAttacked(common.Enabled) bool
}

// Get terrain type at the specified room position by (x,y) coordinates.
func (t Terrain) Get(x, y int) int {
	return t.Call("get", x, y).Int()
}

// SerializePath serializes a path array into a short string representation,
// which is suitable to store in memory.
func (r Room) SerializePath(path []roomposition.RoomPosition) string {
	return r.Call("serializePath", path).String()
}

// DeserializePath deserializes a short string path representation into an array
// form.
func (r Room) DeserializePath(path string) []roomposition.RoomPosition {
	return r.Call("deserializePath", path).Interface().([]roomposition.RoomPosition)
}

// CreateContstructionSite creates new ConstructionSite at the specified
// location.
func (r Room) CreateContstructionSite(pos roomposition.RoomPosition, structureType string, name string) error {
	return common.ErrT(r.Call("createConstructionSite", pos, structureType, name).Int())
}

// CreateFlag creates new Flag at the specified location.
func (r Room) CreateFlag(pos roomposition.RoomPosition, name string, color string, secondary string) error {
	return common.ErrT(r.Call("createFlag", pos, name, color, secondary).Int())
}

// Find all objects of the specified type in the room.
func (r Room) Find(args ...interface{}) *js.Object {
	return r.Call("find", args...)
}

// FindExitTo finds the exit direction en route to another room.
func (r Room) FindExitTo(room Room) int {
	return r.Call("findExitTo", room).Int()
}

// FindPath finds an optimal path inside the room between fromPos and toPos
// using Jump Point Search algorithm.
func (r Room) FindPath(fromPos, toPos roomposition.RoomPosition, opts string) []common.PathOp {
	return r.Call("findPath", fromPos, toPos, opts).Interface().([]common.PathOp)
}

// GetEventLog returns an array of events happened on the previous tick in this
// room.
func (r Room) GetEventLog(raw bool) []Event {
	return r.Call("getEventLog", raw).Interface().([]Event)
}

// GetPositionAt creates a RoomPosition object at the specified location.
func (r Room) GetPositionAt(x, y int) roomposition.RoomPosition {
	return roomposition.RoomPosition{Object: r.Call("getPositionAt", x, y)}
}

// GetTerrain gets a Terrain object which provides fast access to static terrain
// data.
func (r Room) GetTerrain() Terrain {
	return r.Call("getTerrain").Interface().(Terrain)
}
