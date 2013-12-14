package typefun

type Myint int

func (i Myint) Double() Myint {
	return i * 2
}

type Anotherint Myint

func (i Anotherint) Treble() Anotherint {
	return i * 3
}
