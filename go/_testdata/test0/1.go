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
