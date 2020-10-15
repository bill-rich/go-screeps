package cpu

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/gopherjs/gopherjs/js"
)

// CPU is an object containing information about your CPU usage.
type CPU struct {
	*js.Object
	Limit        int     `js:"limit"`
	TickLimit    int     `js:"tickLimit"`
	Bucket       int     `js:"bucket"`
	SharedLimits []Limit `js:"limit"`
	Unlocked     bool    `js:"unlocked"`
	UnlockedTime int     `js:"unlockedTime"`
}

// Limit contains CPU limits on each shard.
type Limit map[string]int

// GetHeapStatistics is a method to get heap statistics for your virtual machine.
func (g CPU) GetHeapStatistics() map[string]int {
	return g.Call("getHeapStatistics").Interface().(map[string]int)
}

// GetUsed gets amount of CPU time used from the beginning of the current game tick.
func (g CPU) GetUsed() int {
	return g.Call("getUsed").Int()
}

// Halt resets your runtime environment and wipe all data in heap memory.
func (g CPU) Halt() {
	g.Call("halt")
}

// SetSharedLimits allocates CPU limits to different shards.
func (g CPU) SetSharedLimits(limits []Limit) error {
	return common.ErrT(g.Call("setSharedLimits", limits).Int())
}

// Unlock unlocks full CPU for your account for additional 24 hours.
func (g CPU) Unlock() error {
	return common.ErrT(g.Call("unlock").Int())
}

// GeneratePixel generates 1 pixel resource unit for 5000 CPU from your bucket.
func (g CPU) GeneratePixel() error {
	return common.ErrT(g.Call("generatePixel").Int())
}
