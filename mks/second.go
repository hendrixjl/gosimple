package mks

type Seconds struct {
	s float64
}

func MakeSeconds(s float64) Seconds {
	return Seconds{s}
}

func (this Seconds) scale(s float64) Seconds {
	return Seconds{this.s * s}
}
