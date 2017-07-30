package game

// Tile enum
const (
	_ Tile = iota
	_
	_
	_
	_
	_
	_
	_
	_
	_
	// We start at 10, so the csvs can be more easily read (evenly spaced columns)
	// Subtract 7 from the line number for the csv number
	Sand
	Coral
	PurpleCoralGate
	PurpleCoralSwitch
	BlueCoralGate
	BlueCoralSwitch
	TealCoralGate
	TealCoralSwitch
	Sandtrap
	Sandglob
	Sandgeyser
	JeremyTile
	Crab
	Treasure
	SandKey
	CoralExit
	GreenCoralGate
	GreenCoralSwitch
	GreenCoralGateOff
	PurpleCoralGateOff
	BlueCoralGateOff
	TealCoralGateOff
)

type Tile int
