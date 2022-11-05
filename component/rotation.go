package component

type Rot struct {
	Radians float64
}

func NewRotF(rotation float64) Rot {
	return Rot{rotation}
}
