package structure

import (
	"github.com/bill-rich/go-screeps/common"
)

type Structure struct {
	common.Object
	Hits int `js:"hits"`
}

type Controller struct {
	Structure
}

type Spawn struct {
	Structure
}

func (s Structure) Destroy() error {
	return common.ErrT(s.Call("destroy").Int())
}
