package console

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// Console provides access to the in game console.
type Console struct {
	*js.Object
}

// Log logs a message to the console.
func (c Console) Log(messages ...interface{}) {
	c.Call("log", fmt.Sprintf("%+v\n", messages...))
}

// GetConsole returns the game's console.
func GetConsole() Console {
	return Console{Object: js.Global.Get("console")}
}
