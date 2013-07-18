%{

/*

Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.

Original source text: http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html

Modifications: Copyright 2013 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Grammar for the input to yacc.

*/

// (WIP:TODO)
package parser

import (
	"github.com/cznic/scanner/yacc"
)

%}

%union {
	number int
}

/* Basic entries. The following are recognized by the lexical analyzer. */

%token	_IDENTIFIER      /* Includes identifiers and literals */
%token	_C_IDENTIFIER    /* identifier (but not literal)
                             followed by a :. */
%token	_NUMBER          /* [0-9][0-9]* */

/* Reserved words : %type=>_TYPE %left=>_LEFT, and so on */

%token	_LEFT _RIGHT _NONASSOC _TOKEN _PREC _TYPE _START _UNION

%token	_MARK            /* The %% mark. */
%token	_LCURL           /* The %{ mark. */
%token	_RCURL           /* The %} mark. */

%type	<number> _NUMBER

/* 8-bit character literals stand for themselves; */
/* tokens have to be defined for multi-byte characters. */

%start	spec

%%

spec:
	defs _MARK rules tail
	{
		panic(".y:56")
	}

tail:
	/* Empty; the second _MARK is optional. */
	{
		panic(".y:62")
	}
|	_MARK
	{
        	/* In this action, set up the rest of the file. */
		panic(".y:67")
	}

defs:
	/* Empty. */
	{
		panic(".y:73")
	}
	| defs def
	{
		panic(".y:77")
	}

def:
   	_START _IDENTIFIER
	{
		panic(".y:83")
	}
|	_UNION
	{
        	/* Copy union definition to output. */
		panic(".y:88")
	}
|	_LCURL
	{
		/* Copy Go code to output file. */
		panic(".y:93")
	}
	_RCURL
	{
		panic(".y:97")
	}
|	rword tag nlist
	{
		panic(".y:101")
	}

rword:
	_TOKEN
	{
		panic(".y:107")
	}
|	_LEFT
	{
		panic(".y:111")
	}
|	_RIGHT
	{
		panic(".y:115")
	}
|	_NONASSOC
	{
		panic(".y:119")
	}
|	_TYPE
	{
		panic(".y:123")
	}

tag:
	/* Empty: union tag ID optional. */
	{
		panic(".y:129")
	}
|	'<' _IDENTIFIER '>'
	{
		panic(".y:133")
	}

nlist:
	nmno
	{
		panic(".y:139")
	}
|	nlist nmno
	{
		panic(".y:143")
	}

nmno:
	_IDENTIFIER
	{
		/* Note: literal invalid with % type. */
		panic(".y:150")
	}
|	_IDENTIFIER _NUMBER 
	{
		/* Note: invalid with % type. */
		panic(".y:155")
	}

/* Rule section */

rules:
	_C_IDENTIFIER rbody prec
	{
		panic(".y:163")
	}
|	rules rule
	{
		panic(".y:167")
	}

rule:
	_C_IDENTIFIER rbody prec
	{
		panic(".y:173")
	}
|	'|' rbody prec
	{
		panic(".y:177")
	}

rbody:
	/* empty */
	{
		panic(".y:183")
	}
|	rbody _IDENTIFIER
	{
		panic(".y:187")
	}
|	rbody act
	{
		panic(".y:191")
	}

act:
	'{'
	{
		/* Copy action, translate $$, and so on. */
		panic(".y:198")
	}
	'}'

prec:
	/* Empty */
	{
		panic(".y:205")
	}
|	_PREC _IDENTIFIER
	{
		panic(".y:209")
	}
|	_PREC _IDENTIFIER act
	{
		panic(".y:213")
	}
|	prec ';'
	{
		panic(".y:217")
	}

%%

var xlat = map[scanner.Token]int{
	scanner.C_IDENTIFIER: _C_IDENTIFIER,
	scanner.IDENTIFIER:   _IDENTIFIER,
	scanner.INT:          _NUMBER,
	scanner.LCURL:        _LCURL,
	scanner.LEFT:         _LEFT,
	scanner.MARK:         _MARK,
	scanner.NONASSOC:     _NONASSOC,
	scanner.PREC:         _PREC,
	scanner.RCURL:        _RCURL,
	scanner.RIGHT:        _RIGHT,
	scanner.START:        _START,
	scanner.TOKEN:        _TOKEN,
	scanner.TYPE:         _TYPE,
	scanner.UNION:        _UNION,
}

type lexer struct {
	*scanner.Scanner
}

func (l *lexer) Lex(lval *yySymType) (y int) {
	tok, val := l.Scan()
	switch tok {
	case scanner.INT:
		if n, ok := val.(uint64); ok {
			lval.number = int(n)
		}
		return _NUMBER
	default:
		y = xlat[tok]
	}

	panic(".156")
}

//TODO docs
func Parse(fname string, src []byte) (y []interface{}, err error) {
	l := lexer{scanner.New(src)}
	l.Fname = fname
	if yyParse(&l) != 0 {
		return nil, l.Errors[0] //TODO
	}

	return
}
