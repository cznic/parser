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
	_ANDAND _ANDNOT _ASOP _BODY _BREAK _CASE _CHAN _COLAS _COMM _CONST
	_CONTINUE _DDD _DEC _DEFAULT _DEFER _ELSE _EQ _FALL _FOR _FUNC _GE _GO
	_GOTO _GT _IF _IGNORE _IMPORT _INC _INTERFACE _LE _LITERAL _LSH _LT
	_MAP _NAME _NE _OROR _PACKAGE _RANGE _RETURN _RSH _SELECT _STRUCT
	_SWITCH _TYPE _VAR
	'.' '-' '*' '[' '(' '+' '=' '{'

%type	<node>
	bare_complitexpr
	complitexpr comptype compound_stmt constdcl constdcl1
	dcl_name dotname
	embed expr expr_or_type
	fndcl fnret_type fntype
	import_stmt indcl interfacedcl interfacetype
	keyval
	name name_or_type new_name ntype non_dcl_stmt
	oexpr oliteral othertype
	package packname pexpr pexpr_no_paren pseudocall ptrtype
	simple_stmt structdcl structtype sym
	typedcl typedclname
	uexpr
	xfndcl

%type	<list>
	braced_keyval_list
	common_dcl constdcl_list
	dcl_name_list
	expr_list expr_or_type_list
	fnbody
	import_stmt_list interfacedcl_list
	keyval_list
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
%left	_EQ _NE _LE _GE _LT _GT
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
		$$ = &Package{$1.pos, $2.(*Ident)}
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
		$$ = newImport(yylex, &Ident{$1.pos, "."}, newLiteral($2))
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
		panic(".y:215")
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
		$$ = &TypeDecl{pos($1.Pos()), (*Name)($1.(*Ident)), $2}
	}

simple_stmt:
	expr
|	expr _ASOP expr
	{ //240
		$$ = &Assignment{$2.pos, $2.tok, []Node{$1}, []Node{$3}}
	}
|	expr_list '=' expr_list
	{ //244
		$$ = &Assignment{$2.pos, $2.tok, $1, $3}
	}
|	expr_list _COLAS expr_list
	{ //248
		$$ = &ShortVarDecl{$2.pos, $1, $3}
	}
|	expr _INC
	{ //252
		$$ = &IncDecStmt{$2.pos, $1, $2.tok}
	}
|	expr _DEC
	{ //256
		$$ = &IncDecStmt{$2.pos, $1, $2.tok}
	}

case:
	_CASE expr_or_type_list ':'
	{ //262
		panic(".y:263")
	}
|	_CASE expr_or_type_list '=' expr ':'
	{ //266
		panic(".y:267")
	}
|	_CASE expr_or_type_list _COLAS expr ':'
	{ //270
		panic(".y:271")
	}
|	_DEFAULT ':'
	{ //274
		panic(".y:275")
	}

compound_stmt:
	'{' stmt_list '}'
	{ //280
		$$ = &CompoundStament{$1.pos, $2}
	}

caseblock:
	case
	{ //286
		panic(".y:287")
	}
	stmt_list
	{ //290
		panic(".y:291")
	}

caseblock_list:
	{ //295
		panic(".y:296")
	}
|	caseblock_list caseblock
	{ //299
		panic(".y:300")
	}

loop_body:
	_BODY stmt_list '}'
	{ //305
		panic(".y:306")
	}

range_stmt:
	expr_list '=' _RANGE expr
	{ //311
		panic(".y:312")
	}
|	expr_list _COLAS _RANGE expr
	{ //315
		panic(".y:316")
	}

for_header:
	osimple_stmt ';' osimple_stmt ';' osimple_stmt
	{ //321
		panic(".y:322")
	}
|	osimple_stmt
	{ //325
		panic(".y:326")
	}
|	range_stmt
	{ //329
		panic(".y:330")
	}

for_body:
	for_header loop_body
	{ //335
		panic(".y:336")
	}

for_stmt:
	_FOR for_body
	{ //341
		panic(".y:342")
	}

if_header:
	osimple_stmt
	{ //347
		panic(".y:348")
	}
|	osimple_stmt ';' osimple_stmt
	{ //351
		panic(".y:352")
	}

if_stmt:
	_IF if_header loop_body elseif_list else
	{ //357
		panic(".y:358")
	}

elseif:
	_ELSE _IF if_header loop_body
	{ //363
		panic(".y:364")
	}

elseif_list:
	{ //368
		panic(".y:369")
	}
|	elseif_list elseif
	{ //372
		panic(".y:373")
	}

else:
	{ //377
		panic(".y:378")
	}
|	_ELSE compound_stmt
	{ //381
		panic(".y:382")
	}

switch_stmt:
	_SWITCH if_header _BODY caseblock_list '}'
	{ //387
		panic(".y:388")
	}

select_stmt:
	_SELECT _BODY caseblock_list '}'
	{ //393
		panic(".y:394")
	}

expr:
	uexpr
|	expr _OROR expr
	{ //403
		panic(".y:404")
	}
|	expr _ANDAND expr
	{ //407
		panic(".y:408")
	}
|	expr _EQ expr
	{ //411
		panic(".y:412")
	}
|	expr _NE expr
	{ //415
		panic(".y:416")
	}
|	expr _LT expr
	{ //419
		panic(".y:420")
	}
|	expr _LE expr
	{ //423
		panic(".y:424")
	}
|	expr _GE expr
	{ //427
		panic(".y:428")
	}
|	expr _GT expr
	{ //431
		panic(".y:432")
	}
|	expr '+' expr
	{ //435
		$$ = &BinOp{$2.pos, token.ADD, $1, $3}
	}
|	expr '-' expr
	{ //439
		$$ = &BinOp{$2.pos, token.SUB, $1, $3}
	}
|	expr '|' expr
	{ //443
		panic(".y:444")
	}
|	expr '^' expr
	{ //447
		panic(".y:448")
	}
|	expr '*' expr
	{ //451
		$$ = &BinOp{$2.pos, token.MUL, $1, $3}
	}
|	expr '/' expr
	{ //455
		panic(".y:456")
	}
|	expr '%' expr
	{ //459
		panic(".y:460")
	}
|	expr '&' expr
	{ //463
		panic(".y:464")
	}
|	expr _ANDNOT expr
	{ //467
		panic(".y:468")
	}
|	expr _LSH expr
	{ //471
		$$ = &BinOp{$2.pos, token.SHL, $1, $3}
	}
|	expr _RSH expr
	{ //475
		panic(".y:476")
	}
|	expr _COMM expr
	{ //479
		panic(".y:480")
	}

uexpr:
	pexpr
|	'*' uexpr
	{ //489
		$$ = &UnOp{$1.pos, token.MUL, $2}
	}
|	'&' uexpr
	{ //493
		panic(".y:494")
	}
|	'+' uexpr
	{ //497
		panic(".y:498")
	}
|	'-' uexpr
	{ //501
		$$ = &UnOp{$1.pos, token.SUB, $2}
	}
|	'!' uexpr
	{ //505
		panic(".y:506")
	}
|	'~' uexpr
	{ //509
		panic(".y:510")
	}
|	'^' uexpr
	{ //513
		panic(".y:514")
	}
|	_COMM uexpr
	{ //517
		panic(".y:518")
	}

pseudocall:
	pexpr '(' ')'
	{ //523
		$$ = &CallOp{$2.pos, $1, nil}
	}
|	pexpr '(' expr_or_type_list ocomma ')'
	{ //527
		$$ = &CallOp{$2.pos, $1, $3}
	}
|	pexpr '(' expr_or_type_list _DDD ocomma ')'
	{ //531
		panic(".y:532")
	}

pexpr_no_paren:
	_LITERAL
	{ //537
		$$ = &Literal{$1.pos, $1.tok, $1.lit}
	}
|	name
|	pexpr '.' sym
	{ //545
		$$ = &SelectOp{$2.pos, $1, $3.(*Ident)}
	}
|	pexpr '.' '(' expr_or_type ')'
	{ //549
		panic(".y:550")
	}
|	pexpr '.' '(' _TYPE ')'
	{ //553
		panic(".y:554")
	}
|	pexpr '[' expr ']'
	{ //557
		$$ = &IndexOp{$2.pos, $1, $3}
	}
|	pexpr '[' oexpr ':' oexpr ']'
	{ //561
		panic(".y:562")
	}
|	pexpr '[' oexpr ':' oexpr ':' oexpr ']'
	{ //565
		panic(".y:566")
	}
|	pseudocall
|	convtype '(' expr ocomma ')'
	{ //573
		panic(".y:574")
	}
|	comptype lbrace start_complit braced_keyval_list '}'
	{ //577
		$$ = &CompLit{pos($1.Pos()), $1, elements($4)}
	}
|	pexpr_no_paren '{' start_complit braced_keyval_list '}'
	{ //581
		panic(".y:582")
	}
|	'(' expr_or_type ')' '{' start_complit braced_keyval_list '}'
	{ //585
		panic(".y:586")
	}
|	fnliteral
	{ //589
		panic(".y:590")
	}

start_complit:
	{ //594
	}

keyval:
	expr ':' complitexpr
	{ //600
		$$ = &Element{pos($1.Pos()), $1, $3}
	}

bare_complitexpr:
	expr
	{ //609
		$$ = &Element{pos($1.Pos()), nil, $1}
	}
|	'{' start_complit braced_keyval_list '}'
	{ //610
		panic(".y:611")
	}

complitexpr:
	expr
	{ //616
		$$ = &Element{pos($1.Pos()), nil, $1}
	}
|	'{' start_complit braced_keyval_list '}'
	{ //620
		panic(".y:621")
	}

pexpr:
	pexpr_no_paren
|	'(' expr_or_type ')'
	{ //630
		$$ = &Paren{$1.pos, $2}
	}

expr_or_type:
	expr
|	non_expr_type	%prec preferToRightParen
	{ //640
		panic(".y:641")
	}

name_or_type:
	ntype

lbrace:
	_BODY
	{ //652
		panic(".y:653")
	}
|	'{'

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
		panic(".y:674")
	}
|	new_name
	{ //677
		panic(".y:678")
	}

sym:
	_NAME
	{ //683
		$$ = &Ident{$1.pos, $1.lit}
	}

name:
	sym	%prec notParen

labelname:
	new_name
	{ //695
		panic(".y:696")
	}

dotdotdot:
	_DDD
	{ //701
		yy(yylex).errPos($1.tpos, "final argument in variadic function missing type")
		$$ = &Param{pos: $1.pos, Ddd: true}
	}
|	_DDD ntype
	{ //705
		$$ = &Param{pos: $1.pos, Ddd: true, Type: $2}
	}

ntype:
	recvchantype
	{ //711
		panic(".y:712")
	}
|	fntype
|	othertype
|	ptrtype
|	dotname
	{
		$$ = &NamedType{pos($1.Pos()), $1.(*QualifiedIdent), nil, yyScope(yylex)}
	}
|	'(' ntype ')'
	{ //731
		$$ = &Paren{$1.pos, $2}
	}

non_expr_type:
	recvchantype
	{ //737
		panic(".y:738")
	}
|	fntype
	{ //741
		panic(".y:742")
	}
|	othertype
	{ //745
		panic(".y:746")
	}
|	'*' non_expr_type
	{ //749
		panic(".y:750")
	}

non_recvchantype:
	fntype
	{ //755
		panic(".y:756")
	}
|	othertype
	{ //759
		panic(".y:760")
	}
|	ptrtype
	{ //763
		panic(".y:764")
	}
|	dotname
	{ //767
		panic(".y:768")
	}
|	'(' ntype ')'
	{ //771
		panic(".y:772")
	}

convtype:
	fntype
	{ //777
		panic(".y:778")
	}
|	othertype
	{ //781
		panic(".y:782")
	}

comptype:
	othertype

fnret_type:
	recvchantype
	{ //793
		panic(".y:794")
	}
|	fntype
	{ //797
		panic(".y:798")
	}
|	othertype
	{ //801
		panic(".y:802")
	}
|	ptrtype
	{ //805
		panic(".y:806")
	}
|	dotname
	{ //790
		$$ = &NamedType{pos($1.Pos()), $1.(*QualifiedIdent), nil, yyScope(yylex)}
	}

dotname:
	name
	{ //815
		$$ = &QualifiedIdent{pos($1.Pos()), nil, $1.(*Ident)}
	}
|	name '.' sym
	{ //819
		$$ = &QualifiedIdent{pos($1.Pos()), $1.(*Ident), $3.(*Ident)}
	}

othertype:
	'[' oexpr ']' ntype
	{ //825
		switch { //TODO + resolve scope
		case $2 != nil:
			$$ = &ArrayType{$1.pos, $2, $4}
		default:
			$$ = &SliceType{$1.pos, $4}
		}
	}
|	'[' _DDD ']' ntype
	{ //829
		panic(".y:830")
	}
|	_CHAN non_recvchantype
	{ //833
		panic(".y:834")
	}
|	_CHAN _COMM ntype
	{ //837
		panic(".y:838")
	}
|	_MAP '[' ntype ']' ntype
	{ //841
		panic(".y:842")
	}
|	structtype
|	interfacetype

ptrtype:
	'*' ntype
	{ //855
		$$ = &PtrType{$1.pos, $2}
	}

recvchantype:
	_COMM _CHAN ntype
	{ //861
		panic(".y:862")
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
		x.pos = $1.pos
		$$ = x
	}
|	_INTERFACE lbrace '}'
	{ //881
		panic(".y:882")
	}

xfndcl:
	_FUNC fndcl fnbody
	{ //887
		x := $2.(*FuncDecl)
		x.pos, x.Body = $1.pos, $3
		$$ = x
	}

fndcl:
	sym '(' oarg_type_list_ocomma ')' fnres
	{ //893
		$$ = &FuncDecl{Name:(*Name)($1.(*Ident)), Type: newFuncType(yylex, $2.pos, $3, $5)}
	}
|	'(' oarg_type_list_ocomma ')' sym '(' oarg_type_list_ocomma ')' fnres
	{ //897
		panic(".y:898")
	}

fntype:
	_FUNC '(' oarg_type_list_ocomma ')' fnres
	{ //903
		$$ = newFuncType(yylex, $1.pos, $3, $5)
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
		$$ = []*Param{{pos: pos($1.Pos()), Type: $1}}
	}
|	'(' oarg_type_list_ocomma ')'
	{ //926
		$$ = $2
	}

fnlitdcl:
	fntype
	{ //932
		panic(".y:933")
	}

fnliteral:
	fnlitdcl lbrace stmt_list '}'
	{ //938
		panic(".y:939")
	}
|	fnlitdcl error
	{ //942
		panic(".y:943")
	}

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
		$$ = newFields([]Node{q.I}, true, &NamedType{q.pos, q, nil, yyScope(yylex)}, $2)
	}
|	'(' embed ')' oliteral
	{ //1015
		panic(".y:1016")
	}
|	'*' embed oliteral
	{ //1019
		q := $2.(*QualifiedIdent)
		$$ = newFields([]Node{q.I}, true, &PtrType{$1.pos, &NamedType{q.pos, q, nil, yyScope(yylex)}}, $3)
	}
|	'(' '*' embed ')' oliteral
	{ //1023
		panic(".y:1024")
	}
|	'*' '(' embed ')' oliteral
	{ //1027
		panic(".y:1028")
	}

packname:
	_NAME
	{ //1033
		$$ = &QualifiedIdent{$1.pos, nil, &Ident{$1.pos, $1.lit}}
	}
|	_NAME '.' sym
	{ //1037
		$$ = &QualifiedIdent{$1.pos, &Ident{$1.pos, $1.lit}, $3.(*Ident)}
	}

embed:
	packname

interfacedcl:
	new_name indcl
	{ //1049
		$$ = &MethodSpec{pos($1.Pos()), $1.(*Ident), $2.(*FuncType)}
	}
|	packname
	{ //1053
		panic(".y:1054")
	}
|	'(' packname ')'
	{ //1057
		panic(".y:1058")
		//yyErrPos(yylex, $2, "cannot parenthesize embedded type");
	}

indcl:
	'(' oarg_type_list_ocomma ')' fnres
	{ //1063
		$$ = newFuncType(yylex, $1.pos, $2, $4)
	}

arg_type:
	name_or_type
	{ //1069
		$$ = &Param{pos: pos($1.Pos()), Type: $1}
	}
|	sym name_or_type
	{ //1073
		$$ = &Param{pos: pos($1.Pos()), Name: $1.(*Ident), Type: $2}
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
	{ //1131
		panic(".y:1132")
	}
|	switch_stmt
	{ //1135
		panic(".y:1136")
	}
|	select_stmt
	{ //1139
		panic(".y:1140")
	}
|	if_stmt
	{ //1143
		panic(".y:1144")
	}
|	labelname ':'
	{ //1147
		panic(".y:1148")
	}
	stmt
	{ //1151
		panic(".y:1152")
	}
|	_FALL
	{ //1155
		panic(".y:1156")
	}
|	_BREAK onew_name
	{ //1159
		panic(".y:1160")
	}
|	_CONTINUE onew_name
	{ //1163
		panic(".y:1164")
	}
|	_GO pseudocall
	{ //1167
		panic(".y:1168")
	}
|	_DEFER pseudocall
	{ //1171
		panic(".y:1172")
	}
|	_GOTO new_name
	{ //1175
		panic(".y:1176")
	}
|	_RETURN oexpr_list
	{ //1179
		$$ = &ReturnStmt{$1.pos, $2}
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
		panic(".y:1230")
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
		panic(".y:1298")
	}
|	simple_stmt
	{ //1301
		panic(".y:1302")
	}

oliteral:
	{ //1306
		$$ = (*Literal)(nil)
	}
|	_LITERAL
	{ //1310
		panic(".y:1311")
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
	if d, ok := n.(Declaration); ok {
		p.pkgScope.declare(p, d.DeclName(), n)
	}
}

func yyTLDs(y yyLexer, l []Node) {
	p := yy(y)
	for _, v := range l {
		if d, ok := v.(Declaration); ok {
			p.pkgScope.declare(p, d.DeclName(), v)
		}
	}
	p.ast = append(p.ast, l...)
}
