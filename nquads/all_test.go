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

// Tests from http://www.w3.org/2013/N-QuadsTests/
func Test(t *testing.T) {
	for i, test := range testSuite {
		_, err := Parse("", []byte(test.src))
		if g, e := err != nil, test.emsg != ""; g != e {
			t.Errorf("%d: err != nil: %v, test.emsg != ``: %v", i, g, e)
			continue
		}

		if err != nil {
			if g, e := err.Error(), test.emsg; g != e {
				t.Errorf("%d:\n%s\n%q vs %q", i, test.src, g, e)
			}
		}
	}
}

// http://www.w3.org/2013/N-QuadsTests/manifest.ttl
var testSuite = []struct {
	src  string
	emsg string
}{

	// # N-Quads Syntax tests
	//
	// @prefix rdfs:    <http://www.w3.org/2000/01/rdf-schema#> .
	// @prefix mf: <http://www.w3.org/2001/sw/DataAccess/tests/test-manifest#> .
	// @prefix qt:     <http://www.w3.org/2001/sw/DataAccess/tests/test-query#> .
	//
	// @prefix rdft:   <http://www.w3.org/ns/rdftest#> .
	//
	// <>  a mf:Manifest ;
	//     mf:name "N-Quads tests" ;
	//     mf:entries
	//     (
	//     <#nq-syntax-uri-01>
	//     <#nq-syntax-uri-02>
	//     <#nq-syntax-uri-03>
	//     <#nq-syntax-uri-04>
	//     <#nq-syntax-uri-05>
	//     <#nq-syntax-uri-06>
	//     <#nq-syntax-bnode-01>
	//     <#nq-syntax-bnode-02>
	//     <#nq-syntax-bnode-03>
	//     <#nq-syntax-bnode-04>
	//     <#nq-syntax-bnode-05>
	//     <#nq-syntax-bnode-06>
	//     <#nq-syntax-bad-literal-01>
	//     <#nq-syntax-bad-literal-02>
	//     <#nq-syntax-bad-literal-03>
	//     <#nq-syntax-bad-uri-01>
	//     <#nq-syntax-bad-quint-01>
	//     <#nt-syntax-file-01>
	//     <#nt-syntax-file-02>
	//     <#nt-syntax-file-03>
	//     <#nt-syntax-uri-01>
	//     <#nt-syntax-uri-02>
	//     <#nt-syntax-uri-03>
	//     <#nt-syntax-uri-04>
	//     <#nt-syntax-string-01>
	//     <#nt-syntax-string-02>
	//     <#nt-syntax-string-03>
	//     <#nt-syntax-str-esc-01>
	//     <#nt-syntax-str-esc-02>
	//     <#nt-syntax-str-esc-03>
	//     <#nt-syntax-bnode-01>
	//     <#nt-syntax-bnode-02>
	//     <#nt-syntax-bnode-03>
	//     <#nt-syntax-datatypes-01>
	//     <#nt-syntax-datatypes-02>
	//     <#nt-syntax-bad-uri-01>
	//     <#nt-syntax-bad-uri-02>
	//     <#nt-syntax-bad-uri-03>
	//     <#nt-syntax-bad-uri-04>
	//     <#nt-syntax-bad-uri-05>
	//     <#nt-syntax-bad-uri-06>
	//     <#nt-syntax-bad-uri-07>
	//     <#nt-syntax-bad-uri-08>
	//     <#nt-syntax-bad-uri-09>
	//     <#nt-syntax-bad-prefix-01>
	//     <#nt-syntax-bad-base-01>
	//     <#nt-syntax-bad-struct-01>
	//     <#nt-syntax-bad-struct-02>
	//     <#nt-syntax-bad-lang-01>
	//     <#nt-syntax-bad-esc-01>
	//     <#nt-syntax-bad-esc-02>
	//     <#nt-syntax-bad-esc-03>
	//     <#nt-syntax-bad-string-01>
	//     <#nt-syntax-bad-string-02>
	//     <#nt-syntax-bad-string-03>
	//     <#nt-syntax-bad-string-04>
	//     <#nt-syntax-bad-string-05>
	//     <#nt-syntax-bad-string-06>
	//     <#nt-syntax-bad-string-07>
	//     <#nt-syntax-bad-num-01>
	//     <#nt-syntax-bad-num-02>
	//     <#nt-syntax-bad-num-03>
	//     <#nt-syntax-subm-01>
	//     <#comment_following_triple>
	//     <#literal>
	//     <#literal_all_controls>
	//     <#literal_all_punctuation>
	//     <#literal_ascii_boundaries>
	//     <#literal_with_2_dquotes>
	//     <#literal_with_2_squotes>
	//     <#literal_with_BACKSPACE>
	//     <#literal_with_CARRIAGE_RETURN>
	//     <#literal_with_CHARACTER_TABULATION>
	//     <#literal_with_dquote>
	//     <#literal_with_FORM_FEED>
	//     <#literal_with_LINE_FEED>
	//     <#literal_with_numeric_escape4>
	//     <#literal_with_numeric_escape8>
	//     <#literal_with_REVERSE_SOLIDUS>
	//     <#literal_with_REVERSE_SOLIDUS2>
	//     <#literal_with_squote>
	//     <#literal_with_UTF8_boundaries>
	//     <#langtagged_string>
	//     <#lantag_with_subtag>
	//     <#minimal_whitespace>
	//     ) .
	//
	// <#nq-syntax-uri-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-uri-01" ;
	//    rdfs:comment "URI graph with URI triple" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-uri-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o> <http://example/g> .`, ""}, // 0

	//
	// <#nq-syntax-uri-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-uri-02" ;
	//    rdfs:comment "URI graph with BNode subject" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-uri-02.nq> ;
	//    .

	{`_:s <http://example/p> <http://example/o> <http://example/g> .`, ""}, // 1

	//
	// <#nq-syntax-uri-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-uri-03" ;
	//    rdfs:comment "URI graph with BNode object" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-uri-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> _:o <http://example/g> .`, ""}, // 2

	//
	// <#nq-syntax-uri-04> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-uri-04" ;
	//    rdfs:comment "URI graph with simple literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-uri-04.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "o" <http://example/g> .`, ""}, // 3

	//
	// <#nq-syntax-uri-05> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-uri-05" ;
	//    rdfs:comment "URI graph with language tagged literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-uri-05.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "o"@en <http://example/g> .`, ""}, // 4

	//
	// <#nq-syntax-uri-06> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-uri-06" ;
	//    rdfs:comment "URI graph with datatyped literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-uri-06.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "o"^^<http://www.w3.org/2001/XMLSchema#string> <http://example/g> .`, ""}, // 5

	//
	// <#nq-syntax-bnode-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-bnode-01" ;
	//    rdfs:comment "BNode graph with URI triple" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bnode-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o> _:g .`, ""}, // 6

	//
	// <#nq-syntax-bnode-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-bnode-02" ;
	//    rdfs:comment "BNode graph with BNode subject" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bnode-02.nq> ;
	//    .

	{`_:s <http://example/p> <http://example/o> _:g .`, ""}, // 7

	//
	// <#nq-syntax-bnode-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-bnode-03" ;
	//    rdfs:comment "BNode graph with BNode object" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bnode-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> _:o _:g .`, ""}, // 8

	//
	// <#nq-syntax-bnode-04> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-bnode-04" ;
	//    rdfs:comment "BNode graph with simple literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bnode-04.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "o" _:g .`, ""}, // 9

	//
	// <#nq-syntax-bnode-05> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-bnode-05" ;
	//    rdfs:comment "BNode graph with language tagged literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bnode-05.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "o"@en _:g .`, ""}, // 10

	//
	// <#nq-syntax-bnode-06> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nq-syntax-bnode-06" ;
	//    rdfs:comment "BNode graph with datatyped literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bnode-06.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "o"^^<http://www.w3.org/2001/XMLSchema#string> _:g .`, ""}, // 11

	//
	// <#nq-syntax-bad-literal-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nq-syntax-bad-literal-01" ;
	//    rdfs:comment "Graph name may not be a simple literal (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bad-literal-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o> "o" .`,
		":1:58 graph name may not be a simple literal"}, // 12

	//
	// <#nq-syntax-bad-literal-02> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nq-syntax-bad-literal-02" ;
	//    rdfs:comment "Graph name may not be a language tagged literal (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bad-literal-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o> "o"@en .`,
		":1:58 graph name may not be a language tagged literal"}, // 13

	//
	// <#nq-syntax-bad-literal-03> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nq-syntax-bad-literal-03" ;
	//    rdfs:comment "Graph name may not be a datatyped literal (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bad-literal-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o> "o"^^<http://www.w3.org/2001/XMLSchema#string> .`,
		":1:58 graph name may not be a datatyped literal"}, // 14

	//
	// <#nq-syntax-bad-uri-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nq-syntax-bad-uri-01" ;
	//    rdfs:comment "Graph name URI must be absolute (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bad-uri-01.nq> ;
	//    .

	{`# No relative IRIs in N-Quads
<http://example/s> <http://example/p> <http://example/o> <g>.`,
		":2:58 graph name URI must be absolute"}, // 15
	//
	// <#nq-syntax-bad-quint-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nq-syntax-bad-quint-01" ;
	//    rdfs:comment "N-Quads does not have a fifth element (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nq-syntax-bad-quint-01.nq> ;
	//    .

	{`# N-Quads rejects a quint
<http://example/s> <http://example/p> <http://example/o> <http://example/g> <http://example/n> .`,
		":2:77 n-quads does not have a fifth element"}, // 16

	//
	// <#nt-syntax-file-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-file-01" ;
	//    rdfs:comment "Empty file" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-file-01.nq> ;
	//    .

	{``, ""}, // 17

	//
	// <#nt-syntax-file-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-file-02" ;
	//    rdfs:comment "Only comment" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-file-02.nq> ;
	//    .

	{`#Empty file.`, ""}, // 18

	//
	// <#nt-syntax-file-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-file-03" ;
	//    rdfs:comment "One comment, one empty line" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-file-03.nq> ;
	//    .

	{`#One comment, one empty line.
`, ""}, // 19

	//
	// <#nt-syntax-uri-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-uri-01" ;
	//    rdfs:comment "Only IRIs" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-uri-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o> .`, ""}, // 20

	//
	// <#nt-syntax-uri-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-uri-02" ;
	//    rdfs:comment "IRIs with Unicode escape" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-uri-02.nq> ;
	//    .

	{`# x53 is capital S
<http://example/\u0053> <http://example/p> <http://example/o> .`, ""}, // 21

	//
	// <#nt-syntax-uri-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-uri-03" ;
	//    rdfs:comment "IRIs with long Unicode escape" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-uri-03.nq> ;
	//    .

	{`# x53 is capital S
<http://example/\U00000053> <http://example/p> <http://example/o> .`, ""}, // 22

	//
	// <#nt-syntax-uri-04> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-uri-04" ;
	//    rdfs:comment "Legal IRIs" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-uri-04.nq> ;
	//    .

	{`# IRI with all chars in it.
<http://example/s> <http://example/p> <scheme:!$%25&'()*+,-./0123456789:/@ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~?#> .`, ""}, // 23

	//
	// <#nt-syntax-string-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-string-01" ;
	//    rdfs:comment "string literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-string-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "string" .`, ""}, // 24

	//
	// <#nt-syntax-string-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-string-02" ;
	//    rdfs:comment "langString literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-string-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "string"@en .`, ""}, // 25

	//
	// <#nt-syntax-string-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-string-03" ;
	//    rdfs:comment "langString literal with region" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-string-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "string"@en-uk .`, ""}, // 26

	//
	// <#nt-syntax-str-esc-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-str-esc-01" ;
	//    rdfs:comment "string literal with escaped newline" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-str-esc-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "a\n" .`, ""}, // 27

	//
	// <#nt-syntax-str-esc-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-str-esc-02" ;
	//    rdfs:comment "string literal with Unicode escape" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-str-esc-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "a\u0020b" .`, ""}, // 28

	//
	// <#nt-syntax-str-esc-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-str-esc-03" ;
	//    rdfs:comment "string literal with long Unicode escape" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-str-esc-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "a\U00000020b" .`, ""}, // 29

	//
	// <#nt-syntax-bnode-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-bnode-01" ;
	//    rdfs:comment "bnode subject" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bnode-01.nq> ;
	//    .

	{`_:a  <http://example/p> <http://example/o> .`, ""}, // 30

	//
	// <#nt-syntax-bnode-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-bnode-02" ;
	//    rdfs:comment "bnode object" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bnode-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> _:a .
_:a  <http://example/p> <http://example/o> .`, ""}, // 31

	//
	// <#nt-syntax-bnode-03> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-bnode-03" ;
	//    rdfs:comment "Blank node labels may start with a digit" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bnode-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> _:1a .
_:1a  <http://example/p> <http://example/o> .`, ""}, // 32

	//
	// <#nt-syntax-datatypes-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-datatypes-01" ;
	//    rdfs:comment "xsd:byte literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-datatypes-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "123"^^<http://www.w3.org/2001/XMLSchema#byte> .`, ""}, // 33

	//
	// <#nt-syntax-datatypes-02> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-datatypes-02" ;
	//    rdfs:comment "integer as xsd:string" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-datatypes-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "123"^^<http://www.w3.org/2001/XMLSchema#string> .`, ""}, // 34

	//
	// <#nt-syntax-bad-uri-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-01" ;
	//    rdfs:comment "Bad IRI : space (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-01.nq> ;
	//    .

	{`# Bad IRI : space.
<http://example/ space> <http://example/p> <http://example/o> .`,
		":2:17 lexical grammar error"}, // 35

	//
	// <#nt-syntax-bad-uri-02> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-02" ;
	//    rdfs:comment "Bad IRI : bad escape (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-02.nq> ;
	//    .

	{`# Bad IRI : bad escape
<http://example/\u00ZZ11> <http://example/p> <http://example/o> .`,
		":2:21 lexical grammar error"}, // 36

	//
	// <#nt-syntax-bad-uri-03> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-03" ;
	//    rdfs:comment "Bad IRI : bad long escape (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-03.nq> ;
	//    .

	{`# Bad IRI : bad escape
<http://example/\U00ZZ1111> <http://example/p> <http://example/o> .`,
		":2:21 lexical grammar error"}, // 37

	//
	// <#nt-syntax-bad-uri-04> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-04" ;
	//    rdfs:comment "Bad IRI : character escapes not allowed (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-04.nq> ;
	//    .

	{`# Bad IRI : character escapes not allowed.
<http://example/\n> <http://example/p> <http://example/o> .`,
		":2:18 lexical grammar error"}, // 38

	//
	// <#nt-syntax-bad-uri-05> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-05" ;
	//    rdfs:comment "Bad IRI : character escapes not allowed (2) (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-05.nq> ;
	//    .

	{`# Bad IRI : character escapes not allowed.
<http://example/\/> <http://example/p> <http://example/o> .`,
		":2:18 lexical grammar error"}, // 39

	//
	// <#nt-syntax-bad-uri-06> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-06" ;
	//    rdfs:comment "Bad IRI : relative IRI not allowed in subject (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-06.nq> ;
	//    .

	{`# No relative IRIs in N-Triples
<s> <http://example/p> <http://example/o> .`,
		":2:1 bad IRI : relative IRI not allowed in subject"}, // 40

	//
	// <#nt-syntax-bad-uri-07> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-07" ;
	//    rdfs:comment "Bad IRI : relative IRI not allowed in predicate (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-07.nq> ;
	//    .

	{`# No relative IRIs in N-Triples
<http://example/s> <p> <http://example/o> .`,
		":2:20 bad IRI : relative IRI not allowed in predicate"}, // 41

	//
	// <#nt-syntax-bad-uri-08> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-08" ;
	//    rdfs:comment "Bad IRI : relative IRI not allowed in object (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-08.nq> ;
	//    .

	{`# No relative IRIs in N-Triples
<http://example/s> <http://example/p> <o> .`,
		":2:39 bad IRI : relative IRI not allowed in object"}, // 42

	//
	// <#nt-syntax-bad-uri-09> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-uri-09" ;
	//    rdfs:comment "Bad IRI : relative IRI not allowed in datatype (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-uri-09.nq> ;
	//    .

	{`# No relative IRIs in N-Triples
<http://example/s> <http://example/p> "foo"^^<dt> .`,
		":2:46 bad IRI : relative IRI not allowed in datatype"}, // 43

	//
	// <#nt-syntax-bad-prefix-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-prefix-01" ;
	//    rdfs:comment "@prefix not allowed in n-triples (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-prefix-01.nq> ;
	//    .

	{`@prefix : <http://example/> .`, ":1:1 syntax error"}, // 44

	//
	// <#nt-syntax-bad-base-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-base-01" ;
	//    rdfs:comment "@base not allowed in N-Triples (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-base-01.nq> ;
	//    .

	{`@base <http://example/> .`, ":1:1 syntax error"}, // 45

	//
	// <#nt-syntax-bad-struct-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-struct-01" ;
	//    rdfs:comment "N-Triples does not have objectList (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-struct-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o>, <http://example/o2> .`,
		":1:57 lexical grammar error"}, // 46

	//
	// <#nt-syntax-bad-struct-02> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-struct-02" ;
	//    rdfs:comment "N-Triples does not have predicateObjectList (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-struct-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> <http://example/o>; <http://example/p2>, <http://example/o2> .`,
		":1:57 lexical grammar error"}, // 47

	//
	// <#nt-syntax-bad-lang-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-lang-01" ;
	//    rdfs:comment "langString with bad lang (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-lang-01.nq> ;
	//    .

	{`# Bad lang tag
<http://example/s> <http://example/p> "string"@1 .`,
		":2:48 lexical grammar error"}, // 48
	//
	// <#nt-syntax-bad-esc-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-esc-01" ;
	//    rdfs:comment "Bad string escape (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-esc-01.nq> ;
	//    .

	{`# Bad string escape
<http://example/s> <http://example/p> "a\zb" .`,
		":2:42 lexical grammar error"}, // 49
	//
	// <#nt-syntax-bad-esc-02> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-esc-02" ;
	//    rdfs:comment "Bad string escape (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-esc-02.nq> ;
	//    .

	{`# Bad string escape
<http://example/s> <http://example/p> "\uWXYZ" .`,
		":2:42 lexical grammar error"}, // 50

	//
	// <#nt-syntax-bad-esc-03> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-esc-03" ;
	//    rdfs:comment "Bad string escape (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-esc-03.nq> ;
	//    .

	{`# Bad string escape
<http://example/s> <http://example/p> "\U0000WXYZ" .`,
		":2:46 lexical grammar error"}, // 51

	//
	// <#nt-syntax-bad-string-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-01" ;
	//    rdfs:comment "mismatching string literal open/close (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "abc' .`,
		":1:44 lexical grammar error"}, // 52

	//
	// <#nt-syntax-bad-string-02> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-02" ;
	//    rdfs:comment "mismatching string literal open/close (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> 1.0 .`,
		":1:39 lexical grammar error"}, // 53

	//
	// <#nt-syntax-bad-string-03> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-03" ;
	//    rdfs:comment "single quotes (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-03.nq> ;
	//    .

	{`<http://example/s> <http://example/p> 1.0e1 .`,
		":1:39 lexical grammar error"}, // 54

	//
	// <#nt-syntax-bad-string-04> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-04" ;
	//    rdfs:comment "long single string literal (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-04.nq> ;
	//    .

	{`<http://example/s> <http://example/p> '''abc''' .`,
		":1:39 lexical grammar error"}, // 55

	//
	// <#nt-syntax-bad-string-05> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-05" ;
	//    rdfs:comment "long double string literal (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-05.nq> ;
	//    .

	{`<http://example/s> <http://example/p> """abc""" .`,
		":1:40 lexical grammar error"}, // 56

	//
	// <#nt-syntax-bad-string-06> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-06" ;
	//    rdfs:comment "string literal with no end (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-06.nq> ;
	//    .

	{`<http://example/s> <http://example/p> "abc .`,
		":1:43 lexical grammar error"}, // 57

	//
	// <#nt-syntax-bad-string-07> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-string-07" ;
	//    rdfs:comment "string literal with no start (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-string-07.nq> ;
	//    .

	{`<http://example/s> <http://example/p> abc" .`,
		":1:39 lexical grammar error"}, // 58

	//
	// <#nt-syntax-bad-num-01> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-num-01" ;
	//    rdfs:comment "no numbers in N-Triples (integer) (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-num-01.nq> ;
	//    .

	{`<http://example/s> <http://example/p> 1 .`,
		":1:39 lexical grammar error"}, // 59

	//
	// <#nt-syntax-bad-num-02> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-num-02" ;
	//    rdfs:comment "no numbers in N-Triples (decimal) (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-num-02.nq> ;
	//    .

	{`<http://example/s> <http://example/p> 1.0 .`,
		":1:39 lexical grammar error"}, // 60

	//
	// <#nt-syntax-bad-num-03> a rdft:TestNQuadsNegativeSyntax ;
	//    mf:name    "nt-syntax-bad-num-03" ;
	//    rdfs:comment "no numbers in N-Triples (float) (negative test)" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-bad-num-03.nq> ;
	//    .
	//
	// <#nt-syntax-subm-01> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "nt-syntax-subm-01" ;
	//    rdfs:comment "Submission test from Original RDF Test Cases" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <nt-syntax-subm-01.nq> ;
	//    .
	//
	// <#comment_following_triple> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "comment_following_triple" ;
	//    rdfs:comment "Tests comments after a triple" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <comment_following_triple.nq> ;
	//    .
	//
	// <#literal_ascii_boundaries> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_ascii_boundaries" ;
	//    rdfs:comment "literal_ascii_boundaries '\\x00\\x26\\x28...'" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_ascii_boundaries.nq> ;
	//    .
	//
	// <#literal_with_UTF8_boundaries> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_UTF8_boundaries" ;
	//    rdfs:comment "literal_with_UTF8_boundaries '\\x80\\x7ff\\x800\\xfff...'" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_UTF8_boundaries.nq> ;
	//    .
	//
	// <#literal_all_controls> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_all_controls" ;
	//    rdfs:comment "literal_all_controls '\\x00\\x01\\x02\\x03\\x04...'" ;
	//    rdft:approval rdft:Approved ;
	//    rdft:approval rdft:Approved ;
	//    mf:action   <literal_all_controls.nq> ;
	//    .
	//
	// <#literal_all_punctuation> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_all_punctuation" ;
	//    rdfs:comment "literal_all_punctuation '!\"#$%&()...'" ;
	//    rdft:approval rdft:Approved ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_all_punctuation.nq> ;
	//    .
	//
	// <#literal_with_squote> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_squote" ;
	//    rdfs:comment "literal with squote \"x'y\"" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_squote.nq> ;
	//    .
	//
	// <#literal_with_2_squotes> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_2_squotes" ;
	//    rdfs:comment "literal with 2 squotes \"x''y\"" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_2_squotes.nq> ;
	//    .
	//
	// <#literal> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal" ;
	//    rdfs:comment "literal \"\"\"x\"\"\"" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal.nq> ;
	//    .
	//
	// <#literal_with_dquote> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_dquote" ;
	//    rdfs:comment 'literal with dquote "x\"y"' ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_dquote.nq> ;
	//    .
	//
	// <#literal_with_2_dquotes> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_2_dquotes" ;
	//    rdfs:comment "literal with 2 squotes \"\"\"a\"\"b\"\"\"" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_2_dquotes.nq> ;
	//    .
	//
	// <#literal_with_REVERSE_SOLIDUS2> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name    "literal_with_REVERSE_SOLIDUS2" ;
	//    rdfs:comment "REVERSE SOLIDUS at end of literal" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_REVERSE_SOLIDUS2.nq> ;
	//    .
	//
	// <#literal_with_CHARACTER_TABULATION> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_CHARACTER_TABULATION" ;
	//    rdfs:comment "literal with CHARACTER TABULATION" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_CHARACTER_TABULATION.nq> ;
	//    .
	//
	// <#literal_with_BACKSPACE> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_BACKSPACE" ;
	//    rdfs:comment "literal with BACKSPACE" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_BACKSPACE.nq> ;
	//    .
	//
	// <#literal_with_LINE_FEED> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_LINE_FEED" ;
	//    rdfs:comment "literal with LINE FEED" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_LINE_FEED.nq> ;
	//    .
	//
	// <#literal_with_CARRIAGE_RETURN> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_CARRIAGE_RETURN" ;
	//    rdfs:comment "literal with CARRIAGE RETURN" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_CARRIAGE_RETURN.nq> ;
	//    .
	//
	// <#literal_with_FORM_FEED> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_FORM_FEED" ;
	//    rdfs:comment "literal with FORM FEED" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_FORM_FEED.nq> ;
	//    .
	//
	// <#literal_with_REVERSE_SOLIDUS> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_REVERSE_SOLIDUS" ;
	//    rdfs:comment "literal with REVERSE SOLIDUS" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_REVERSE_SOLIDUS.nq> ;
	//    .
	//
	// <#literal_with_numeric_escape4> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_numeric_escape4" ;
	//    rdfs:comment "literal with numeric escape4 \\u" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_numeric_escape4.nq> ;
	//    .
	//
	// <#literal_with_numeric_escape8> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "literal_with_numeric_escape8" ;
	//    rdfs:comment "literal with numeric escape8 \\U" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <literal_with_numeric_escape8.nq> ;
	//    .
	//
	// <#langtagged_string> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "langtagged_string" ;
	//    rdfs:comment "langtagged string \"x\"@en" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <langtagged_string.nq> ;
	//    .
	//
	// <#lantag_with_subtag> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "lantag_with_subtag" ;
	//    rdfs:comment "lantag with subtag \"x\"@en-us" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <lantag_with_subtag.nq> ;
	//    .
	//
	// <#minimal_whitespace> a rdft:TestNQuadsPositiveSyntax ;
	//    mf:name      "minimal_whitespace" ;
	//    rdfs:comment "tests absense of whitespace between subject, predicate, object and end-of-statement" ;
	//    rdft:approval rdft:Approved ;
	//    mf:action    <minimal_whitespace.nq> ;
	//    .
}
