package main

//type ()
//type (
//	t struct{}
//	t struct { *a }
//	t struct { *a.b }
//	t struct { a }
//	t struct { a.b }
//	t struct { a int }
//	t struct { a, b int }
//
//	t struct { *a; *a }
//	t struct { *a.b; *a.b }
//	//t struct { a; a }
//	t struct { a.b; a.b }
//	t struct { a int; a int }
//	t struct { a, b int; a, b int }
//)

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
const (	j = 1; k, l = 9, 10 )

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

//TODO ParameterDecl 
