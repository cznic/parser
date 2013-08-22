// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms; the output is an abstract syntax tree (AST)
// representing the Go source.
package parser

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/cznic/scanner/go"
)

// ParseFile parses the source code of a single Go source file and returns the
// corresponding AST. The source code may be provided via the filename of the
// source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is only
// used when recording position information. The type of the argument for the
// src parameter must be string, []byte, or io.Reader. If src == nil, ParseFile
// parses the file specified by filename.
//
// The mode parameter controls the amount of source text parsed and other
// optional parser functionality.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure.
func ParseFile(filename string, src interface{}) (f interface{}, err error) {
	var bsrc []byte
	switch x := src.(type) {
	case nil:
		if bsrc, err = ioutil.ReadFile(filename); err != nil {
			return
		}
	case string:
		bsrc = []byte(x)
	case []byte:
		bsrc = x
	case io.Reader:
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, x); err != nil {
			return nil, err
		}

		bsrc = buf.Bytes()
	}

	lx := lx{
		Scanner: scanner.New(bsrc),
	}
	if yyParse(&lx) != 0 {
		err = lx.Errors[0]
	}
	return
}

//TODO idlist_colas	= . // identifier { "," identifier } colas .
//TODO identifier_list = . // Manually enabled in proper contexts.
//TODO lbr = .

type lx struct {
	*scanner.Scanner
}

func (x *lx) Lex(lval *yySymType) int {
	return 0
}
