// meter.go
package mks

// simple meters
type Meters struct {
	m float64
}

func MakeMeters(m float64) Meters {
	return Meters{m}
}

func (this Meters) scale(s float64) Meters {
	return Meters{this.m * s}
}

// square meters
type Meters2 struct {
	m float64
}

func (this Meters) Mult(s Meters) Meters2 {
	return Meters2{this.m * s.m}
}

// cubic meters
type Meters3 struct {
	m float64
}

func (this Meters2) mult(s Meters) Meters3 {
	return Meters3{this.m * s.m}
}
