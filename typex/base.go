package typex

type Integer interface {
	int | int8 | int16 | int32 | int64 | UnsignedInteger
}

type UnsignedInteger interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	Integer | Float
}
