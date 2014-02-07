// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

%{

package parser

%}

%union	{
	TODO int //TODO
}

%token	LLITERAL
%token  LASOP LCOLAS
%token	LBREAK LCASE LCHAN LCONST LCONTINUE LDDD
%token	LDEFAULT LDEFER LELSE LFALL LFOR LFUNC LGO LGOTO
%token	LIF LIMPORT LINTERFACE LMAP LNAME
%token	LPACKAGE LRANGE LRETURN LSELECT LSTRUCT LSWITCH
%token	LTYPE LVAR

%token	LANDAND LANDNOT LBODY LCOMM LDEC LEQ LGE LGT
%token	LIGNORE LINC LLE LLSH LLT LNE LOROR LRSH

%left		LCOMM

%left		LOROR
%left		LANDAND
%left		LEQ LNE LLE LGE LLT LGT
%left		'+' '-' '|' '^'
%left		'*' '/' '%' '&' LLSH LRSH LANDNOT

%left		NotPackage
%left		LPACKAGE

%left		NotParen
%left		'('

%left		')'
%left		PreferToRightParen

%%
file:
	package
	{
	}
	imports
	{
	}
	xdcl_list
	{
	}

package:
	%prec NotPackage
	{
	}
|	LPACKAGE sym ';'
	{
	}

imports:
	{
	}
|	imports import ';'
	{
	}

import:
	LIMPORT import_stmt
	{
	}
|	LIMPORT '(' import_stmt_list osemi ')'
	{
	}
|	LIMPORT '(' ')'
	{
	}

import_stmt:
	LLITERAL
	{
	}
|	sym LLITERAL
	{
	}
|	'.' LLITERAL
	{
	}

import_stmt_list:
	import_stmt
	{
	}
|	import_stmt_list ';' import_stmt
	{
	}

xdcl:
	{
	}
|	common_dcl
	{
	}
|	xfndcl
	{
	}
|	non_dcl_stmt
	{
	}
|	error
	{
	}

common_dcl:
	LVAR vardcl
	{
	}
|	LVAR '(' vardcl_list osemi ')'
	{
	}
|	LVAR '(' ')'
	{
	}
|	lconst constdcl
	{
	}
|	lconst '(' constdcl osemi ')'
	{
	}
|	lconst '(' constdcl ';' constdcl_list osemi ')'
	{
	}
|	lconst '(' ')'
	{
	}
|	LTYPE typedcl
	{
	}
|	LTYPE '(' typedcl_list osemi ')'
	{
	}
|	LTYPE '(' ')'
	{
	}

lconst:
	LCONST
	{
	}

vardcl:
	dcl_name_list ntype
	{
	}
|	dcl_name_list ntype '=' expr_list
	{
	}
|	dcl_name_list '=' expr_list
	{
	}

constdcl:
	dcl_name_list ntype '=' expr_list
	{
	}
|	dcl_name_list '=' expr_list
	{
	}

constdcl1:
	constdcl
	{
	}
|	dcl_name_list ntype
	{
	}
|	dcl_name_list
	{
	}

typedclname:
	sym
	{
	}

typedcl:
	typedclname ntype
	{
	}

simple_stmt:
	expr
	{
	}
|	expr LASOP expr
	{
	}
|	expr_list '=' expr_list
	{
	}
|	expr_list LCOLAS expr_list
	{
	}
|	expr LINC
	{
	}
|	expr LDEC
	{
	}

case:
	LCASE expr_or_type_list ':'
	{
	}
|	LCASE expr_or_type_list '=' expr ':'
	{
	}
|	LCASE expr_or_type_list LCOLAS expr ':'
	{
	}
|	LDEFAULT ':'
	{
	}

compound_stmt:
	'{' stmt_list '}'
	{
	}

caseblock:
	case
	{
	}
	stmt_list
	{
	}

caseblock_list:
	{
	}
|	caseblock_list caseblock
	{
	}

loop_body:
	LBODY stmt_list '}'
	{
	}

range_stmt:
	expr_list '=' LRANGE expr
	{
	}
|	expr_list LCOLAS LRANGE expr
	{
	}

for_header:
	osimple_stmt ';' osimple_stmt ';' osimple_stmt
	{
	}
|	osimple_stmt
	{
	}
|	range_stmt
	{
	}

for_body:
	for_header loop_body
	{
	}

for_stmt:
	LFOR for_body
	{
	}

if_header:
	osimple_stmt
	{
	}
|	osimple_stmt ';' osimple_stmt
	{
	}

if_stmt:
	LIF if_header loop_body elseif_list else
	{
	}

elseif:
	LELSE LIF if_header loop_body
	{
	}

elseif_list:
	{
	}
|	elseif_list elseif
	{
	}

else:
	{
	}
|	LELSE compound_stmt
	{
	}

switch_stmt:
	LSWITCH if_header LBODY caseblock_list '}'
	{
	}

select_stmt:
	LSELECT LBODY caseblock_list '}'
	{
	}

expr:
	uexpr
	{
	}
|	expr LOROR expr
	{
	}
|	expr LANDAND expr
	{
	}
|	expr LEQ expr
	{
	}
|	expr LNE expr
	{
	}
|	expr LLT expr
	{
	}
|	expr LLE expr
	{
	}
|	expr LGE expr
	{
	}
|	expr LGT expr
	{
	}
|	expr '+' expr
	{
	}
|	expr '-' expr
	{
	}
|	expr '|' expr
	{
	}
|	expr '^' expr
	{
	}
|	expr '*' expr
	{
	}
|	expr '/' expr
	{
	}
|	expr '%' expr
	{
	}
|	expr '&' expr
	{
	}
|	expr LANDNOT expr
	{
	}
|	expr LLSH expr
	{
	}
|	expr LRSH expr
	{
	}
|	expr LCOMM expr
	{
	}

uexpr:
	pexpr
	{
	}
|	'*' uexpr
	{
	}
|	'&' uexpr
	{
	}
|	'+' uexpr
	{
	}
|	'-' uexpr
	{
	}
|	'!' uexpr
	{
	}
|	'~' uexpr
	{
	}
|	'^' uexpr
	{
	}
|	LCOMM uexpr
	{
	}

pseudocall:
	pexpr '(' ')'
	{
	}
|	pexpr '(' expr_or_type_list ocomma ')'
	{
	}
|	pexpr '(' expr_or_type_list LDDD ocomma ')'
	{
	}

pexpr_no_paren:
	LLITERAL
	{
	}
|	name
	{
	}
|	pexpr '.' sym
	{
	}
|	pexpr '.' '(' expr_or_type ')'
	{
	}
|	pexpr '.' '(' LTYPE ')'
	{
	}
|	pexpr '[' expr ']'
	{
	}
|	pexpr '[' oexpr ':' oexpr ']'
	{
	}
|	pexpr '[' oexpr ':' oexpr ':' oexpr ']'
	{
	}
|	pseudocall
	{
	}
|	convtype '(' expr ocomma ')'
	{
	}
|	comptype lbrace start_complit braced_keyval_list '}'
	{
	}
|	pexpr_no_paren '{' start_complit braced_keyval_list '}'
	{
	}
|	'(' expr_or_type ')' '{' start_complit braced_keyval_list '}'
	{
	}
|	fnliteral
	{
	}

start_complit:
	{
	}

keyval:
	expr ':' complitexpr
	{
	}

bare_complitexpr:
	expr
	{
	}
|	'{' start_complit braced_keyval_list '}'
	{
	}

complitexpr:
	expr
	{
	}
|	'{' start_complit braced_keyval_list '}'
	{
	}

pexpr:
	pexpr_no_paren
	{
	}
|	'(' expr_or_type ')'
	{
	}

expr_or_type:
	expr
	{
	}
|	non_expr_type	%prec PreferToRightParen
	{
	}

name_or_type:
	ntype
	{
	}

lbrace:
	LBODY
	{
	}
|	'{'
	{
	}

new_name:
	sym
	{
	}

dcl_name:
	sym
	{
	}

onew_name:
	{
	}
|	new_name
	{
	}

sym:
	LNAME
	{
	}

name:
	sym	%prec NotParen
	{
	}

labelname:
	new_name
	{
	}

dotdotdot:
	LDDD
	{
	}
|	LDDD ntype
	{
	}

ntype:
	recvchantype
	{
	}
|	fntype
	{
	}
|	othertype
	{
	}
|	ptrtype
	{
	}
|	dotname
	{
	}
|	'(' ntype ')'
	{
	}

non_expr_type:
	recvchantype
	{
	}
|	fntype
	{
	}
|	othertype
	{
	}
|	'*' non_expr_type
	{
	}

non_recvchantype:
	fntype
	{
	}
|	othertype
	{
	}
|	ptrtype
	{
	}
|	dotname
	{
	}
|	'(' ntype ')'
	{
	}

convtype:
	fntype
	{
	}
|	othertype
	{
	}

comptype:
	othertype
	{
	}

fnret_type:
	recvchantype
	{
	}
|	fntype
	{
	}
|	othertype
	{
	}
|	ptrtype
	{
	}
|	dotname
	{
	}

dotname:
	name
	{
	}
|	name '.' sym
	{
	}

othertype:
	'[' oexpr ']' ntype
	{
	}
|	'[' LDDD ']' ntype
	{
	}
|	LCHAN non_recvchantype
	{
	}
|	LCHAN LCOMM ntype
	{
	}
|	LMAP '[' ntype ']' ntype
	{
	}
|	structtype
	{
	}
|	interfacetype
	{
	}

ptrtype:
	'*' ntype
	{
	}

recvchantype:
	LCOMM LCHAN ntype
	{
	}

structtype:
	LSTRUCT lbrace structdcl_list osemi '}'
	{
	}
|	LSTRUCT lbrace '}'
	{
	}

interfacetype:
	LINTERFACE lbrace interfacedcl_list osemi '}'
	{
	}
|	LINTERFACE lbrace '}'
	{
	}

xfndcl:
	LFUNC fndcl fnbody
	{
	}

fndcl:
	sym '(' oarg_type_list_ocomma ')' fnres
	{
	}
|	'(' oarg_type_list_ocomma ')' sym '(' oarg_type_list_ocomma ')' fnres
	{
	}

fntype:
	LFUNC '(' oarg_type_list_ocomma ')' fnres
	{
	}

fnbody:
	{
	}
|	'{' stmt_list '}'
	{
	}

fnres:
	%prec NotParen
	{
	}
|	fnret_type
	{
	}
|	'(' oarg_type_list_ocomma ')'
	{
	}

fnlitdcl:
	fntype
	{
	}

fnliteral:
	fnlitdcl lbrace stmt_list '}'
	{
	}
|	fnlitdcl error
	{
	}

xdcl_list:
	{
	}
|	xdcl_list xdcl ';'
	{
	}

vardcl_list:
	vardcl
	{
	}
|	vardcl_list ';' vardcl
	{
	}

constdcl_list:
	constdcl1
	{
	}
|	constdcl_list ';' constdcl1
	{
	}

typedcl_list:
	typedcl
	{
	}
|	typedcl_list ';' typedcl
	{
	}

structdcl_list:
	structdcl
	{
	}
|	structdcl_list ';' structdcl
	{
	}

interfacedcl_list:
	interfacedcl
	{
	}
|	interfacedcl_list ';' interfacedcl
	{
	}

structdcl:
	new_name_list ntype oliteral
	{
	}
|	embed oliteral
	{
	}
|	'(' embed ')' oliteral
	{
	}
|	'*' embed oliteral
	{
	}
|	'(' '*' embed ')' oliteral
	{
	}
|	'*' '(' embed ')' oliteral
	{
	}

packname:
	LNAME
	{
	}
|	LNAME '.' sym
	{
	}

embed:
	packname
	{
	}

interfacedcl:
	new_name indcl
	{
	}
|	packname
	{
	}
|	'(' packname ')'
	{
	}

indcl:
	'(' oarg_type_list_ocomma ')' fnres
	{
	}

arg_type:
	name_or_type
	{
	}
|	sym name_or_type
	{
	}
|	sym dotdotdot
	{
	}
|	dotdotdot
	{
	}

arg_type_list:
	arg_type
	{
	}
|	arg_type_list ',' arg_type
	{
	}

oarg_type_list_ocomma:
	{
	}
|	arg_type_list ocomma
	{
	}

stmt:
	{
	}
|	compound_stmt
	{
	}
|	common_dcl
	{
	}
|	non_dcl_stmt
	{
	}
|	error
	{
	}

non_dcl_stmt:
	simple_stmt
	{
	}
|	for_stmt
	{
	}
|	switch_stmt
	{
	}
|	select_stmt
	{
	}
|	if_stmt
	{
	}
|	labelname ':'
	{
	}
	stmt
	{
	}
|	LFALL
	{
	}
|	LBREAK onew_name
	{
	}
|	LCONTINUE onew_name
	{
	}
|	LGO pseudocall
	{
	}
|	LDEFER pseudocall
	{
	}
|	LGOTO new_name
	{
	}
|	LRETURN oexpr_list
	{
	}

stmt_list:
	stmt
	{
	}
|	stmt_list ';' stmt
	{
	}

new_name_list:
	new_name
	{
	}
|	new_name_list ',' new_name
	{
	}

dcl_name_list:
	dcl_name
	{
	}
|	dcl_name_list ',' dcl_name
	{
	}

expr_list:
	expr
	{
	}
|	expr_list ',' expr
	{
	}

expr_or_type_list:
	expr_or_type
	{
	}
|	expr_or_type_list ',' expr_or_type
	{
	}

keyval_list:
	keyval
	{
	}
|	bare_complitexpr
	{
	}
|	keyval_list ',' keyval
	{
	}
|	keyval_list ',' bare_complitexpr
	{
	}

braced_keyval_list:
	{
	}
|	keyval_list ocomma
	{
	}

osemi:
	{
	}
|	';'
	{
	}

ocomma:
	{
	}
|	','
	{
	}

oexpr:
	{
	}
|	expr
	{
	}

oexpr_list:
	{
	}
|	expr_list
	{
	}

osimple_stmt:
	{
	}
|	simple_stmt
	{
	}

oliteral:
	{
	}
|	LLITERAL
	{
	}

%%
