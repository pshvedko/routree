%{

package router

import (
 	"fmt"
 	"unicode"
)

var regs = make([]int, 26)
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	val int
	pos map[int]bool
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr number

// same for terminals
%token <val> DIGIT LETTER

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%%

list	: /* empty */
	| list stat '\n'
	;

stat	:    expr
		{
			fmt.Printf( "= %d\n", $1 );
		}
	|    LETTER '=' expr
		{
			regs[$1]  =  $3
		}
	;

expr	:    '(' expr ')'
		{ $$  =  $2 }
	|    expr '+' expr
		{ $$  =  $1 + $3 }
	|    expr '-' expr
		{ $$  =  $1 - $3 }
	|    expr '*' expr
		{ $$  =  $1 * $3 }
	|    expr '/' expr
		{ $$  =  $1 / $3 }
	|    expr '%' expr
		{ $$  =  $1 % $3 }
	|    expr '&' expr
		{ $$  =  $1 & $3 }
	|    expr '|' expr
		{ $$  =  $1 | $3 }
	|    '-'  expr        %prec  UMINUS
		{ $$  = -$2  }
	|    LETTER
		{ $$  = regs[$1] }
	|    number
	;

number	:    DIGIT
		{
			$$ = $1;
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
		}
	|    number DIGIT
		{ $$ = base * $1 + $2 }
	;

%%      /*  start  of  programs  */

type CalcLex struct {
	E string
	pos int
}


func (l *CalcLex) Lex(lval *CalcSymType) int {
	var c rune = ' '
	for c == ' ' {
		if l.pos == len(l.E) {
			return 0
		}
		c = rune(l.E[l.pos])
		l.pos += 1
	}

	if unicode.IsDigit(c) {
		lval.val = int(c) - '0'
		return DIGIT
	} else if unicode.IsLower(c) {
		lval.val = int(c) - 'a'
		return LETTER
	}
	return int(c)
}

func (l *CalcLex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}
