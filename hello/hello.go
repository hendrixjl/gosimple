package main

import (
	"fmt"

	. "github.com/hendrixjl/gosimple/mks"
	"github.com/hendrixjl/gosimple/newmath"
	"github.com/hendrixjl/gosimple/typefun"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))

	i := typefun.Myint(1)
	fmt.Printf("i=%v\n", i)
	fmt.Printf("i.Double()=%v\n", i.Double())
	i = i * 2
	fmt.Printf("i=%v\n", i)

	ii := int(5)
	ii = ii * ii
	fmt.Printf("ii=%v\n", ii)

	j := typefun.Anotherint(2)
	fmt.Printf("j.Treble()=%v\n", j.Treble())

	k := typefun.Myintplus{3}
	fmt.Printf("k.Double()=%v\n", k.Double())
	fmt.Printf("k.Treble()=%v\n", k.Treble())
	fmt.Printf("k.Double().Double()=%v\n", k.Double().Double())

	fmt.Printf("5*k.Myint=%v\n", 5*k.Myint)

	m := MakeMeters(5.0)
	s := MakeSeconds(2.0)
	mps := m.Divide(s)
	fmt.Printf("mps=%v\n", mps)
	mps2 := mps.Divide(s)
	fmt.Printf("mps2=%v\n", mps2)
}
