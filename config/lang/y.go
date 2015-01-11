//line lang.y:6
package lang

import __yyfmt__ "fmt"

//line lang.y:6
import (
	"github.com/hashicorp/terraform/config/lang/ast"
)

//line lang.y:14
type parserSymType struct {
	yys      int
	node     ast.Node
	nodeList []ast.Node
	str      string
}

const STRING = 57346
const IDENTIFIER = 57347
const PROGRAM_BRACKET_LEFT = 57348
const PROGRAM_BRACKET_RIGHT = 57349
const PROGRAM_STRING_START = 57350
const PROGRAM_STRING_END = 57351
const PAREN_LEFT = 57352
const PAREN_RIGHT = 57353
const COMMA = 57354

var parserToknames = []string{
	"STRING",
	"IDENTIFIER",
	"PROGRAM_BRACKET_LEFT",
	"PROGRAM_BRACKET_RIGHT",
	"PROGRAM_STRING_START",
	"PROGRAM_STRING_END",
	"PAREN_LEFT",
	"PAREN_RIGHT",
	"COMMA",
}
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line lang.y:103

//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 14
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 21

var parserAct = []int{

	9, 16, 17, 13, 3, 12, 1, 8, 6, 11,
	7, 6, 14, 7, 15, 8, 10, 2, 18, 4,
	5,
}
var parserPact = []int{

	7, -1000, 7, -1000, -1000, -1000, -1000, 4, -1000, -2,
	7, -7, -1000, 4, -10, -1000, -1000, 4, -1000,
}
var parserPgo = []int{

	0, 0, 20, 19, 16, 4, 12, 6,
}
var parserR1 = []int{

	0, 7, 4, 4, 5, 5, 2, 1, 1, 1,
	6, 6, 6, 3,
}
var parserR2 = []int{

	0, 1, 1, 2, 1, 1, 3, 1, 1, 4,
	0, 3, 1, 1,
}
var parserChk = []int{

	-1000, -7, -4, -5, -3, -2, 4, 6, -5, -1,
	-4, 5, 7, 10, -6, -1, 11, 12, -1,
}
var parserDef = []int{

	0, -2, 1, 2, 4, 5, 13, 0, 3, 0,
	7, 8, 6, 10, 0, 12, 9, 0, 11,
}
var parserTok1 = []int{

	1,
}
var parserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
}
var parserTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var parserDebug = 0

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

const parserFlag = -1000

func parserTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(parserToknames) {
		if parserToknames[c-4] != "" {
			return parserToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserlex1(lex parserLexer, lval *parserSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		c = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			c = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		c = parserTok3[i+0]
		if c == char {
			c = parserTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(c), uint(char))
	}
	return c
}

func parserParse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parserchar), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserchar < 0 {
		parserchar = parserlex1(parserlex, &parserlval)
	}
	parsern += parserchar
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parserchar { /* valid shift */
		parserchar = -1
		parserVAL = parserlval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserchar < 0 {
			parserchar = parserlex1(parserlex, &parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parserchar {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error("syntax error")
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parserchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parserchar))
			}
			if parserchar == parserEofCode {
				goto ret1
			}
			parserchar = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		//line lang.y:31
		{
			parserResult = parserS[parserpt-0].node
		}
	case 2:
		//line lang.y:37
		{
			parserVAL.node = parserS[parserpt-0].node
		}
	case 3:
		//line lang.y:41
		{
			var result []ast.Node
			if c, ok := parserS[parserpt-1].node.(*ast.Concat); ok {
				result = append(c.Exprs, parserS[parserpt-0].node)
			} else {
				result = []ast.Node{parserS[parserpt-1].node, parserS[parserpt-0].node}
			}

			parserVAL.node = &ast.Concat{
				Exprs: result,
			}
		}
	case 4:
		//line lang.y:56
		{
			parserVAL.node = parserS[parserpt-0].node
		}
	case 5:
		//line lang.y:60
		{
			parserVAL.node = parserS[parserpt-0].node
		}
	case 6:
		//line lang.y:66
		{
			parserVAL.node = parserS[parserpt-1].node
		}
	case 7:
		//line lang.y:72
		{
			parserVAL.node = parserS[parserpt-0].node
		}
	case 8:
		//line lang.y:76
		{
			parserVAL.node = &ast.VariableAccess{Name: parserS[parserpt-0].str}
		}
	case 9:
		//line lang.y:80
		{
			parserVAL.node = &ast.Call{Func: parserS[parserpt-3].str, Args: parserS[parserpt-1].nodeList}
		}
	case 10:
		//line lang.y:85
		{
			parserVAL.nodeList = nil
		}
	case 11:
		//line lang.y:89
		{
			parserVAL.nodeList = append(parserS[parserpt-2].nodeList, parserS[parserpt-0].node)
		}
	case 12:
		//line lang.y:93
		{
			parserVAL.nodeList = append(parserVAL.nodeList, parserS[parserpt-0].node)
		}
	case 13:
		//line lang.y:99
		{
			parserVAL.node = &ast.LiteralNode{Value: parserS[parserpt-0].str, Type: ast.TypeString}
		}
	}
	goto parserstack /* stack new state and value */
}