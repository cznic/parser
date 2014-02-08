// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"go/token"
)

var (
	_ Node = (*Ident)(nil)
	_ Node = (*Package)(nil)
)

type Node interface {
	Pos() token.Pos
}

type pos token.Pos

func (p pos) Pos() token.Pos { return token.Pos(p) }

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

func newImport(y yyLexer, pos pos, nm *Ident, pth *Literal) *Import {
	switch {
	case pth.Kind != token.STRING:
		yy(y).errPos(pth.pos, "import statement not a string")
	case pth.Lit == `""`:
		yy(y).errPos(pth.pos, "import path is empty")
	}
	return &Import{pos, nm, pth}
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

// -------------------------------------------------------------------- Package

type Package struct {
	pos
	Name *Ident
}
