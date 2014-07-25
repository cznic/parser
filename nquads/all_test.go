// Copyright (c) 2014 Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"testing"
)

func caller(s string, va ...interface{}) {
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Printf("caller: %s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Printf("\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Println()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Printf("dbg %s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
}

func use(...interface{}) {}

// ============================================================================

func Test(t *testing.T) {
	_, err := Parse("test", []byte(`
<a> <b> <c> <d> . # comments here
# or on a line by themselves
<e> <f> <g> .
_:0 <h> <i> .
<j> <k> <l> _:1 .
<m> <n> "A" .
<m> <n> "A" ^^ <x> .
<m> <n> "A" @en-US .
<e> <f> _:2 .
`))
	if err != nil {
		t.Fatal(err)
	}
}

func ExampleParse() {
	ast, err := Parse("test", []byte(`# L 1
<a> <b> <c> <d> . # L2 comments here
# L3 or on a line by themselves
<e> <f> <g> .        # L 4
_:0 <h> <i> .        # L 5
<j> <k> <l> _:1 .    # L 6
<m> <n> "A" .        # L 7
<m> <n> "A" ^^ <x> . # L 8
<m> <n> "A" @en-US . # L 9
<e> <f> _:2 .        # L 10
`))
	if err != nil {
		panic(err)
	}

	for i, v := range ast {
		fmt.Printf("%d: %v\n", i, v)
	}
	// Output:
	// 0: stmt@2:1{subj@2:1{IRIRef="a"}, pred@2:5{"b"}, obj@2:9{IRIRef="c"}, graph@2:13{IRIRef="d"}}
	// 1: stmt@4:1{subj@4:1{IRIRef="e"}, pred@4:5{"f"}, obj@4:9{IRIRef="g"}}
	// 2: stmt@5:1{subj@5:1{BlankNodeLabel="0"}, pred@5:5{"h"}, obj@5:9{IRIRef="i"}}
	// 3: stmt@6:1{subj@6:1{IRIRef="j"}, pred@6:5{"k"}, obj@6:9{IRIRef="l"}, graph@6:13{BlankNodeLabel="1"}}
	// 4: stmt@7:1{subj@7:1{IRIRef="m"}, pred@7:5{"n"}, obj@7:9{Literal="A"}}
	// 5: stmt@8:1{subj@8:1{IRIRef="m"}, pred@8:5{"n"}, obj@8:9{Literal="A", IRIRef="x"}}
	// 6: stmt@9:1{subj@9:1{IRIRef="m"}, pred@9:5{"n"}, obj@9:9{Literal="A", LangTag="en-US"}}
	// 7: stmt@10:1{subj@10:1{IRIRef="e"}, pred@10:5{"f"}, obj@10:9{BlankNodeLabel="2"}}
}
