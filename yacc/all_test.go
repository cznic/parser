// Copyright 2013 The parser Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"go/token"
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

func caller(s string, va ...interface{}) {
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Printf("caller: %s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Printf("\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Println()
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
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

		ast, err := Parse(token.NewFileSet(), pth, src)
		if err != nil {
			t.Fatal(err)
		}

		if ast == nil {
			t.Fatal(ast)
		}

		t.Logf("%s\n%s", pth, ast)
		return nil

	}); err != nil {
		t.Fatal(err)
	}
}

func Test0(t *testing.T) {
	test0(t, ".")
	test0(t, filepath.Join(runtime.GOROOT(), "src"))
}

func ExampleDef_start() {
	ast, err := Parse(token.NewFileSet(), "start.y", []byte(`

%start Foo

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@start.y:3:1{
	// · · · Rword: Start, Tag: "Foo", Nlist: []*parser.Nmno{
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@start.y:7:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_union() {
	ast, err := Parse(token.NewFileSet(), "union.y", []byte(`

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

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@union.y:3:1{
	// · · · Rword: Union, Tag: "{\n        bar int\n        baz struct{a, b int}\n}", Nlist: []*parser.Nmno{
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@union.y:10:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_copy() {
	ast, err := Parse(token.NewFileSet(), "copy.y", []byte(`

%{

package main

%}

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@copy.y:3:1{
	// · · · Rword: Copy, Tag: "\n\npackage main\n\n", Nlist: []*parser.Nmno{
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@copy.y:11:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_token() {
	ast, err := Parse(token.NewFileSet(), "token.y", []byte(`

%token foo
%token bar 1234
%token <typ> qux

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@token.y:3:1{
	// · · · Rword: Token, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@token.y:3:8{Identifier: "foo", Number: -1}
	// · · · }
	// · · }
	// · · *parser.Def@token.y:4:1{
	// · · · Rword: Token, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@token.y:4:8{Identifier: "bar", Number: 1234}
	// · · · }
	// · · }
	// · · *parser.Def@token.y:5:1{
	// · · · Rword: Token, Tag: "typ", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@token.y:5:14{Identifier: "qux", Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@token.y:9:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_left() {
	ast, err := Parse(token.NewFileSet(), "left.y", []byte(`

%left foo '+' '-' 1234 'L'
%left <typ> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@left.y:3:1{
	// · · · Rword: Left, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@left.y:3:7{Identifier: "foo", Number: -1}
	// · · · · *parser.Nmno@left.y:3:11{Identifier: '+', Number: -1}
	// · · · · *parser.Nmno@left.y:3:15{Identifier: '-', Number: 1234}
	// · · · · *parser.Nmno@left.y:3:24{Identifier: 'L', Number: -1}
	// · · · }
	// · · }
	// · · *parser.Def@left.y:4:1{
	// · · · Rword: Left, Tag: "typ", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@left.y:4:13{Identifier: '?', Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@left.y:8:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_right() {
	ast, err := Parse(token.NewFileSet(), "right.y", []byte(`

%right foo '+' '-' 1234 'L'
%right <typ> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@right.y:3:1{
	// · · · Rword: Right, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@right.y:3:8{Identifier: "foo", Number: -1}
	// · · · · *parser.Nmno@right.y:3:12{Identifier: '+', Number: -1}
	// · · · · *parser.Nmno@right.y:3:16{Identifier: '-', Number: 1234}
	// · · · · *parser.Nmno@right.y:3:25{Identifier: 'L', Number: -1}
	// · · · }
	// · · }
	// · · *parser.Def@right.y:4:1{
	// · · · Rword: Right, Tag: "typ", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@right.y:4:14{Identifier: '?', Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@right.y:8:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_nonassoc() {
	ast, err := Parse(token.NewFileSet(), "nonassoc.y", []byte(`

%nonassoc foo '+' '-' 1234 'L'
%nonassoc <typ> '?'

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@nonassoc.y:3:1{
	// · · · Rword: Nonassoc, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@nonassoc.y:3:11{Identifier: "foo", Number: -1}
	// · · · · *parser.Nmno@nonassoc.y:3:15{Identifier: '+', Number: -1}
	// · · · · *parser.Nmno@nonassoc.y:3:19{Identifier: '-', Number: 1234}
	// · · · · *parser.Nmno@nonassoc.y:3:28{Identifier: 'L', Number: -1}
	// · · · }
	// · · }
	// · · *parser.Def@nonassoc.y:4:1{
	// · · · Rword: Nonassoc, Tag: "typ", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@nonassoc.y:4:17{Identifier: '?', Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@nonassoc.y:8:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_errVerbose() {
	ast, err := Parse(token.NewFileSet(), "errVerbose.y", []byte(`

%error-verbose

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@errVerbose.y:3:1{
	// · · · Rword: ErrorVerbose, Tag: "", Nlist: []*parser.Nmno{
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@errVerbose.y:7:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleDef_type() {
	ast, err := Parse(token.NewFileSet(), "type.y", []byte(`

%type	<typ>	foo bar
%type	<list>	baz

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@type.y:3:1{
	// · · · Rword: Type, Tag: "typ", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@type.y:3:13{Identifier: "foo", Number: -1}
	// · · · · *parser.Nmno@type.y:3:17{Identifier: "bar", Number: -1}
	// · · · }
	// · · }
	// · · *parser.Def@type.y:4:1{
	// · · · Rword: Type, Tag: "list", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@type.y:4:14{Identifier: "baz", Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@type.y:8:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleSpec_tail() {
	ast, err := Parse(token.NewFileSet(), "tail.y", []byte(`

%%

Foo:

%%

	fmt.Println("Hello")

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@tail.y:5:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: "\n\n\tfmt.Println(\"Hello\")\n\n"
	// }
}

func ExampleRule() {
	ast, err := Parse(token.NewFileSet(), "rule.y", []byte(`

%%

Foo:
        Bar
|   '1' Baz foo
        {
            $$ = "abc"
        }
|   '2' Qux lol
        {
            $$ = "def"
        }
        '2'
        {
            fmt.Println([]t{2})
        }

Bar:
        /* Empty */
|   Bar IDENT

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@rule.y:5:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · · "Bar"
	// · · · }
	// · · }
	// · · *parser.Rule@rule.y:7:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · · '1'
	// · · · · "Baz"
	// · · · · "foo"
	// · · · · []*parser.Act{
	// · · · · · *parser.Act@rule.y:9:13{
	// · · · · · · Src: "\n            "
	// · · · · · · Tok: DLR_DLR, Tag: "", Num: 0
	// · · · · · }
	// · · · · · *parser.Act@rule.y:9:16{
	// · · · · · · Src: "= \"abc\"\n        "
	// · · · · · }
	// · · · · }
	// · · · }
	// · · }
	// · · *parser.Rule@rule.y:11:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · · '2'
	// · · · · "Qux"
	// · · · · "lol"
	// · · · · []*parser.Act{
	// · · · · · *parser.Act@rule.y:13:13{
	// · · · · · · Src: "\n            "
	// · · · · · · Tok: DLR_DLR, Tag: "", Num: 0
	// · · · · · }
	// · · · · · *parser.Act@rule.y:13:16{
	// · · · · · · Src: "= \"def\"\n        "
	// · · · · · }
	// · · · · }
	// · · · · '2'
	// · · · · []*parser.Act{
	// · · · · · *parser.Act@rule.y:17:13{
	// · · · · · · Src: "\n            fmt.Println([]t{2})\n        "
	// · · · · · }
	// · · · · }
	// · · · }
	// · · }
	// · · *parser.Rule@rule.y:20:1{
	// · · · Name: "Bar", Body: []interface {}{
	// · · · }
	// · · }
	// · · *parser.Rule@rule.y:22:1{
	// · · · Name: "Bar", Body: []interface {}{
	// · · · · "Bar"
	// · · · · "IDENT"
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExampleAct() {
	ast, err := Parse(token.NewFileSet(), "act.y", []byte(`

%%

StatementList:
        /* Empty */
        {
            $$ = nil
        }
|   StatementList Statement
        {
            $$ = append($1, $2)
        }

%%

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@act.y:5:1{
	// · · · Name: "StatementList", Body: []interface {}{
	// · · · · []*parser.Act{
	// · · · · · *parser.Act@act.y:8:13{
	// · · · · · · Src: "\n            "
	// · · · · · · Tok: DLR_DLR, Tag: "", Num: 0
	// · · · · · }
	// · · · · · *parser.Act@act.y:8:16{
	// · · · · · · Src: "= nil\n        "
	// · · · · · }
	// · · · · }
	// · · · }
	// · · }
	// · · *parser.Rule@act.y:10:1{
	// · · · Name: "StatementList", Body: []interface {}{
	// · · · · "StatementList"
	// · · · · "Statement"
	// · · · · []*parser.Act{
	// · · · · · *parser.Act@act.y:12:13{
	// · · · · · · Src: "\n            "
	// · · · · · · Tok: DLR_DLR, Tag: "", Num: 0
	// · · · · · }
	// · · · · · *parser.Act@act.y:12:25{
	// · · · · · · Src: "= append("
	// · · · · · · Tok: DLR_NUM, Tag: "", Num: 1
	// · · · · · }
	// · · · · · *parser.Act@act.y:12:29{
	// · · · · · · Src: ", "
	// · · · · · · Tok: DLR_NUM, Tag: "", Num: 2
	// · · · · · }
	// · · · · · *parser.Act@act.y:12:31{
	// · · · · · · Src: ")\n        "
	// · · · · · }
	// · · · · }
	// · · · }
	// · · }
	// · }
	// · Tail: "\n\n"
	// }
}

func ExampleNmno() {
	ast, err := Parse(token.NewFileSet(), "nmno.y", []byte(`

%token abc '+' def 123 ghi

%%

Foo:

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@nmno.y:3:1{
	// · · · Rword: Token, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@nmno.y:3:8{Identifier: "abc", Number: -1}
	// · · · · *parser.Nmno@nmno.y:3:12{Identifier: '+', Number: -1}
	// · · · · *parser.Nmno@nmno.y:3:16{Identifier: "def", Number: 123}
	// · · · · *parser.Nmno@nmno.y:3:24{Identifier: "ghi", Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@nmno.y:7:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func ExamplePrec() {
	ast, err := Parse(token.NewFileSet(), "prec.y", []byte(`

%%

Foo:
        bar %prec A
|   foo %prec B
        {
            qux($1)
        }

`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@prec.y:5:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · · "bar"
	// · · · }
	// · · · Prec: *parser.Prec@prec.y:6:13{
	// · · · · Identifier: "A"
	// · · · · []*parser.Act{
	// · · · · }
	// · · · }
	// · · }
	// · · *parser.Rule@prec.y:7:1{
	// · · · Name: "Foo", Body: []interface {}{
	// · · · · "foo"
	// · · · }
	// · · · Prec: *parser.Prec@prec.y:7:9{
	// · · · · Identifier: "B"
	// · · · · []*parser.Act{
	// · · · · · *parser.Act@prec.y:9:17{
	// · · · · · · Src: "\n            qux("
	// · · · · · · Tok: DLR_NUM, Tag: "", Num: 1
	// · · · · · }
	// · · · · · *parser.Act@prec.y:9:19{
	// · · · · · · Src: ")\n        "
	// · · · · · }
	// · · · · }
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}

func Example_yacc() {
	ast, err := Parse(token.NewFileSet(), "prec.y", []byte(`
// http://dinosaur.compilertools.net/yacc/
%token  DING  DONG  DELL
%%
rhyme   :       sound  place
        ;
sound   :       DING  DONG
        ;
place   :       DELL
        ;
`))
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
	// Output:
	// *parser.AST{
	// · Defs: []*parser.Def{
	// · · *parser.Def@prec.y:3:1{
	// · · · Rword: Token, Tag: "", Nlist: []*parser.Nmno{
	// · · · · *parser.Nmno@prec.y:3:9{Identifier: "DING", Number: -1}
	// · · · · *parser.Nmno@prec.y:3:15{Identifier: "DONG", Number: -1}
	// · · · · *parser.Nmno@prec.y:3:21{Identifier: "DELL", Number: -1}
	// · · · }
	// · · }
	// · }
	// · Rules: []*parser.Rule{
	// · · *parser.Rule@prec.y:5:1{
	// · · · Name: "rhyme", Body: []interface {}{
	// · · · · "sound"
	// · · · · "place"
	// · · · }
	// · · }
	// · · *parser.Rule@prec.y:7:1{
	// · · · Name: "sound", Body: []interface {}{
	// · · · · "DING"
	// · · · · "DONG"
	// · · · }
	// · · }
	// · · *parser.Rule@prec.y:9:1{
	// · · · Name: "place", Body: []interface {}{
	// · · · · "DELL"
	// · · · }
	// · · }
	// · }
	// · Tail: ""
	// }
}
