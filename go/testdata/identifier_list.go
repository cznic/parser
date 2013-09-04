package main

//TODO ParameterDecl 

var (
	_ = func(int) (i int) {} //TODO-
	_ = func() (i int) {} //TODO-
	_ = func() {}
	_ = func(int) {}
	_ = func(int, uint) {}
	_ = func(int, uint, float64) {}

	_ = func() {}
	_ = func(int,) {}
	_ = func(int, uint,) {}
	_ = func(int, uint, float64,) {}
)

type (
	_ func()
	_ func(int)
	_ func(int, uint)
	_ func(int, uint, float64)

	_ func()
	_ func(int,)
	_ func(int, uint,)
	_ func(int, uint, float64,)
)

var (
	_ = func(...int) {}
	_ = func(...T, U) {}
	_ = func(T, ...U) {}
	_ = func(int, uint, float64) {}
	_ = func(int, uint, ...float64) {}
	_ = func(int, ...uint, float64) {}
	_ = func(int, ...uint, ...float64) {}
	_ = func(...int, uint, float64) {}
	_ = func(...int, uint, ...float64) {}
	_ = func(...int, ...uint, float64) {}
	_ = func(...int, ...uint, ...float64) {}
)

type (
	_ func(...int)
	_ func(...T, U)
	_ func(T, ...U)
	_ func(int, uint, float64)
	_ func(int, uint, ...float64)
	_ func(int, ...uint, float64)
	_ func(int, ...uint, ...float64)
	_ func(...int, uint, float64)
	_ func(...int, uint, ...float64)
	_ func(...int, ...uint, float64)
	_ func(...int, ...uint, ...float64)
)

type (
	_ func()
	_ func(T)
	_ func(...T)
	_ func(id T)
	_ func(id ...T)
	_ func(id, id2 T)
	_ func(id, id2 ...T)
	_ func(id, id2 T, id3 U)
	_ func(id, id2 T, id3, id4 U)
	_ func(T1, T2)
)

var (
	_ = func() {}
	_ = func(T) {}
	_ = func(...T) {}
	_ = func(id T) {}
	_ = func(id ...T) {}
	_ = func(id, id2 T) {}
	_ = func(id, id2 ...T) {}
	_ = func(id, id2 T, id3 U) {}
	_ = func(id, id2 T, id3, id4 U) {}
	_ = func(T1, T2) {}
)

func _()
func _(T)
func _(...T)
func _(id T)
func _(id ...T)
func _(id, id2 T)
func _(id, id2 ...T)
func _(id, id2 T, id3 U)
func _(id, id2 T, id3, id4 U)
func _(T1, T2)

func _() {}
func _(T) {}
func _(...T) {}
func _(id T) {}
func _(id ...T) {}
func _(id, id2 T) {}
func _(id, id2 ...T) {}
func _(id, id2 T, id3 U) {}
func _(id, id2 T, id3, id4 U) {}
func _(T1, T2) {}

func (T) _()
func (*T) _()
func (id T) _()
func (id *T) _()

// ----------------------------------------------------------------------------

func M(f uint64) (in, out T) {
	in = make(T, 100)
	out = make(T, 100)
	go func(in, out T, f uint64) {
		for {
			out <- f*<-in
		}
	}(in, out, f)
	return in, out
}

// ----------------------------------------------------------------------------

const a = 1
const b, c = 2, 3
const ()
const (
	e = 4
)
const ( f = 5 )
const (
	g    = 6
	h, i = 7, 8
)
const ( j = 1; k, l = 9, 10 )

const ()
const (
	e = 4;
)
const (	f = 5; )
const (
	g    = 6;
	h, i = 7, 8;
)
const ( j = 1; k, l = 9, 10; )

var m = 11
var n, o = 12, 13
var ()
var (
	p = 14
)
var ( q = 15 )
var (
	r    = 16
	s, t = 17, 18
)
var ( u = 19; v, w = 19, 20 )

var ()
var (
	p = 14;
)
var ( q = 15; )
var (
	r    = 16;
	s, t = 17, 18;
)
var ( u = 19; v, w = 19, 20; )

type ()
type (
	t struct{}
	t struct { *a }
	t struct { *b.c }
	t struct { d }
	t struct { e.f }
	t struct { g int }
	t struct { h, i int }

	t struct { *a; *b }
	t struct { *c.d; *e.f }
	t struct { g; h }
	t struct { i.j; k.l }
	t struct { m int; n int }
	//t struct { o, p int; q, r int }
)

//type ()
//type (
//	t struct{}
//	t struct { *a; }
//	t struct { *b.c; }
//	t struct { d; }
//	t struct { e.f; }
//	t struct { g int; }
//	t struct { h, i int; }
//
//	t struct { *a; *b; }
//	t struct { *c.d; *e.f; }
//	t struct { g; h; }
//	t struct { i.j; k.l; }
//	t struct { m int; n int; }
//	t struct { o, p int; q, r int; }
//)
