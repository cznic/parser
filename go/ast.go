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

// -------------------------------------------------------------------- Package
type Package struct {
	pos
	Name *Ident
}
