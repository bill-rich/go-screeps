package deposit

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/room"
)

// Deposit is a rare resource deposit needed for producing commodities.
type Deposit struct {
	common.Object
	Effects      []common.Effect `js:"effects"`
	Room         room.Room       `js:"room"`
	Cooldown     int             `js:"cooldown"`
	DepositType  string          `js:"depositType"`
	ID           string          `js:"id"`
	LastCooldown int             `js:"lastCooldown"`
	TicksToDecay int             `js:"ticksToDecay"`
}
