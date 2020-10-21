package creep

import (
	"errors"

	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/constructionsite"
	"github.com/bill-rich/go-screeps/object"
	"github.com/bill-rich/go-screeps/resource"
	"github.com/bill-rich/go-screeps/room"
	"github.com/bill-rich/go-screeps/roomposition"
	"github.com/bill-rich/go-screeps/source"
)

// Creep is an object representing one of your units.
type Creep struct {
	object.Object
	Store       common.Store           `js:"store"`
	Room        room.Room              `js:"room"`
	Body        []string               `js:"body"`
	Fatigue     int                    `js:"fatigue"`
	Hits        int                    `js:"hits"`
	HitsMax     int                    `js:"hitsMax"`
	Memory      map[string]interface{} `js:"memory"`
	My          bool                   `js:"my"`
	Name        string                 `js:"name"`
	Owner       common.Owner           `js:"owner"`
	Saying      string                 `js:"saying"`
	Spawning    bool                   `js:"spawning"`
	TicksToLive int                    `js:"ticksToLive"`
}

// StructureInterface is a bit of a hack to get around an import cycle.
type StructureInterface interface {
	Destroy() error
	IsActive() bool
	NotifyWhenAttacked(common.Enabled) bool
}

// Attack another creep, power creep, or structure in a short-ranged attack.
func (c Creep) Attack(target common.AttackTarget) error {
	return common.ErrT(c.Call("attack", target).Int())
}

// AttackController decreases the controller's downgrade timer by 300 ticks per
// every CLAIM body part.
func (c Creep) AttackController(target StructureInterface) error {
	return common.ErrT(c.Call("attackController", target).Int())
}

// Build a structure at the target construction site using carried energy.
func (c Creep) Build(target constructionsite.ConstructionSite) error {
	return common.ErrT(c.Call("build", target).Int())
}

// CancelOrder cancels the order given during the current game tick.
func (c Creep) CancelOrder(name string) error {
	return common.ErrT(c.Call("removeOrder", name).Int())
}

// ClaimController claims a neutral controller under your control.
func (c Creep) ClaimController(target StructureInterface) error {
	return common.ErrT(c.Call("claimController", target).Int())
}

// Dismantle dismantles any structure that can be constructed.
func (c Creep) Dismantle(target StructureInterface) error {
	return common.ErrT(c.Call("dismantle", target).Int())
}

// Drop this resource on the ground.
func (c Creep) Drop(resourceType string, amount int) error {
	return common.ErrT(c.Call("drop", resourceType, amount).Int())
}

// AddSafeMode adds one more available safe mode activation to a room
// controller.
func (c Creep) AddSafeMode(target StructureInterface) error {
	return common.ErrT(c.Call("addSafeMode", target).Int())
}

// GetActiveBodyParts gets the quantity of live body parts of the given type.
func (c Creep) GetActiveBodyParts(bodyType string) int {
	return c.Call("getActiveBodyparts", bodyType).Int()
}

// Harvest energy from the source or resources from minerals and deposits.
func (c Creep) Harvest(target source.Source) error {
	return common.ErrT(c.Call("harvest", target).Int())
}

// Heal self or another creep.
func (c Creep) Heal(target Creep) error {
	return common.ErrT(c.Call("heal", target).Int())
}

// Move moves the creep one square in the specified direction.
func (c Creep) Move(direction int) error {
	if direction < 1 || direction > 8 {
		return errors.New("ERR_INVALID_ARGS")
	}
	return common.ErrT(c.Call("move", direction).Int())
}

// MoveByPath moves the creep using the specified predefined path.
func (c Creep) MoveByPath(path []roomposition.RoomPosition) error {
	return common.ErrT(c.Call("moveByPath", path).Int())
}

// MoveTo finds the optimal path to the target within the same room and move
// to it.
func (c Creep) MoveTo(target roomposition.RoomPosition) error {
	return common.ErrT(c.Call("moveTo", target).Int())
}

// NotifyWhenAttacked toggle auto notification when the creep is under attack.
func (c Creep) NotifyWhenAttacked(enabled bool) common.Enabled {
	return common.Enabled{
		Enabled: c.Call("notifyWhenAttacked", enabled).Bool(),
	}
}

// Pickup an item (a dropped piece of energy).
func (c Creep) Pickup(target resource.Resource) error {
	return common.ErrT(c.Call("pickup", target).Int())
}

// Pull another creep.
func (c Creep) Pull(target Creep) error {
	return common.ErrT(c.Call("pull", target).Int())
}

// RangedAttack attacks against another creep or structure.
func (c Creep) RangedAttack(target common.AttackTarget) error {
	return common.ErrT(c.Call("rangedAttack", target).Int())
}

// RangedHeal heals another creep at a distance.
func (c Creep) RangedHeal(target Creep) error {
	return common.ErrT(c.Call("rangedHeal", target).Int())
}

// RangedMassAttack attacks against all hostile creeps or structures within 3
// squares range.
func (c Creep) RangedMassAttack() error {
	return common.ErrT(c.Call("rangedMassAttack").Int())
}

// Repair a damaged structure using carried energy.
func (c Creep) Repair(target Creep) error {
	return common.ErrT(c.Call("repair", target).Int())
}

// ReserveController temporarily block a neutral controller from claiming by
// other players.
func (c Creep) ReserveController(target StructureInterface) error {
	return common.ErrT(c.Call("reserveController", target).Int())
}

// Say display a visual speech balloon above the creep with the specified
// message.
func (c Creep) Say(comment string) error {
	return common.ErrT(c.Call("say", comment).Int())
}

// SignController signs a controller with an arbitrary text visible to all
// players.
func (c Creep) SignController(target StructureInterface) error {
	return common.ErrT(c.Call("signController", target).Int())
}

// Suicide kills the creep immediately.
func (c Creep) Suicide() error {
	return common.ErrT(c.Call("suicide").Int())
}

// Transfer resource from the creep to another object.
func (c Creep) Transfer(target common.Store, resourceType string, amount int) error {
	return common.ErrT(c.Call("pickup", target, resourceType, amount).Int())
}

// UpgradeController upgrades your controller to the next level using carried
// energy.
func (c Creep) UpgradeController(target StructureInterface) error {
	return common.ErrT(c.Call("upgradeController", target).Int())
}

// Withdraw resources from a structure or tombstone.
func (c Creep) Withdraw(target common.Store, resourceType string, amount int) error {
	return common.ErrT(c.Call("withdraw", target, resourceType, amount).Int())
}
