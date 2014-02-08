// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"bytes"
	"fmt"
	"go/token"
	"reflect"

	"github.com/cznic/strutil"
)

func String(fs *token.FileSet, v interface{}) (r string) {
	var b bytes.Buffer
	f := strutil.IndentFormatter(&b, "·  ")

	defer func() {
		if err := recover(); err != nil {
			f.Format("\n---- [recovered error] ---- %q\n", err)
			r = b.String()
		}
	}()

	visited := map[uintptr]bool{}
	var s func(string, reflect.Value)
	s = func(pre string, rv reflect.Value) {
		v := rv.Interface()
		switch rv.Kind() {
		case reflect.Interface:
			if rv.IsNil() {
				f.Format("%s: nil\n", pre)
				return
			}

			s(pre, rv.Elem())
		case reflect.Ptr:
			if rv.IsNil() {
				f.Format("%s: nil\n", pre)
				return
			}

			p := rv.Pointer()
			visited[p] = true
			s(fmt.Sprintf("%s*", pre), rv.Elem())
		case reflect.Slice:
			n := rv.Len()
			f.Format("%s%T{ // len: %d%i\n", pre, v, n)
			for i := 0; i < n; i++ {
				s(fmt.Sprintf("[%d]: ", i), rv.Index(i))
			}
			f.Format("%u}\n")
		case reflect.String:
			f.Format("%s%q\n", pre, v.(string))
		case reflect.Struct:
			t := rv.Type()
			f.Format("%s%T{%i\n", pre, v)
			switch t.Name() {
			default:
				for i := 0; i < rv.NumField(); i++ {
					fld := rv.Field(i)
					fldt := fld.Type()
					switch tnm := fldt.Name(); tnm {
					case "pos":
						p := token.Pos(fld.Interface().(pos))
						f.Format("%s: %s\n", tnm, fs.Position(p))
					default:
						s(fmt.Sprintf("%s: ", t.Field(i).Name), rv.Field(i))
					}
				}
			}
			f.Format("%u}\n")
		default:
			panic(fmt.Sprintf("internal error %s %T(%v)", rv.Kind(), v, v))
		}
	}

	s("", reflect.ValueOf(v))
	return b.String()
}
