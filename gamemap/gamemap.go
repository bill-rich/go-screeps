package gamemap

import (
	"github.com/bill-rich/go-screeps/room"
	"github.com/gopherjs/gopherjs/js"
)

// Map is a global object representing world map.
type Map struct {
	*js.Object
}

// Exits describe possible exits from a room.
type Exits struct {
	*js.Object
	Top    string `js:"1"`
	Right  string `js:"3"`
	Bottom string `js:"5"`
	Left   string `js:"7"`
}

// ExitOrder provides a step in moving from one room to another.
type ExitOrder struct {
	*js.Object
	Exit int    `js:"exit"`
	Room string `js:"room"`
}

// Status describes the status of an object.
type Status struct {
	*js.Object
	Status    string `js:"status"`
	Timestamp int    `js:"timestamp"`
}

// DescribeExits finds the exits to a room.
func (m Map) DescribeExits(room *room.Room) Exits {
	return m.Call("describeExits", room.Name).Interface().(Exits)
}

// FindExit finds the exit from one room to another.
func (m Map) FindExit(fromRoom, toRoom *room.Room, opts string) int {
	return m.Call("findExit", fromRoom.Name, toRoom.Name, opts).Int()
}

// FindRoute finds all the exits and rooms from one room to another.
func (m Map) FindRoute(fromRoom, toRoom *room.Room, opts string) []ExitOrder {
	orders := []ExitOrder{}
	rawOrders := m.Call("findRoute", fromRoom.Name, toRoom.Name, opts).Interface().([]interface{})
	for _, order := range rawOrders {
		newOrder := order.(*js.Object).Interface().(ExitOrder)
		orders = append(orders, newOrder)
	}
	return orders
}

// GetRoomLinearDistance gets the linear distance (in rooms) between two rooms.
func (m Map) GetRoomLinearDistance(fromRoom, toRoom *room.Room, continuous bool) int {
	return m.Call("getRoomLinearDistance", fromRoom.Name, toRoom.Name, continuous).Int()
}

// GetWorldSize returns the world size as a number of rooms.
func (m Map) GetWorldSize() int {
	return m.Call("getWorldSize").Int()
}

// GetRoomStatus gets availablity status of the room.
func (m Map) GetRoomStatus(room *room.Room) Status {
	return m.Call("getRoomStatus", room.Name).Interface().(Status)
}
