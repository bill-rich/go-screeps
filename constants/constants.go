package constants

var (
	FIND_EXIT_TOP                   = 1
	FIND_EXIT_RIGHT                 = 3
	FIND_EXIT_BOTTOM                = 5
	FIND_EXIT_LEFT                  = 7
	FIND_EXIT                       = 10
	FIND_CREEPS                     = 101
	FIND_MY_CREEPS                  = 102
	FIND_HOSTILE_CREEPS             = 103
	FIND_SOURCES_ACTIVE             = 104
	FIND_SOURCES                    = 105
	FIND_DROPPED_RESOURCES          = 106
	FIND_STRUCTURES                 = 107
	FIND_MY_STRUCTURES              = 108
	FIND_HOSTILE_STRUCTURES         = 109
	FIND_FLAGS                      = 110
	FIND_CONSTRUCTION_SITES         = 111
	FIND_MY_SPAWNS                  = 112
	FIND_HOSTILE_SPAWNS             = 113
	FIND_MY_CONSTRUCTION_SITES      = 114
	FIND_HOSTILE_CONSTRUCTION_SITES = 115
	FIND_MINERALS                   = 116
	FIND_NUKES                      = 117
	FIND_TOMBSTONES                 = 118
	FIND_POWER_CREEPS               = 119
	FIND_MY_POWER_CREEPS            = 120
	FIND_HOSTILE_POWER_CREEPS       = 121
	FIND_DEPOSITS                   = 122
	FIND_RUINS                      = 123

	TOP          = 1
	TOP_RIGHT    = 2
	RIGHT        = 3
	BOTTOM_RIGHT = 4
	BOTTOM       = 5
	BOTTOM_LEFT  = 6
	LEFT         = 7
	TOP_LEFT     = 8

	COLOR_RED    = 1
	COLOR_PURPLE = 2
	COLOR_BLUE   = 3
	COLOR_CYAN   = 4
	COLOR_GREEN  = 5
	COLOR_YELLOW = 6
	COLOR_ORANGE = 7
	COLOR_BROWN  = 8
	COLOR_GREY   = 9
	COLOR_WHITE  = 10

	LOOK_CREEPS             = "creep"
	LOOK_ENERGY             = "energy"
	LOOK_RESOURCES          = "resource"
	LOOK_SOURCES            = "source"
	LOOK_MINERALS           = "mineral"
	LOOK_DEPOSITS           = "deposit"
	LOOK_STRUCTURES         = "structure"
	LOOK_FLAGS              = "flag"
	LOOK_CONSTRUCTION_SITES = "constructionSite"
	LOOK_NUKES              = "nuke"
	LOOK_TERRAIN            = "terrain"
	LOOK_TOMBSTONES         = "tombstone"
	LOOK_POWER_CREEPS       = "powerCreep"
	LOOK_RUINS              = "ruin"

	OBSTACLE_OBJECT_TYPES = []string{"spawn", "creep", "powerCreep", "source", "mineral", "deposit", "controller", "constructedWall", "extension", "link", "storage", "tower", "observer", "powerSpawn", "powerBank", "lab", "terminal", "nuker", "factory", "invaderCore"}

	MOVE          = "move"
	WORK          = "work"
	CARRY         = "carry"
	ATTACK        = "attack"
	RANGED_ATTACK = "ranged_attack"
	TOUGH         = "tough"
	HEAL          = "heal"
	CLAIM         = "claim"

	BODYPART_COST = map[string]int{
		"move":          50,
		"work":          100,
		"attack":        80,
		"carry":         50,
		"heal":          250,
		"ranged_attack": 150,
		"tough":         10,
		"claim":         600,
	}

	// WORLD_WIDTH and WORLD_HEIGHT constants are deprecated, please use Game.map.getWorldSize() instead
	WORLD_WIDTH  = 202
	WORLD_HEIGHT = 202

	CREEP_LIFE_TIME       = 1500
	CREEP_CLAIM_LIFE_TIME = 600
	CREEP_CORPSE_RATE     = 0.2
	CREEP_PART_MAX_ENERGY = 125

	CARRY_CAPACITY           = 50
	HARVEST_POWER            = 2
	HARVEST_MINERAL_POWER    = 1
	HARVEST_DEPOSIT_POWER    = 1
	REPAIR_POWER             = 100
	DISMANTLE_POWER          = 50
	BUILD_POWER              = 5
	ATTACK_POWER             = 30
	UPGRADE_CONTROLLER_POWER = 1
	RANGED_ATTACK_POWER      = 10
	HEAL_POWER               = 12
	RANGED_HEAL_POWER        = 4
	REPAIR_COST              = 0.01
	DISMANTLE_COST           = 0.005

	RAMPART_DECAY_AMOUNT = 300
	RAMPART_DECAY_TIME   = 100
	RAMPART_HITS         = 1
	RAMPART_HITS_MAX     = map[int]int{
		2: 300000,
		3: 1000000,
		4: 3000000,
		5: 10000000,
		6: 30000000,
		7: 100000000,
		8: 300000000,
	}

	ENERGY_REGEN_TIME = 300
	ENERGY_DECAY      = 1000

	SPAWN_HITS            = 5000
	SPAWN_ENERGY_START    = 300
	SPAWN_ENERGY_CAPACITY = 300
	CREEP_SPAWN_TIME      = 3
	SPAWN_RENEW_RATIO     = 1.2

	SOURCE_ENERGY_CAPACITY         = 3000
	SOURCE_ENERGY_NEUTRAL_CAPACITY = 1500
	SOURCE_ENERGY_KEEPER_CAPACITY  = 4000

	WALL_HITS     = 1
	WALL_HITS_MAX = 300000000

	EXTENSION_HITS            = 1000
	EXTENSION_ENERGY_CAPACITY = map[int]int{
		0: 50,
		1: 50,
		2: 50,
		3: 50,
		4: 50,
		5: 50,
		6: 50,
		7: 100,
		8: 200,
	}

	ROAD_HITS                = 5000
	ROAD_WEAROUT             = 1
	ROAD_WEAROUT_POWER_CREEP = 100
	ROAD_DECAY_AMOUNT        = 100
	ROAD_DECAY_TIME          = 1000

	LINK_HITS       = 1000
	LINK_HITS_MAX   = 1000
	LINK_CAPACITY   = 800
	LINK_COOLDOWN   = 1
	LINK_LOSS_RATIO = 0.03

	STORAGE_CAPACITY = 1000000
	STORAGE_HITS     = 10000

	STRUCTURE_SPAWN        = "spawn"
	STRUCTURE_EXTENSION    = "extension"
	STRUCTURE_ROAD         = "road"
	STRUCTURE_WALL         = "constructedWall"
	STRUCTURE_RAMPART      = "rampart"
	STRUCTURE_KEEPER_LAIR  = "keeperLair"
	STRUCTURE_PORTAL       = "portal"
	STRUCTURE_CONTROLLER   = "controller"
	STRUCTURE_LINK         = "link"
	STRUCTURE_STORAGE      = "storage"
	STRUCTURE_TOWER        = "tower"
	STRUCTURE_OBSERVER     = "observer"
	STRUCTURE_POWER_BANK   = "powerBank"
	STRUCTURE_POWER_SPAWN  = "powerSpawn"
	STRUCTURE_EXTRACTOR    = "extractor"
	STRUCTURE_LAB          = "lab"
	STRUCTURE_TERMINAL     = "terminal"
	STRUCTURE_CONTAINER    = "container"
	STRUCTURE_NUKER        = "nuker"
	STRUCTURE_FACTORY      = "factory"
	STRUCTURE_INVADER_CORE = "invaderCore"

	CONSTRUCTION_COST = map[string]int{
		"spawn":           15000,
		"extension":       3000,
		"road":            300,
		"constructedWall": 1,
		"rampart":         1,
		"link":            5000,
		"storage":         30000,
		"tower":           5000,
		"observer":        8000,
		"powerSpawn":      100000,
		"extractor":       5000,
		"lab":             50000,
		"terminal":        100000,
		"container":       5000,
		"nuker":           100000,
		"factory":         100000,
	}
	CONSTRUCTION_COST_ROAD_SWAMP_RATIO = 5
	CONSTRUCTION_COST_ROAD_WALL_RATIO  = 150

	CONTROLLER_LEVELS = map[int]int{
		1: 200,
		2: 45000,
		3: 135000,
		4: 405000,
		5: 1215000,
		6: 3645000,
		7: 10935000,
	}
	CONTROLLER_DOWNGRADE = map[int]int{
		1: 20000,
		2: 10000,
		3: 20000,
		4: 40000,
		5: 80000,
		6: 120000,
		7: 150000,
		8: 200000,
	}
	CONTROLLER_DOWNGRADE_RESTORE            = 100
	CONTROLLER_DOWNGRADE_SAFEMODE_THRESHOLD = 5000
	CONTROLLER_CLAIM_DOWNGRADE              = 300
	CONTROLLER_RESERVE                      = 1
	CONTROLLER_RESERVE_MAX                  = 5000
	CONTROLLER_MAX_UPGRADE_PER_TICK         = 15
	CONTROLLER_ATTACK_BLOCKED_UPGRADE       = 1000
	CONTROLLER_NUKE_BLOCKED_UPGRADE         = 200

	SAFE_MODE_DURATION = 20000
	SAFE_MODE_COOLDOWN = 50000
	SAFE_MODE_COST     = 1000

	TOWER_HITS          = 3000
	TOWER_CAPACITY      = 1000
	TOWER_ENERGY_COST   = 10
	TOWER_POWER_ATTACK  = 600
	TOWER_POWER_HEAL    = 400
	TOWER_POWER_REPAIR  = 800
	TOWER_OPTIMAL_RANGE = 5
	TOWER_FALLOFF_RANGE = 20
	TOWER_FALLOFF       = 0.75

	OBSERVER_HITS  = 500
	OBSERVER_RANGE = 10

	POWER_BANK_HITS          = 2000000
	POWER_BANK_CAPACITY_MAX  = 5000
	POWER_BANK_CAPACITY_MIN  = 500
	POWER_BANK_CAPACITY_CRIT = 0.3
	POWER_BANK_DECAY         = 5000
	POWER_BANK_HIT_BACK      = 0.5

	POWER_SPAWN_HITS            = 5000
	POWER_SPAWN_ENERGY_CAPACITY = 5000
	POWER_SPAWN_POWER_CAPACITY  = 100
	POWER_SPAWN_ENERGY_RATIO    = 50

	EXTRACTOR_HITS     = 500
	EXTRACTOR_COOLDOWN = 5

	LAB_HITS             = 500
	LAB_MINERAL_CAPACITY = 3000
	LAB_ENERGY_CAPACITY  = 2000
	LAB_BOOST_ENERGY     = 20
	LAB_BOOST_MINERAL    = 30
	LAB_COOLDOWN         = 10
	LAB_REACTION_AMOUNT  = 5
	LAB_UNBOOST_ENERGY   = 0
	LAB_UNBOOST_MINERAL  = 15

	GCL_POW      = 2.4
	GCL_MULTIPLY = 1000000
	GCL_NOVICE   = 3

	TERRAIN_MASK_WALL  = 1
	TERRAIN_MASK_SWAMP = 2
	TERRAIN_MASK_LAVA  = 4

	MAX_CONSTRUCTION_SITES = 100
	MAX_CREEP_SIZE         = 50

	MINERAL_REGEN_TIME = 50000
	MINERAL_MIN_AMOUNT = map[string]int{
		"H": 35000,
		"O": 35000,
		"L": 35000,
		"K": 35000,
		"Z": 35000,
		"U": 35000,
		"X": 35000,
	}
	MINERAL_RANDOM_FACTOR = 2

	MINERAL_DENSITY = map[int]int{
		1: 15000,
		2: 35000,
		3: 70000,
		4: 100000,
	}
	MINERAL_DENSITY_PROBABILITY = map[int]float32{
		1: 0.1,
		2: 0.5,
		3: 0.9,
		4: 1.0,
	}
	MINERAL_DENSITY_CHANGE = 0.05

	DENSITY_LOW      = 1
	DENSITY_MODERATE = 2
	DENSITY_HIGH     = 3
	DENSITY_ULTRA    = 4

	DEPOSIT_EXHAUST_MULTIPLY = 0.001
	DEPOSIT_EXHAUST_POW      = 1.2
	DEPOSIT_DECAY_TIME       = 50000

	TERMINAL_CAPACITY  = 300000
	TERMINAL_HITS      = 3000
	TERMINAL_SEND_COST = 0.1
	TERMINAL_MIN_SEND  = 100
	TERMINAL_COOLDOWN  = 10

	CONTAINER_HITS             = 250000
	CONTAINER_CAPACITY         = 2000
	CONTAINER_DECAY            = 5000
	CONTAINER_DECAY_TIME       = 100
	CONTAINER_DECAY_TIME_OWNED = 500

	NUKER_HITS             = 1000
	NUKER_COOLDOWN         = 100000
	NUKER_ENERGY_CAPACITY  = 300000
	NUKER_GHODIUM_CAPACITY = 5000
	NUKE_LAND_TIME         = 50000
	NUKE_RANGE             = 10
	NUKE_DAMAGE            = map[int]int{
		0: 10000000,
		2: 5000000,
	}

	FACTORY_HITS     = 1000
	FACTORY_CAPACITY = 50000

	TOMBSTONE_DECAY_PER_PART    = 5
	TOMBSTONE_DECAY_POWER_CREEP = 500

	RUIN_DECAY            = 500
	RUIN_DECAY_STRUCTURES = map[string]int{
		"powerBank": 10,
	}

	PORTAL_DECAY = 30000

	ORDER_SELL = "sell"
	ORDER_BUY  = "buy"

	MARKET_FEE = 0.05

	MARKET_MAX_ORDERS = 300

	FLAGS_LIMIT = 10000

	SUBSCRIPTION_TOKEN = "token"
	CPU_UNLOCK         = "cpuUnlock"
	PIXEL              = "pixel"
	ACCESS_KEY         = "accessKey"

	PIXEL_CPU_COST = 5000

	RESOURCE_ENERGY = "energy"
	RESOURCE_POWER  = "power"

	RESOURCE_HYDROGEN  = "H"
	RESOURCE_OXYGEN    = "O"
	RESOURCE_UTRIUM    = "U"
	RESOURCE_LEMERGIUM = "L"
	RESOURCE_KEANIUM   = "K"
	RESOURCE_ZYNTHIUM  = "Z"
	RESOURCE_CATALYST  = "X"
	RESOURCE_GHODIUM   = "G"

	RESOURCE_SILICON = "silicon"
	RESOURCE_METAL   = "metal"
	RESOURCE_BIOMASS = "biomass"
	RESOURCE_MIST    = "mist"

	RESOURCE_HYDROXIDE        = "OH"
	RESOURCE_ZYNTHIUM_KEANITE = "ZK"
	RESOURCE_UTRIUM_LEMERGITE = "UL"

	RESOURCE_UTRIUM_HYDRIDE    = "UH"
	RESOURCE_UTRIUM_OXIDE      = "UO"
	RESOURCE_KEANIUM_HYDRIDE   = "KH"
	RESOURCE_KEANIUM_OXIDE     = "KO"
	RESOURCE_LEMERGIUM_HYDRIDE = "LH"
	RESOURCE_LEMERGIUM_OXIDE   = "LO"
	RESOURCE_ZYNTHIUM_HYDRIDE  = "ZH"
	RESOURCE_ZYNTHIUM_OXIDE    = "ZO"
	RESOURCE_GHODIUM_HYDRIDE   = "GH"
	RESOURCE_GHODIUM_OXIDE     = "GO"

	RESOURCE_UTRIUM_ACID        = "UH2O"
	RESOURCE_UTRIUM_ALKALIDE    = "UHO2"
	RESOURCE_KEANIUM_ACID       = "KH2O"
	RESOURCE_KEANIUM_ALKALIDE   = "KHO2"
	RESOURCE_LEMERGIUM_ACID     = "LH2O"
	RESOURCE_LEMERGIUM_ALKALIDE = "LHO2"
	RESOURCE_ZYNTHIUM_ACID      = "ZH2O"
	RESOURCE_ZYNTHIUM_ALKALIDE  = "ZHO2"
	RESOURCE_GHODIUM_ACID       = "GH2O"
	RESOURCE_GHODIUM_ALKALIDE   = "GHO2"

	RESOURCE_CATALYZED_UTRIUM_ACID        = "XUH2O"
	RESOURCE_CATALYZED_UTRIUM_ALKALIDE    = "XUHO2"
	RESOURCE_CATALYZED_KEANIUM_ACID       = "XKH2O"
	RESOURCE_CATALYZED_KEANIUM_ALKALIDE   = "XKHO2"
	RESOURCE_CATALYZED_LEMERGIUM_ACID     = "XLH2O"
	RESOURCE_CATALYZED_LEMERGIUM_ALKALIDE = "XLHO2"
	RESOURCE_CATALYZED_ZYNTHIUM_ACID      = "XZH2O"
	RESOURCE_CATALYZED_ZYNTHIUM_ALKALIDE  = "XZHO2"
	RESOURCE_CATALYZED_GHODIUM_ACID       = "XGH2O"
	RESOURCE_CATALYZED_GHODIUM_ALKALIDE   = "XGHO2"

	RESOURCE_OPS = "ops"

	RESOURCE_UTRIUM_BAR    = "utrium_bar"
	RESOURCE_LEMERGIUM_BAR = "lemergium_bar"
	RESOURCE_ZYNTHIUM_BAR  = "zynthium_bar"
	RESOURCE_KEANIUM_BAR   = "keanium_bar"
	RESOURCE_GHODIUM_MELT  = "ghodium_melt"
	RESOURCE_OXIDANT       = "oxidant"
	RESOURCE_REDUCTANT     = "reductant"
	RESOURCE_PURIFIER      = "purifier"
	RESOURCE_BATTERY       = "battery"

	RESOURCE_COMPOSITE = "composite"
	RESOURCE_CRYSTAL   = "crystal"
	RESOURCE_LIQUID    = "liquid"

	RESOURCE_WIRE       = "wire"
	RESOURCE_SWITCH     = "switch"
	RESOURCE_TRANSISTOR = "transistor"
	RESOURCE_MICROCHIP  = "microchip"
	RESOURCE_CIRCUIT    = "circuit"
	RESOURCE_DEVICE     = "device"

	RESOURCE_CELL     = "cell"
	RESOURCE_PHLEGM   = "phlegm"
	RESOURCE_TISSUE   = "tissue"
	RESOURCE_MUSCLE   = "muscle"
	RESOURCE_ORGANOID = "organoid"
	RESOURCE_ORGANISM = "organism"

	RESOURCE_ALLOY      = "alloy"
	RESOURCE_TUBE       = "tube"
	RESOURCE_FIXTURES   = "fixtures"
	RESOURCE_FRAME      = "frame"
	RESOURCE_HYDRAULICS = "hydraulics"
	RESOURCE_MACHINE    = "machine"

	RESOURCE_CONDENSATE  = "condensate"
	RESOURCE_CONCENTRATE = "concentrate"
	RESOURCE_EXTRACT     = "extract"
	RESOURCE_SPIRIT      = "spirit"
	RESOURCE_EMANATION   = "emanation"
	RESOURCE_ESSENCE     = "essence"
)

// ErrorCodes is a list of ErrorCodes used in Screeps.
var ErrorCodes = map[string]int{
	"OK":                        0,
	"ERR_NOT_OWNER":             -1,
	"ERR_NO_PATH":               -2,
	"ERR_NAME_EXISTS":           -3,
	"ERR_BUSY":                  -4,
	"ERR_NOT_FOUND":             -5,
	"ERR_NOT_ENOUGH_ENERGY":     -6,
	"ERR_NOT_ENOUGH_RESOURCES":  -6,
	"ERR_INVALID_TARGET":        -7,
	"ERR_FULL":                  -8,
	"ERR_NOT_IN_RANGE":          -9,
	"ERR_INVALID_ARGS":          -10,
	"ERR_TIRED":                 -11,
	"ERR_NO_BODYPART":           -12,
	"ERR_NOT_ENOUGH_EXTENSIONS": -6,
	"ERR_RCL_NOT_ENOUGH":        -14,
	"ERR_GCL_NOT_ENOUGH":        -15,
}
