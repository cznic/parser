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

// ---------------------------------------------------------------------- BinOp

type BinOp struct {
	pos
	Op   token.Token
	L, R Node
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

// ---------------------------------------------------------------------- Ident

type Ident struct {
	pos
	Lit string
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

// --------------------------------------------------------------------- Literal

type Literal struct {
	pos
	Kind token.Token
	Lit  string
}

func newLiteral(lit tkn) *Literal {
	return &Literal{lit.pos, lit.tok, lit.lit}
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
