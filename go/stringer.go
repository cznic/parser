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

	var s func(string, reflect.Value)
	s = func(pre string, rv reflect.Value) {
		v := rv.Interface()
		switch rv.Kind() {
		case reflect.Bool:
			f.Format("%s%v\n", pre, v)
		case reflect.Int:
			f.Format("%s%v\n", pre, v)
		case reflect.Interface:
			if rv.IsNil() {
				//f.Format("%s: nil\n", pre)
				return
			}

			s(pre, rv.Elem())
		case reflect.Map:
			return
			//f.Format("%s%T{%i\n", pre, v)
			//for _, v := range rv.MapKeys() {
			//	s(fmt.Sprintf("%v: ", v), rv.MapIndex(v))
			//}
			//f.Format("%u}\n")
		case reflect.Ptr:
			if rv.IsNil() {
				//f.Format("%s: nil\n", pre)
				return
			}

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
			switch structType.Name() {
			case "Ident":
				id := rv.Interface().(Ident)
				f.Format("%s%T{%s %q}\n", pre, v, fs.Position(id.Pos()), id.Lit)
			case "Literal":
				id := rv.Interface().(Literal)
				f.Format("%s%T{%s %s %q}\n", pre, v, fs.Position(id.Pos()), id.Kind, id.Lit)
			case "QualifiedIdent":
				id := rv.Interface().(QualifiedIdent)
				switch {
				case id.Q == nil:
					f.Format("%s%T{%s %q}\n", pre, v, fs.Position(id.Pos()), id.I.Lit)
				default:
					f.Format("%s%T{%s \"%s.%s\"}\n", pre, v, fs.Position(id.Pos()), id.I.Lit, id.Q.Lit)
				}
			case "Scope":
				return
			default:
				f.Format("%s%T{%i\n", pre, v)
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
						if !ast.IsExported(fieldName) {
							break
						}

						s(fmt.Sprintf("%s: ", fieldName), rv.Field(i))
					}
				}
				f.Format("%u}\n")
			}
		default:
			panic(fmt.Sprintf("internal error %s %T(%v)", rv.Kind(), v, v))
		}
	}

	s("", reflect.ValueOf(v))
	return b.String()
}
