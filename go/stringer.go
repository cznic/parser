// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"bytes"
	"fmt"
	"reflect"
)

func stringer(v interface{}) string {
	b := &bytes.Buffer{}

	var f func(reflect.Value)
	f = func(rv reflect.Value) {
		switch rv.Kind() {
		default:
			panic(fmt.Sprintf("internal error %s", rv.Kind()))
		}
	}

	f(reflect.ValueOf(v))
	return b.String()
}
