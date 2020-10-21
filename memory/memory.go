package memory

import "github.com/gopherjs/gopherjs/js"

// Memory is a global plain object which can contain arbitrary data.
type Memory map[string]interface{}

// GetMemory returns the game's memory.
func GetMemory() Memory {
	memory := Memory{}
	memory = js.Global.Get("memory").Interface().(map[string]interface{})
	return memory
}

type RawMemory struct {
	*js.Object
}

// GetRawMemory returns the game's memory.
func GetRawMemory() RawMemory {
	memory := RawMemory{js.Global.Get("rawMemory")}
	return memory
}
