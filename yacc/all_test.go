// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func dbg(s string, va ...interface{}) {
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Printf("%s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
}

func test0(t *testing.T, root string) {
	if err := filepath.Walk(root, func(pth string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ok, err := filepath.Match("*.y", filepath.Base(pth))
		if err != nil {
			t.Fatal(err)
		}

		if !ok {
			return nil
		}

		file, err := os.Open(pth)
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()
		src, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatal(err)
		}

		spec, err := Parse(pth, src)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("\n%s", spec)
		return nil

	}); err != nil {
		t.Fatal(err)
	}
}

func Test0(t *testing.T) {
	test0(t, "testdata")
}
