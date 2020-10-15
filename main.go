package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("run", run)
}

func run() {
	return
}
