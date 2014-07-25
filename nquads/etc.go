// Copyright (c) 2014 Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scanner implements a scanner for N-Quads[0] source text.
package parser

import (
	"github.com/cznic/scanner/nquads"
)

type lexer struct {
	*scanner.Scanner
	ast []*Statement
}

type Statement struct{}

// Parse parses src as a single N-Quads source file fname and returns the
// corresponding AST. If the source couldn't be parsed, the returned AST is
// nil and the error indicates all of the specific failures.
func Parse(fname string, src []byte) ([]Statement, error) {
	l := lexer{
		Scanner: scanner.New(fname, src),
	}
	//l.Fname = fname
	//defer func() {
	//	if e := recover(); e != nil {
	//		l.Error(fmt.Sprintf("%v", e))
	//		err = errList(l.Errors)
	//		return
	//	}
	//}()
	//if yyParse(&l) != 0 {
	//	return nil, errList(l.Errors)
	//}

	//return l.spec, nil
	panic(TODO(l))
}
