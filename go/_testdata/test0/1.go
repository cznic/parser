package foo

import ()
import "foo4"
import foo5 "bar5"
import . "foo6"
import (
	"foo8"
)
import (
	"bar12"
	"foo11"
)

const ()
const c16 = 3e8
const c17 int = 3e8
const Pi18 float64 = 3.14159265358979323846
const zero19 = 0.0 // untyped floating-point constant
const (
	size21 int64 = 1024
)
const (
	size24 int64 = 1024
	eof25        = -1 // untyped integer constant
)
const a27, b27, c27 = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
const u28, v28 float32 = 0, 3     // u = 0.0, v = 3.0
const (
	Sunday30 = iota
	Monday31
	Tuesday32
	Wednesday33
	Thursday34
	Friday35
	Partyday36
	numberOfDays37 // this constant is not exported
)
const (  // iota is reset to 0
	c040 = iota  // c0 == 0
	c141 = iota  // c1 == 1
	c242 = iota  // c2 == 2
)

const (
	a46 = 1 << iota  // a == 1 (iota has been reset)
	b47 = 1 << iota  // b == 2
	c48 = 1 << iota  // c == 4
)

//const (
//	u52         = iota * 42  // u == 0     (untyped integer constant)
//	v53 float64 = iota * 42  // v == 42.0  (float64 constant)
//	w54         = iota * 42  // w == 84    (untyped integer constant)
//)
//
//const x57 = iota  // x == 0 (iota has been reset)
//const y58 = iota  // y == 0 (iota has been reset)
//Within an ExpressionList, the value of each iota is the same because it is only incremented after each ConstSpec:
//
//const (
//	bit062, mask062 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
//	bit163, mask163                           // bit1 == 2, mask1 == 1
//	_, _                                  // skips iota == 2
//	bit365, mask365                           // bit3 == 8, mask3 == 7
//)
