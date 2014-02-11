// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"go/token"

	"github.com/cznic/mathutil"
)

type Node interface {
	Pos() token.Pos
}

type pos token.Pos

func (p pos) Pos() token.Pos { return token.Pos(p) }

type Declaration interface {
	DeclName() string
}

type Name Ident

func (n *Name) DeclName() string { return n.Lit }

// ------------------------------------------------------------------ ArrayType

type ArrayType struct {
	pos
	Expr Node
	Type Node
}

// ----------------------------------------------------------------- Assignment
type Assignment struct {
	pos
	Op   token.Token
	L, R []Node
}

// ---------------------------------------------------------------------- BinOp

type BinOp struct {
	pos
	Op   token.Token
	L, R Node
}

// ------------------------------------------------------------------ BreakStmt

type BreakStmt struct {
	pos
	Label *Ident
}

// --------------------------------------------------------------------- CallOp

type CallOp struct {
	pos
	Calee Node
	Args  []Node
}

// -------------------------------------------------------------------- CompLit

type CompLit struct {
	pos
	Type Node
	Val  []*Element
}

// ------------------------------------------------------------ CompoundStament

type CompoundStament struct {
	pos
	Stmts []Node
}

// ------------------------------------------------------------------ ConstDecl

type ConstDecl struct {
	pos
	Iota int
	*Name
	Type Node
	Expr Node
}

func newConstDecls(y yyLexer, lst []Node) (r []Node) {
	for _, v0 := range lst {
		v := v0.(*constSpec)
		if g, e := len(v.Names), len(v.Expr); g != e {
			switch {
			case g > e:
				yyErrPos(y, v.Names[e], "missing value in const declaration")
			default:
				yyErrPos(y, v.Expr[g], "extra expression in const declaration")
			}
			continue
		}

		for j, nm := range v.Names[:mathutil.Min(len(v.Names), len(v.Expr))] {
			id := nm.(*Ident)
			d := &ConstDecl{pos(nm.Pos()), v.Iota, (*Name)(id), v.Type, v.Expr[j]}
			r = append(r, d)
		}
	}
	return
}

// ------------------------------------------------------------------ constSpec
type constSpec struct {
	pos
	Iota  int
	Names []Node
	Type  Node
	Expr  []Node
}

func newConstSpec(y yyLexer, names []Node, typ Node, expr []Node) (c *constSpec) {
	p := yy(y)
	switch {
	case len(expr) == 0:
		c = &constSpec{Iota: p.constIota, Names: names, Type: p.constType, Expr: p.constExpr}
	default:
		c = &constSpec{Iota: p.constIota, Names: names, Type: typ, Expr: expr}
		p.constType, p.constExpr = typ, expr
	}
	p.constIota++
	return
}

// --------------------------------------------------------------- ContinueStmt

type ContinueStmt struct {
	pos
	Label *Ident
}

// -------------------------------------------------------------------- Element

type Element struct {
	pos
	Key Node
	Val Node
}

// ------------------------------------------------------------ FallthroughStmt

type FallthroughStmt struct {
	pos
}

// ---------------------------------------------------------------------- Field

type Field struct {
	pos
	Name     *Ident
	Embedded bool
	Type     Node
	Tag      *Literal
	Scope    *Scope
}

// --------------------------------------------------------------------- fields
type fields struct {
	pos
	Names    []Node
	Embedded bool
	Type     Node
	Tag      *Literal
}

func newFields(l []Node, emb bool, typ, tag Node) *fields {
	return &fields{Names: l, Embedded: emb, Type: typ, Tag: tag.(*Literal)}
}

// ------------------------------------------------------------------- FuncDecl
type FuncDecl struct {
	pos
	*Name
	RxName *Ident
	RxType Node
	Type   *FuncType
	Body   []Node
}

// -------------------------------------------------------------------- FuncLit

type FuncLit struct {
	pos
	Type *FuncType
	Body []Node
}

// ------------------------------------------------------------------- FuncType
type FuncType struct {
	pos
	Ddd     bool // in
	In, Out []*Param
}

func newFuncType(y yyLexer, pos pos, in, out []*Param) (r *FuncType) {
	ps := yy(y)
	r = &FuncType{pos: pos, In: in, Out: out}
	r.Ddd = newParams(ps, pos, in)
	if newParams(ps, pos, out) {
		ps.errPos(token.Pos(pos), "cannot use ... in output argument list")
	}
	return
}

// ------------------------------------------------------------------- GotoStmt

type GotoStmt struct {
	pos
	Label *Ident
}

// ---------------------------------------------------------------------- Ident

type Ident struct {
	pos
	Lit string
}

// --------------------------------------------------------------------- IfStmt

type IfStmt struct {
	pos
	Init Node
	Cond Node
	Body []Node
	Elif []*IfStmt
	Else *CompoundStament
}

// --------------------------------------------------------------------- Import

type Import struct {
	pos
	Name *Ident
	Path *Literal
}

func newImport(y yyLexer, id Node, pth *Literal) (r *Import) {
	ps := yy(y)
	ident := id.(*Ident)
	r = &Import{Name: ident, Path: pth}
	switch {
	case pth.Kind != token.STRING:
		ps.errPos(pth.Pos(), "import statement not a string")
	case pth.Lit == `""`:
		ps.errPos(pth.Pos(), "import path is empty")
	}

	if ident != nil {
		nm := ident.Lit
		if nm != "." {
			ps.pkgScope.declare(ps, nm, id)
		}
	}
	return
}

// ----------------------------------------------------------------- IncDecStmt

type IncDecStmt struct {
	pos
	Expr Node
	Op   token.Token
}

// -------------------------------------------------------------------- IndexOp

type IndexOp struct {
	pos
	Expr  Node
	Index Node
}

// -------------------------------------------------------------- InterfaceType
type InterfaceType struct {
	pos
	Methods []*MethodSpec
}

func newInterfaceType(y yyLexer, l []Node) *InterfaceType {
	i := &InterfaceType{Methods: make([]*MethodSpec, len(l))}
	ps := yy(y)
	sc := ps.currentScope.New()
	for j, v := range l {
		m := v.(*MethodSpec)
		i.Methods[j] = m
		sc.declare(ps, m.Name.Lit, m)
	}
	return i
}

// --------------------------------------------------------------------- Literal

type Literal struct {
	pos
	Kind token.Token
	Lit  string
}

func newLiteral(lit tkn) *Literal {
	return &Literal{lit.pos, lit.tok, lit.lit}
}

// ----------------------------------------------------------------- MethodSpec
type MethodSpec struct {
	pos
	Name *Ident
	Type *FuncType
}

// ------------------------------------------------------------------ NamedType
type NamedType struct {
	pos
	Name  *QualifiedIdent
	Type  Node // resolved type
	Scope *Scope
}

// -------------------------------------------------------------------- Package

type Package struct {
	pos
	Name *Ident
}

// ---------------------------------------------------------------------- Param
type Param struct {
	pos
	Name  *Ident
	Ddd   bool
	Type  Node
	Scope *Scope
}

func newParams(y yyLexer, pos pos, a []*Param) (ddd bool) {
	if len(a) == 0 {
		return
	}

	p := []int{}
	for i, v := range a {
		if v.Name != nil {
			p = append(p, i)
		}
	}

	defer func() {
		for i, v := range a {
			ddd = ddd || v.Ddd
			if ddd && i != len(a)-1 {
				yyErrPos(y, v, "can only use ... as final argument in list")
			}
		}
	}()

	if len(p) == 0 { // only types
		return
	}

	for pi, ai := range p { // a[i] is (name, type)
		from := 0
		if pi != 0 {
			from = p[pi-1] + 1
		}
		typ := a[ai].Type
		for i := from; i < ai; i++ {
			id, ok := a[i].Type.(*NamedType)
			if !ok {
				yyErrPos(y, a[i], "mixed named and unnamed function parameters")
				continue
			}

			if id.Name.Q != nil {
				yyErrPos(y, id, "mixed named and unnamed function parameters")
				continue
			}

			a[i].Name, a[i].Type = id.Name.I, typ
		}
	}
	return
}

// ---------------------------------------------------------------------- Paren

type Paren struct {
	pos
	ExprOrType Node
}

// -------------------------------------------------------------------- PtrType

type PtrType struct {
	pos
	Type Node
}

// ------------------------------------------------------------- QualifiedIdent

type QualifiedIdent struct {
	pos
	Q, I *Ident
}

// ----------------------------------------------------------------- ReturnStmt

type ReturnStmt struct {
	pos
	Expr []Node
}

// ------------------------------------------------------------------- SelectOp

type SelectOp struct {
	pos
	Expr     Node
	Selector *Ident
}

// --------------------------------------------------------------- ShortVarDecl

type ShortVarDecl struct {
	pos
	Names []Node
	Expr  []Node
}

// ------------------------------------------------------------------ SliceType

type SliceType struct {
	pos
	Type Node
}

// ----------------------------------------------------------------- StructType

type StructType struct {
	pos
	Fields []*Field
}

func newStructType(y yyLexer, n tkn, l []Node) (r *StructType) {
	ps := yy(y)
	r = &StructType{pos: n.pos}
	cs := ps.currentScope
	ns := cs.New()
	for _, v := range l {
		fields := v.(*fields)
		for _, v := range fields.Names {
			id := v.(*Ident)
			ns.declare(ps, id.Lit, id)
			r.Fields = append(r.Fields, &Field{
				pos:      pos(v.Pos()),
				Name:     id,
				Embedded: fields.Embedded,
				Type:     fields.Type,
				Tag:      fields.Tag,
				Scope:    cs,
			})
		}
	}
	return
}

// ------------------------------------------------------------------- TypeDecl

type TypeDecl struct {
	pos
	*Name
	Type Node
}

// ----------------------------------------------------------------------- UnOp

type UnOp struct {
	pos
	Op token.Token
	R  Node
}

// -------------------------------------------------------------------- VarDecl

type VarDecl struct {
	pos
	*Name
	Type Node
	Expr Node
}

func newVarDecls(ids []Node, t Node, expr []Node) (r []Node) {
	for _, v := range ids {
		id := v.(*Ident)
		d := &VarDecl{pos(id.Pos()), (*Name)(id), t, nil}
		r = append(r, d)
	}
	for i, v := range expr[:mathutil.Min(len(ids), len(expr))] {
		r[i].(*VarDecl).Expr = v
	}
	return
}
