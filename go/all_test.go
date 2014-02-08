// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"go/scanner"
	"go/token"
	"path/filepath"
	"testing"
)

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
