%{

//TODO Put your favorite license here
		
// yacc source generated by ebnf2y[1]
// at 2013-08-23 17:07:36.910561255 +0200 CEST.
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
// 
//   [1]: http://github.com/cznic/ebnf2y

package parser //TODO real package name

//TODO required only be the demo _dump function
import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cznic/strutil"
)

%}

%union {
	val  interface{}
	pos  pos
	item interface{} //TODO insert real field(s)
}

%token	ANDAND
%token	ANDNOT
%token	ASSIGN_OP
%token	CHANCOMM
%token	COMM
%token	COMMCHAN
%token	DDD
%token	DEC
%token	EQ
%token	FLOAT_LIT
%token	GE
%token	IDENTIFIER
%token	IDENTIFIER_LIST
%token	IDLIST_COLAS
%token	IMAGINARY_LIT
%token	INC
%token	INT_LIT
%token	LBR
%token	LE
%token	LSH
%token	NE
%token	OROR
%token	RSH
%token	RUNE_LIT
%token	STRING_LIT

%type	<item> 	/*TODO real type(s), if/where applicable */
	ANDAND
	ANDNOT
	ASSIGN_OP
	CHANCOMM
	COMM
	COMMCHAN
	DDD
	DEC
	EQ
	FLOAT_LIT
	GE
	IDENTIFIER
	IDENTIFIER_LIST
	IDLIST_COLAS
	IMAGINARY_LIT
	INC
	INT_LIT
	LBR
	LE
	LSH
	NE
	OROR
	RSH
	RUNE_LIT
	STRING_LIT

%token BREAK
%token CASE
%token CHAN
%token CONST
%token CONTINUE
%token DEFAULT
%token DEFER
%token ELSE
%token FALLTHROUGH
%token FOR
%token FUNC
%token GO
%token GOTO
%token IF
%token IMPORT
%token INTERFACE
%token MAP
%token PACKAGE
%token RANGE
%token RETURN
%token SELECT
%token STRUCT
%token SWITCH
%token TYPE
%token VAR

%type	<item> 	/*TODO real type(s), if/where applicable */
	ArgumentList
	ArrayType
	BaseTypeName
	Block
	Condition
	ConstDecl111
	ConstSpec
	Declaration
	Element
	ElementList1
	ElementType
	ExprSwitchStmt3
	Expression
	ExpressionList
	ExpressionList1
	FieldDecl
	FieldDecl1
	Function
	IfStmt
	ImportDecl111
	ImportSpec
	InterfaceType11
	Label
	LiteralValue
	MapType
	MethodName
	MethodSpec
	Name
	PackageName
	ParameterDecl
	ParameterDecl2
	ParameterList1
	Parameters
	PrimaryExpr
	Receiver
	SelectStmt1
	SendStmt
	Signature
	SimpleStmt
	Slice2
	SliceType
	SourceFile
	SourceFile1
	SourceFile2
	Start
	Statement
	StatementList
	StatementList1
	StructType
	StructType11
	Type
	TypeDecl111
	TypeList1
	TypeLit
	TypeNoName
	TypeSpec
	TypeSwitchGuard
	TypeSwitchStmt2
	UnaryExpr
	Value
	VarDecl111
	VarSpec

/*TODO %left, %right, ... declarations */

%left	OROR
%left	ANDAND
%left	'<' '>' EQ GE LE NE
%left	'+' '-' '^' '|'
%left	'%' '&' '*' '/' ANDNOT LSH RSH

%left	notDot // Name
%left	'.'

%left	notPackage
%left	PACKAGE

%start Start

%%

ArgumentList:
	ExpressionList
	{
		$$ = $1 //TODO 1
	}
|	ExpressionList DDD
	{
		$$ = []ArgumentList{$1, $2} //TODO 2
	}

ArrayType:
	'[' Expression ']' ElementType
	{
		$$ = []ArrayType{"[", $2, "]", $4} //TODO 3
	}

BaseTypeName:
	IDENTIFIER
	{
		$$ = $1 //TODO 4
	}

Block:
	'{' StatementList '}'
	{
		$$ = []Block{"{", $2, "}"} //TODO 5
	}

Condition:
	Expression
	{
		$$ = $1 //TODO 6
	}

ConstDecl111:
	/* EMPTY */
	{
		$$ = []ConstDecl111(nil) //TODO 7
	}
|	ConstDecl111 ';'
	{
		lx := yylex.(*lx)
		lx.toks, lx.state = nil, st2 //TODO named state alias
	}
	ConstSpec
	{
		//TODO $$ = append($1.([]ConstDecl111), ";", $3) //TODO 8
	}

ConstSpec:
	IDENTIFIER_LIST
	{
		$$ = $1 //TODO 9
	}
|	IDENTIFIER_LIST '=' ExpressionList
	{
		$$ = []ConstSpec{$1, "=", $3} //TODO 10
	}
|	IDENTIFIER_LIST Type '=' ExpressionList
	{
		$$ = []ConstSpec{$1, $2, "=", $4} //TODO 11
	}

Declaration:
	CONST ConstSpec
	{
		$$ = []Declaration{"const", $2} //TODO 12
	}
|	CONST '(' ')'
	{
		$$ = []Declaration{"const", "(", ")"} //TODO 13
	}
|	CONST '(' ConstSpec ConstDecl111 ')'
	{
		$$ = []Declaration{"const", "(", $3, $4, ")"} //TODO 14
	}
|	TYPE TypeSpec
	{
		$$ = []Declaration{"type", $2} //TODO 15
	}
|	TYPE '(' ')'
	{
		$$ = []Declaration{"type", "(", ")"} //TODO 16
	}
|	TYPE '(' TypeSpec TypeDecl111 ')'
	{
		$$ = []Declaration{"type", "(", $3, $4, ")"} //TODO 17
	}
|	VAR VarSpec
	{
		$$ = []Declaration{"var", $2} //TODO 18
	}
|	VAR '(' ')'
	{
		$$ = []Declaration{"var", "(", ")"} //TODO 19
	}
|	VAR '(' VarSpec VarDecl111 ')'
	{
		$$ = []Declaration{"var", "(", $3, $4, ")"} //TODO 20
	}

Element:
	Value
	{
		$$ = $1 //TODO 21
	}
|	Expression ':' Value
	{
		$$ = []Element{$1, ":", $3} //TODO 22
	}

ElementList1:
	/* EMPTY */
	{
		$$ = []ElementList1(nil) //TODO 23
	}
|	ElementList1 ',' Element
	{
		$$ = append($1.([]ElementList1), ",", $3) //TODO 24
	}

ElementType:
	Type
	{
		$$ = $1 //TODO 25
	}

ExprSwitchStmt3:
	/* EMPTY */
	{
		$$ = []ExprSwitchStmt3(nil) //TODO 26
	}
|	ExprSwitchStmt3 CASE ExpressionList ':' StatementList
	{
		$$ = append($1.([]ExprSwitchStmt3), "case", $3, ":", $5) //TODO 27
	}
|	ExprSwitchStmt3 DEFAULT ':' StatementList
	{
		$$ = append($1.([]ExprSwitchStmt3), "default", ":", $4) //TODO 28
	}

Expression:
	UnaryExpr
	{
		$$ = $1 //TODO 29
	}
|	Expression OROR Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 30
	}
|	Expression ANDAND Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 31
	}
|	Expression EQ Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 32
	}
|	Expression NE Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 33
	}
|	Expression '<' Expression
	{
		$$ = []Expression{$1, "<", $3} //TODO 34
	}
|	Expression LE Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 35
	}
|	Expression '>' Expression
	{
		$$ = []Expression{$1, ">", $3} //TODO 36
	}
|	Expression GE Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 37
	}
|	Expression '+' Expression
	{
		$$ = []Expression{$1, "+", $3} //TODO 38
	}
|	Expression '-' Expression
	{
		$$ = []Expression{$1, "-", $3} //TODO 39
	}
|	Expression '|' Expression
	{
		$$ = []Expression{$1, "|", $3} //TODO 40
	}
|	Expression '^' Expression
	{
		$$ = []Expression{$1, "^", $3} //TODO 41
	}
|	Expression '*' Expression
	{
		$$ = []Expression{$1, "*", $3} //TODO 42
	}
|	Expression '/' Expression
	{
		$$ = []Expression{$1, "/", $3} //TODO 43
	}
|	Expression '%' Expression
	{
		$$ = []Expression{$1, "%", $3} //TODO 44
	}
|	Expression LSH Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 45
	}
|	Expression RSH Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 46
	}
|	Expression '&' Expression
	{
		$$ = []Expression{$1, "&", $3} //TODO 47
	}
|	Expression ANDNOT Expression
	{
		$$ = []Expression{$1, $2, $3} //TODO 48
	}

ExpressionList:
	Expression ExpressionList1
	{
		$$ = []ExpressionList{$1, $2} //TODO 49
	}

ExpressionList1:
	/* EMPTY */
	{
		$$ = []ExpressionList1(nil) //TODO 50
	}
|	ExpressionList1 ',' Expression
	{
		$$ = append($1.([]ExpressionList1), ",", $3) //TODO 51
	}

FieldDecl:
	FieldDecl1
	{
		$$ = $1 //TODO 52
	}
|	FieldDecl1 STRING_LIT
	{
		$$ = []FieldDecl{$1, $2} //TODO 53
	}

FieldDecl1:
	IDENTIFIER_LIST Type
	{
		$$ = []FieldDecl1{$1, $2} //TODO 54
	}
|	Name
	{
		$$ = $1 //TODO 55
	}
|	'*' Name
	{
		$$ = []FieldDecl1{"*", $2} //TODO 56
	}

Function:
	Signature Block
	{
		$$ = []Function{$1, $2} //TODO 57
	}

IfStmt:
	IF Expression LBR StatementList '}'
	{
		$$ = []IfStmt{"if", $2, $3, $4, "}"} //TODO 58
	}
|	IF SimpleStmt ';' Expression LBR StatementList '}'
	{
		$$ = []IfStmt{"if", $2, ";", $4, $5, $6, "}"} //TODO 59
	}
|	IF Expression LBR StatementList '}' ELSE IfStmt
	{
		$$ = []IfStmt{"if", $2, $3, $4, "}", "else", $7} //TODO 60
	}
|	IF SimpleStmt ';' Expression LBR StatementList '}' ELSE IfStmt
	{
		$$ = []IfStmt{"if", $2, ";", $4, $5, $6, "}", "else", $9} //TODO 61
	}
|	IF Expression LBR StatementList '}' ELSE Block
	{
		$$ = []IfStmt{"if", $2, $3, $4, "}", "else", $7} //TODO 62
	}
|	IF SimpleStmt ';' Expression LBR StatementList '}' ELSE Block
	{
		$$ = []IfStmt{"if", $2, ";", $4, $5, $6, "}", "else", $9} //TODO 63
	}

ImportDecl111:
	/* EMPTY */
	{
		$$ = []ImportDecl111(nil) //TODO 64
	}
|	ImportDecl111 ';' ImportSpec
	{
		$$ = append($1.([]ImportDecl111), ";", $3) //TODO 65
	}

ImportSpec:
	STRING_LIT
	{
		$$ = $1 //TODO 66
	}
|	'.' STRING_LIT
	{
		$$ = []ImportSpec{".", $2} //TODO 67
	}
|	PackageName STRING_LIT
	{
		$$ = []ImportSpec{$1, $2} //TODO 68
	}

InterfaceType11:
	/* EMPTY */
	{
		$$ = []InterfaceType11(nil) //TODO 69
	}
|	InterfaceType11 ';' MethodSpec
	{
		$$ = append($1.([]InterfaceType11), ";", $3) //TODO 70
	}

Label:
	IDENTIFIER
	{
		$$ = $1 //TODO 71
	}

LiteralValue:
	'{' '}'
	{
		$$ = []LiteralValue{"{", "}"} //TODO 72
	}
|	'{' Element ElementList1 '}'
	{
		$$ = []LiteralValue{"{", $2, $3, "}"} //TODO 73
	}

MapType:
	MAP '[' Type ']' ElementType
	{
		$$ = []MapType{"map", "[", $3, "]", $5} //TODO 74
	}

MethodName:
	IDENTIFIER
	{
		$$ = $1 //TODO 75
	}

MethodSpec:
	MethodName Signature
	{
		$$ = []MethodSpec{$1, $2} //TODO 76
	}
|	Name
	{
		$$ = $1 //TODO 77
	}

Name:
	IDENTIFIER %prec notDot
	{
		$$ = $1 //TODO 78
	}
|	IDENTIFIER '.' IDENTIFIER
	{
		$$ = []Name{$1, ".", $3} //TODO 79
	}

PackageName:
	IDENTIFIER
	{
		$$ = $1 //TODO 80
	}

ParameterDecl:
	ParameterDecl2 Type
	{
		$$ = []ParameterDecl{$1, $2} //TODO 81
	}
|	IDENTIFIER_LIST ParameterDecl2 Type
	{
		$$ = []ParameterDecl{$1, $2, $3} //TODO 82
	}

ParameterDecl2:
	/* EMPTY */
	{
		$$ = nil //TODO 83
	}
|	DDD
	{
		$$ = $1 //TODO 84
	}

ParameterList1:
	/* EMPTY */
	{
		$$ = []ParameterList1(nil) //TODO 85
	}
|	ParameterList1 ',' ParameterDecl
	{
		$$ = append($1.([]ParameterList1), ",", $3) //TODO 86
	}

Parameters:
	'(' ')'
	{
		$$ = []Parameters{"(", ")"} //TODO 87
	}
|	'(' ParameterDecl ParameterList1 ')'
	{
		$$ = []Parameters{"(", $2, $3, ")"} //TODO 88
	}

PrimaryExpr:
	INT_LIT
	{
		$$ = $1 //TODO 89
	}
|	FLOAT_LIT
	{
		$$ = $1 //TODO 90
	}
|	IMAGINARY_LIT
	{
		$$ = $1 //TODO 91
	}
|	RUNE_LIT
	{
		$$ = $1 //TODO 92
	}
|	STRING_LIT
	{
		$$ = $1 //TODO 93
	}
|	StructType LiteralValue
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 94
	}
|	ArrayType LiteralValue
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 95
	}
|	'[' DDD ']' ElementType LiteralValue
	{
		$$ = []PrimaryExpr{"[", $2, "]", $4, $5} //TODO 96
	}
|	SliceType LiteralValue
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 97
	}
|	MapType LiteralValue
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 98
	}
|	FUNC Function
	{
		$$ = []PrimaryExpr{"func", $2} //TODO 99
	}
|	Name
	{
		$$ = $1 //TODO 100
	}
|	Name LiteralValue
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 101
	}
|	'(' Expression ')'
	{
		$$ = []PrimaryExpr{"(", $2, ")"} //TODO 102
	}
|	PrimaryExpr '.' IDENTIFIER
	{
		$$ = []PrimaryExpr{$1, ".", $3} //TODO 103
	}
|	PrimaryExpr '[' Expression ']'
	{
		$$ = []PrimaryExpr{$1, "[", $3, "]"} //TODO 104
	}
|	PrimaryExpr '[' ':' Slice2 ']'
	{
		$$ = []PrimaryExpr{$1, "[", ":", $4, "]"} //TODO 105
	}
|	PrimaryExpr '[' Expression ':' Slice2 ']'
	{
		$$ = []PrimaryExpr{$1, "[", $3, ":", $5, "]"} //TODO 106
	}
|	PrimaryExpr '.' '(' Type ')'
	{
		$$ = []PrimaryExpr{$1, ".", "(", $4, ")"} //TODO 107
	}
|	PrimaryExpr '(' ')'
	{
		$$ = []PrimaryExpr{$1, "(", ")"} //TODO 108
	}
|	PrimaryExpr '(' ArgumentList ')'
	{
		$$ = []PrimaryExpr{$1, "(", $3, ")"} //TODO 109
	}
|	PrimaryExpr '(' TypeNoName ')'
	{
		$$ = []PrimaryExpr{$1, "(", $3, ")"} //TODO 110
	}
|	PrimaryExpr '(' TypeNoName ',' ArgumentList ')'
	{
		$$ = []PrimaryExpr{$1, "(", $3, ",", $5, ")"} //TODO 111
	}

Receiver:
	'(' BaseTypeName ')'
	{
		$$ = []Receiver{"(", $2, ")"} //TODO 112
	}
|	'(' IDENTIFIER BaseTypeName ')'
	{
		$$ = []Receiver{"(", $2, $3, ")"} //TODO 113
	}
|	'(' '*' BaseTypeName ')'
	{
		$$ = []Receiver{"(", "*", $3, ")"} //TODO 114
	}
|	'(' IDENTIFIER '*' BaseTypeName ')'
	{
		$$ = []Receiver{"(", $2, "*", $4, ")"} //TODO 115
	}

SelectStmt1:
	/* EMPTY */
	{
		$$ = []SelectStmt1(nil) //TODO 116
	}
|	SelectStmt1 CASE SendStmt ':' StatementList
	{
		$$ = append($1.([]SelectStmt1), "case", $3, ":", $5) //TODO 117
	}
|	SelectStmt1 CASE Expression ':' StatementList
	{
		$$ = append($1.([]SelectStmt1), "case", $3, ":", $5) //TODO 118
	}
|	SelectStmt1 CASE ExpressionList '=' Expression ':' StatementList
	{
		$$ = append($1.([]SelectStmt1), "case", $3, "=", $5, ":", $7) //TODO 119
	}
|	SelectStmt1 CASE IDLIST_COLAS Expression ':' StatementList
	{
		$$ = append($1.([]SelectStmt1), "case", $3, $4, ":", $6) //TODO 120
	}
|	SelectStmt1 DEFAULT ':' StatementList
	{
		$$ = append($1.([]SelectStmt1), "default", ":", $4) //TODO 121
	}

SendStmt:
	Expression COMM Expression
	{
		$$ = []SendStmt{$1, $2, $3} //TODO 122
	}

Signature:
	Parameters
	{
		$$ = $1 //TODO 123
	}
|	Parameters Parameters
	{
		$$ = []Signature{$1, $2} //TODO 124
	}
|	Parameters Name
	{
		$$ = []Signature{$1, $2} //TODO 125
	}
|	Parameters TypeLit
	{
		$$ = []Signature{$1, $2} //TODO 126
	}

SimpleStmt:
	/* EMPTY */
	{
		$$ = nil //TODO 127
	}
|	Expression
	{
		$$ = $1 //TODO 128
	}
|	SendStmt
	{
		$$ = $1 //TODO 129
	}
|	Expression INC
	{
		$$ = []SimpleStmt{$1, $2} //TODO 130
	}
|	Expression DEC
	{
		$$ = []SimpleStmt{$1, $2} //TODO 131
	}
|	ExpressionList ASSIGN_OP ExpressionList
	{
		$$ = []SimpleStmt{$1, $2, $3} //TODO 132
	}
|	IDLIST_COLAS ExpressionList
	{
		$$ = []SimpleStmt{$1, $2} //TODO 133
	}

Slice2:
	/* EMPTY */
	{
		$$ = nil //TODO 134
	}
|	Expression
	{
		$$ = $1 //TODO 135
	}

SliceType:
	'[' ']' ElementType
	{
		$$ = []SliceType{"[", "]", $3} //TODO 136
	}

SourceFile:
	%prec notPackage
	{
		yylex.(*lx).Error("package statement must be first")
		goto ret1
	}
|	PACKAGE PackageName ';' SourceFile1 SourceFile2
	{
		$$ = []SourceFile{"package", $2, ";", $4, $5} //TODO 137
	}

SourceFile1:
	/* EMPTY */
	{
		$$ = []SourceFile1(nil) //TODO 138
	}
|	SourceFile1 IMPORT ImportSpec ';'
	{
		$$ = append($1.([]SourceFile1), "import", $3, ";") //TODO 139
	}
|	SourceFile1 IMPORT '(' ')' ';'
	{
		$$ = append($1.([]SourceFile1), "import", "(", ")", ";") //TODO 140
	}
|	SourceFile1 IMPORT '(' ImportSpec ImportDecl111 ')' ';'
	{
		$$ = append($1.([]SourceFile1), "import", "(", $4, $5, ")", ";") //TODO 141
	}

SourceFile2:
	/* EMPTY */
	{
		$$ = []SourceFile2(nil) //TODO 142
	}
|	SourceFile2 Declaration ';'
	{
		$$ = append($1.([]SourceFile2), $2, ";") //TODO 143
	}
|	SourceFile2 FUNC IDENTIFIER Function ';'
	{
		$$ = append($1.([]SourceFile2), "func", $3, $4, ";") //TODO 144
	}
|	SourceFile2 FUNC IDENTIFIER Signature ';'
	{
		$$ = append($1.([]SourceFile2), "func", $3, $4, ";") //TODO 145
	}
|	SourceFile2 FUNC Receiver MethodName Function ';'
	{
		$$ = append($1.([]SourceFile2), "func", $3, $4, $5, ";") //TODO 146
	}
|	SourceFile2 FUNC Receiver MethodName Signature ';'
	{
		$$ = append($1.([]SourceFile2), "func", $3, $4, $5, ";") //TODO 147
	}

Start:
	SourceFile
	{
		_parserResult = $1 //TODO 148
	}

Statement:
	Declaration
	{
		$$ = $1 //TODO 149
	}
|	Label ':' Statement
	{
		$$ = []Statement{$1, ":", $3} //TODO 150
	}
|	SimpleStmt
	{
		$$ = $1 //TODO 151
	}
|	GO Expression
	{
		$$ = []Statement{"go", $2} //TODO 152
	}
|	RETURN
	{
		$$ = "return" //TODO 153
	}
|	RETURN ExpressionList
	{
		$$ = []Statement{"return", $2} //TODO 154
	}
|	BREAK
	{
		$$ = "break" //TODO 155
	}
|	BREAK Label
	{
		$$ = []Statement{"break", $2} //TODO 156
	}
|	CONTINUE
	{
		$$ = "continue" //TODO 157
	}
|	CONTINUE Label
	{
		$$ = []Statement{"continue", $2} //TODO 158
	}
|	GOTO Label
	{
		$$ = []Statement{"goto", $2} //TODO 159
	}
|	FALLTHROUGH
	{
		$$ = "fallthrough" //TODO 160
	}
|	Block
	{
		$$ = $1 //TODO 161
	}
|	IfStmt
	{
		$$ = $1 //TODO 162
	}
|	SWITCH LBR ExprSwitchStmt3 '}'
	{
		$$ = []Statement{"switch", $2, $3, "}"} //TODO 163
	}
|	SWITCH SimpleStmt ';' LBR ExprSwitchStmt3 '}'
	{
		$$ = []Statement{"switch", $2, ";", $4, $5, "}"} //TODO 164
	}
|	SWITCH Expression LBR ExprSwitchStmt3 '}'
	{
		$$ = []Statement{"switch", $2, $3, $4, "}"} //TODO 165
	}
|	SWITCH SimpleStmt ';' Expression LBR ExprSwitchStmt3 '}'
	{
		$$ = []Statement{"switch", $2, ";", $4, $5, $6, "}"} //TODO 166
	}
|	SWITCH TypeSwitchGuard LBR TypeSwitchStmt2 '}'
	{
		$$ = []Statement{"switch", $2, $3, $4, "}"} //TODO 167
	}
|	SWITCH SimpleStmt ';' TypeSwitchGuard LBR TypeSwitchStmt2 '}'
	{
		$$ = []Statement{"switch", $2, ";", $4, $5, $6, "}"} //TODO 168
	}
|	SELECT '{' SelectStmt1 '}'
	{
		$$ = []Statement{"select", "{", $3, "}"} //TODO 169
	}
|	FOR LBR StatementList '}'
	{
		$$ = []Statement{"for", $2, $3, "}"} //TODO 170
	}
|	FOR Condition LBR StatementList '}'
	{
		$$ = []Statement{"for", $2, $3, $4, "}"} //TODO 171
	}
|	FOR SimpleStmt ';' ';' SimpleStmt LBR StatementList '}'
	{
		$$ = []Statement{"for", $2, ";", ";", $5, $6, $7, "}"} //TODO 172
	}
|	FOR SimpleStmt ';' Condition ';' SimpleStmt LBR StatementList '}'
	{
		$$ = []Statement{"for", $2, ";", $4, ";", $6, $7, $8, "}"} //TODO 173
	}
|	FOR ExpressionList '=' RANGE Expression LBR StatementList '}'
	{
		$$ = []Statement{"for", $2, "=", "range", $5, $6, $7, "}"} //TODO 174
	}
|	FOR IDLIST_COLAS RANGE Expression LBR StatementList '}'
	{
		$$ = []Statement{"for", $2, "range", $4, $5, $6, "}"} //TODO 175
	}
|	DEFER Expression
	{
		$$ = []Statement{"defer", $2} //TODO 176
	}

StatementList:
	StatementList1
	{
		$$ = $1 //TODO 177
	}

StatementList1:
	/* EMPTY */
	{
		$$ = []StatementList1(nil) //TODO 178
	}
|	StatementList1 Statement ';'
	{
		$$ = append($1.([]StatementList1), $2, ";") //TODO 179
	}

StructType:
	STRUCT '{' '}'
	{
		$$ = []StructType{"struct", "{", "}"} //TODO 180
	}
|	STRUCT '{' FieldDecl StructType11 '}'
	{
		$$ = []StructType{"struct", "{", $3, $4, "}"} //TODO 181
	}

StructType11:
	/* EMPTY */
	{
		$$ = []StructType11(nil) //TODO 182
	}
|	StructType11 ';' FieldDecl
	{
		$$ = append($1.([]StructType11), ";", $3) //TODO 183
	}

Type:
	Name
	{
		$$ = $1 //TODO 184
	}
|	TypeLit
	{
		$$ = $1 //TODO 185
	}
|	'(' Type ')'
	{
		$$ = []Type{"(", $2, ")"} //TODO 186
	}

TypeDecl111:
	/* EMPTY */
	{
		$$ = []TypeDecl111(nil) //TODO 187
	}
|	TypeDecl111 ';' TypeSpec
	{
		$$ = append($1.([]TypeDecl111), ";", $3) //TODO 188
	}

TypeList1:
	/* EMPTY */
	{
		$$ = []TypeList1(nil) //TODO 189
	}
|	TypeList1 ',' Type
	{
		$$ = append($1.([]TypeList1), ",", $3) //TODO 190
	}

TypeLit:
	ArrayType
	{
		$$ = $1 //TODO 191
	}
|	StructType
	{
		$$ = $1 //TODO 192
	}
|	'*' TypeNoName
	{
		$$ = []TypeLit{"*", $2} //TODO 193
	}
|	FUNC Signature
	{
		$$ = []TypeLit{"func", $2} //TODO 194
	}
|	INTERFACE '{' '}'
	{
		$$ = []TypeLit{"interface", "{", "}"} //TODO 195
	}
|	INTERFACE '{' MethodSpec InterfaceType11 '}'
	{
		$$ = []TypeLit{"interface", "{", $3, $4, "}"} //TODO 196
	}
|	SliceType
	{
		$$ = $1 //TODO 197
	}
|	MapType
	{
		$$ = $1 //TODO 198
	}
|	CHAN ElementType
	{
		$$ = []TypeLit{"chan", $2} //TODO 199
	}
|	CHANCOMM ElementType
	{
		$$ = []TypeLit{$1, $2} //TODO 200
	}
|	COMMCHAN ElementType
	{
		$$ = []TypeLit{$1, $2} //TODO 201
	}

TypeNoName:
	TypeLit
	{
		$$ = $1 //TODO 202
	}
|	'(' TypeNoName ')'
	{
		$$ = []TypeNoName{"(", $2, ")"} //TODO 203
	}

TypeSpec:
	IDENTIFIER Type
	{
		$$ = []TypeSpec{$1, $2} //TODO 204
	}

TypeSwitchGuard:
	PrimaryExpr '.' '(' TYPE ')'
	{
		$$ = []TypeSwitchGuard{$1, ".", "(", "type", ")"} //TODO 205
	}
|	IDLIST_COLAS PrimaryExpr '.' '(' TYPE ')'
	{
		$$ = []TypeSwitchGuard{$1, $2, ".", "(", "type", ")"} //TODO 206
	}

TypeSwitchStmt2:
	/* EMPTY */
	{
		$$ = []TypeSwitchStmt2(nil) //TODO 207
	}
|	TypeSwitchStmt2 CASE Type TypeList1 ':' StatementList
	{
		$$ = append($1.([]TypeSwitchStmt2), "case", $3, $4, ":", $6) //TODO 208
	}
|	TypeSwitchStmt2 DEFAULT ':' StatementList
	{
		$$ = append($1.([]TypeSwitchStmt2), "default", ":", $4) //TODO 209
	}

UnaryExpr:
	PrimaryExpr
	{
		$$ = $1 //TODO 210
	}
|	'+' UnaryExpr
	{
		$$ = []UnaryExpr{"+", $2} //TODO 211
	}
|	'-' UnaryExpr
	{
		$$ = []UnaryExpr{"-", $2} //TODO 212
	}
|	'!' UnaryExpr
	{
		$$ = []UnaryExpr{"!", $2} //TODO 213
	}
|	'^' UnaryExpr
	{
		$$ = []UnaryExpr{"^", $2} //TODO 214
	}
|	'*' UnaryExpr
	{
		$$ = []UnaryExpr{"*", $2} //TODO 215
	}
|	'&' UnaryExpr
	{
		$$ = []UnaryExpr{"&", $2} //TODO 216
	}
|	COMM UnaryExpr
	{
		$$ = []UnaryExpr{$1, $2} //TODO 217
	}

Value:
	Expression
	{
		$$ = $1 //TODO 218
	}
|	LiteralValue
	{
		$$ = $1 //TODO 219
	}

VarDecl111:
	/* EMPTY */
	{
		$$ = []VarDecl111(nil) //TODO 220
	}
|	VarDecl111 ';'
	{
		lx := yylex.(*lx)
		lx.toks, lx.state = nil, st2 //TODO named state alias
	}
	VarSpec
	{
		//TODO $$ = append($1.([]VarDecl111), ";", $3) //TODO 221
	}

VarSpec:
	IDENTIFIER_LIST Type
	{
		$$ = []VarSpec{$1, $2} //TODO 222
	}
|	IDENTIFIER_LIST Type '=' ExpressionList
	{
		$$ = []VarSpec{$1, $2, "=", $4} //TODO 223
	}
|	IDENTIFIER_LIST '=' ExpressionList
	{
		$$ = []VarSpec{$1, "=", $3} //TODO 224
	}

%%

//TODO remove demo stuff below

var _parserResult interface{}

type (
	ArgumentList interface{}
	ArrayType interface{}
	BaseTypeName interface{}
	Block interface{}
	Condition interface{}
	ConstDecl111 interface{}
	ConstSpec interface{}
	Declaration interface{}
	Element interface{}
	ElementList1 interface{}
	ElementType interface{}
	ExprSwitchStmt3 interface{}
	Expression interface{}
	ExpressionList interface{}
	ExpressionList1 interface{}
	FieldDecl interface{}
	FieldDecl1 interface{}
	Function interface{}
	IfStmt interface{}
	ImportDecl111 interface{}
	ImportSpec interface{}
	InterfaceType11 interface{}
	Label interface{}
	LiteralValue interface{}
	MapType interface{}
	MethodName interface{}
	MethodSpec interface{}
	Name interface{}
	PackageName interface{}
	ParameterDecl interface{}
	ParameterDecl2 interface{}
	ParameterList1 interface{}
	Parameters interface{}
	PrimaryExpr interface{}
	Receiver interface{}
	SelectStmt1 interface{}
	SendStmt interface{}
	Signature interface{}
	SimpleStmt interface{}
	Slice2 interface{}
	SliceType interface{}
	SourceFile interface{}
	SourceFile1 interface{}
	SourceFile2 interface{}
	Start interface{}
	Statement interface{}
	StatementList interface{}
	StatementList1 interface{}
	StructType interface{}
	StructType11 interface{}
	Type interface{}
	TypeDecl111 interface{}
	TypeList1 interface{}
	TypeLit interface{}
	TypeNoName interface{}
	TypeSpec interface{}
	TypeSwitchGuard interface{}
	TypeSwitchStmt2 interface{}
	UnaryExpr interface{}
	Value interface{}
	VarDecl111 interface{}
	VarSpec interface{}
)
	
func _dump() {
	s := fmt.Sprintf("%#v", _parserResult)
	s = strings.Replace(s, "%", "%%", -1)
	s = strings.Replace(s, "{", "{%i\n", -1)
	s = strings.Replace(s, "}", "%u\n}", -1)
	s = strings.Replace(s, ", ", ",\n", -1)
	var buf bytes.Buffer
	strutil.IndentFormatter(&buf, ". ").Format(s)
	buf.WriteString("\n")
	a := strings.Split(buf.String(), "\n")
	for _, v := range a {
		if strings.HasSuffix(v, "(nil)") || strings.HasSuffix(v, "(nil),") {
			continue
		}
	
		fmt.Println(v)
	}
}

// End of demo stuff
