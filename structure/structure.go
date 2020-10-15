package structure

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/creep"
	"github.com/bill-rich/go-screeps/object"
	"github.com/bill-rich/go-screeps/room"
)

// Structure is the base prototype object of all structures.
type Structure struct {
	object.Object
	Room          room.Room `js:"room"`
	Hits          int       `js:"hits"`
	HitsMax       int       `js:"hitsMax"`
	StructureType string    `js:"structureType"`
}

// OwnedStructure is the base prototype for a structure that has an owner.
type OwnedStructure struct {
	Structure
	My    bool         `js:"my"`
	Owner common.Owner `js:"owner"`
}

// Container is a small container that can be used to store resources.
type Container struct {
	Structure
	Store common.Store `js:"store"`
}

// Controller is the controller structure.
type Controller struct {
	Structure
}

// Extension contains energy which can be spent on spawning bigger creeps.
type Extension struct {
	OwnedStructure
	Store common.Store `js:"store"`
}

// Extractor allows to harvest a mineral deposit.
type Extractor struct {
	OwnedStructure
	Cooldown int `js:"cooldown"`
}

// Factory produces trade commodities from base minerals and other commodities.
type Factory struct {
	OwnedStructure
	Store    common.Store `js:"store"`
	Cooldown int          `js:"cooldown"`
	Level    int          `js:"level"`
}

// InvaderCore is a control center of NPC Strongholds, and also rules all invaders in the sector.
type InvaderCore struct {
	OwnedStructure
	TicksToDeploy int `js:"ticksToDeploy"`
	Level         int `js:"level"`
}

// KeeperLair spawns NPC Source Keepers that guards energy sources and minerals in some rooms.
type KeeperLair struct {
	OwnedStructure
	TicksToSpawn int `js:"ticksToSpawn"`
	Level        int `js:"level"`
}

// Lab produces mineral compounds from base minerals, boosts and unboosts creeps.
type Lab struct {
	OwnedStructure
	Store    common.Store `js:"store"`
	Cooldown int          `js:"cooldown"`
}

// Link remotely transfers energy to another Link in the same room.
type Link struct {
	OwnedStructure
	Store    common.Store `js:"store"`
	Cooldown int          `js:"cooldown"`
}

// Nuker launches a nuke to another room dealing huge damage to the landing area.
type Nuker struct {
	OwnedStructure
	Store    common.Store `js:"store"`
	Cooldown int          `js:"cooldown"`
}

// Spawn is the spawn structure.
type Spawn struct {
	Structure
}

// Storage is a structure that can store huge amount of resource units.
type Storage struct {
	OwnedStructure
	Store common.Store `js:"store"`
}

// Terminal sends any resources to a Terminal in another room.
type Terminal struct {
	OwnedStructure
	Store    common.Store `js:"store"`
	Cooldown int          `js:"cooldown"`
}

// Send sends resource to a Terminal in another room with the specified name.
func (t Terminal) Send(resourceType string, amount int, destination string, description string) error {
	return common.ErrT(t.Call("send", resourceType, amount, destination, description).Int())
}

// TransferEnergy remotely transfer energy to another link at any location in the same room.
func (l Link) TransferEnergy(target Link, amount int) error {
	return common.ErrT(l.Call("transferEnergy", target, amount).Int())
}

// LaunchNuke launches a nuke to the specified position.
func (l Link) LaunchNuke(pos common.RoomPosition) error {
	return common.ErrT(l.Call("launchNuke", pos).Int())
}

// BoostCreep boosts creep body parts using the containing mineral compound.
func (l Lab) BoostCreep(creep creep.Creep, bodyPartCount int) error {
	return common.ErrT(l.Call("boostCreep", creep, bodyPartCount).Int())
}

// ReverseReaction breaks mineral compounds back into reagents.
func (l Lab) ReverseReaction(lab1, lab2 Lab) error {
	return common.ErrT(l.Call("reverseReaction", lab1, lab2).Int())
}

// RunReaction produces mineral compounds using reagents from two other labs.
func (l Lab) RunReaction(lab1, lab2 Lab) error {
	return common.ErrT(l.Call("runReaction", lab1, lab2).Int())
}

// UnboostCreep immediately remove boosts from the creep and drop 50% of the mineral compounds used to boost it onto the ground regardless of the creep's remaining time to live.
func (l Lab) UnboostCreep(creep creep.Creep) error {
	return common.ErrT(l.Call("unboostCreep", creep).Int())
}

// Produce produces the specified commodity.
func (f Factory) Produce(resourceType string) error {
	return common.ErrT(f.Call("produce", resourceType).Int())
}

// Destroy will destroy this structure immediately.
func (s Structure) Destroy() error {
	return common.ErrT(s.Call("destroy").Int())
}

// IsActive checks whether this structure can be used.
func (s Structure) IsActive() bool {
	return s.Call("isActive").Bool()
}

// NotifyWhenAttacked will toggle auto notification when the structure is under attack.
func (s Structure) NotifyWhenAttacked(enabled common.Enabled) bool {
	return s.Call("notifyWhenAttacked", enabled).Bool()
}

// SpawnCreep spawns a creep.
func (s Spawn) SpawnCreep(body []string, name string) error {
	return common.ErrT(s.Call("spawnCreep", body, name).Int())
}
