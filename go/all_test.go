// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"path/filepath"
	"testing"
)

func Test0(t *testing.T) {
	return //TODO-
	ast, err := ParseFile(filepath.FromSlash("_testdata/test0/1.go"), nil)
	if err != nil {
		t.Fatal(err)
	}

	if ast == nil {
		t.Fatal(ast)
	}

	t.Log(ast)
}
