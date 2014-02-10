// ---------------------------------------------------------------------Package
package foo

// --------------------------------------------------------------------- Import
import ()
import "foo4"
import foo5 "bar5"
import . "foo6"
import (
	"foo8"
)
import (
	foo12 "bar12"
	_ "foo11"
)

// ------------------------------------------------------------------ ConstDecl
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
const ( // iota is reset to 0
	c040 = iota // c0 == 0
	c141 = iota // c1 == 1
	c242 = iota // c2 == 2
)

const (
	a46 = 1 << iota // a == 1 (iota has been reset)
	b47 = 1 << iota // b == 2
	c48 = 1 << iota // c == 4
)

const (
	u52         = iota * 42 // u == 0     (untyped integer constant)
	v53 float64 = iota * 42 // v == 42.0  (float64 constant)
	w54         = iota * 42 // w == 84    (untyped integer constant)
)

const x57 = iota // x == 0 (iota has been reset)
const y58 = iota // y == 0 (iota has been reset)
const (
	bit062, mask062 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
	bit163, mask163                          // bit1 == 2, mask1 == 1
	_, _                                     // skips iota == 2
	bit365, mask365                          // bit3 == 8, mask3 == 7
)

// ------------------------------------------------------------------- TypeDecl
type IntArray70 [16]int
type (
	Point72 struct{ x, y float64 }
)
type (
	Point75 struct{ x, y float64 }
	Polar76 Point
)
type TreeNode78 struct {
	left, right *TreeNode
	value       *Comparable
}
type Block82 interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
type Mutex87 struct         { /* Mutex fields */ }
type NewMutex88 Mutex
type PtrMutex89 *Mutex
type PrintableMutex90 struct {
	Mutex
}
type MyBlock93 Block
type TimeZone94 int

const (
	EST97 TimeZone = -(5 + iota)
	CST98
	MST99
	PST100
)

// -------------------------------------------------------------------- VarDecl
var i104 int
var U105, V105, W105 float64
var k106 = 0
var x107, y107 float32 = -1, -2
var (
	i109       int
	u110, v110, s110 = 2.0, 3.0, "bar"
)
var re112, im112 = complexSqrt(-1)
var _, found113 = entries[name]  // map lookup; only interested in "found"
