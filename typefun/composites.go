package typefun

type Myintplus struct {
	Myint
}

func (this Myintplus) Treble() Myint {
	return this.Myint * 3
}
