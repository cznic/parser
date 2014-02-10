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

import "x"
import "y"
import `z`
import ()
import ( "x16" )
import ( "x17"; )
import ( "x18"; "y18" )
import ( "x19"; "y19"; )
import (
)
import (
	"x23"
)
import (
	"x26";
)
import (
	"x29"
	"y30"
)
import (
	"x33"
	"y34";
)
import (
	"x37";
	"y38"
)
import (
	"x41";
	"y42";
)
import a44 "x"
import a45 "y"
import a46_ `z`
import ()
import ( a48 "x" )
import ( a49 "x"; )
import ( a50 "x"; a50b "y" )
import ( a51 "x"; a51b "y"; )
import (
)
import (
	a55 "x"
)
import (
	a58 "x";
)
import (
	a61 "x"
	a62 "y"
)
import (
	a65 "x"
	a66 "y";
)
import (
	a69 "x";
	a70 "y"
)
import (
	a73 "x";
	a74 "y";
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

// ============================================================================
type u77 uint;
type u78 uint
type t78 u
type t79 u;
type u_V int;
type t80 u.V
type t81 u.V;
type ()
type ( t85 u )
type ( t86 u; )
type w88 int
type ( t87 u; v87 w )
type ( t89 u; v89 w; )
type (
	t92 u
)
type (
	t95 u;
)
type (
	t98 u
	v99 w
)
type (
	t102 u
	v103 w;
)
type (
	t106 u;
	v107 w
)
type (
	t110 u;
	v111 w;
)
type t113 (u)
type t114 *u
type t115 (*u)
type t116 *(u)
type t118 struct{}
type t119 struct{ u }
type t119a struct{ *u }
type t119b struct{ int }
type t119c struct{ *int }
type t120 struct{ u; }
type t121 struct{
	u
}
type t124 struct{
	u;
}
type t127 struct{
	u.V
}
type t130 struct{
	*u
}
type t133 struct{
	*u.V
}
type X int
type y_Z int
type t136 struct{
	i int
	j, k uint
	l,
	m,
	n *uint
	o struct{
		p, q bool
		*X
		*y.Z
	}
}
type t136b *struct{
	i int
	j, k uint
	l,
	m,
	n *uint
	o struct{
		p, q bool
		*X
		*y.Z
	}
}
type t136c **struct{
	i int
	j, k uint
	l,
	m,
	n *uint
	o struct{
		p, q bool
		*X
		*y.Z
	}
}
type t140 int		
type t141 *int		
type t142 **int		

type t150 [3]int
type t151 [3]*int
type t152 [3]**int	
type t153 *[3]int
type t154 *[3]*int	
type t155 *[3]**int	
type t156 **[3]int	
type t157 **[3]*int	
type t158 **[3]**int	

type t188 bool;
type t189 int8;
type t190 int16;
type t191 int32;
type t192 int64;
type t193 int;
type t194 uint8;
type t195 uint16;
type t196 uint32;
type t197 uint64;
type t198 uint;
type t199 float32;
type t200 float64;
type t201 complex64;
type t202 complex128;

type t188_b *bool;
type t189_b *int8;
type t190_b *int16;
type t191_b *int32;
type t192_b *int64;
type t193_b *int;
type t194_b *uint8;
type t195_b *uint16;
type t196_b *uint32;
type t197_b *uint64;
type t198_b *uint;
type t199_b *float32;
type t200_b *float64;
type t201_b *complex64;
type t202_b *complex128;

type t188_c **bool;
type t189_c **int8;
type t190_c **int16;
type t191_c **int32;
type t192_c **int64;
type t193_c **int;
type t194_c **uint8;
type t195_c **uint16;
type t196_c **uint32;
type t197_c **uint64;
type t198_c **uint;
type t199_c **float32;
type t200_c **float64;
type t201_c **complex64;
type t202_c **complex128;

type t236 func()			
type t237 func() x		
type t238 func() (x)		
type t239 func() (x y)		
type t240 func(int)		
type t241 func(int, uint)		
type t242 func(...int)		
type t243 func(int, ...int)	
type t244 func(int, int, ...int)	
type t245 func(x int)		
type t246 func(x ...int)		
type t247 func(x int, y uint)	
type t248 func(x, y int)		
type t250 func(x, y int, z ...uint)

type t255 func()()
type t256 func() int

const c258 = 3
const c259 = 3;
const c260 t = 3
const c261 t = 3;
const ()
const (
	_ = 1 << iota
	a268
	b269
	c270
)
const (
	bit0371, mask0371 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
	bit1372, mask1372                           // bit1 == 2, mask1 == 1
	_, _                                  // skips iota == 2
	bit3374, mask3374                           // bit3 == 8, mask3 == 7
)

var v279 t
var v280 t;
var v283 = 3
var v284 = 3;
var v285 t = 3
var v286 t = 3;
var ()
var ( v289 t )
var ( v290 t; )
var ( v291 = 3 )
var ( v292 = 3; )
var ( v293 t = 3 )
var ( v294 t = 3; )
var (
	v296 t
)
var (
	v299 t;
	u300 *struct{
		bar
		baz.Q
		*int
	}
	w305 [3]int
	w306 [3]*int
	w307 *[3]int
	w308 *[3]*int
)
var v310 = [1]int{2}
var v311 = [1][3]int{2}
var v312 = [1]struct{}{2}
var v313 = [1]struct{i int}{2}
var v314 = [2]int{1: 3}
var v315 = [2]int{1: 3, 2: 4}
var v316 = [2]int{1+2: 3}
var v317 = [2]int{"bad": 3}
var v318 = [2]int{ident: 3}
var v319 = struct{}{}
var v320 = struct{in, out int}{}
var v321 = struct{in, out int}{1, 2}
var v322 = struct{in, out int}{in: 1}
var v323 = [2]struct{in, out int}{in: 1}[0]
var a324, b324, c324 = d, e+1, f-2
var v325 = [3]int{4, 5}
var v326 = [3]int{4, 5}[1]
var v327 = struct{in, out int}{1, 2}.in

func f329()
func f330() x
func f331() (x)
func f332() (x y)
func f333(int)
func f334(int, int)
func f335(...int)
func f336(int, ...int)
func f337(int, int, ...int)
func f338(x int)
func f339(x ...int)
func f340(x int, y int)
func f341(x, y int)
func f343(x, y int, z ...int)
func f344() {}
func f345() {;}
func f346() {2}
func f347() {2;}
func f348() {;}
func f349() {;;}
func f350() {;2}
func f351() {;2;}
func f352() {1}
func f353() {1;}
func f354() {12}
func f355() {12;}
func f356() {1;}
func f357() {1;;}
func f358() {1;2}
func f359() {1;2;}
func f360() {;1;2;}
func f361() { a = 1 }
func f364() { a, b = 1, 2 }
func f365() { a += 1 }
func f366() { a -= 1 }
func f367() { a *= 1 }
func f368() { a /= 1 }
func f369() { a %= 1 }
func f370() { a >>= 1 }
func f371() { a <<= 1 }
func f372() { a &= 1 }
func f373() { a &^= 1 }
func f374() { a ^= 1 }
func f375() { v = 42 }
//TODOfunc f376() {
//TODO	v = 42
//TODO}
//TODOfunc f379() {
//TODO	v = 42
//TODO	var ()
//TODO}
//TODOfunc f383() {
//TODO	v = 42
//TODO	var (x int)
//TODO}
//TODOfunc f387() {
//TODO	v = 42
//TODO	var (x int; y bool)
//TODO}
//TODOfunc f391() {
//TODO	var ()
//TODO	v = 42
//TODO}
//TODOfunc f395() {
//TODO	var (x int)
//TODO	v = 42
//TODO}
//TODOfunc f399() {
//TODO	var (x int; y bool)
//TODO	v = 42
//TODO}
//TODOfunc f403() {
//TODO	v = 42
//TODO	{
//TODO		w = 314
//TODO	}
//TODO}
//TODOvar v410 = f()
//TODOvar v411 = f(1)
//TODOvar v412 = f(1,)
//TODOvar v416 = f(
//TODO	1,
//TODO)
//TODOvar v419 = uint(3)
//TODOvar v420 = (uint)(3)
//TODOvar v421 = *uint(3)
//TODOvar v422 = (*uint)(3)
//TODOvar v423 = **uint(3)
//TODOvar v424 = (**uint)(3)
//TODOvar v425 = 3.(uint) // callOp{float_lit: 3, args{"uint"}}
//TODOvar v426 *int32
//TODOvar w427, x427, y427 = uint(*a), *(*uint)(b), (uint)(*c)
//TODO
//TODOvar v429 = (int)(x)
//TODOvar v430 = int(x)
//TODO
//TODOvar v432 = (*int)(x)
//TODOvar v433 = *(int)(x)
//TODOvar v434 = *int(x)
//TODO
//TODOvar v436 = (**int)(x)
//TODOvar v437 = *(*int)(x)
//TODOvar v438 = **int(x)
//TODO
//TODOfunc f423() {
//TODO	a++
//TODO	b--
//TODO	a++;
//TODO	b--;
//TODO	a = b
//TODO	c, d *= 2, 4
//TODO	f, err := open(name)
//TODO	f, err = open(name)
//TODO	var f432, err432 = open(name)
//TODO	return
//TODO	return;
//TODO	return foo
//TODO	return foo;
//TODO	return foo, bar, 3
//TODO	return foo, bar, 3;
//TODO	return 42, "blah", 3.14
//TODO	return 42, "blah", 3.14;
//TODO	break
//TODO	break;
//TODO	break loop
//TODO	break loop2;
//TODO	continue
//TODO	continue;
//TODO	continue loop
//TODO	continue loop2;
//TODO	goto ident
//TODO	goto ident2;
//TODO	fallthrough
//TODO	fallthrough;
//TODO	if 1 {
//TODO		println(42)
//TODO	}
//TODO	if (1) {
//TODO		println(42)
//TODO	}
//TODO	if x {
//TODO		println(42)
//TODO	}
//TODO	if (x) {
//TODO		println(42)
//TODO	}
//TODO	if x.y {}
//TODO	if (x.y) {}
//TODO	if (x{}) {
//TODO		println(42)
//TODO	}
//TODO	if (int{}) {
//TODO		println(42)
//TODO	}
//TODO	if (int{1}) {
//TODO		println(42)
//TODO	}
//TODO	if ([...]int{1}) {
//TODO		println(42)
//TODO	}
//TODO	if ([...]int{1}[0]) {
//TODO		println(42)
//TODO	}
//TODO	if ([...]int{1}[0] == 0) {
//TODO		println(42)
//TODO	}
//TODO	if ([...]int{1}[0]) == 0 {
//TODO		println(42)
//TODO	}
//TODO	if ([1]int{1}) {
//TODO	     println(42)
//TODO	}
//TODO	if ([1]int{1}[0]) {
//TODO	     println(42)
//TODO	}
//TODO	if ([1]int{1}[0] == 0) {
//TODO	     println(42)
//TODO	}
//TODO	if ([1]int{1}[0]) == 0 {
//TODO		println(42)
//TODO	}
//TODO	if x == (T{a,b,c}[i]) {
//TODO		println(42)
//TODO	}
//TODO	if (x == T{a,b,c}[i]) {
//TODO		println(42)
//TODO	}
//TODO	if (x.y{false, true}[0]) {}
//TODO	if x == (p.T{a,b,c}[i]) {
//TODO		println(42)
//TODO	}
//TODO	if (x == p.T{a,b,c}[i]) {
//TODO		println(42)
//TODO	}
//TODO	if n = len(l.stack); n != 0 {
//TODO		l.stack[n-1]++
//TODO	}
//TODO	if (func() bool { return true }()) { //LATER lift parens requirement in this special case (?)
//TODO		fmt.Println("Hello, playground")
//TODO	}
//TODO}
//TODO
//TODOvar v = func() bool {}
//TODO
//TODOfunc f521() {
//TODO	for {}
//TODO	for x {}
//TODO	for x > 0 {}
//TODO	for (x{1}[0]) != 0 {}
//TODO	for (x{1}[0] != 0) {}
//TODO	for ; ; {}
//TODO	for ; ; i++ {}
//TODO	for ; i < 10 ; {}
//TODO	for ; i < 10 ; i++ {}
//TODO	for i = 0 ; ; {}
//TODO	for i = 0 ; ; i++ {}
//TODO	for i = 0 ; i < 10 ; {}
//TODO	for i = 0 ; i < 10 ; i++ {}
//TODO	a = b
//TODO	a := b
//TODO	for i = range x {}
//TODO	for i, v = range x {}
//TODO	for *p = range x {}
//TODO	for *p, l.x = range x {}
//TODO	for i := range x {}
//TODO	for i, v := range x {}
//TODO	for i := 0; i < 10; i++ {}
//TODO}
//TODO
//TODOfunc f547() {
//TODO	switch{}
//TODO	switch x {}
//TODO	switch (x{1})[0] {}
//TODO	switch (x{1}[0]) {}
//TODO	switch n := 2*m; {}
//TODO
//TODO	switch n := 2*m; {
//TODO	case n > b:
//TODO		break
//TODO	case n <= c:
//TODO		f()
//TODO		fallthrough
//TODO	default:
//TODO		fmt.Println(42)
//TODO	}
//TODO
//TODO	switch n := 2*m; n+3 {
//TODO	case 42:
//TODO		break
//TODO	case 255:
//TODO		float32(f(int64(n)))
//TODO		fallthrough
//TODO	default:
//TODO		fmt.Println(42)
//TODO	}
//TODOlabel:;
//TODOlabel2:
//TODOloop:
//TODO	for x {
//TODO		if y {
//TODO			break loop
//TODO		}
//TODO	}
//TODO
//TODO	f := float64(3.14)
//TODO	i := int16(f) // conversion, i == 3
//TODO	bits := *(*uint64)(&f) // cast, bits == IEE754 pattern for 64bit 3.14
//TODO	half := *(*uint32)(&f) // ugly cast
//TODO
//TODO	var arr [32]byte
//TODO	a591 := [math.MaxInt64]byte(arr)
//TODO	a592 := ([math.MaxInt64]byte(arr))
//TODO	a593 := *([math.MaxInt64]byte(arr))
//TODO	a594 := *[math.MaxInt64]byte(arr)
//TODO	a595 := (*[math.MaxInt64]byte(arr))
//TODO	a596 := *(*[math.MaxInt64]byte(arr))
//TODO	a597 := **[math.MaxInt64]byte(arr)
//TODO	a598 := (**[math.MaxInt64]byte(arr))
//TODO	a599 := *(**[math.MaxInt64]byte(arr))
//TODO}
//TODO
//TODO// Go
//TODOfunc Float32bits603(f float32) uint32 { return *(*uint32)(unsafe.Pointer(&f)) }
//TODOfunc Float32bits604(f float32) uint32 { return *(*uint32(unsafe.Pointer(&f))) }
//TODOfunc Float32frombits605(b uint32) float32 { return *(*float32)(unsafe.Pointer(&b)) }
//TODOfunc Float32frombits606(b uint32) float32 { return *(*float32(unsafe.Pointer(&b))) }
//TODOfunc Float64bits607(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }
//TODOfunc Float64bits608(f float64) uint64 { return *(*uint64(unsafe.Pointer(&f))) }
//TODOfunc Float64frombits609(b uint64) float64 { return *(*float64)(unsafe.Pointer(&b)) }
//TODOfunc Float64frombits610(b uint64) float64 { return *(*float64(unsafe.Pointer(&b))) }
//TODO// Run
//TODOfunc Float32bits612(f float32) uint32 { return *(*uint32(&f)) }
//TODOfunc Float32bits613(f float32) uint32 { return *(*uint32)(&f) }
//TODOfunc Float32frombits614(b uint32) float32 { return *(*float32(&b)) }
//TODOfunc Float32frombits615(b uint32) float32 { return *(*float32)(&b) }
//TODOfunc Float64bits616(f float64) uint64 { return *(*uint64)(&f) }
//TODOfunc Float64bits617(f float64) uint64 { return *(*uint64(&f)) }
//TODOfunc Float64frombits618(b uint64) float64 { return *(*float64)(&b) }
//TODOfunc Float64frombits619(b uint64) float64 { return *(*float64(&b)) }
//TODO
//TODOvar (
//TODO	fp622 *float64
//TODO	up623 *uint64 = *uint64(fp)
//TODO	up624 = *uint64(fp)
//TODO	v625 [n]t
//TODO	v626 = struct{
//TODO		u, v struct{
//TODO			int
//TODO		}
//TODO	}{
//TODO		u: {3},
//TODO	}
//TODO	v633 = struct{
//TODO		u struct{
//TODO			int
//TODO		}
//TODO	}{
//TODO		{3},
//TODO	}
//TODO)
//TODOfunc f641() {
//TODO	i := a || b && c
//TODO	j := a && b || c
//TODO	k := struct{int}(x)
//TODO	l := (struct{int}(y))
//TODO	m := (*[math.MaxInt64]byte(p))
//TODO	n := (*(struct{int}(z)))
//TODO	o := (*struct{int}(z))
//TODO	p := *(struct{int}(z))
//TODO	q := a >= b
//TODO}
//TODO
//TODOfunc (t) f653()
//TODOfunc (*t) f654()
//TODOfunc (x t) f655()
//TODOfunc (x *t) f656()
//TODOfunc f657() {
//TODO	if x {
//TODO		a
//TODO	} else if y {
//TODO		b
//TODO	}
//TODO	if x {
//TODO		a
//TODO	} else {
//TODO		b
//TODO	}
//TODO	i := 1i
//TODO	nl := '\n'
//TODO	v := [...]int{1, 2, 3}
//TODO	a = b ^ c
//TODO	a = b | c
//TODO	a = b &^ c
//TODO	a = b & c
//TODO	a = b >> c
//TODO	a = b % c
//TODO	a = b / c
//TODO	a |= b
//TODO}
//TODOfunc f680() t
//TODOfunc f681() *t
//TODOfunc f682() [3]int
//TODOfunc f683() *[3]int
//TODOfunc f684() func(int) bool
//TODOfunc f685() *func(int) bool
//TODOfunc f686() q.I
//TODOfunc f687() *q.I
//TODOfunc f688() struct{i int}
//TODOfunc f689() *struct{i int}
//TODO
//TODOfunc f691() (t)
//TODOfunc f692() (*t)
//TODOfunc f693() ([3]int)
//TODOfunc f694() (*[3]int)
//TODOfunc f695() (func(int) bool)
//TODOfunc f696() (*func(int) bool)
//TODOfunc f697() (q.I)
//TODOfunc f698() (*q.I)
//TODOfunc f699() (struct{i int})
//TODOfunc f700() (*struct{i int})
//TODO
//TODOfunc f714() (r t)
//TODOfunc f715() (r *t)
//TODOfunc f716() (r [3]int)
//TODOfunc f717() (r *[3]int)
//TODOfunc f718() (r func(int) bool)
//TODOfunc f719() (r *func(int) bool)
//TODOfunc f720() (r q.I)
//TODOfunc f721() (r *q.I)
//TODOfunc f722() (r struct{i int})
//TODOfunc f723() (r *struct{i int})
//TODO
//TODOfunc f737() (r t, err error)
//TODOfunc f738() (r *t, err error)
//TODOfunc f739() (r [3]int, err error)
//TODOfunc f740() (r *[3]int, err error)
//TODOfunc f741() (r func(int) bool, err error)
//TODOfunc f742() (r *func(int) bool, err error)
//TODOfunc f743() (r q.I, err error)
//TODOfunc f744() (r *q.I, err error)
//TODOfunc f745() (r struct{i int}, err error)
//TODOfunc f746() (r *struct{i int}, err error)
//TODO
//TODOfunc f748() {
//TODO	a = ^b
//TODO	a = ^b ^ c
//TODO	a = !b && !c
//TODO	a = -b + -c
//TODO	a = +b - +c
//TODO	g()
//TODO	g(42)
//TODO	g(42, 314)
//TODO}
//TODO
//TODOconst (
//TODO	a760 = 1
//TODO	b761, c761 = 2, 3
//TODO	d762, e762, f762 = 4, 5, 6
//TODO)
//TODO
//TODOfunc f765() {
//TODO	v := union{ u uint64; f float64 }{f: 3.14}.u // poor man's bits from float64
//TODO	v741 := union{ uint64; float64 }{float64: 3.14}.uint64
//TODO	var v742 union{
//TODO		u uint64
//TODO		f float64
//TODO	}
//TODO	type t union{
//TODO		u uint64
//TODO		f float64
//TODO	}
//TODO	type g func() union{a b; c d}
//TODO	v = (union{u uint64; f float64}(x)).u
//TODO	v = (*union{u uint64; f float64}(x)).u
//TODO	v753 := [3]int(x)
//TODO	v754 := *[3]int(x)
//TODO	v755 := ([3]int(x))
//TODO	v756 := (*[3]int(x))
//TODO	v757 := *([3]int(x))
//TODO	v758 :=  (struct{i int}(x))
//TODO	v759 := (*struct{i int}(x))
//TODO	v760 := *(struct{i int}(x))
//TODO	v761 :=  (union{i int; u uint}(x))
//TODO	v762 := (*union{i int; u uint}(x))
//TODO	v763 := *(union{i int; u uint}(x))
//TODO	v764 := t{1}(x) //LATER invalid
//TODO	v765 := int(x)
//TODO	v766 := *int(x)
//TODO	v767 := (*int)(x)
//TODO	v768 := [3]int(x)
//TODO	v769 := *[3]int(x)
//TODO	v770 := (*[3]int(x))
//TODO	v771 := struct{int}(x)
//TODO	v772 := *struct{int}(x)
//TODO	v773 := (*struct{int}(x))
//TODO	v774 := union{int}(x)
//TODO	v775 := *union{int}(x)
//TODO	v776 := (*union{int}(x))
//TODO
//TODO	v778 := func()()(x)
//TODO	v779 := (func()()(x))
//TODO
//TODO	v781 := func()()(c)
//TODO	v782 := func()(b)(c)
//TODO	v783 := func(a)()(c)
//TODO	v784 := func(a)(b)(c)
//TODO
//TODO	v786 := func(a)(b)(c)()  // call
//TODO	v787 := func(a)(b)(c)(d) // call
//TODO	v788 := func(a)()(c)()   // call
//TODO	v789 := func()(b)(c)()   // call
//TODO	v790 := func()()(o)()    // call
//TODO	v791 := func()(func())(o)()
//TODO	v = 1
//TODO}
//TODO
//TODOtype t827 union{}
//TODO
//TODOvar v797 int
//TODOfunc f798() {
//TODO	var v797 int
//TODO	{
//TODO		var v797 int
//TODO	}
//TODO}
//TODO
//TODO
//TODO//DONE scpecial scoping of ELSE
//TODO//	Turns out there's nothing special: http://play.golang.org/p/imXmTYTlEl
//TODO//	and: http://play.golang.org/p/IHTOPj0IYK
//TODOfunc f807() {
//TODO	if a := true; a {
//TODO		a := 42
//TODO	} else {
//TODO		a := "foo"
//TODO	}
//TODO}
//TODOfunc f818() {
//TODO	if a := 42; a != 42 {
//TODO		_ = a + 1
//TODO	} else {
//TODO		_ = a - 1
//TODO	}
//TODO}
//TODO
//TODOfunc (r) f826(r)
//TODO
//TODOtype t828 int
//TODOtype u829 uint
//TODO
//TODOfunc (t828) f()
//TODOfunc (t829) f()
