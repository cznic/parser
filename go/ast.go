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
	Name *Ident
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
			d := &ConstDecl{pos(nm.Pos()), v.Iota, nm.(*Ident), v.Type, v.Expr[j]}
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
}

// --------------------------------------------------------------------- fields
type fields struct {
	pos
	Names    []*Ident
	Embedded bool
	Type     Node
	Tag      *Literal
}

func newFields(l []Node, emb bool, typ, tag Node) *fields {
	return &fields{Names: idList(l), Embedded: emb, Type: typ, Tag: tag.(*Literal)}
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

func newImport(y yyLexer, nm Node, pth *Literal) *Import {
	switch {
	case pth.Kind != token.STRING:
		yyErrPos(y, pth, "import statement not a string")
	case pth.Lit == `""`:
		yyErrPos(y, pth, "import path is empty")
	}
	return &Import{Name: nm.(*Ident), Path: pth}
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
	Name *QualifiedIdent
	Type Node // resolved type
}

// -------------------------------------------------------------------- Package

type Package struct {
	pos
	Name *Ident
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

func newStructType(n tkn, l []Node) (r *StructType) {
	r = &StructType{pos: n.pos}
	for _, v := range l {
		fields := v.(*fields)
		for _, v := range fields.Names {
			r.Fields = append(r.Fields, &Field{
				pos:      v.pos,
				Name:     v,
				Embedded: fields.Embedded,
				Type:     fields.Type,
				Tag:      fields.Tag,
			})
		}
	}
	return
}

// ------------------------------------------------------------------- TypeDecl

type TypeDecl struct {
	pos
	Name *Ident
	Type Node
}

// ----------------------------------------------------------------------- UnOp

type UnOp struct {
	pos
	Op token.Token
	R  Node
}
