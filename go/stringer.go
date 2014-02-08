// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/cznic/strutil"
)

// String pretty prints Nodes or lists of Nodes.
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
				//f.Format("%s: nil\n", pre)
				return
			}

			s(pre, rv.Elem())
		case reflect.Ptr:
			if rv.IsNil() {
				//f.Format("%s: nil\n", pre)
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
			structType := rv.Type()
			f.Format("%s%T{%i\n", pre, v)
			switch structType.Name() {
			default:
				for i := 0; i < rv.NumField(); i++ {
					field := rv.Field(i)
					fieldName := structType.Field(i).Name
					switch fldTypeName := field.Type().Name(); fldTypeName {
					case "pos":
						p := token.Pos(field.Interface().(pos))
						f.Format("%s: %s\n", fldTypeName, fs.Position(p))
					case "Token":
						f.Format("%s: %s\n", fieldName, field.Interface())
					default:
						fldName := structType.Field(i).Name
						if !ast.IsExported(fldName) {
							break
						}

						s(fmt.Sprintf("%s: ", fieldName), rv.Field(i))
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
