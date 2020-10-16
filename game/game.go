package game

import (
	"flag"

	"github.com/bill-rich/go-screeps/constructionsite"
	"github.com/bill-rich/go-screeps/cpu"
	"github.com/bill-rich/go-screeps/creep"
	"github.com/bill-rich/go-screeps/gamemap"
	"github.com/bill-rich/go-screeps/resource"
	"github.com/bill-rich/go-screeps/room"
	"github.com/bill-rich/go-screeps/structure"
	"github.com/gopherjs/gopherjs/js"
)

// Game object.
type Game struct {
	*js.Object
	CPU               cpu.CPU                                      `js:"cpu"`
	ConstructionSites map[string]constructionsite.ConstructionSite `js:"constructionSites"`
	Rooms             map[string]room.Room                         `js:"rooms"`
	Creeps            map[string]creep.Creep                       `js:"creeps"`
	Flags             map[string]flag.Flag                         `js:"flags"`
	GCL               Level                                        `js:"gcl"`
	GPL               Level                                        `js:"gpl"`
	Map               gamemap.Map                                  `js:"map"`
	PowerCreeps       map[string]creep.Creep                       `js:"powerCreeps"`
	Resources         map[string]resource.Resource                 `js:"resources"`
	Shard             Shard                                        `js:"shard"`
	Spawns            map[string]structure.Spawn                   `js:"spawns"`
	Structures        map[string]structure.Structure               `js:"structures"`
	Time              int                                          `js:"time"`
}

// Shard is an object describing the world shard where your script is currently being executed in.
type Shard struct {
	Name string `js:"name"`
	Type string `js:"type"`
	Ptr  bool   `js:"ptr"`
}

// Level provides Global Control Level information.
type Level struct {
	*js.Object
	Level         int `js:"level"`
	Progress      int `js:"progress"`
	ProgressTotal int `js:"progressTotal"`
}

// GetGame returns the game object.
func GetGame() Game {
	return Game{Object: js.Global.Get("game")}
}

// GetObjectByID gets an object with the specified unique ID.
//func (g Game) GetObjectByID(id string) interface{} {
func (g Game) GetObjectByID(id string) *js.Object {
	return g.Call("getObjectById", id)
}

// Notify sends a custom message at your profile email.
func (g Game) Notify(message string, groupInterval int) {
	g.Call("notify", message, groupInterval)
}
