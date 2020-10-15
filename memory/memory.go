package memory

import "github.com/gopherjs/gopherjs/js"

// Memory is a global plain object which can contain arbitrary data.
type Memory struct {
	*js.Object
}
