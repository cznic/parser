// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms; the output is an abstract syntax tree (AST)
// representing the Go source.
package parser

import (
	"bytes"
	"fmt"
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
	if yyParse(&lx) != 0 || len(lx.Errors) != 0 {
		err = lx.Errors[0]
		//dbg("%v", err)
	}
	return
}

const (
	COLAS = -1
)

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
	token.ADD_ASSIGN:     ADD_ASSIGN,
	token.SUB_ASSIGN:     SUB_ASSIGN,
	token.MUL_ASSIGN:     MUL_ASSIGN,
	token.QUO_ASSIGN:     QUO_ASSIGN,
	token.REM_ASSIGN:     REM_ASSIGN,
	token.AND_ASSIGN:     AND_ASSIGN,
	token.OR_ASSIGN:      OR_ASSIGN,
	token.XOR_ASSIGN:     XOR_ASSIGN,
	token.SHL_ASSIGN:     SHL_ASSIGN,
	token.SHR_ASSIGN:     SHR_ASSIGN,
	token.AND_NOT_ASSIGN: AND_NOT_ASSIGN,
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
	token.DEFINE:         COLAS,
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
//TODO lbrHunt = .

const (
	st1 = iota
	st2
	st3
	st4
	st5
	st6
	st7
	st8
	st9
	st10
	st11
	st12
	st13
	st14
	st15
	st16
	st17
	st18
	st19
)

type pos struct {
	line, col int
}

type tok struct {
	tk  int
	val interface{}
	pos pos
}

type lx struct {
	*scanner.Scanner
	state      int
	dump       []tok
	toks       []tok
	ids        []tok
	preamble   int
	buf        tok
	bufValid   bool
	lbrHunt    bool
	lbrBalance int
	lparHunt   bool
}

func (x *lx) Lex(lval *yySymType) (r int) {
	dbg("\n<<<< Lex state st%d", x.state+1)
	defer func() {
		if e := recover(); e != nil {
			dbg(
				"---------------------------\n\n!!! RECOVERED: %s:%d:%d: %v\n\n---------------------------",
				x.Fname, x.Line, x.Col, e,
			)
			r = 0
			x.error("")
			return
		}

		if x.lparHunt && r == '(' {
			//dbg("LPAR")
			r, x.preamble, x.toks, x.ids, x.state = LPAR, 0, nil, nil, st11
		}
		x.lparHunt = false
		var s string
		if r < 128 {
			s = string(r)
		} else {
			s = yyToknames[r-ADD_ASSIGN]
		}
		dbg(">>>> %d:%d: returning %q, (%v)\n", lval.pos.line, lval.pos.col, s, lval.val)
		//TODO-
	}()

dump:
	if len(x.dump) != 0 {
		tk := x.dump[0]
		r, lval.pos, lval.val = tk.tk, tk.pos, tk.val
		x.dump = x.dump[1:]
		if len(x.dump) == 0 {
			x.dump = x.dump[:0]
		}
		x.state = st1
		return
	}

	for {
		if len(x.dump) != 0 {
			goto dump
		}

		dbg("[state st%d]", x.state+1)
		tk := x.lex0()

		switch r = tk.tk; x.state {
		case st1:
			x.preamble, x.toks, x.ids = -1, x.toks[:0], nil
			switch r {
			case CONST, VAR:
				x.toks, x.state = append(x.toks, tk), st2
			case FUNC:
				x.toks, x.state = append(x.toks, tk), st5
			case IDENTIFIER:
				x.toks, x.ids, x.state = append(x.toks, tk), append(x.ids, tk), st14
			case STRUCT:
				panic("st1 struct")
			default:
				x.dump = append(x.dump, tk)
			}
		case st2:
			switch r {
			case '(':
				x.toks, x.state = append(x.toks, tk), st3
			case IDENTIFIER:
				panic("st2 identifier")
			default:
				panic("st2 default")
			}
		case st3:
			switch r {
			case IDENTIFIER:
				if x.preamble < 0 {
					x.preamble = len(x.toks)
				}
				x.toks, x.ids, x.state = append(x.toks, tk), append(x.ids, tk), st4
			default:
				x.dump = append(x.toks, tk)
			}
		case st4: // state 4 accepts rule 1: const var
			switch r {
			case ',':
				panic("st4 ,")
			default:
				x.dump = append(x.toks[:x.preamble], tok{IDENTIFIER_LIST, x.ids, x.ids[0].pos}, tk)
			}
		case st5:
			switch r {
			case '(':
				x.toks, x.state = append(x.toks, tk), st6
			case IDENTIFIER:
				panic("st5 identifier")
			default:
				panic("st5 default")
			}
		case st6:
			switch r {
			case '*':
				panic("st6 *")
			case IDENTIFIER:
				if x.preamble < 0 {
					x.preamble = len(x.toks)
				}
				x.toks, x.ids, x.state = append(x.toks, tk), append(x.ids, tk), st13
			default:
				x.dump = append(x.toks, tk)
			}
		case st7:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st8:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st9:
			switch r {
			case IDENTIFIER:
				panic("st9 identifier")
			default:
				x.dump = append(x.toks, tk)
			}
		case st10:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st11:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st12: // state 12 accepts rule 2: func
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st13: // state 13 accepts rule 2: func
			switch r {
			case '*':
				panic("st13 *")
			case IDENTIFIER:
				panic("st13 identifier")
			case ')':
				x.preamble, x.toks, x.ids, x.state = -1, append(x.toks, tk), nil, st9
			case ',':
				panic("st13 ,")
			default:
				panic("st13 default")
			}
		case st14:
			switch r {
			case ',':
				panic("str14 ,")
			case COLAS:
				panic("str14 :=")
			default:
				x.dump = append(x.toks, tk)
			}
		case st15:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st16: // state 16 accepts rule 4: idlist_colas
			panic(fmt.Sprintf("internal error st%d", x.state+1))
		case st17:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st18:
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st19: // state 19 accepts rule 3: struct
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		default:
			panic(fmt.Sprintf("internal error st%d", x.state+1))
		}
	}
}

func (x *lx) lex0() (y tok) {
	defer func() {
		var s string
		if y.tk < 128 {
			s = string(y.tk)
		} else {
			s = yyToknames[y.tk-ADD_ASSIGN]
		}
		dbg("........ %d:%d returning %q", x.Line, x.Col, s)
	}()
	for {
		t, val := x.ScanSemis()
		if t == token.COMMENT {
			continue
		}

		tok := tok{xlat[t], val, pos{x.Line, x.Col}}
		//dbg("ScanSemis %v (%v : %v, %#x, %q)", t, val, tok.tk, tok.tk, string(tok.tk))
		if !x.bufValid {
			x.buf, x.bufValid = tok, true
			continue
		}

		if p, n := x.buf.tk, tok.tk; (p == ',' || p == ';') && (n == ')' || n == '}') {
			tok.val, x.bufValid = x.buf, true
			x.buf = tok
			continue
		}

		if p, n := x.buf.tk, tok.tk; p == '.' && n == '(' {
			tok.tk, tok.pos, x.bufValid = DOT_LPAR, x.buf.pos, true
			x.buf = tok
			continue
		}

		if p, n := x.buf.tk, tok.tk; p == DOT_LPAR && n == TYPE {
			tok.tk, tok.pos, x.bufValid = DOT_LPAR_TYPE, x.buf.pos, true
			x.buf = tok
			continue
		}

		y, x.buf = x.buf, tok
		break
	}

	switch y.tk {
	case FOR, IF, SWITCH:
		x.lbrHunt, x.lbrBalance = true, 0
	case '(':
		if x.lbrHunt {
			x.lbrBalance++
		}
	case ')':
		if x.lbrHunt {
			x.lbrBalance--
		}
	case '{':
		if x.lbrHunt && x.lbrBalance <= 0 {
			y.tk, x.lbrHunt = LBR, false
		}
	}
	return
}

func (x *lx) peek() int {
	if !x.bufValid {
		panic("internal error")
	}

	return x.buf.tk
}

func (x *lx) err(pos pos, format string, arg ...interface{}) {
	x.Error(fmt.Sprintf("%s:%d:%d: "+format, append([]interface{}{x.Fname, pos.line, pos.col}, arg...)...))
}

func (x *lx) error(format string, arg ...interface{}) {
	x.err(pos{x.Line, x.Col}, format, arg...)
}
