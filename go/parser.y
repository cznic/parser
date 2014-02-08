%{

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

%}

%union	{
	tk   tk
	node Node
}

%token	<tk>
	_ANDAND _ANDNOT _ASOP _BODY _BREAK _CASE _CHAN _COLAS _COMM _CONST
	_CONTINUE _DDD _DEC _DEFAULT _DEFER _ELSE _EQ _FALL _FOR _FUNC _GE _GO
	_GOTO _GT _IF _IGNORE _IMPORT _INC _INTERFACE _LE _LITERAL _LSH _LT
	_MAP _NAME _NE _OROR _PACKAGE _RANGE _RETURN _RSH _SELECT _STRUCT
	_SWITCH _TYPE _VAR

%type	<node>
	package
	sym

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
	tlds

package:
	%prec notPackage
	{ //60
		panic(".y:61")
	}
|	_PACKAGE sym ';'
	{ //64
		$$ = &Package{$1.pos, $2.(*Ident)}
	}

imports:
|	imports import ';'
	{ //73
		panic(".y:74")
	}

import:
	_IMPORT import_stmt
	{ //79
		panic(".y:80")
	}
|	_IMPORT '(' import_stmt_list osemi ')'
	{ //83
		panic(".y:84")
	}
|	_IMPORT '(' ')'
	{ //87
		panic(".y:88")
	}

import_stmt:
	_LITERAL
	{ //93
		panic(".y:94")
	}
|	sym _LITERAL
	{ //97
		panic(".y:98")
	}
|	'.' _LITERAL
	{ //101
		panic(".y:102")
	}

import_stmt_list:
	import_stmt
	{ //107
		panic(".y:108")
	}
|	import_stmt_list ';' import_stmt
	{ //111
		panic(".y:112")
	}

tld:
	{ //116
		panic(".y:117")
	}
|	common_dcl
	{ //120
		panic(".y:121")
	}
|	xfndcl
	{ //124
		panic(".y:125")
	}
|	non_dcl_stmt
	{ //128
		panic(".y:129")
	}
|	error
	{ //132
		panic(".y:133")
	}

common_dcl:
	_VAR vardcl
	{ //138
		panic(".y:139")
	}
|	_VAR '(' vardcl_list osemi ')'
	{ //142
		panic(".y:143")
	}
|	_VAR '(' ')'
	{ //146
		panic(".y:147")
	}
|	lconst constdcl
	{ //150
		panic(".y:151")
	}
|	lconst '(' constdcl osemi ')'
	{ //154
		panic(".y:155")
	}
|	lconst '(' constdcl ';' constdcl_list osemi ')'
	{ //158
		panic(".y:159")
	}
|	lconst '(' ')'
	{ //162
		panic(".y:163")
	}
|	_TYPE typedcl
	{ //166
		panic(".y:167")
	}
|	_TYPE '(' typedcl_list osemi ')'
	{ //170
		panic(".y:171")
	}
|	_TYPE '(' ')'
	{ //174
		panic(".y:175")
	}

lconst:
	_CONST
	{ //180
		panic(".y:181")
	}

vardcl:
	dcl_name_list ntype
	{ //186
		panic(".y:187")
	}
|	dcl_name_list ntype '=' expr_list
	{ //190
		panic(".y:191")
	}
|	dcl_name_list '=' expr_list
	{ //194
		panic(".y:195")
	}

constdcl:
	dcl_name_list ntype '=' expr_list
	{ //200
		panic(".y:201")
	}
|	dcl_name_list '=' expr_list
	{ //204
		panic(".y:205")
	}

constdcl1:
	constdcl
	{ //210
		panic(".y:211")
	}
|	dcl_name_list ntype
	{ //214
		panic(".y:215")
	}
|	dcl_name_list
	{ //218
		panic(".y:219")
	}

typedclname:
	sym
	{ //224
		panic(".y:225")
	}

typedcl:
	typedclname ntype
	{ //230
		panic(".y:231")
	}

simple_stmt:
	expr
	{ //236
		panic(".y:237")
	}
|	expr _ASOP expr
	{ //240
		panic(".y:241")
	}
|	expr_list '=' expr_list
	{ //244
		panic(".y:245")
	}
|	expr_list _COLAS expr_list
	{ //248
		panic(".y:249")
	}
|	expr _INC
	{ //252
		panic(".y:253")
	}
|	expr _DEC
	{ //256
		panic(".y:257")
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
		panic(".y:281")
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
	{ //399
		panic(".y:400")
	}
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
		panic(".y:436")
	}
|	expr '-' expr
	{ //439
		panic(".y:440")
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
		panic(".y:452")
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
		panic(".y:472")
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
	{ //485
		panic(".y:486")
	}
|	'*' uexpr
	{ //489
		panic(".y:490")
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
		panic(".y:502")
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
		panic(".y:524")
	}
|	pexpr '(' expr_or_type_list ocomma ')'
	{ //527
		panic(".y:528")
	}
|	pexpr '(' expr_or_type_list _DDD ocomma ')'
	{ //531
		panic(".y:532")
	}

pexpr_no_paren:
	_LITERAL
	{ //537
		panic(".y:538")
	}
|	name
	{ //541
		panic(".y:542")
	}
|	pexpr '.' sym
	{ //545
		panic(".y:546")
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
		panic(".y:558")
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
	{ //569
		panic(".y:570")
	}
|	convtype '(' expr ocomma ')'
	{ //573
		panic(".y:574")
	}
|	comptype lbrace start_complit braced_keyval_list '}'
	{ //577
		panic(".y:578")
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
		panic(".y:595")
	}

keyval:
	expr ':' complitexpr
	{ //600
		panic(".y:601")
	}

bare_complitexpr:
	expr
	{ //606
		panic(".y:607")
	}
|	'{' start_complit braced_keyval_list '}'
	{ //610
		panic(".y:611")
	}

complitexpr:
	expr
	{ //616
		panic(".y:617")
	}
|	'{' start_complit braced_keyval_list '}'
	{ //620
		panic(".y:621")
	}

pexpr:
	pexpr_no_paren
	{ //626
		panic(".y:627")
	}
|	'(' expr_or_type ')'
	{ //630
		panic(".y:631")
	}

expr_or_type:
	expr
	{ //636
		panic(".y:637")
	}
|	non_expr_type	%prec preferToRightParen
	{ //640
		panic(".y:641")
	}

name_or_type:
	ntype
	{ //646
		panic(".y:647")
	}

lbrace:
	_BODY
	{ //652
		panic(".y:653")
	}
|	'{'
	{ //656
		panic(".y:657")
	}

new_name:
	sym
	{ //662
		panic(".y:663")
	}

dcl_name:
	sym
	{ //668
		panic(".y:669")
	}

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
		$$ = &Ident{$1.pos, $1.lit, -1}
	}

name:
	sym	%prec notParen
	{ //689
		panic(".y:690")
	}

labelname:
	new_name
	{ //695
		panic(".y:696")
	}

dotdotdot:
	_DDD
	{ //701
		panic(".y:702")
	}
|	_DDD ntype
	{ //705
		panic(".y:706")
	}

ntype:
	recvchantype
	{ //711
		panic(".y:712")
	}
|	fntype
	{ //715
		panic(".y:716")
	}
|	othertype
	{ //719
		panic(".y:720")
	}
|	ptrtype
	{ //723
		panic(".y:724")
	}
|	dotname
	{ //727
		panic(".y:728")
	}
|	'(' ntype ')'
	{ //731
		panic(".y:732")
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
	{ //787
		panic(".y:788")
	}

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
	{ //809
		panic(".y:810")
	}

dotname:
	name
	{ //815
		panic(".y:816")
	}
|	name '.' sym
	{ //819
		panic(".y:820")
	}

othertype:
	'[' oexpr ']' ntype
	{ //825
		panic(".y:826")
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
	{ //845
		panic(".y:846")
	}
|	interfacetype
	{ //849
		panic(".y:850")
	}

ptrtype:
	'*' ntype
	{ //855
		panic(".y:856")
	}

recvchantype:
	_COMM _CHAN ntype
	{ //861
		panic(".y:862")
	}

structtype:
	_STRUCT lbrace structdcl_list osemi '}'
	{ //867
		panic(".y:868")
	}
|	_STRUCT lbrace '}'
	{ //871
		panic(".y:872")
	}

interfacetype:
	_INTERFACE lbrace interfacedcl_list osemi '}'
	{ //877
		panic(".y:878")
	}
|	_INTERFACE lbrace '}'
	{ //881
		panic(".y:882")
	}

xfndcl:
	_FUNC fndcl fnbody
	{ //887
		panic(".y:888")
	}

fndcl:
	sym '(' oarg_type_list_ocomma ')' fnres
	{ //893
		panic(".y:894")
	}
|	'(' oarg_type_list_ocomma ')' sym '(' oarg_type_list_ocomma ')' fnres
	{ //897
		panic(".y:898")
	}

fntype:
	_FUNC '(' oarg_type_list_ocomma ')' fnres
	{ //903
		panic(".y:904")
	}

fnbody:
	{ //908
		panic(".y:909")
	}
|	'{' stmt_list '}'
	{ //912
		panic(".y:913")
	}

fnres:
	%prec notParen
	{ //918
		panic(".y:919")
	}
|	fnret_type
	{ //922
		panic(".y:923")
	}
|	'(' oarg_type_list_ocomma ')'
	{ //926
		panic(".y:927")
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

tlds:
|	tlds tld ';'
	{ //951
		panic(".y:952")
	}

vardcl_list:
	vardcl
	{ //957
		panic(".y:958")
	}
|	vardcl_list ';' vardcl
	{ //961
		panic(".y:962")
	}

constdcl_list:
	constdcl1
	{ //967
		panic(".y:968")
	}
|	constdcl_list ';' constdcl1
	{ //971
		panic(".y:972")
	}

typedcl_list:
	typedcl
	{ //977
		panic(".y:978")
	}
|	typedcl_list ';' typedcl
	{ //981
		panic(".y:982")
	}

structdcl_list:
	structdcl
	{ //987
		panic(".y:988")
	}
|	structdcl_list ';' structdcl
	{ //991
		panic(".y:992")
	}

interfacedcl_list:
	interfacedcl
	{ //997
		panic(".y:998")
	}
|	interfacedcl_list ';' interfacedcl
	{ //1001
		panic(".y:1002")
	}

structdcl:
	new_name_list ntype oliteral
	{ //1007
		panic(".y:1008")
	}
|	embed oliteral
	{ //1011
		panic(".y:1012")
	}
|	'(' embed ')' oliteral
	{ //1015
		panic(".y:1016")
	}
|	'*' embed oliteral
	{ //1019
		panic(".y:1020")
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
		panic(".y:1034")
	}
|	_NAME '.' sym
	{ //1037
		panic(".y:1038")
	}

embed:
	packname
	{ //1043
		panic(".y:1044")
	}

interfacedcl:
	new_name indcl
	{ //1049
		panic(".y:1050")
	}
|	packname
	{ //1053
		panic(".y:1054")
	}
|	'(' packname ')'
	{ //1057
		panic(".y:1058")
	}

indcl:
	'(' oarg_type_list_ocomma ')' fnres
	{ //1063
		panic(".y:1064")
	}

arg_type:
	name_or_type
	{ //1069
		panic(".y:1070")
	}
|	sym name_or_type
	{ //1073
		panic(".y:1074")
	}
|	sym dotdotdot
	{ //1077
		panic(".y:1078")
	}
|	dotdotdot
	{ //1081
		panic(".y:1082")
	}

arg_type_list:
	arg_type
	{ //1087
		panic(".y:1088")
	}
|	arg_type_list ',' arg_type
	{ //1091
		panic(".y:1092")
	}

oarg_type_list_ocomma:
	{ //1096
		panic(".y:1097")
	}
|	arg_type_list ocomma
	{ //1100
		panic(".y:1101")
	}

stmt:
	{ //1105
		panic(".y:1106")
	}
|	compound_stmt
	{ //1109
		panic(".y:1110")
	}
|	common_dcl
	{ //1113
		panic(".y:1114")
	}
|	non_dcl_stmt
	{ //1117
		panic(".y:1118")
	}
|	error
	{ //1121
		panic(".y:1122")
	}

non_dcl_stmt:
	simple_stmt
	{ //1127
		panic(".y:1128")
	}
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
		panic(".y:1180")
	}

stmt_list:
	stmt
	{ //1185
		panic(".y:1186")
	}
|	stmt_list ';' stmt
	{ //1189
		panic(".y:1190")
	}

new_name_list:
	new_name
	{ //1195
		panic(".y:1196")
	}
|	new_name_list ',' new_name
	{ //1199
		panic(".y:1200")
	}

dcl_name_list:
	dcl_name
	{ //1205
		panic(".y:1206")
	}
|	dcl_name_list ',' dcl_name
	{ //1209
		panic(".y:1210")
	}

expr_list:
	expr
	{ //1215
		panic(".y:1216")
	}
|	expr_list ',' expr
	{ //1219
		panic(".y:1220")
	}

expr_or_type_list:
	expr_or_type
	{ //1225
		panic(".y:1226")
	}
|	expr_or_type_list ',' expr_or_type
	{ //1229
		panic(".y:1230")
	}

keyval_list:
	keyval
	{ //1235
		panic(".y:1236")
	}
|	bare_complitexpr
	{ //1239
		panic(".y:1240")
	}
|	keyval_list ',' keyval
	{ //1243
		panic(".y:1244")
	}
|	keyval_list ',' bare_complitexpr
	{ //1247
		panic(".y:1248")
	}

braced_keyval_list:
	{ //1252
		panic(".y:1253")
	}
|	keyval_list ocomma
	{ //1256
		panic(".y:1257")
	}

osemi:
	{ //1261
		panic(".y:1262")
	}
|	';'
	{ //1265
		panic(".y:1266")
	}

ocomma:
	{ //1270
		panic(".y:1271")
	}
|	','
	{ //1274
		panic(".y:1275")
	}

oexpr:
	{ //1279
		panic(".y:1280")
	}
|	expr
	{ //1283
		panic(".y:1284")
	}

oexpr_list:
	{ //1288
		panic(".y:1289")
	}
|	expr_list
	{ //1292
		panic(".y:1293")
	}

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
		panic(".y:1307")
	}
|	_LITERAL
	{ //1310
		panic(".y:1311")
	}

%%

func yy(y yyLexer) *parser { return y.(*parser) }

func yyTLD(y yyLexer, n Node) {
	p := yy(y)
	p.ast = append(p.ast, n)
}
