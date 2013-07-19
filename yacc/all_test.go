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

		if spec == nil {
			t.Fatal(spec)
		}

		t.Logf("%s\n%s", pth, spec)
		return nil

	}); err != nil {
		t.Fatal(err)
	}
}

func Test0(t *testing.T) {
	test0(t, ".")
}

func ExampleDef_start() {
	spec, err := Parse("start.y", []byte(`

%start Foo

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Start, Tag: "Foo", Nlist: []*parser.Nmno{
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_union() {
	spec, err := Parse("union.y", []byte(`

%union{
	bar int
	baz struct{a, b int}
}

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Union, Tag: "{\n\tbar int\n\tbaz struct{a, b int}\n}", Nlist: []*parser.Nmno{
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_copy() {
	spec, err := Parse("copy.y", []byte(`

%{

package main

%}

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Copy, Tag: "\n\npackage main\n\n", Nlist: []*parser.Nmno{
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_token() {
	spec, err := Parse("token.y", []byte(`

%token foo
%token bar 1234
%token <typ> qux

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Token, Tag: "", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "foo", Number: -1}
	// . . . }
	// . . }
	// . . *parser.Def{
	// . . . Rword: Token, Tag: "", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "bar", Number: 1234}
	// . . . }
	// . . }
	// . . *parser.Def{
	// . . . Rword: Token, Tag: "typ", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "qux", Number: -1}
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_left() {
	spec, err := Parse("left.y", []byte(`

%left foo '+' '-' 1234 'L'
%left <typ> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Left, Tag: "", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "foo", Number: -1}
	// . . . . *parser.Nmno{Identifier: '+', Number: -1}
	// . . . . *parser.Nmno{Identifier: '-', Number: 1234}
	// . . . . *parser.Nmno{Identifier: 'L', Number: -1}
	// . . . }
	// . . }
	// . . *parser.Def{
	// . . . Rword: Left, Tag: "typ", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: '?', Number: -1}
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_right() {
	spec, err := Parse("right.y", []byte(`

%right foo '+' '-' 1234 'L'
%right <typ> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Right, Tag: "", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "foo", Number: -1}
	// . . . . *parser.Nmno{Identifier: '+', Number: -1}
	// . . . . *parser.Nmno{Identifier: '-', Number: 1234}
	// . . . . *parser.Nmno{Identifier: 'L', Number: -1}
	// . . . }
	// . . }
	// . . *parser.Def{
	// . . . Rword: Right, Tag: "typ", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: '?', Number: -1}
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_nonassoc() {
	spec, err := Parse("right.y", []byte(`

%nonassoc foo '+' '-' 1234 'L'
%nonassoc <typ> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Nonassoc, Tag: "", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "foo", Number: -1}
	// . . . . *parser.Nmno{Identifier: '+', Number: -1}
	// . . . . *parser.Nmno{Identifier: '-', Number: 1234}
	// . . . . *parser.Nmno{Identifier: 'L', Number: -1}
	// . . . }
	// . . }
	// . . *parser.Def{
	// . . . Rword: Nonassoc, Tag: "typ", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: '?', Number: -1}
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleDef_type() {
	spec, err := Parse("right.y", []byte(`

%type <typ> foo '+' '-' 1234 'L'
%type <list> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . . *parser.Def{
	// . . . Rword: Type, Tag: "typ", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: "foo", Number: -1}
	// . . . . *parser.Nmno{Identifier: '+', Number: -1}
	// . . . . *parser.Nmno{Identifier: '-', Number: 1234}
	// . . . . *parser.Nmno{Identifier: 'L', Number: -1}
	// . . . }
	// . . }
	// . . *parser.Def{
	// . . . Rword: Type, Tag: "list", Nlist: []*parser.Nmno{
	// . . . . *parser.Nmno{Identifier: '?', Number: -1}
	// . . . }
	// . . }
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}

func ExampleSpec_tail() {
	spec, err := Parse("right.y", []byte(`

%%

Foo:

%%

	fmt.Println("Hello")

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . }
	// . . }
	// . }
	// . Tail: "\n\n\tfmt.Println(\"Hello\")\n\n"
	// }
}

func ExampleRule() {
	spec, err := Parse("right.y", []byte(`

%%

Foo:
	Bar
|	'1' Baz foo
	{
		$$ = "abc"
	}
|	'1' Baz foo
	{
	$$ = "def"
	}
	'2'
	{
		fmt.Println([]t{2})
	}

Bar:
	/* Empty */
|	Bar IDENT	

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(spec)
	// Output:
	// *parser.Spec{
	// . Defs: []*parser.Def{
	// . }
	// . Rules: []*parser.Rule{
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . . "Bar"
	// . . . }
	// . . }
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . . '1'
	// . . . . "Baz"
	// . . . . "foo"
	// . . . . *parser.Act{Src: "\n\t\t$$ = \"abc\"\n\t"}
	// . . . }
	// . . }
	// . . *parser.Rule{
	// . . . Name: "Foo", Body: []interface {}{
	// . . . . '1'
	// . . . . "Baz"
	// . . . . "foo"
	// . . . . *parser.Act{Src: "\n\t$$ = \"def\"\n\t"}
	// . . . . '2'
	// . . . . *parser.Act{Src: "\n\t\tfmt.Println([]t{2})\n\t"}
	// . . . }
	// . . }
	// . . *parser.Rule{
	// . . . Name: "Bar", Body: []interface {}{
	// . . . }
	// . . }
	// . . *parser.Rule{
	// . . . Name: "Bar", Body: []interface {}{
	// . . . . "Bar"
	// . . . . "IDENT"
	// . . . }
	// . . }
	// . }
	// . Tail: ""
	// }
}
