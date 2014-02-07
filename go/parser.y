%{

//TODO Put your favorite license here
		
// yacc source generated by ebnf2y[1]
// at 2014-02-07 14:33:37.410633479 +0100 CET
//
//  $ ebnf2y -o parser.y -pkg parser parser.ebnf
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
	item interface{} //TODO insert real field(s)
}

%token	ADD_ASGN
%token	ANDAND
%token	AND_ASGN
%token	AND_NOT
%token	AND_NOT_ASGN
%token	COLAS
%token	COMM
%token	DDD
%token	DEC
%token	DIV_ASGN
%token	EQ
%token	FLOAT_LIT
%token	GT
%token	IDENTIFIER
%token	IMAGINARY_LIT
%token	INC
%token	INT_LIT
%token	LBR
%token	LSH
%token	LSH_ASGN
%token	LT
%token	MOD_ASGN
%token	MUL_ASGN
%token	NE
%token	OROR
%token	OR_ASGN
%token	RSH
%token	RSH_ASGN
%token	RUNE_LIT
%token	STRING_LIT
%token	SUB_ASGN
%token	XOR_ASGN

%type	<item> 	/*TODO real type(s), if/where applicable */
	ADD_ASGN
	ANDAND
	AND_ASGN
	AND_NOT
	AND_NOT_ASGN
	COLAS
	COMM
	DDD
	DEC
	DIV_ASGN
	EQ
	FLOAT_LIT
	GT
	IDENTIFIER
	IMAGINARY_LIT
	INC
	INT_LIT
	LBR
	LSH
	LSH_ASGN
	LT
	MOD_ASGN
	MUL_ASGN
	NE
	OROR
	OR_ASGN
	RSH
	RSH_ASGN
	RUNE_LIT
	STRING_LIT
	SUB_ASGN
	XOR_ASGN

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
	AnonymousField
	ArgumentList
	ArgumentList1
	ArrayLength
	ArrayType
	Assignment
	Assignment1
	BaseType
	BasicLit
	Block
	Block2
	BreakStmt
	BreakStmt1
	BuiltinArgs
	BuiltinArgs1
	BuiltinCall
	BuiltinCall1
	BuiltinCall11
	Call
	Call1
	Call11
	Channel
	ChannelType
	ChannelType1
	ChannelType11
	CommCase
	CommCase1
	CommClause
	CompositeLit
	Condition
	ConstDecl
	ConstDecl1
	ConstDecl11
	ConstDecl111
	ConstDecl112
	ConstSpec
	ConstSpec1
	ConstSpec11
	ContinueStmt
	ContinueStmt1
	Conversion
	Conversion1
	Declaration
	DeferStmt
	Element
	ElementList
	ElementList1
	ElementType
	ExprCaseClause
	ExprSwitchCase
	ExprSwitchStmt
	ExprSwitchStmt1
	ExprSwitchStmt2
	ExprSwitchStmt3
	ExprSwitchStmt4
	Expression
	Expression1
	ExpressionList
	ExpressionList1
	ExpressionStmt
	FallthroughStmt
	FieldDecl
	FieldDecl1
	FieldDecl2
	ForClause
	ForClause1
	ForClause2
	ForClause3
	ForStmt
	ForStmt1
	ForStmt11
	Function
	FunctionBody
	FunctionDecl
	FunctionDecl1
	FunctionLit
	FunctionType
	GoStmt
	GotoStmt
	IdentifierList
	IdentifierList1
	IfStmt
	IfStmt1
	IfStmt11
	IfStmt2
	IfStmt21
	ImportDecl
	ImportDecl1
	ImportDecl11
	ImportDecl111
	ImportDecl112
	ImportPath
	ImportSpec
	ImportSpec1
	ImportSpec11
	IncDecStmt
	IncDecStmt1
	Index
	InitStmt
	InterfaceType
	InterfaceType1
	InterfaceTypeName
	KeyType
	LabeledStmt
	Literal
	LiteralType
	LiteralValue
	LiteralValue1
	LiteralValue11
	MapType
	MethodDecl
	MethodDecl1
	MethodExpr
	MethodSpec
	Operand
	OperandName
	PackageClause
	ParameterDecl
	ParameterList
	ParameterList1
	Parameters
	Parameters1
	Parameters11
	PointerType
	PostStmt
	PrimaryExpr
	QualifiedIdent
	RangeClause
	RangeClause1
	Receiver
	ReceiverType
	RecvExpr
	RecvStmt
	RecvStmt1
	Result
	ReturnStmt
	ReturnStmt1
	SelectStmt
	SelectStmt1
	Selector
	SendStmt
	ShortVarDecl
	Signature
	Signature1
	SimpleStmt
	Slice
	Slice1
	Slice11
	Slice12
	Slice2
	Slice21
	SliceType
	SourceFile
	SourceFile1
	SourceFile2
	Start
	Statement
	StatementList
	StatementList1
	StatementList2
	StatementList21
	StructType
	StructType1
	SwitchStmt
	Tag
	TopLevelDecl
	Type
	TypeAssertion
	TypeCaseClause
	TypeDecl
	TypeDecl1
	TypeDecl11
	TypeDecl111
	TypeDecl112
	TypeList
	TypeList1
	TypeLit
	TypeName
	TypeSpec
	TypeSwitchCase
	TypeSwitchGuard
	TypeSwitchStmt
	TypeSwitchStmt1
	TypeSwitchStmt2
	UnaryExpr
	UnaryExpr1
	Value
	VarDecl
	VarDecl1
	VarDecl11
	VarDecl111
	VarDecl112
	VarSpec
	VarSpec1
	VarSpec11

/*TODO %left, %right, ... declarations */

%start Start

%%

AnonymousField:
	TypeName
	{
		$$ = $1 //TODO 1
	}
|	'*' TypeName
	{
		$$ = []AnonymousField{"*", $2} //TODO 2
	}

ArgumentList:
	ExpressionList ArgumentList1
	{
		$$ = []ArgumentList{$1, $2} //TODO 3
	}

ArgumentList1:
	/* EMPTY */
	{
		$$ = nil //TODO 4
	}
|	DDD
	{
		$$ = $1 //TODO 5
	}

ArrayLength:
	Expression
	{
		$$ = $1 //TODO 6
	}

ArrayType:
	'[' ArrayLength ']' ElementType
	{
		$$ = []ArrayType{"[", $2, "]", $4} //TODO 7
	}

Assignment:
	ExpressionList Assignment1 ExpressionList
	{
		$$ = []Assignment{$1, $2, $3} //TODO 8
	}

Assignment1:
	MOD_ASGN
	{
		$$ = $1 //TODO 9
	}
|	AND_ASGN
	{
		$$ = $1 //TODO 10
	}
|	AND_NOT_ASGN
	{
		$$ = $1 //TODO 11
	}
|	MUL_ASGN
	{
		$$ = $1 //TODO 12
	}
|	ADD_ASGN
	{
		$$ = $1 //TODO 13
	}
|	SUB_ASGN
	{
		$$ = $1 //TODO 14
	}
|	DIV_ASGN
	{
		$$ = $1 //TODO 15
	}
|	LSH_ASGN
	{
		$$ = $1 //TODO 16
	}
|	'='
	{
		$$ = "=" //TODO 17
	}
|	RSH_ASGN
	{
		$$ = $1 //TODO 18
	}
|	XOR_ASGN
	{
		$$ = $1 //TODO 19
	}
|	OR_ASGN
	{
		$$ = $1 //TODO 20
	}

BaseType:
	Type
	{
		$$ = $1 //TODO 21
	}

BasicLit:
	INT_LIT
	{
		$$ = $1 //TODO 22
	}
|	FLOAT_LIT
	{
		$$ = $1 //TODO 23
	}
|	IMAGINARY_LIT
	{
		$$ = $1 //TODO 24
	}
|	RUNE_LIT
	{
		$$ = $1 //TODO 25
	}
|	STRING_LIT
	{
		$$ = $1 //TODO 26
	}

Block:
	'{' StatementList '}'
	{
		$$ = []Block{"{", $2, "}"} //TODO 27
	}

Block2:
	LBR StatementList '}'
	{
		$$ = []Block2{$1, $2, "}"} //TODO 28
	}

BreakStmt:
	BREAK BreakStmt1
	{
		$$ = []BreakStmt{"break", $2} //TODO 29
	}

BreakStmt1:
	/* EMPTY */
	{
		$$ = nil //TODO 30
	}
|	IDENTIFIER
	{
		$$ = $1 //TODO 31
	}

BuiltinArgs:
	Type BuiltinArgs1
	{
		$$ = []BuiltinArgs{$1, $2} //TODO 32
	}
|	ArgumentList
	{
		$$ = $1 //TODO 33
	}

BuiltinArgs1:
	/* EMPTY */
	{
		$$ = nil //TODO 34
	}
|	',' ArgumentList
	{
		$$ = []BuiltinArgs1{",", $2} //TODO 35
	}

BuiltinCall:
	IDENTIFIER '(' BuiltinCall1 ')'
	{
		$$ = []BuiltinCall{$1, "(", $3, ")"} //TODO 36
	}

BuiltinCall1:
	/* EMPTY */
	{
		$$ = nil //TODO 37
	}
|	BuiltinArgs BuiltinCall11
	{
		$$ = []BuiltinCall1{$1, $2} //TODO 38
	}

BuiltinCall11:
	/* EMPTY */
	{
		$$ = nil //TODO 39
	}
|	','
	{
		$$ = "," //TODO 40
	}

Call:
	'(' Call1 ')'
	{
		$$ = []Call{"(", $2, ")"} //TODO 41
	}

Call1:
	/* EMPTY */
	{
		$$ = nil //TODO 42
	}
|	ArgumentList Call11
	{
		$$ = []Call1{$1, $2} //TODO 43
	}

Call11:
	/* EMPTY */
	{
		$$ = nil //TODO 44
	}
|	','
	{
		$$ = "," //TODO 45
	}

Channel:
	Expression
	{
		$$ = $1 //TODO 46
	}

ChannelType:
	ChannelType1 ElementType
	{
		$$ = []ChannelType{$1, $2} //TODO 47
	}

ChannelType1:
	CHAN ChannelType11
	{
		$$ = []ChannelType1{"chan", $2} //TODO 48
	}
|	COMM CHAN
	{
		$$ = []ChannelType1{$1, "chan"} //TODO 49
	}

ChannelType11:
	/* EMPTY */
	{
		$$ = nil //TODO 50
	}
|	COMM
	{
		$$ = $1 //TODO 51
	}

CommCase:
	CASE CommCase1
	{
		$$ = []CommCase{"case", $2} //TODO 52
	}
|	DEFAULT
	{
		$$ = "default" //TODO 53
	}

CommCase1:
	SendStmt
	{
		$$ = $1 //TODO 54
	}
|	RecvStmt
	{
		$$ = $1 //TODO 55
	}

CommClause:
	CommCase ':' StatementList
	{
		$$ = []CommClause{$1, ":", $3} //TODO 56
	}

CompositeLit:
	LiteralType LiteralValue
	{
		$$ = []CompositeLit{$1, $2} //TODO 57
	}

Condition:
	Expression
	{
		$$ = $1 //TODO 58
	}

ConstDecl:
	CONST ConstDecl1
	{
		$$ = []ConstDecl{"const", $2} //TODO 59
	}

ConstDecl1:
	ConstSpec
	{
		$$ = $1 //TODO 60
	}
|	'(' ConstDecl11 ')'
	{
		$$ = []ConstDecl1{"(", $2, ")"} //TODO 61
	}

ConstDecl11:
	/* EMPTY */
	{
		$$ = nil //TODO 62
	}
|	ConstSpec ConstDecl111 ConstDecl112
	{
		$$ = []ConstDecl11{$1, $2, $3} //TODO 63
	}

ConstDecl111:
	/* EMPTY */
	{
		$$ = []ConstDecl111(nil) //TODO 64
	}
|	ConstDecl111 ';' ConstSpec
	{
		$$ = append($1.([]ConstDecl111), ";", $3) //TODO 65
	}

ConstDecl112:
	/* EMPTY */
	{
		$$ = nil //TODO 66
	}
|	';'
	{
		$$ = ";" //TODO 67
	}

ConstSpec:
	IdentifierList ConstSpec1
	{
		$$ = []ConstSpec{$1, $2} //TODO 68
	}

ConstSpec1:
	/* EMPTY */
	{
		$$ = nil //TODO 69
	}
|	ConstSpec11 '=' ExpressionList
	{
		$$ = []ConstSpec1{$1, "=", $3} //TODO 70
	}

ConstSpec11:
	/* EMPTY */
	{
		$$ = nil //TODO 71
	}
|	Type
	{
		$$ = $1 //TODO 72
	}

ContinueStmt:
	CONTINUE ContinueStmt1
	{
		$$ = []ContinueStmt{"continue", $2} //TODO 73
	}

ContinueStmt1:
	/* EMPTY */
	{
		$$ = nil //TODO 74
	}
|	IDENTIFIER
	{
		$$ = $1 //TODO 75
	}

Conversion:
	Type '(' Expression Conversion1 ')'
	{
		$$ = []Conversion{$1, "(", $3, $4, ")"} //TODO 76
	}

Conversion1:
	/* EMPTY */
	{
		$$ = nil //TODO 77
	}
|	','
	{
		$$ = "," //TODO 78
	}

Declaration:
	ConstDecl
	{
		$$ = $1 //TODO 79
	}
|	TypeDecl
	{
		$$ = $1 //TODO 80
	}
|	VarDecl
	{
		$$ = $1 //TODO 81
	}

DeferStmt:
	DEFER Expression
	{
		$$ = []DeferStmt{"defer", $2} //TODO 82
	}

Element:
	Value
	{
		$$ = $1 //TODO 83
	}
|	Expression ':' Value
	{
		$$ = []Element{$1, ":", $3} //TODO 84
	}

ElementList:
	Element ElementList1
	{
		$$ = []ElementList{$1, $2} //TODO 85
	}

ElementList1:
	/* EMPTY */
	{
		$$ = []ElementList1(nil) //TODO 86
	}
|	ElementList1 ',' Element
	{
		$$ = append($1.([]ElementList1), ",", $3) //TODO 87
	}

ElementType:
	Type
	{
		$$ = $1 //TODO 88
	}

ExprCaseClause:
	ExprSwitchCase ':' StatementList
	{
		$$ = []ExprCaseClause{$1, ":", $3} //TODO 89
	}

ExprSwitchCase:
	CASE ExpressionList
	{
		$$ = []ExprSwitchCase{"case", $2} //TODO 90
	}
|	DEFAULT
	{
		$$ = "default" //TODO 91
	}

ExprSwitchStmt:
	SWITCH LBR ExprSwitchStmt1 '}'
	{
		$$ = []ExprSwitchStmt{"switch", $2, $3, "}"} //TODO 92
	}
|	SWITCH Expression LBR ExprSwitchStmt2 '}'
	{
		$$ = []ExprSwitchStmt{"switch", $2, $3, $4, "}"} //TODO 93
	}
|	SWITCH SimpleStmt ';' LBR ExprSwitchStmt3 '}'
	{
		$$ = []ExprSwitchStmt{"switch", $2, ";", $4, $5, "}"} //TODO 94
	}
|	SWITCH SimpleStmt ';' Expression LBR ExprSwitchStmt4 '}'
	{
		$$ = []ExprSwitchStmt{"switch", $2, ";", $4, $5, $6, "}"} //TODO 95
	}

ExprSwitchStmt1:
	/* EMPTY */
	{
		$$ = []ExprSwitchStmt1(nil) //TODO 96
	}
|	ExprSwitchStmt1 ExprCaseClause
	{
		$$ = append($1.([]ExprSwitchStmt1), $2) //TODO 97
	}

ExprSwitchStmt2:
	/* EMPTY */
	{
		$$ = []ExprSwitchStmt2(nil) //TODO 98
	}
|	ExprSwitchStmt2 ExprCaseClause
	{
		$$ = append($1.([]ExprSwitchStmt2), $2) //TODO 99
	}

ExprSwitchStmt3:
	/* EMPTY */
	{
		$$ = []ExprSwitchStmt3(nil) //TODO 100
	}
|	ExprSwitchStmt3 ExprCaseClause
	{
		$$ = append($1.([]ExprSwitchStmt3), $2) //TODO 101
	}

ExprSwitchStmt4:
	/* EMPTY */
	{
		$$ = []ExprSwitchStmt4(nil) //TODO 102
	}
|	ExprSwitchStmt4 ExprCaseClause
	{
		$$ = append($1.([]ExprSwitchStmt4), $2) //TODO 103
	}

Expression:
	UnaryExpr
	{
		$$ = $1 //TODO 104
	}
|	Expression Expression1 UnaryExpr
	{
		$$ = []Expression{$1, $2, $3} //TODO 105
	}

Expression1:
	NE
	{
		$$ = $1 //TODO 106
	}
|	'%'
	{
		$$ = "%" //TODO 107
	}
|	'&'
	{
		$$ = "&" //TODO 108
	}
|	ANDAND
	{
		$$ = $1 //TODO 109
	}
|	AND_NOT
	{
		$$ = $1 //TODO 110
	}
|	'*'
	{
		$$ = "*" //TODO 111
	}
|	'+'
	{
		$$ = "+" //TODO 112
	}
|	'-'
	{
		$$ = "-" //TODO 113
	}
|	'/'
	{
		$$ = "/" //TODO 114
	}
|	'<'
	{
		$$ = "<" //TODO 115
	}
|	LSH
	{
		$$ = $1 //TODO 116
	}
|	LT
	{
		$$ = $1 //TODO 117
	}
|	EQ
	{
		$$ = $1 //TODO 118
	}
|	'>'
	{
		$$ = ">" //TODO 119
	}
|	GT
	{
		$$ = $1 //TODO 120
	}
|	RSH
	{
		$$ = $1 //TODO 121
	}
|	'^'
	{
		$$ = "^" //TODO 122
	}
|	'|'
	{
		$$ = "|" //TODO 123
	}
|	OROR
	{
		$$ = $1 //TODO 124
	}

ExpressionList:
	Expression ExpressionList1
	{
		$$ = []ExpressionList{$1, $2} //TODO 125
	}

ExpressionList1:
	/* EMPTY */
	{
		$$ = []ExpressionList1(nil) //TODO 126
	}
|	ExpressionList1 ',' Expression
	{
		$$ = append($1.([]ExpressionList1), ",", $3) //TODO 127
	}

ExpressionStmt:
	Expression
	{
		$$ = $1 //TODO 128
	}

FallthroughStmt:
	FALLTHROUGH
	{
		$$ = "fallthrough" //TODO 129
	}

FieldDecl:
	FieldDecl1 FieldDecl2
	{
		$$ = []FieldDecl{$1, $2} //TODO 130
	}

FieldDecl1:
	IdentifierList Type
	{
		$$ = []FieldDecl1{$1, $2} //TODO 131
	}
|	AnonymousField
	{
		$$ = $1 //TODO 132
	}

FieldDecl2:
	/* EMPTY */
	{
		$$ = nil //TODO 133
	}
|	Tag
	{
		$$ = $1 //TODO 134
	}

ForClause:
	ForClause1 ';' ForClause2 ';' ForClause3
	{
		$$ = []ForClause{$1, ";", $3, ";", $5} //TODO 135
	}

ForClause1:
	/* EMPTY */
	{
		$$ = nil //TODO 136
	}
|	InitStmt
	{
		$$ = $1 //TODO 137
	}

ForClause2:
	/* EMPTY */
	{
		$$ = nil //TODO 138
	}
|	Condition
	{
		$$ = $1 //TODO 139
	}

ForClause3:
	/* EMPTY */
	{
		$$ = nil //TODO 140
	}
|	PostStmt
	{
		$$ = $1 //TODO 141
	}

ForStmt:
	FOR ForStmt1 Block2
	{
		$$ = []ForStmt{"for", $2, $3} //TODO 142
	}

ForStmt1:
	/* EMPTY */
	{
		$$ = nil //TODO 143
	}
|	ForStmt11
	{
		$$ = $1 //TODO 144
	}

ForStmt11:
	Condition
	{
		$$ = $1 //TODO 145
	}
|	ForClause
	{
		$$ = $1 //TODO 146
	}
|	RangeClause
	{
		$$ = $1 //TODO 147
	}

Function:
	Signature FunctionBody
	{
		$$ = []Function{$1, $2} //TODO 148
	}

FunctionBody:
	Block
	{
		$$ = $1 //TODO 149
	}

FunctionDecl:
	FUNC IDENTIFIER FunctionDecl1
	{
		$$ = []FunctionDecl{"func", $2, $3} //TODO 150
	}

FunctionDecl1:
	Function
	{
		$$ = $1 //TODO 151
	}
|	Signature
	{
		$$ = $1 //TODO 152
	}

FunctionLit:
	FUNC Function
	{
		$$ = []FunctionLit{"func", $2} //TODO 153
	}

FunctionType:
	FUNC Signature
	{
		$$ = []FunctionType{"func", $2} //TODO 154
	}

GoStmt:
	GO Expression
	{
		$$ = []GoStmt{"go", $2} //TODO 155
	}

GotoStmt:
	GOTO IDENTIFIER
	{
		$$ = []GotoStmt{"goto", $2} //TODO 156
	}

IdentifierList:
	IDENTIFIER IdentifierList1
	{
		$$ = []IdentifierList{$1, $2} //TODO 157
	}

IdentifierList1:
	/* EMPTY */
	{
		$$ = []IdentifierList1(nil) //TODO 158
	}
|	IdentifierList1 ',' IDENTIFIER
	{
		$$ = append($1.([]IdentifierList1), ",", $3) //TODO 159
	}

IfStmt:
	IF Expression Block2 IfStmt1
	{
		$$ = []IfStmt{"if", $2, $3, $4} //TODO 160
	}
|	IF SimpleStmt ';' Expression Block2 IfStmt2
	{
		$$ = []IfStmt{"if", $2, ";", $4, $5, $6} //TODO 161
	}

IfStmt1:
	/* EMPTY */
	{
		$$ = nil //TODO 162
	}
|	ELSE IfStmt11
	{
		$$ = []IfStmt1{"else", $2} //TODO 163
	}

IfStmt11:
	IfStmt
	{
		$$ = $1 //TODO 164
	}
|	Block
	{
		$$ = $1 //TODO 165
	}

IfStmt2:
	/* EMPTY */
	{
		$$ = nil //TODO 166
	}
|	ELSE IfStmt21
	{
		$$ = []IfStmt2{"else", $2} //TODO 167
	}

IfStmt21:
	IfStmt
	{
		$$ = $1 //TODO 168
	}
|	Block
	{
		$$ = $1 //TODO 169
	}

ImportDecl:
	IMPORT ImportDecl1
	{
		$$ = []ImportDecl{"import", $2} //TODO 170
	}

ImportDecl1:
	ImportSpec
	{
		$$ = $1 //TODO 171
	}
|	'(' ImportDecl11 ')'
	{
		$$ = []ImportDecl1{"(", $2, ")"} //TODO 172
	}

ImportDecl11:
	/* EMPTY */
	{
		$$ = nil //TODO 173
	}
|	ImportSpec ImportDecl111 ImportDecl112
	{
		$$ = []ImportDecl11{$1, $2, $3} //TODO 174
	}

ImportDecl111:
	/* EMPTY */
	{
		$$ = []ImportDecl111(nil) //TODO 175
	}
|	ImportDecl111 ';' ImportSpec
	{
		$$ = append($1.([]ImportDecl111), ";", $3) //TODO 176
	}

ImportDecl112:
	/* EMPTY */
	{
		$$ = nil //TODO 177
	}
|	';'
	{
		$$ = ";" //TODO 178
	}

ImportPath:
	STRING_LIT
	{
		$$ = $1 //TODO 179
	}

ImportSpec:
	ImportSpec1 ImportPath
	{
		$$ = []ImportSpec{$1, $2} //TODO 180
	}

ImportSpec1:
	/* EMPTY */
	{
		$$ = nil //TODO 181
	}
|	ImportSpec11
	{
		$$ = $1 //TODO 182
	}

ImportSpec11:
	'.'
	{
		$$ = "." //TODO 183
	}
|	IDENTIFIER
	{
		$$ = $1 //TODO 184
	}

IncDecStmt:
	Expression IncDecStmt1
	{
		$$ = []IncDecStmt{$1, $2} //TODO 185
	}

IncDecStmt1:
	INC
	{
		$$ = $1 //TODO 186
	}
|	DEC
	{
		$$ = $1 //TODO 187
	}

Index:
	'[' Expression ']'
	{
		$$ = []Index{"[", $2, "]"} //TODO 188
	}

InitStmt:
	SimpleStmt
	{
		$$ = $1 //TODO 189
	}

InterfaceType:
	INTERFACE '{' InterfaceType1 '}'
	{
		$$ = []InterfaceType{"interface", "{", $3, "}"} //TODO 190
	}

InterfaceType1:
	/* EMPTY */
	{
		$$ = []InterfaceType1(nil) //TODO 191
	}
|	InterfaceType1 MethodSpec ';'
	{
		$$ = append($1.([]InterfaceType1), $2, ";") //TODO 192
	}

InterfaceTypeName:
	TypeName
	{
		$$ = $1 //TODO 193
	}

KeyType:
	Type
	{
		$$ = $1 //TODO 194
	}

LabeledStmt:
	IDENTIFIER ':' Statement
	{
		$$ = []LabeledStmt{$1, ":", $3} //TODO 195
	}

Literal:
	BasicLit
	{
		$$ = $1 //TODO 196
	}
|	CompositeLit
	{
		$$ = $1 //TODO 197
	}
|	FunctionLit
	{
		$$ = $1 //TODO 198
	}

LiteralType:
	StructType
	{
		$$ = $1 //TODO 199
	}
|	ArrayType
	{
		$$ = $1 //TODO 200
	}
|	'[' DDD ']' ElementType
	{
		$$ = []LiteralType{"[", $2, "]", $4} //TODO 201
	}
|	SliceType
	{
		$$ = $1 //TODO 202
	}
|	MapType
	{
		$$ = $1 //TODO 203
	}
|	TypeName
	{
		$$ = $1 //TODO 204
	}

LiteralValue:
	'{' LiteralValue1 '}'
	{
		$$ = []LiteralValue{"{", $2, "}"} //TODO 205
	}

LiteralValue1:
	/* EMPTY */
	{
		$$ = nil //TODO 206
	}
|	ElementList LiteralValue11
	{
		$$ = []LiteralValue1{$1, $2} //TODO 207
	}

LiteralValue11:
	/* EMPTY */
	{
		$$ = nil //TODO 208
	}
|	','
	{
		$$ = "," //TODO 209
	}

MapType:
	MAP '[' KeyType ']' ElementType
	{
		$$ = []MapType{"map", "[", $3, "]", $5} //TODO 210
	}

MethodDecl:
	FUNC Receiver IDENTIFIER MethodDecl1
	{
		$$ = []MethodDecl{"func", $2, $3, $4} //TODO 211
	}

MethodDecl1:
	Function
	{
		$$ = $1 //TODO 212
	}
|	Signature
	{
		$$ = $1 //TODO 213
	}

MethodExpr:
	ReceiverType '.' IDENTIFIER
	{
		$$ = []MethodExpr{$1, ".", $3} //TODO 214
	}

MethodSpec:
	IDENTIFIER Signature
	{
		$$ = []MethodSpec{$1, $2} //TODO 215
	}
|	InterfaceTypeName
	{
		$$ = $1 //TODO 216
	}

Operand:
	Literal
	{
		$$ = $1 //TODO 217
	}
|	OperandName
	{
		$$ = $1 //TODO 218
	}
|	MethodExpr
	{
		$$ = $1 //TODO 219
	}
|	'(' Expression ')'
	{
		$$ = []Operand{"(", $2, ")"} //TODO 220
	}

OperandName:
	IDENTIFIER
	{
		$$ = $1 //TODO 221
	}
|	QualifiedIdent
	{
		$$ = $1 //TODO 222
	}

PackageClause:
	PACKAGE IDENTIFIER
	{
		$$ = []PackageClause{"package", $2} //TODO 223
	}

ParameterDecl:
	Type
	{
		$$ = $1 //TODO 224
	}
|	IDENTIFIER Type
	{
		$$ = []ParameterDecl{$1, $2} //TODO 225
	}
|	IDENTIFIER DDD Type
	{
		$$ = []ParameterDecl{$1, $2, $3} //TODO 226
	}
|	DDD Type
	{
		$$ = []ParameterDecl{$1, $2} //TODO 227
	}

ParameterList:
	ParameterDecl ParameterList1
	{
		$$ = []ParameterList{$1, $2} //TODO 228
	}

ParameterList1:
	/* EMPTY */
	{
		$$ = []ParameterList1(nil) //TODO 229
	}
|	ParameterList1 ',' ParameterDecl
	{
		$$ = append($1.([]ParameterList1), ",", $3) //TODO 230
	}

Parameters:
	'(' Parameters1 ')'
	{
		$$ = []Parameters{"(", $2, ")"} //TODO 231
	}

Parameters1:
	/* EMPTY */
	{
		$$ = nil //TODO 232
	}
|	ParameterList Parameters11
	{
		$$ = []Parameters1{$1, $2} //TODO 233
	}

Parameters11:
	/* EMPTY */
	{
		$$ = nil //TODO 234
	}
|	','
	{
		$$ = "," //TODO 235
	}

PointerType:
	'*' BaseType
	{
		$$ = []PointerType{"*", $2} //TODO 236
	}

PostStmt:
	SimpleStmt
	{
		$$ = $1 //TODO 237
	}

PrimaryExpr:
	Operand
	{
		$$ = $1 //TODO 238
	}
|	Conversion
	{
		$$ = $1 //TODO 239
	}
|	BuiltinCall
	{
		$$ = $1 //TODO 240
	}
|	PrimaryExpr Selector
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 241
	}
|	PrimaryExpr Index
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 242
	}
|	PrimaryExpr Slice
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 243
	}
|	PrimaryExpr TypeAssertion
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 244
	}
|	PrimaryExpr Call
	{
		$$ = []PrimaryExpr{$1, $2} //TODO 245
	}

QualifiedIdent:
	IDENTIFIER '.' IDENTIFIER
	{
		$$ = []QualifiedIdent{$1, ".", $3} //TODO 246
	}

RangeClause:
	RangeClause1 RANGE Expression
	{
		$$ = []RangeClause{$1, "range", $3} //TODO 247
	}

RangeClause1:
	ExpressionList '='
	{
		$$ = []RangeClause1{$1, "="} //TODO 248
	}
|	IdentifierList COLAS
	{
		$$ = []RangeClause1{$1, $2} //TODO 249
	}

Receiver:
	'(' IDENTIFIER ')'
	{
		$$ = []Receiver{"(", $2, ")"} //TODO 250
	}
|	'(' '*' IDENTIFIER ')'
	{
		$$ = []Receiver{"(", "*", $3, ")"} //TODO 251
	}
|	'(' IDENTIFIER IDENTIFIER ')'
	{
		$$ = []Receiver{"(", $2, $3, ")"} //TODO 252
	}
|	'(' IDENTIFIER '*' IDENTIFIER ')'
	{
		$$ = []Receiver{"(", $2, "*", $4, ")"} //TODO 253
	}

ReceiverType:
	TypeName
	{
		$$ = $1 //TODO 254
	}
|	'(' '*' TypeName ')'
	{
		$$ = []ReceiverType{"(", "*", $3, ")"} //TODO 255
	}
|	'(' ReceiverType ')'
	{
		$$ = []ReceiverType{"(", $2, ")"} //TODO 256
	}

RecvExpr:
	Expression
	{
		$$ = $1 //TODO 257
	}

RecvStmt:
	RecvExpr
	{
		$$ = $1 //TODO 258
	}
|	RecvStmt1 RecvExpr
	{
		$$ = []RecvStmt{$1, $2} //TODO 259
	}

RecvStmt1:
	ExpressionList '='
	{
		$$ = []RecvStmt1{$1, "="} //TODO 260
	}
|	IdentifierList COLAS
	{
		$$ = []RecvStmt1{$1, $2} //TODO 261
	}

Result:
	Parameters
	{
		$$ = $1 //TODO 262
	}
|	Type
	{
		$$ = $1 //TODO 263
	}

ReturnStmt:
	RETURN ReturnStmt1
	{
		$$ = []ReturnStmt{"return", $2} //TODO 264
	}

ReturnStmt1:
	/* EMPTY */
	{
		$$ = nil //TODO 265
	}
|	ExpressionList
	{
		$$ = $1 //TODO 266
	}

SelectStmt:
	SELECT '{' SelectStmt1 '}'
	{
		$$ = []SelectStmt{"select", "{", $3, "}"} //TODO 267
	}

SelectStmt1:
	/* EMPTY */
	{
		$$ = []SelectStmt1(nil) //TODO 268
	}
|	SelectStmt1 CommClause
	{
		$$ = append($1.([]SelectStmt1), $2) //TODO 269
	}

Selector:
	'.' IDENTIFIER
	{
		$$ = []Selector{".", $2} //TODO 270
	}

SendStmt:
	Channel COMM Expression
	{
		$$ = []SendStmt{$1, $2, $3} //TODO 271
	}

ShortVarDecl:
	IdentifierList COLAS ExpressionList
	{
		$$ = []ShortVarDecl{$1, $2, $3} //TODO 272
	}

Signature:
	Parameters Signature1
	{
		$$ = []Signature{$1, $2} //TODO 273
	}

Signature1:
	/* EMPTY */
	{
		$$ = nil //TODO 274
	}
|	Result
	{
		$$ = $1 //TODO 275
	}

SimpleStmt:
	ExpressionStmt
	{
		$$ = $1 //TODO 276
	}
|	SendStmt
	{
		$$ = $1 //TODO 277
	}
|	IncDecStmt
	{
		$$ = $1 //TODO 278
	}
|	Assignment
	{
		$$ = $1 //TODO 279
	}
|	ShortVarDecl
	{
		$$ = $1 //TODO 280
	}

Slice:
	'[' Slice1
	{
		$$ = []Slice{"[", $2} //TODO 281
	}
|	Slice2 ']'
	{
		$$ = []Slice{$1, "]"} //TODO 282
	}

Slice1:
	Slice11 ':' Slice12
	{
		$$ = []Slice1{$1, ":", $3} //TODO 283
	}

Slice11:
	/* EMPTY */
	{
		$$ = nil //TODO 284
	}
|	Expression
	{
		$$ = $1 //TODO 285
	}

Slice12:
	/* EMPTY */
	{
		$$ = nil //TODO 286
	}
|	Expression
	{
		$$ = $1 //TODO 287
	}

Slice2:
	Slice21 ':' Expression ':' Expression
	{
		$$ = []Slice2{$1, ":", $3, ":", $5} //TODO 288
	}

Slice21:
	/* EMPTY */
	{
		$$ = nil //TODO 289
	}
|	Expression
	{
		$$ = $1 //TODO 290
	}

SliceType:
	'[' ']' ElementType
	{
		$$ = []SliceType{"[", "]", $3} //TODO 291
	}

SourceFile:
	PackageClause ';' SourceFile1 SourceFile2
	{
		$$ = []SourceFile{$1, ";", $3, $4} //TODO 292
	}

SourceFile1:
	/* EMPTY */
	{
		$$ = []SourceFile1(nil) //TODO 293
	}
|	SourceFile1 ImportDecl ';'
	{
		$$ = append($1.([]SourceFile1), $2, ";") //TODO 294
	}

SourceFile2:
	/* EMPTY */
	{
		$$ = []SourceFile2(nil) //TODO 295
	}
|	SourceFile2 TopLevelDecl ';'
	{
		$$ = append($1.([]SourceFile2), $2, ";") //TODO 296
	}

Start:
	SourceFile
	{
		_parserResult = $1 //TODO 297
	}

Statement:
	Declaration
	{
		$$ = $1 //TODO 298
	}
|	LabeledStmt
	{
		$$ = $1 //TODO 299
	}
|	SimpleStmt
	{
		$$ = $1 //TODO 300
	}
|	GoStmt
	{
		$$ = $1 //TODO 301
	}
|	ReturnStmt
	{
		$$ = $1 //TODO 302
	}
|	BreakStmt
	{
		$$ = $1 //TODO 303
	}
|	ContinueStmt
	{
		$$ = $1 //TODO 304
	}
|	GotoStmt
	{
		$$ = $1 //TODO 305
	}
|	FallthroughStmt
	{
		$$ = $1 //TODO 306
	}
|	Block
	{
		$$ = $1 //TODO 307
	}
|	IfStmt
	{
		$$ = $1 //TODO 308
	}
|	SwitchStmt
	{
		$$ = $1 //TODO 309
	}
|	SelectStmt
	{
		$$ = $1 //TODO 310
	}
|	ForStmt
	{
		$$ = $1 //TODO 311
	}
|	DeferStmt
	{
		$$ = $1 //TODO 312
	}

StatementList:
	StatementList1 StatementList2
	{
		$$ = []StatementList{$1, $2} //TODO 313
	}

StatementList1:
	/* EMPTY */
	{
		$$ = nil //TODO 314
	}
|	Statement
	{
		$$ = $1 //TODO 315
	}

StatementList2:
	/* EMPTY */
	{
		$$ = []StatementList2(nil) //TODO 316
	}
|	StatementList2 ';' StatementList21
	{
		$$ = append($1.([]StatementList2), ";", $3) //TODO 317
	}

StatementList21:
	/* EMPTY */
	{
		$$ = nil //TODO 318
	}
|	Statement
	{
		$$ = $1 //TODO 319
	}

StructType:
	STRUCT '{' StructType1 '}'
	{
		$$ = []StructType{"struct", "{", $3, "}"} //TODO 320
	}

StructType1:
	/* EMPTY */
	{
		$$ = []StructType1(nil) //TODO 321
	}
|	StructType1 FieldDecl ';'
	{
		$$ = append($1.([]StructType1), $2, ";") //TODO 322
	}

SwitchStmt:
	ExprSwitchStmt
	{
		$$ = $1 //TODO 323
	}
|	TypeSwitchStmt
	{
		$$ = $1 //TODO 324
	}

Tag:
	STRING_LIT
	{
		$$ = $1 //TODO 325
	}

TopLevelDecl:
	Declaration
	{
		$$ = $1 //TODO 326
	}
|	FunctionDecl
	{
		$$ = $1 //TODO 327
	}
|	MethodDecl
	{
		$$ = $1 //TODO 328
	}

Type:
	TypeName
	{
		$$ = $1 //TODO 329
	}
|	TypeLit
	{
		$$ = $1 //TODO 330
	}
|	'(' Type ')'
	{
		$$ = []Type{"(", $2, ")"} //TODO 331
	}

TypeAssertion:
	'.' '(' Type ')'
	{
		$$ = []TypeAssertion{".", "(", $3, ")"} //TODO 332
	}

TypeCaseClause:
	TypeSwitchCase ':' StatementList
	{
		$$ = []TypeCaseClause{$1, ":", $3} //TODO 333
	}

TypeDecl:
	TYPE TypeDecl1
	{
		$$ = []TypeDecl{"type", $2} //TODO 334
	}

TypeDecl1:
	TypeSpec
	{
		$$ = $1 //TODO 335
	}
|	'(' TypeDecl11 ')'
	{
		$$ = []TypeDecl1{"(", $2, ")"} //TODO 336
	}

TypeDecl11:
	/* EMPTY */
	{
		$$ = nil //TODO 337
	}
|	TypeSpec TypeDecl111 TypeDecl112
	{
		$$ = []TypeDecl11{$1, $2, $3} //TODO 338
	}

TypeDecl111:
	/* EMPTY */
	{
		$$ = []TypeDecl111(nil) //TODO 339
	}
|	TypeDecl111 ';' TypeSpec
	{
		$$ = append($1.([]TypeDecl111), ";", $3) //TODO 340
	}

TypeDecl112:
	/* EMPTY */
	{
		$$ = nil //TODO 341
	}
|	';'
	{
		$$ = ";" //TODO 342
	}

TypeList:
	Type TypeList1
	{
		$$ = []TypeList{$1, $2} //TODO 343
	}

TypeList1:
	/* EMPTY */
	{
		$$ = []TypeList1(nil) //TODO 344
	}
|	TypeList1 ',' Type
	{
		$$ = append($1.([]TypeList1), ",", $3) //TODO 345
	}

TypeLit:
	ArrayType
	{
		$$ = $1 //TODO 346
	}
|	StructType
	{
		$$ = $1 //TODO 347
	}
|	PointerType
	{
		$$ = $1 //TODO 348
	}
|	FunctionType
	{
		$$ = $1 //TODO 349
	}
|	InterfaceType
	{
		$$ = $1 //TODO 350
	}
|	SliceType
	{
		$$ = $1 //TODO 351
	}
|	MapType
	{
		$$ = $1 //TODO 352
	}
|	ChannelType
	{
		$$ = $1 //TODO 353
	}

TypeName:
	IDENTIFIER
	{
		$$ = $1 //TODO 354
	}
|	QualifiedIdent
	{
		$$ = $1 //TODO 355
	}

TypeSpec:
	IDENTIFIER Type
	{
		$$ = []TypeSpec{$1, $2} //TODO 356
	}

TypeSwitchCase:
	CASE TypeList
	{
		$$ = []TypeSwitchCase{"case", $2} //TODO 357
	}
|	DEFAULT
	{
		$$ = "default" //TODO 358
	}

TypeSwitchGuard:
	PrimaryExpr '.' '(' TYPE ')'
	{
		$$ = []TypeSwitchGuard{$1, ".", "(", "type", ")"} //TODO 359
	}
|	IDENTIFIER COLAS PrimaryExpr '.' '(' TYPE ')'
	{
		$$ = []TypeSwitchGuard{$1, $2, $3, ".", "(", "type", ")"} //TODO 360
	}

TypeSwitchStmt:
	SWITCH TypeSwitchGuard LBR TypeSwitchStmt1 '}'
	{
		$$ = []TypeSwitchStmt{"switch", $2, $3, $4, "}"} //TODO 361
	}
|	SWITCH SimpleStmt ';' TypeSwitchGuard LBR TypeSwitchStmt2 '}'
	{
		$$ = []TypeSwitchStmt{"switch", $2, ";", $4, $5, $6, "}"} //TODO 362
	}

TypeSwitchStmt1:
	/* EMPTY */
	{
		$$ = []TypeSwitchStmt1(nil) //TODO 363
	}
|	TypeSwitchStmt1 TypeCaseClause
	{
		$$ = append($1.([]TypeSwitchStmt1), $2) //TODO 364
	}

TypeSwitchStmt2:
	/* EMPTY */
	{
		$$ = []TypeSwitchStmt2(nil) //TODO 365
	}
|	TypeSwitchStmt2 TypeCaseClause
	{
		$$ = append($1.([]TypeSwitchStmt2), $2) //TODO 366
	}

UnaryExpr:
	PrimaryExpr
	{
		$$ = $1 //TODO 367
	}
|	UnaryExpr1 UnaryExpr
	{
		$$ = []UnaryExpr{$1, $2} //TODO 368
	}

UnaryExpr1:
	'!'
	{
		$$ = "!" //TODO 369
	}
|	'&'
	{
		$$ = "&" //TODO 370
	}
|	'*'
	{
		$$ = "*" //TODO 371
	}
|	'+'
	{
		$$ = "+" //TODO 372
	}
|	'-'
	{
		$$ = "-" //TODO 373
	}
|	COMM
	{
		$$ = $1 //TODO 374
	}
|	'^'
	{
		$$ = "^" //TODO 375
	}

Value:
	Expression
	{
		$$ = $1 //TODO 376
	}
|	LiteralValue
	{
		$$ = $1 //TODO 377
	}

VarDecl:
	VAR VarDecl1
	{
		$$ = []VarDecl{"var", $2} //TODO 378
	}

VarDecl1:
	VarSpec
	{
		$$ = $1 //TODO 379
	}
|	'(' VarDecl11 ')'
	{
		$$ = []VarDecl1{"(", $2, ")"} //TODO 380
	}

VarDecl11:
	/* EMPTY */
	{
		$$ = nil //TODO 381
	}
|	VarSpec VarDecl111 VarDecl112
	{
		$$ = []VarDecl11{$1, $2, $3} //TODO 382
	}

VarDecl111:
	/* EMPTY */
	{
		$$ = []VarDecl111(nil) //TODO 383
	}
|	VarDecl111 ';' VarSpec
	{
		$$ = append($1.([]VarDecl111), ";", $3) //TODO 384
	}

VarDecl112:
	/* EMPTY */
	{
		$$ = nil //TODO 385
	}
|	';'
	{
		$$ = ";" //TODO 386
	}

VarSpec:
	IdentifierList VarSpec1
	{
		$$ = []VarSpec{$1, $2} //TODO 387
	}

VarSpec1:
	Type VarSpec11
	{
		$$ = []VarSpec1{$1, $2} //TODO 388
	}
|	'=' ExpressionList
	{
		$$ = []VarSpec1{"=", $2} //TODO 389
	}

VarSpec11:
	/* EMPTY */
	{
		$$ = nil //TODO 390
	}
|	'=' ExpressionList
	{
		$$ = []VarSpec11{"=", $2} //TODO 391
	}

%%

//TODO remove demo stuff below

var _parserResult interface{}

type (
	AnonymousField interface{}
	ArgumentList interface{}
	ArgumentList1 interface{}
	ArrayLength interface{}
	ArrayType interface{}
	Assignment interface{}
	Assignment1 interface{}
	BaseType interface{}
	BasicLit interface{}
	Block interface{}
	Block2 interface{}
	BreakStmt interface{}
	BreakStmt1 interface{}
	BuiltinArgs interface{}
	BuiltinArgs1 interface{}
	BuiltinCall interface{}
	BuiltinCall1 interface{}
	BuiltinCall11 interface{}
	Call interface{}
	Call1 interface{}
	Call11 interface{}
	Channel interface{}
	ChannelType interface{}
	ChannelType1 interface{}
	ChannelType11 interface{}
	CommCase interface{}
	CommCase1 interface{}
	CommClause interface{}
	CompositeLit interface{}
	Condition interface{}
	ConstDecl interface{}
	ConstDecl1 interface{}
	ConstDecl11 interface{}
	ConstDecl111 interface{}
	ConstDecl112 interface{}
	ConstSpec interface{}
	ConstSpec1 interface{}
	ConstSpec11 interface{}
	ContinueStmt interface{}
	ContinueStmt1 interface{}
	Conversion interface{}
	Conversion1 interface{}
	Declaration interface{}
	DeferStmt interface{}
	Element interface{}
	ElementList interface{}
	ElementList1 interface{}
	ElementType interface{}
	ExprCaseClause interface{}
	ExprSwitchCase interface{}
	ExprSwitchStmt interface{}
	ExprSwitchStmt1 interface{}
	ExprSwitchStmt2 interface{}
	ExprSwitchStmt3 interface{}
	ExprSwitchStmt4 interface{}
	Expression interface{}
	Expression1 interface{}
	ExpressionList interface{}
	ExpressionList1 interface{}
	ExpressionStmt interface{}
	FallthroughStmt interface{}
	FieldDecl interface{}
	FieldDecl1 interface{}
	FieldDecl2 interface{}
	ForClause interface{}
	ForClause1 interface{}
	ForClause2 interface{}
	ForClause3 interface{}
	ForStmt interface{}
	ForStmt1 interface{}
	ForStmt11 interface{}
	Function interface{}
	FunctionBody interface{}
	FunctionDecl interface{}
	FunctionDecl1 interface{}
	FunctionLit interface{}
	FunctionType interface{}
	GoStmt interface{}
	GotoStmt interface{}
	IdentifierList interface{}
	IdentifierList1 interface{}
	IfStmt interface{}
	IfStmt1 interface{}
	IfStmt11 interface{}
	IfStmt2 interface{}
	IfStmt21 interface{}
	ImportDecl interface{}
	ImportDecl1 interface{}
	ImportDecl11 interface{}
	ImportDecl111 interface{}
	ImportDecl112 interface{}
	ImportPath interface{}
	ImportSpec interface{}
	ImportSpec1 interface{}
	ImportSpec11 interface{}
	IncDecStmt interface{}
	IncDecStmt1 interface{}
	Index interface{}
	InitStmt interface{}
	InterfaceType interface{}
	InterfaceType1 interface{}
	InterfaceTypeName interface{}
	KeyType interface{}
	LabeledStmt interface{}
	Literal interface{}
	LiteralType interface{}
	LiteralValue interface{}
	LiteralValue1 interface{}
	LiteralValue11 interface{}
	MapType interface{}
	MethodDecl interface{}
	MethodDecl1 interface{}
	MethodExpr interface{}
	MethodSpec interface{}
	Operand interface{}
	OperandName interface{}
	PackageClause interface{}
	ParameterDecl interface{}
	ParameterList interface{}
	ParameterList1 interface{}
	Parameters interface{}
	Parameters1 interface{}
	Parameters11 interface{}
	PointerType interface{}
	PostStmt interface{}
	PrimaryExpr interface{}
	QualifiedIdent interface{}
	RangeClause interface{}
	RangeClause1 interface{}
	Receiver interface{}
	ReceiverType interface{}
	RecvExpr interface{}
	RecvStmt interface{}
	RecvStmt1 interface{}
	Result interface{}
	ReturnStmt interface{}
	ReturnStmt1 interface{}
	SelectStmt interface{}
	SelectStmt1 interface{}
	Selector interface{}
	SendStmt interface{}
	ShortVarDecl interface{}
	Signature interface{}
	Signature1 interface{}
	SimpleStmt interface{}
	Slice interface{}
	Slice1 interface{}
	Slice11 interface{}
	Slice12 interface{}
	Slice2 interface{}
	Slice21 interface{}
	SliceType interface{}
	SourceFile interface{}
	SourceFile1 interface{}
	SourceFile2 interface{}
	Start interface{}
	Statement interface{}
	StatementList interface{}
	StatementList1 interface{}
	StatementList2 interface{}
	StatementList21 interface{}
	StructType interface{}
	StructType1 interface{}
	SwitchStmt interface{}
	Tag interface{}
	TopLevelDecl interface{}
	Type interface{}
	TypeAssertion interface{}
	TypeCaseClause interface{}
	TypeDecl interface{}
	TypeDecl1 interface{}
	TypeDecl11 interface{}
	TypeDecl111 interface{}
	TypeDecl112 interface{}
	TypeList interface{}
	TypeList1 interface{}
	TypeLit interface{}
	TypeName interface{}
	TypeSpec interface{}
	TypeSwitchCase interface{}
	TypeSwitchGuard interface{}
	TypeSwitchStmt interface{}
	TypeSwitchStmt1 interface{}
	TypeSwitchStmt2 interface{}
	UnaryExpr interface{}
	UnaryExpr1 interface{}
	Value interface{}
	VarDecl interface{}
	VarDecl1 interface{}
	VarDecl11 interface{}
	VarDecl111 interface{}
	VarDecl112 interface{}
	VarSpec interface{}
	VarSpec1 interface{}
	VarSpec11 interface{}
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
