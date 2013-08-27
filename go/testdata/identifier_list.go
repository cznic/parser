package main

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
	t struct { o, p int; q, r int }
)

type ()
type (
	t struct{}
	t struct { *a; }
	t struct { *b.c; }
	t struct { d; }
	t struct { e.f; }
	t struct { g int; }
	t struct { h, i int; }

	t struct { *a; *b; }
	t struct { *c.d; *e.f; }
	t struct { g; h; }
	t struct { i.j; k.l; }
	t struct { m int; n int; }
	t struct { o, p int; q, r int; }
)

//TODO ParameterDecl 

var (
	_ = func() {}
	_ = func(int) {}
	_ = func(int, uint) {}
	_ = func(int, uint, float64) {}
)

type (
	_ func()
	_ func(int)
	_ func(int, uint)
	_ func(int, uint, float64)
)

var (
	_104 = func(...int) {}
	_105 = func(...T, U) {}
	//_ = func(int, uint, float64) {}
)

//func a()
//func idtA(T)
//func idtB(...T)
//func idtC(id T)
//func idtD(id ...T)
//func idtE(id, id2 T)
//func idtF(id, id2 ...T)
//func idtG(T1, T2)
//func idtH(T1, T2)
