// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms; the output is an abstract syntax tree (AST)
// representing the Go source.
package parser

import (
	"bytes"
	"go/token"
	"io"
	"io/ioutil"

	"github.com/cznic/scanner/go"
)

// ParseFile parses the source code of a single Go source file and returns the
// corresponding AST. The source code may be provided via the filename of the
// source file, or via the src parameter.
//
// If src != nil, ParseFile parses the source from src and the filename is only
// used when recording position information. The type of the argument for the
// src parameter must be string, []byte, or io.Reader. If src == nil, ParseFile
// parses the file specified by filename.
//
// The mode parameter controls the amount of source text parsed and other
// optional parser functionality.
//
// If the source couldn't be read, the returned AST is nil and the error
// indicates the specific failure.
func ParseFile(filename string, src interface{}) (f interface{}, err error) {
	var bsrc []byte
	switch x := src.(type) {
	case nil:
		if bsrc, err = ioutil.ReadFile(filename); err != nil {
			return
		}
	case string:
		bsrc = []byte(x)
	case []byte:
		bsrc = x
	case io.Reader:
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, x); err != nil {
			return nil, err
		}

		bsrc = buf.Bytes()
	}

	lx := lx{
		Scanner: scanner.New(bsrc),
	}
	lx.Scanner.Fname = filename
	if yyParse(&lx) != 0 {
		err = lx.Errors[0]
	}
	return
}

var xlat = []int{
	//token.ILLEGAL:        ILLEGAL,
	token.EOF: 0,
	//token.COMMENT:        COMMENT,
	token.IDENT:          IDENTIFIER,
	token.INT:            INT_LIT,
	token.FLOAT:          FLOAT_LIT,
	token.IMAG:           IMAGINARY_LIT,
	token.CHAR:           RUNE_LIT,
	token.STRING:         STRING_LIT,
	token.ADD:            '+',
	token.SUB:            '-',
	token.MUL:            '*',
	token.QUO:            '/',
	token.REM:            '%',
	token.AND:            '&',
	token.OR:             '|',
	token.XOR:            '^',
	token.SHL:            LSH,
	token.SHR:            RSH,
	token.AND_NOT:        ANDNOT,
	token.ADD_ASSIGN:     ASSIGN_OP, // ADD_ASSIGN,
	token.SUB_ASSIGN:     ASSIGN_OP, // SUB_ASSIGN,
	token.MUL_ASSIGN:     ASSIGN_OP, // MUL_ASSIGN,
	token.QUO_ASSIGN:     ASSIGN_OP, // QUO_ASSIGN,
	token.REM_ASSIGN:     ASSIGN_OP, // REM_ASSIGN,
	token.AND_ASSIGN:     ASSIGN_OP, // AND_ASSIGN,
	token.OR_ASSIGN:      ASSIGN_OP, // OR_ASSIGN,
	token.XOR_ASSIGN:     ASSIGN_OP, // XOR_ASSIGN,
	token.SHL_ASSIGN:     ASSIGN_OP, // SHL_ASSIGN,
	token.SHR_ASSIGN:     ASSIGN_OP, // SHR_ASSIGN,
	token.AND_NOT_ASSIGN: ASSIGN_OP, // AND_NOT_ASSIGN,
	token.LAND:           ANDAND,
	token.LOR:            OROR,
	token.ARROW:          COMM,
	token.INC:            INC,
	token.DEC:            DEC,
	token.EQL:            EQ,
	token.LSS:            '<',
	token.GTR:            '>',
	token.ASSIGN:         '=',
	token.NOT:            '!',
	token.NEQ:            NE,
	token.LEQ:            LE,
	token.GEQ:            GE,
	token.DEFINE:         int(token.DEFINE),
	token.ELLIPSIS:       DDD,
	token.LPAREN:         '(',
	token.LBRACK:         '[',
	token.LBRACE:         '{',
	token.COMMA:          ',',
	token.PERIOD:         '.',
	token.RPAREN:         ')',
	token.RBRACK:         ']',
	token.RBRACE:         '}',
	token.SEMICOLON:      ';',
	token.COLON:          ':',
	token.BREAK:          BREAK,
	token.CASE:           CASE,
	token.CHAN:           CHAN,
	token.CONST:          CONST,
	token.CONTINUE:       CONTINUE,
	token.DEFAULT:        DEFAULT,
	token.DEFER:          DEFER,
	token.ELSE:           ELSE,
	token.FALLTHROUGH:    FALLTHROUGH,
	token.FOR:            FOR,
	token.FUNC:           FUNC,
	token.GO:             GO,
	token.GOTO:           GOTO,
	token.IF:             IF,
	token.IMPORT:         IMPORT,
	token.INTERFACE:      INTERFACE,
	token.MAP:            MAP,
	token.PACKAGE:        PACKAGE,
	token.RANGE:          RANGE,
	token.RETURN:         RETURN,
	token.SELECT:         SELECT,
	token.STRUCT:         STRUCT,
	token.SWITCH:         SWITCH,
	token.TYPE:           TYPE,
	token.VAR:            VAR,
}

//TODO idlist_colas	= . // identifier { "," identifier } colas .
//TODO identifier_list = . // Manually enabled in proper contexts.
//TODO lbr = .

const (
	st1 = iota
	st2
	st3
	st4
	st5
	st6
	stDump
)

type tok struct {
	tk        int
	val       interface{}
	line, col int
}

type lx struct {
	*scanner.Scanner
	state int
	toks  []tok
}

/*
(13:25) jnml@fsc-r550:~/src/github.com/cznic/parser/go$ cat fsm
const	C
ident	I
colas	S

%%

{const}{ident}(,{ident})*
{ident}(,{ident})*{colas}

(13:26) jnml@fsc-r550:~/src/github.com/cznic/parser/go$ golex -DFA fsm
StartConditions:
	INITIAL, scId:0, stateId:1
DFA:
[1]
	"C"--> 2
	"I"--> 4
[2]
	"I"--> 3
[3]
	","--> 2
[4]
	","--> 5
	"S"--> 6
[5]
	"I"--> 4
[6]
state 3 accepts rule 1
state 6 accepts rule 2

(13:26) jnml@fsc-r550:~/src/github.com/cznic/parser/go$

*/

func (x *lx) Lex(lval *yySymType) (r int) {
	dbg("Lex state st%d", x.state+1)
	defer func() {
		var s string
		if r < 128 {
			s = string(r)
		} else {
			s = yyToknames[r-ANDAND]
		}
		dbg("returning %q", s)
	}()
	var val interface{}
	var tk token.Token
	if x.state != stDump {
		r = -1
	}
	for {
		if r < 0 {
			tk, val = x.ScanSemis()
			if tk == token.COMMENT {
				continue
			}

			dbg("scanned:%d:%d %v", x.Line, x.Col, tk)
			r = xlat[tk]
		}

		switch x.state {
		case st1:
			switch r {
			case CONST,VAR:
				x.state = st2
			case IDENTIFIER:
				x.state = st4
				x.toks = append([]tok(nil), tok{tk: r, val: val, line: x.Line, col: x.Col})
				r = -1
				continue
			}
			return
		case st2:
			switch r {
			case '(':
				return
			case IDENTIFIER:
				x.toks = append([]tok(nil), tok{tk: r, val: val, line: x.Line, col: x.Col})
				x.state = st3
				r = -1
				continue
			default:
				x.toks = append(x.toks, tok{tk: r, val: val, line: x.Line, col: x.Col})
				x.state = stDump
				continue
			}
		case st3:
			switch r {
			case ',':
				r = -1
				x.state = st2
			default:
				x.toks = append([]tok(nil), tok{tk: r, val: val, line: x.Line, col: x.Col})
				r = IDENTIFIER_LIST
				x.state = stDump
				return
			}
		case st4:
			switch r {
			case ',':
				r = -1
				x.state = st2
				continue
			case int(token.DEFINE):
				r = IDLIST_COLAS
				//TODO pack tokens into lval
				x.toks = nil
				x.state = st1
				return
			default:
				x.toks = append(x.toks, tok{tk: r, val: val, line: x.Line, col: x.Col})
				x.state = stDump
				continue
			}
		case stDump:
			t := x.toks[0]
			r = t.tk
			//TODO lval.val = t.val
			x.Line, x.Col = t.line, t.col
			x.toks = x.toks[1:]
			if len(x.toks) == 0 {
				x.toks = nil
				x.state = st1
			}
			return
		default:
			dbg("TODO st%d", x.state+1)
			return
		}
	}
}
