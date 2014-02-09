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
//TODOconst c16 = 3e8
//TODOconst c17 int = 3e8
//TODOconst Pi18 float64 = 3.14159265358979323846
//TODOconst zero19 = 0.0 // untyped floating-point constant
//TODOconst (
//TODO	size21 int64 = 1024
//TODO)
//TODOconst (
//TODO	size24 int64 = 1024
//TODO	eof25        = -1 // untyped integer constant
//TODO)
//TODOconst a27, b27, c27 = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
//TODOconst u28, v28 float32 = 0, 3     // u = 0.0, v = 3.0
//TODOconst (
//TODO	Sunday30 = iota
//TODO	Monday31
//TODO	Tuesday32
//TODO	Wednesday33
//TODO	Thursday34
//TODO	Friday35
//TODO	Partyday36
//TODO	numberOfDays37 // this constant is not exported
//TODO)
//TODOconst (  // iota is reset to 0
//TODO	c040 = iota  // c0 == 0
//TODO	c141 = iota  // c1 == 1
//TODO	c242 = iota  // c2 == 2
//TODO)
//TODO
//TODOconst (
//TODO	a46 = 1 << iota  // a == 1 (iota has been reset)
//TODO	b47 = 1 << iota  // b == 2
//TODO	c48 = 1 << iota  // c == 4
//TODO)
//TODO
//TODOconst (
//TODO	u52         = iota * 42  // u == 0     (untyped integer constant)
//TODO	v53 float64 = iota * 42  // v == 42.0  (float64 constant)
//TODO	w54         = iota * 42  // w == 84    (untyped integer constant)
//TODO)
//TODO
//TODOconst x57 = iota  // x == 0 (iota has been reset)
//TODOconst y58 = iota  // y == 0 (iota has been reset)
//TODOconst (
//TODO	bit062, mask062 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
//TODO	bit163, mask163                           // bit1 == 2, mask1 == 1
//TODO	_, _                                  // skips iota == 2
//TODO	bit365, mask365                           // bit3 == 8, mask3 == 7
//TODO)
//TODO
//TODO// ------------------------------------------------------------------- TypeDecl
//TODOtype IntArray70 [16]int
//TODOtype (
//TODO	Point72 struct{ x, y float64 }
//TODO)
//TODOtype (
//TODO	Point75 struct{ x, y float64 }
//TODO	Polar76 Point
//TODO)
//TODO//type TreeNode78 struct {
//TODO//	left, right *TreeNode
//TODO//	value *Comparable
//TODO//}
//TODO//type Block82 interface {
//TODO//	BlockSize() int
//TODO//	Encrypt(src, dst []byte)
//TODO//	Decrypt(src, dst []byte)
//TODO//}
//TODO//type Mutex87 struct         { /* Mutex fields */ }
//TODO//type NewMutex88 Mutex
//TODO//type PtrMutex89 *Mutex
//TODO//type PrintableMutex90 struct {
//TODO//	Mutex
//TODO//}
//TODO//type MyBlock93 Block
//TODO//type TimeZone94 int
//TODO//
//TODO//const (
//TODO//	EST TimeZone = -(5 + iota)
//TODO//	CST
//TODO//	MST
//TODO//	PST
//TODO//)
