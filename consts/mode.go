package consts

const (
	_ Mode = iota
	DebugMode
	TestingMode
	ReleaseMode
	ProductionMode
)

type Mode uint8
