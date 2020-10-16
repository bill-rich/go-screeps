This will set up a `run` function and get Game, Memory, and Pathfinder.

```
func main() {
	js.Global.Set("run", run)
}

func run() {
	console := console.Console{Object: js.Global.Get("consol")}
	game := game.Game{Object: js.Global.Get("game")}
	mem := memory.Memory{}
	mem = js.Global.Get("memory").Interface().(map[string]interface{})
  pathfinder = Pathfinder{Object: js.Global.Get("pathfinder")}
	return
}
```

You'll also need to create a JS file with the following:
```
require('your_converted_js')

Object.defineProperty(StructureController.prototype, 'room', {
 get: function() {
     return
 },
 enumerable: false,
 configurable: true
})

Object.defineProperty(StructureTerminal.prototype, 'room', {
 get: function() {
     return
 },
 enumerable: false,
 configurable: true
})

module.exports.loop = function () {
    global.pathfinder = Pathfinder
    global.game = Game
    global.consol = console
    global.memory = Memory
    run()
}
```

The two property overrides are required to avoid an infinite loop when
converting Rooms, StructureContollers, and StructureTerminals because they
reference each other. eg room.controller.room.controller.room.controller...
