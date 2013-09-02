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

	"github.com/cznic/mathutil"
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
//TODO lbr = .

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
	st20
	st21
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
	prev       tok
	prevValid  bool
	rxFix      int
	lbrHunt    bool
	lbrBalance int
}

func (x *lx) Lex(lval *yySymType) (r int) {
	dbg("\n<<<< Lex state st%d", x.state+1)
	defer func() {
		var s string
		if r < 128 {
			s = string(r)
		} else {
			s = yyToknames[r-ANDAND]
		}
		dbg(">>>> %d:%d: returning %q, (%v)\n", lval.pos.line, lval.pos.col, s, lval.val)
	}()

dump:
	if len(x.dump) != 0 {
		tk := x.dump[0]
		r, lval.pos, lval.val = tk.tk, tk.pos, tk.val
		x.dump = x.dump[1:]
		if len(x.dump) == 0 {
			x.dump = nil
		}
		x.state = st1
		return
	}

	for {
		if len(x.dump) != 0 {
			goto dump
		}

		dbg("[state st%d]", x.state+1)
		tk := x.lex()

		switch r = tk.tk; x.state {
		case st1:
			x.rxFix = 0
			switch r {
			case CONST, VAR:
				x.preamble, x.toks, x.ids, x.state = -1, []tok{tk}, nil, st2
			case FUNC:
				x.toks, x.state = []tok{tk}, st5
			case IDENTIFIER:
				x.toks, x.ids, x.state = []tok{tk}, []tok{tk}, st14
			case STRUCT:
				panic("st1 struct")
			default:
				lval.pos, lval.val = tk.pos, tk.val
				return
			}
		case st2:
			switch r {
			case '(':
				x.toks, x.state = append(x.toks, tk), st3
			case IDENTIFIER:
				x.preamble, x.toks, x.ids, x.state = len(x.toks), append(x.toks, tk), []tok{tk}, st4
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
		case st4: // state 4 accepts rule 1 // const, var
			switch r {
			case ',':
				x.toks, x.state = append(x.toks, tk), st3
			default:
				x.dump = append(x.toks[:mathutil.Max(0, x.preamble)], tok{IDENTIFIER_LIST, x.ids, x.ids[0].pos}, tk)
			}
		case st5:
			switch r {
			case '(':
				x.toks, x.state = append(x.toks, tk), st6
			case IDENTIFIER:
				x.toks, x.state = append(x.toks, tk), st10
			default:
				panic("st5 default")
			}
		case st6:
			switch r {
			case '*':
				x.toks, x.state = append(x.toks, tk), st7
			case IDENTIFIER:
				x.preamble, x.toks, x.ids, x.state = len(x.toks), append(x.toks, tk), []tok{tk}, st13
			default:
				x.dump = append(x.toks, tk)
			}
		case st7:
			switch r {
			case IDENTIFIER:
				x.toks, x.state = append(x.toks, tk), st8
			default:
				panic("st7 default")
			}
		case st8:
			switch r {
			case ')':
				x.toks, x.state = append(x.toks, tk), st9
			default:
				panic("st8 default")
			}
		case st9:
			switch r {
			case IDENTIFIER:
				x.toks, x.state = append(x.toks, tk), st10
			default:
				if i := x.rxFix; i != 0 {
					t := x.toks[i]
					x.toks[i] = tok{IDENTIFIER_LIST, []tok{t}, t.pos}
				}
				x.dump = append(x.toks, tk)
			}
		case st10:
			switch r {
			case '(':
				x.preamble, x.toks, x.ids, x.state = len(x.toks)+1, append(x.toks, tk), nil, st11
			default:
				panic("st10 default")
			}
		case st11:
			switch r {
			case IDENTIFIER:
				x.toks, x.ids, x.state = append(x.toks, tk), append(x.ids, tk), st12
			default:
				x.dump = append(x.toks, tk)
			}
		case st12: // state 12 accepts rule 2 // func
			switch r {
			case ',':
				x.toks, x.state = append(x.toks, tk), st11
			case ')':
				x.dump = append(x.toks, tk)
			default:
				x.dump = append(x.toks[:x.preamble], tok{IDENTIFIER_LIST, x.ids, x.ids[0].pos}, tk)
			}
		case st13: // state 13 accepts rule 2 // func
			switch r {
			case '*':
				x.toks, x.state = append(x.toks, tk), st7
			case IDENTIFIER:
				x.rxFix, x.toks, x.state = len(x.toks)-1, append(x.toks, tk), st8
			case ')':
				x.toks, x.state = append(x.toks, tk), st9
			case ',':
				x.toks, x.state = append(x.toks, tk), st11
			default:
				x.dump = append(x.toks[:x.preamble], tok{IDENTIFIER_LIST, x.ids, x.ids[0].pos}, tk)
			}
		case st14:
			switch r {
			case CONST, VAR:
				panic("st14 const var")
			case FUNC:
				x.toks, x.state = append(x.toks, tk), st5
			case ',':
				x.toks, x.state = append(x.toks, tk), st15
			case COLAS:
				panic("st14 colas")
			case STRUCT:
				x.toks, x.state = append(x.toks, tk), st18
			default:
				x.dump = append(x.toks, tk)
			}
		case st15:
			switch r {
			case IDENTIFIER:
				x.toks, x.ids, x.state = append(x.toks, tk), append(x.ids, tk), st16
			default:
				x.dump = append(x.toks, tk)
			}
		case st16:
			switch r {
			case ',':
				x.toks, x.state = append(x.toks, tk), st15
			case COLAS:
				panic("st16 :=")
			default:
				x.dump = append(x.toks, tk)
			}
		case st17: // state 17 accepts rule 4 // colas
			panic(fmt.Sprintf("TODO st%d", x.state+1))
		case st18:
			switch r {
			case '{':
				x.preamble, x.toks, x.ids, x.state = -1, append(x.toks, tk), nil, st19
			default:
				panic("st18 default")
			}
		case st19:
			switch r {
			case IDENTIFIER:
				if x.preamble < 0 {
					x.preamble = len(x.toks)
				}
				x.toks, x.ids, x.state = append(x.toks, tk), append(x.ids, tk), st20
			default:
				x.dump = append(x.toks, tk)
			}
		case st20: // state 20 accepts rule 3 // struct
			switch r {
			case ',':
				x.toks, x.state = append(x.toks, tk), st19
			case '}', '.':
				x.dump = append(x.toks, tk)
			default:
				x.dump = append(x.toks[:mathutil.Max(0, x.preamble)], tok{IDENTIFIER_LIST, x.ids, x.ids[0].pos}, tk)
			}
		default:
			panic(fmt.Sprintf("internal error st%d", x.state+1))
		}
	}
}

func (x *lx) lex() (y tok) {
	defer func() {
		var s string
		if y.tk < 128 {
			s = string(y.tk)
		} else {
			s = yyToknames[y.tk-ANDAND]
		}
		dbg("........ %d:%d returning %q", x.Line, x.Col, s)
	}()
	for {
		t, val := x.ScanSemis()
		if t == token.COMMENT {
			continue
		}

		tok := tok{xlat[t], val, pos{x.Line, x.Col}}
		//dbg("ScanSemis %v", t)
		if !x.prevValid {
			x.prev, x.prevValid = tok, true
			continue
		}

		if p, n := x.prev.tk, tok.tk; (p == ',' || p == ';') && (n == ')' || n == '}') {
			tok.val, x.prevValid = x.prev, false
			y = tok
			break
		}

		y, x.prev = x.prev, tok
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

func (x *lx) err(pos pos, format string, arg ...interface{}) {
	x.Error(fmt.Sprintf("%s:%d:%d: "+format, append([]interface{}{x.Fname, pos.line, pos.col}, arg...)...))
}

func (x *lx) error(format string, arg ...interface{}) {
	x.err(pos{x.Line, x.Col}, format, arg...)
}
