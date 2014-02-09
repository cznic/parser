// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"go/scanner"
	"go/token"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Printf("dbg %s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
}

func caller(s string, va ...interface{}) {
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Printf("caller: %s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Printf("\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Println()
}

func use(...interface{}) {}

func yyDbg(y yyLexer, v interface{}) {
	caller("yyDbg\n%s", String(yyFset(y), v))
}

func Test0(t *testing.T) {
	fs := token.NewFileSet()
	ast, err := ParseFile(fs, filepath.FromSlash("_testdata/test0/1.go"), nil)
	if err != nil {
		switch x := err.(type) {
		case scanner.ErrorList:
			for _, v := range x {
				t.Error(v)
			}
			return
		default:
			t.Fatal(err)
		}
	}

	if ast == nil {
		t.Fatal(ast)
	}

	t.Logf("\n%s", String(fs, ast))
}
