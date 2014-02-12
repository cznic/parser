%{

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"go/token"
)

%}

%union	{
	token  tkn
	node   Node
	list   []Node
	param  *Param
	params []*Param
}

%token	<token>
	'!' '%' '&' '(' '*' '+' '-' '.' '/' ':' '<' '=' '>' '[' '^' '{' '|' '~'
	_ANDAND _ANDNOT _ASOP _BODY _BREAK _CASE _CHAN _COLAS _COMM _CONST
	_CONTINUE _DDD _DEC _DEFAULT _DEFER _ELSE _EQ _FALL _FOR _FUNC _GE _GO
	_GOTO _IF _IGNORE _IMPORT _INC _INTERFACE _LE _LITERAL _LSH _MAP _NAME
	_NE _OROR _PACKAGE _RANGE _RETURN _RSH _SELECT _STRUCT _SWITCH _TYPE
	_VAR

%type	<node>
	bare_complitexpr
	case caseblock complitexpr comptype compound_stmt constdcl constdcl1
	convtype
	dcl_name dotname
	else elseif embed expr expr_or_type
	fndcl fnlitdcl fnliteral fnret_type fntype for_body for_header for_stmt
	if_header if_stmt import_stmt indcl interfacedcl interfacetype
	keyval
	labelname lbrace
	name name_or_type new_name ntype non_dcl_stmt non_expr_type
	non_recvchantype
	oexpr oliteral onew_name osimple_stmt othertype
	package packname pexpr pexpr_no_paren pseudocall ptrtype
	range_stmt recvchantype
	select_stmt simple_stmt structdcl structtype switch_stmt sym
	typedcl typedclname
	uexpr
	xfndcl

%type	<list>
	braced_keyval_list
	caseblock_list common_dcl constdcl_list
	dcl_name_list
	elseif_list expr_list expr_or_type_list
	fnbody
	import_stmt_list interfacedcl_list
	keyval_list
	loop_body
	new_name_list
	oexpr_list
	stmt stmt_list structdcl_list
	typedcl_list
	vardcl vardcl_list

%type	<param>
	arg_type
	dotdotdot

%type	<params>
	arg_type_list
	fnres
	oarg_type_list_ocomma

%left	_COMM

%left	_OROR
%left	_ANDAND
%left	_EQ _NE _LE _GE '<' '>'
%left	'+' '-' '|' '^'
%left	'*' '/' '%' '&' _LSH _RSH _ANDNOT

%left	notPackage
%left	_PACKAGE

%left	notParen
%left	'('

%left	')'
%left	preferToRightParen

%%

file:
	package
	{ //46
		yyTLD(yylex, $1)
	}
	imports
	xdcl_list

package:
	%prec notPackage
	{ //60
		yyErr(yylex, "package statement must be first")
		goto ret1
	}
|	_PACKAGE sym ';'
	{ //64
		$$ = &Package{$1.p(), $2.(*Ident)}
	}

imports:
|	imports import ';'

import:
	_IMPORT import_stmt
	{
		yyTLD(yylex, $2)
	}
|	_IMPORT '(' import_stmt_list osemi ')'
	{ //83
		yyTLDs(yylex, $3)
	}
|	_IMPORT '(' ')'

import_stmt:
	_LITERAL
	{ //93
		$$ = newImport(yylex, (*Ident)(nil), newLiteral($1))
	}
|	sym _LITERAL
	{ //97
		$$ = newImport(yylex, $1, newLiteral($2))
	}
|	'.' _LITERAL
	{ //101
		$$ = newImport(yylex, &Ident{$1.p(), "."}, newLiteral($2))
	}

import_stmt_list:
	import_stmt
	{ //107
		$$ = []Node{$1}
	}
|	import_stmt_list ';' import_stmt
	{ //111
		$$ = append($1, $3)
	}

// TLD
xdcl:
|	common_dcl
	{ //120
		yyTLDs(yylex, $1)
	}
|	xfndcl
	{ //124
		yyTLD(yylex, $1)
	}
|	non_dcl_stmt
	{ //128
		yyErrPos(yylex, $1, "non-declaration statement outside function body");
	}
|	error

common_dcl:
	_VAR vardcl
	{ //138
		$$ = $2
	}
|	_VAR '(' vardcl_list osemi ')'
	{ //142
		$$ = $3
	}
|	_VAR '(' ')'
	{ //146
		$$ = nil
	}
|	lconst constdcl
	{ //150
		$$ = newConstDecls(yylex, []Node{$2})
	}
|	lconst '(' constdcl osemi ')'
	{ //154
		$$ = newConstDecls(yylex, []Node{$3})
	}
|	lconst '(' constdcl ';' constdcl_list osemi ')'
	{ //158
		$$ = newConstDecls(yylex, append([]Node{$3}, $5...))
	}
|	lconst '(' ')'
	{ //162
		$$ = nil
	}
|	_TYPE typedcl
	{ //166
		$$ = []Node{$2}
	}
|	_TYPE '(' typedcl_list osemi ')'
	{ //170
		$$ = $3
	}
|	_TYPE '(' ')'
	{ //174
		$$ = nil
	}

lconst:
	_CONST
	{ //180
		p := yy(yylex)
		p.constExpr, p.constIota, p.constType = nil, 0, nil
	}

vardcl:
	dcl_name_list ntype
	{ //186
		$$ = newVarDecls($1, $2, nil)
	}
|	dcl_name_list ntype '=' expr_list
	{ //190
		$$ = newVarDecls($1, $2, $4)
	}
|	dcl_name_list '=' expr_list
	{ //194
		$$ = newVarDecls($1, nil, $3)
	}

constdcl:
	dcl_name_list ntype '=' expr_list
	{ //200
		$$ = newConstSpec(yylex, $1, $2, $4)
	}
|	dcl_name_list '=' expr_list
	{ //204
		$$ = newConstSpec(yylex, $1, nil, $3)
	}

constdcl1:
	constdcl
|	dcl_name_list ntype
	{ //214
		panic(".y:215") //TODO ???
		//yyErrPos(yylex, $2, "const declaration cannot have type without expression")
	}
|	dcl_name_list
	{ //218
		$$ = newConstSpec(yylex, $1, nil, nil)
	}

typedclname:
	sym

typedcl:
	typedclname ntype
	{ //230
		$$ = &TypeDecl{$1.p(), $1.(*Ident), $2}
	}

simple_stmt:
	expr
|	expr _ASOP expr
	{ //240
		$$ = &Assignment{$2.p(), $2.tok, []Node{$1}, []Node{$3}}
	}
|	expr_list '=' expr_list
	{ //244
		$$ = &Assignment{$2.p(), $2.tok, $1, $3}
	}
|	expr_list _COLAS expr_list
	{ //248
		$$ = &ShortVarDecl{$2.p(), $1, $3}
	}
|	expr _INC
	{ //252
		$$ = &IncDecStmt{$2.p(), $1, $2.tok}
	}
|	expr _DEC
	{ //256
		$$ = &IncDecStmt{$2.p(), $1, $2.tok}
	}

case:
	_CASE expr_or_type_list ':'
	{ //262
		$$ = &SwitchCase{$1.p(), $2, nil}
	}
|	_CASE expr_or_type_list '=' expr ':'
	{ //266
		$$ = &SwitchCase{$1.p(), []Node{&Assignment{$3.p(), $3.tok, $2, []Node{$4}}}, nil}
	}
|	_CASE expr_or_type_list _COLAS expr ':'
	{ //270
		$$ = &SwitchCase{$1.p(), []Node{&Assignment{$3.p(), $3.tok, $2, []Node{$4}}}, nil}
	}
|	_DEFAULT ':'
	{ //274
		$$ = &SwitchCase{pos: $1.p()}
	}

compound_stmt:
	'{' stmt_list '}'
	{ //280
		$$ = &CompoundStament{$1.p(), $2}
	}

caseblock:
	case stmt_list
	{ //290
		$1.(*SwitchCase).Body = $2
		$$ = $1
	}

caseblock_list:
	{ //295
		$$ = nil
	}
|	caseblock_list caseblock
	{ //299
		$$ = append($1, $2)
	}

loop_body:
	_BODY stmt_list '}'
	{ //305
		$$ = $2
	}

range_stmt:
	expr_list '=' _RANGE expr
	{ //311
		$$ = &ForStmt{Range: &Assignment{$2.p(), $2.tok, $1, []Node{$4}}}
	}
|	expr_list _COLAS _RANGE expr
	{ //315
		$$ = &ForStmt{Range: &Assignment{$2.p(), $2.tok, $1, []Node{$4}}}
	}

for_header:
	osimple_stmt ';' osimple_stmt ';' osimple_stmt
	{ //321
		$$ = &ForStmt{Init: $1, Cond: $3, Post: $5}
	}
|	osimple_stmt
	{ //325
		$$ = &ForStmt{Cond: $1}
	}
|	range_stmt

for_body:
	for_header loop_body
	{ //335
		$1.(*ForStmt).Body = $2
		$$ = $1
	}

for_stmt:
	_FOR for_body
	{ //341
		$2.(*ForStmt).pos = $1.p()
		$$ = $2
	}

if_header:
	osimple_stmt
	{ //347
		$$ = &IfStmt{Cond: $1}
	}
|	osimple_stmt ';' osimple_stmt
	{ //351
		$$ = &IfStmt{Init: $1, Cond: $3}
	}

if_stmt:
	_IF if_header loop_body elseif_list else
	{ //357
		x := $2.(*IfStmt)
		l := make([]*IfStmt, len($4))
		for i, v := range $4 {
			l[i] = v.(*IfStmt)
		}
		x.pos, x.Body, x.Elif, x.Else = $1.p(), $3, l, $5.(*CompoundStament)
		$$ = x
	}

elseif:
	_ELSE _IF if_header loop_body
	{ //363
		x := $3.(*IfStmt)
		x.pos, x.Body = $2.p(), $4
		$$ = x
	}

elseif_list:
	{ //368
		$$ = nil
	}
|	elseif_list elseif
	{ //372
		$$ = append($1, $2)
	}

else:
	{ //377
		$$ = (*CompoundStament)(nil)
	}
|	_ELSE compound_stmt
	{ //381
		$$ = $2
	}

switch_stmt:
	_SWITCH if_header _BODY caseblock_list '}'
	{ //387
		l := make([]*SwitchCase, len($4))
		for i, v := range $4 {
			l[i] = v.(*SwitchCase)
		}
		x := $2.(*IfStmt)
		$$ = &SwitchStmt{$1.p(), x.Init, x.Cond, l}
	}

select_stmt:
	//_SELECT _BODY caseblock_list '}'
	_SELECT '{' caseblock_list '}'
	{ //393
		x := &SelectStmt{pos: $1.p()}
		for _, v := range $3 {
			l := v.(*SwitchCase).Expr
			if len(l) != 1 {
				yyErrPos(yylex, l[1], "select cases cannot be lists")
				continue
			}

			v0 := l[0]
			switch t := v0.(type) {
			case *Assignment:
				if t.Op != token.ASSIGN && t.Op != token.DEFINE {
					break
				}

				if len(t.L) > 2 || len(t.R) != 1 {
					break
				}

				if y, ok := t.R[0].(*UnOp); ok && y.Op == token.ARROW {
					x.Cases = append(x.Cases, &CommCase{v0.p(), v0})
					continue
				}
			case *BinOp:
				if t.Op != token.ARROW {
					break
				}

				x.Cases = append(x.Cases, &CommCase{v0.p(), v0})
				continue
			case *UnOp:
				if t.Op != token.ARROW {
					break
				}

				x.Cases = append(x.Cases, &CommCase{v0.p(), v0})
				continue
			}
			yyErrPos(yylex, v0, "select case must be receive, send or assign recv")
		}
		$$ = x
	}

expr:
	uexpr
|	expr _OROR expr
	{ //403
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _ANDAND expr
	{ //407
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _EQ expr
	{ //411
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _NE expr
	{ //415
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '<' expr
	{ //419
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _LE expr
	{ //423
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _GE expr
	{ //427
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '>' expr
	{ //431
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '+' expr
	{ //435
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '-' expr
	{ //439
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '|' expr
	{ //443
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '^' expr
	{ //447
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '*' expr
	{ //451
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '/' expr
	{ //455
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '%' expr
	{ //459
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr '&' expr
	{ //463
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _ANDNOT expr
	{ //467
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _LSH expr
	{ //471
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _RSH expr
	{ //475
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}
|	expr _COMM expr
	{ //479
		$$ = &BinOp{$2.p(), $2.tok, $1, $3}
	}

uexpr:
	pexpr
|	'*' uexpr
	{ //489
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	'&' uexpr
	{ //493
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	'+' uexpr
	{ //497
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	'-' uexpr
	{ //501
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	'!' uexpr
	{ //505
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	'~' uexpr
	{ //509
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	'^' uexpr
	{ //513
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}
|	_COMM uexpr
	{ //517
		$$ = &UnOp{$1.p(), $1.tok, $2}
	}

pseudocall:
	pexpr '(' ')'
	{ //523
		$$ = &CallOp{$2.p(), $1, nil, false}
	}
|	pexpr '(' expr_or_type_list ocomma ')'
	{ //527
		$$ = &CallOp{$2.p(), $1, $3, false}
	}
|	pexpr '(' expr_or_type_list _DDD ocomma ')'
	{ //531
		$$ = &CallOp{$2.p(), $1, $3, true}
	}

pexpr_no_paren:
	_LITERAL
	{ //537
		$$ = &Literal{$1.p(), $1.tok, $1.lit}
	}
|	name
|	pexpr '.' sym
	{ //545
		$$ = &SelectOp{$2.p(), $1, $3.(*Ident)}
	}
|	pexpr '.' '(' expr_or_type ')'
	{ //549
		$$ = &TypeAssertion{$3.p(), $1, $4}
	}
|	pexpr '.' '(' _TYPE ')'
	{ //553
		$$ = &TypeSwitch{$4.p(), $1}
	}
|	pexpr '[' expr ']'
	{ //557
		$$ = &IndexOp{$2.p(), $1, $3}
	}
|	pexpr '[' oexpr ':' oexpr ']'
	{ //561
		$$ = &SliceOp{$2.p(), $1, $3, $5, nil}
	}
|	pexpr '[' oexpr ':' oexpr ':' oexpr ']'
	{ //565
		if $5 == nil {
			yyErrPos(yylex, $4, "middle index required in 3-index slice")
		}
		if $7 == nil {
			yyErrPos(yylex, $6, "final index required in 3-index slice")
		}
		$$ = &SliceOp{$2.p(), $1, $3, $5, $7}
	}
|	pseudocall
|	convtype '(' expr ocomma ')'
	{ //573
		$$ = &ConvOp{$2.p(), $1, $3}
	}
|	comptype lbrace start_complit braced_keyval_list '}'
	{ //577
		$$ = &CompLit{$2.p(), $1, elements($4)}
	}
|	pexpr_no_paren '{' start_complit braced_keyval_list '}'
	{ //581
		$$ = &CompLit{$2.p(), $1, elements($4)}
	}
|	'(' expr_or_type ')' '{' start_complit braced_keyval_list '}'
	{ //585
		$$ = &CompLit{$4.p(), $2, elements($6)}
	}
|	fnliteral

start_complit:
	{ //594
	}

keyval:
	expr ':' complitexpr
	{ //600
		$$ = &Element{$1.p(), $1, $3}
	}

bare_complitexpr:
	expr
	{ //609
		$$ = &Element{$1.p(), nil, $1}
	}
|	'{' start_complit braced_keyval_list '}'
	{ //610
		$$ = &Element{$1.p(), nil, &CompLit{$1.p(), nil, elements($3)}}
	}

complitexpr:
	expr
	{ //616
		$$ = &Element{$1.p(), nil, $1}
	}
|	'{' start_complit braced_keyval_list '}'
	{ //620
		$$ = &Element{$1.p(), nil, &CompLit{$1.p(), nil, elements($3)}}
	}

pexpr:
	pexpr_no_paren
|	'(' expr_or_type ')'
	{ //630
		$$ = &Paren{$1.p(), $2}
	}

expr_or_type:
	expr
|	non_expr_type	%prec preferToRightParen

name_or_type:
	ntype

lbrace:
	_BODY
	{
		$$ = Node($1)
	}
|	'{'
	{
		$$ = Node($1)
	}

// - field name of a struct type definition
// - label name declaration/reference
// - method name of an interface type defintion
new_name:
	sym

// identifier in the identifier list of a var or const declaration
//TODO declare in current scope
dcl_name:
	sym

onew_name:
	{ //673
		$$ = (*Ident)(nil)
	}
|	new_name

sym:
	_NAME
	{ //683
		$$ = &Ident{$1.p(), $1.lit}
	}

name:
	sym	%prec notParen

labelname:
	new_name

dotdotdot:
	_DDD
	{ //701
		yy(yylex).errPos($1.tpos, "final argument in variadic function missing type")
		$$ = &Param{pos: $1.p(), Ddd: true}
	}
|	_DDD ntype
	{ //705
		$$ = &Param{pos: $1.p(), Ddd: true, Type: $2}
	}

ntype:
	recvchantype
|	fntype
|	othertype
|	ptrtype
|	dotname
	{
		$$ = &NamedType{$1.p(), $1.(*QualifiedIdent), nil, yyScope(yylex)}
	}
|	'(' ntype ')'
	{ //731
		$$ = &Paren{$1.p(), $2}
	}

non_expr_type:
	recvchantype
|	fntype
|	othertype
|	'*' non_expr_type
	{ //749
		$$ = &PtrType{$1.p(), $2}
	}

non_recvchantype:
	fntype
|	othertype
|	ptrtype
|	dotname
|	'(' ntype ')'
	{ //771
		$$ = &Paren{$1.p(), $2}
	}

convtype:
	fntype
|	othertype

comptype:
	othertype

fnret_type:
	recvchantype
|	fntype
|	othertype
|	ptrtype
|	dotname
	{ //790
		$$ = &NamedType{$1.p(), $1.(*QualifiedIdent), nil, yyScope(yylex)}
	}

dotname:
	name
	{ //815
		$$ = &QualifiedIdent{$1.p(), nil, $1.(*Ident)}
	}
|	name '.' sym
	{ //819
		$$ = &QualifiedIdent{$1.p(), $1.(*Ident), $3.(*Ident)}
	}

othertype:
	'[' oexpr ']' ntype
	{ //825
		switch {
		case $2 != nil:
			$$ = &ArrayType{$1.p(), $2, $4}
		default:
			$$ = &SliceType{$1.p(), $4}
		}
	}
|	'[' _DDD ']' ntype
	{ //829
		$$ = &ArrayType{$1.p(), nil, $4}
	}
|	_CHAN non_recvchantype
	{ //833
		$$ = &ChannelType{$1.p(), BidirectionalChannel, $2}
	}
|	_CHAN _COMM ntype
	{ //837
		$$ = &ChannelType{$2.p(), SendOnlyChannel, $3}
	}
|	_MAP '[' ntype ']' ntype
	{ //841
		$$ = &MapType{$1.p(), $3, $5}
	}
|	structtype
|	interfacetype

ptrtype:
	'*' ntype
	{ //855
		$$ = &PtrType{$1.p(), $2}
	}

recvchantype:
	_COMM _CHAN ntype
	{ //861
		$$ = &ChannelType{$1.p(), ReadOnlyChannel, $3}
	}

structtype:
	_STRUCT lbrace structdcl_list osemi '}'
	{ //867
		$$ = newStructType(yylex, $1, $3)
	}
|	_STRUCT lbrace '}'
	{ //871
		$$ = newStructType(yylex, $1, nil)
	}

interfacetype:
	_INTERFACE lbrace interfacedcl_list osemi '}'
	{ //877
		x := newInterfaceType(yylex, $3)
		x.pos = $1.p()
		$$ = x
	}
|	_INTERFACE lbrace '}'
	{ //881
		$$ = &InterfaceType{pos: $1.p()}
	}

xfndcl:
	_FUNC fndcl fnbody
	{ //887
		x := $2.(*FuncDecl)
		x.pos, x.Body = $1.p(), $3
		$$ = x
	}

fndcl:
	sym '(' oarg_type_list_ocomma ')' fnres
	{ //893
		$$ = &FuncDecl{Name: $1.(*Ident), Type: newFuncType(yylex, $2.p(), nil, $3, $5)}
	}
|	'(' oarg_type_list_ocomma ')' sym '(' oarg_type_list_ocomma ')' fnres
	{ //897
		$$ = &FuncDecl{Name: $4.(*Ident), Type: newFuncType(yylex, $1.p(), $2, $6, $8)}
	}

fntype:
	_FUNC '(' oarg_type_list_ocomma ')' fnres
	{ //903
		$$ = newFuncType(yylex, $1.p(), nil, $3, $5)
	}

fnbody:
	{ //908
		$$ = nil
	}
|	'{' stmt_list '}'
	{ //912
		$$ = $2
	}

fnres:
	%prec notParen
	{ //918
		$$ = nil
	}
|	fnret_type
	{ //922
		$$ = []*Param{{pos: $1.p(), Type: $1}}
	}
|	'(' oarg_type_list_ocomma ')'
	{ //926
		$$ = $2
	}

fnlitdcl:
	fntype

fnliteral:
	fnlitdcl lbrace stmt_list '}'
	{ //938
		x := $1.(*FuncType)
		$$ = &FuncLit{x.p(), x, $3}
	}
|	fnlitdcl error

xdcl_list:
|	xdcl_list xdcl ';'

vardcl_list:
	vardcl
	{ //957
		$$ = $1
	}
|	vardcl_list ';' vardcl
	{ //961
		$$ = append($1, $3...)
	}

constdcl_list:
	constdcl1
	{ //967
		$$ = []Node{$1}
	}
|	constdcl_list ';' constdcl1
	{ //971
		$$ = append($1, $3)
	}

typedcl_list:
	typedcl
	{ //977
		$$ = []Node{$1}
	}
|	typedcl_list ';' typedcl
	{ //981
		$$ = append($1, $3)
	}

structdcl_list:
	structdcl
	{ //987
		$$ = []Node{$1}
	}
|	structdcl_list ';' structdcl
	{ //991
		$$ = append($1, $3)
	}

interfacedcl_list:
	interfacedcl
	{ //997
		$$ = []Node{$1}
	}
|	interfacedcl_list ';' interfacedcl
	{ //1001
		$$ = append($1, $3)
	}

structdcl:
	new_name_list ntype oliteral
	{ //1007
		$$ = newFields($1, false, $2, $3)
	}
|	embed oliteral
	{ //1011
		q := $1.(*QualifiedIdent)
		$$ = newFields([]Node{q.I}, true, &NamedType{q.p(), q, nil, yyScope(yylex)}, $2)
	}
|	'(' embed ')' oliteral
	{ //1015
		yyErrPos(yylex, $1.p(), "cannot parenthesize embedded type")
		$$ = &fields{}
	}
|	'*' embed oliteral
	{ //1019
		q := $2.(*QualifiedIdent)
		$$ = newFields([]Node{q.I}, true, &PtrType{$1.p(), &NamedType{q.p(), q, nil, yyScope(yylex)}}, $3)
	}
|	'(' '*' embed ')' oliteral
	{ //1023
		yyErrPos(yylex, $1.p(), "cannot parenthesize embedded type")
		$$ = &fields{}
	}
|	'*' '(' embed ')' oliteral
	{ //1027
		yyErrPos(yylex, $1.p(), "cannot parenthesize embedded type")
		$$ = &fields{}
	}

packname:
	_NAME
	{ //1033
		$$ = &QualifiedIdent{$1.p(), nil, &Ident{$1.p(), $1.lit}}
	}
|	_NAME '.' sym
	{ //1037
		$$ = &QualifiedIdent{$1.p(), &Ident{$1.p(), $1.lit}, $3.(*Ident)}
	}

embed:
	packname

interfacedcl:
	new_name indcl
	{ //1049
		$$ = &MethodSpec{$1.p(), &QualifiedIdent{$1.p(), nil, $1.(*Ident)}, $2.(*FuncType)}
	}
|	packname
	{ //1053
		$$ = &MethodSpec{$1.p(), $1.(*QualifiedIdent), nil}
	}
|	'(' packname ')'
	{ //1057
		$$ = &MethodSpec{$2.p(), $2.(*QualifiedIdent), nil}
		yyErrPos(yylex, $2, "cannot parenthesize embedded type");
	}

indcl:
	'(' oarg_type_list_ocomma ')' fnres
	{ //1063
		$$ = newFuncType(yylex, $1.p(), nil, $2, $4)
	}

arg_type:
	name_or_type
	{ //1069
		$$ = &Param{pos: $1.p(), Type: $1}
	}
|	sym name_or_type
	{ //1073
		$$ = &Param{pos: $1.p(), Name: $1.(*Ident), Type: $2}
	}
|	sym dotdotdot
	{ //1077
		x := $2
		x.Name, x.Ddd = $1.(*Ident), true
		$$ = x
	}
|	dotdotdot

arg_type_list:
	arg_type
	{ //1087
		$$ = []*Param{$1}
	}
|	arg_type_list ',' arg_type
	{ //1091
		$$ = append($1, $3)
	}

oarg_type_list_ocomma:
	{ //1096
		$$ = nil
	}
|	arg_type_list ocomma
	{
		$$ = $1
	}

stmt:
	{ //1105
		$$ = nil
	}
|	compound_stmt
	{ //1109
		$$ = []Node{$1}
	}
|	common_dcl
|	non_dcl_stmt
	{ //1117
		$$ = []Node{$1}
	}
|	error
	{ //1121
		$$ = nil
	}

non_dcl_stmt:
	simple_stmt
|	for_stmt
|	switch_stmt
|	select_stmt
|	if_stmt
|	labelname ':' stmt
	{ //1151
		$$ = &LabeledStmt{$2.p(), $1.(*Ident), $3}
	}
|	_FALL
	{ //1155
		$$ = &FallthroughStmt{$1.p()}
	}
|	_BREAK onew_name
	{ //1159
		$$ = &BreakStmt{$1.p(), $2.(*Ident)}
	}
|	_CONTINUE onew_name
	{ //1163
		$$ = &ContinueStmt{$1.p(), $2.(*Ident)}
	}
|	_GO pseudocall
	{ //1167
		$$ = &GoStmt{$1.p(), $2.(*CallOp)}
	}
|	_DEFER pseudocall
	{ //1171
		$$ = &DeferStmt{$1.p(), $2.(*CallOp)}
	}
|	_GOTO new_name
	{ //1175
		$$ = &GotoStmt{$1.p(), $2.(*Ident)}
	}
|	_RETURN oexpr_list
	{ //1179
		$$ = &ReturnStmt{$1.p(), $2}
	}

stmt_list:
	stmt
|	stmt_list ';' stmt
	{ //1189
		$$ = append($1, $3...)
	}

// field names of a struct type definition
new_name_list:
	new_name
	{ //1195
		$$ = []Node{$1}
	}
|	new_name_list ',' new_name
	{ //1199
		$$ = append($1, $3)
	}

// identifier list of the var or const declarations
dcl_name_list:
	dcl_name
	{ //1205
		$$ = []Node{$1}
	}
|	dcl_name_list ',' dcl_name
	{ //1209
		$$ = append($1, $3)
	}

expr_list:
	expr
	{ //1215
		$$ = []Node{$1}
	}
|	expr_list ',' expr
	{ //1219
		$$ = append($1, $3)
	}

expr_or_type_list:
	expr_or_type
	{ //1225
		$$ = []Node{$1}
	}
|	expr_or_type_list ',' expr_or_type
	{ //1229
		$$ = append($1, $3)
	}

keyval_list:
	keyval
	{ //1235
		$$ = []Node{$1}
	}
|	bare_complitexpr
	{
		$$ = []Node{$1}
	}
|	keyval_list ',' keyval
	{ //1243
		$$ = append($1, $3)
	}
|	keyval_list ',' bare_complitexpr
	{ //1247
		$$ = append($1, $3)
	}

braced_keyval_list:
	{ //1252
		$$ = nil
	}
|	keyval_list ocomma
	{ //1256
		$$ = $1
	}

osemi:
|	';'

ocomma:
|	','

oexpr:
	{ //1279
		$$ = nil
	}
|	expr

oexpr_list:
	{ //1288
		$$ = nil
	}
|	expr_list

osimple_stmt:
	{ //1297
		$$ = nil
	}
|	simple_stmt

oliteral:
	{ //1306
		$$ = (*Literal)(nil)
	}
|	_LITERAL
	{
		$$ = newLiteral($1)
	}

%%

func yy(y yyLexer) *parser                   { return y.(*parser) }
func yyErr(y yyLexer, msg string)            { yy(y).Error(msg) }
func yyErrPos(y yyLexer, n Node, msg string) { yy(y).errPos(n.Pos(), msg) }
func yyFset(y yyLexer) *token.FileSet        { return yy(y).fset }
func yyFScope(y yyLexer) *Scope              { return yy(y).pkgScope }
func yyScope(y yyLexer) *Scope               { return yy(y).currentScope }

func yyTLD(y yyLexer, n Node) {
	p := yy(y)
	p.ast = append(p.ast, n)
	switch x := n.(type) {
	case *ConstDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *VarDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *TypeDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *FuncDecl:
		p.pkgScope.declare(p, x.Name.Lit, n)
	case *Package:
		switch ex0, ok := p.pkgScope.Names[dlrPkgName]; {
		case ok:
			ex := ex0.(*Package)
			if ex.Name.Lit == x.Name.Lit {
				break
			}

			p.errPos(n.Pos(), "%s redeclared, previous declaration at %s", x.Name.Lit, p.fset.Position(ex.Pos()))
		default:
			p.pkgScope.declare(p, dlrPkgName, n)
		}
	case *Import:
		// not handled here
	default:
		panic("internal error")
	}
}

func yyTLDs(y yyLexer, l []Node) {
	for _, v := range l {
		yyTLD(y, v)
	}
}
