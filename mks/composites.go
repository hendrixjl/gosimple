package mks

type MetersPerSec struct {
	mps float64
}

func (this Meters) Divide(s Seconds) MetersPerSec {
	return MetersPerSec{this.m / s.s}
}

func (this MetersPerSec) Mult(s Seconds) Meters {
	return Meters{this.mps / s.s}
}

type MetersPerSec2 struct {
	mps2 float64
}

func (this MetersPerSec) Divide(s Seconds) MetersPerSec2 {
	return MetersPerSec2{this.mps / s.s}
}
