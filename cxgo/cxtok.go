package main

import (
	"flag"
	"fmt"
	"os"
)

type cgoTokOptions struct {
	src   string
	out   string
	force bool
}

func registerFlags(options *cgoTokOptions) {
	flag.StringVar(&options.src, "i", "", "PATH to input file")
	flag.StringVar(&options.out, "o", "", "PATH to output file")
	flag.BoolVar(&options.force, "f", false, "Overwrite output file if it exists")
}

func main() {
	var sym yySymType
	var options cgoTokOptions
	var r *os.File
	var w *os.File
	var err error

	registerFlags(&options)
	flag.Parse()
	if options.src == "" {
		r = os.Stdin
	} else {
		r, err = os.Open(options.src)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening %s : %s", options.src, err)
			os.Exit(1)
		}
		defer r.Close()
	}

	outFlags := os.O_WRONLY | os.O_CREATE
	if options.force {
		outFlags |= os.O_EXCL
	}
	if options.out == "" {
		w = os.Stdout
	} else {
		w, err = os.OpenFile(options.out, outFlags, 0555)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening %s : %s", options.src, err)
			os.Exit(1)
		}
		defer w.Close()
	}

	lex := NewLexer(r)
	token := lex.Lex(&sym)
	for token >= 0 {
		fmt.Fprintln(w, "%s %s", tokenName(token), sym.tok)
		lex.Lex(&sym)
	}
}

func tokenName(token int) string {
	switch token {
	case ADDR:
		return "  ADDR"
	case ADD_ASSIGN:
		return "ADDSET"
	case ADD_OP:
		return " ADDOP"
	case AFF:
		return "   AFF"
	case AFFVAR:
		return "AFFVAR"
	case AND:
		return "   AND"
	case AND_ASSIGN:
		return "ANDSET"
	case AND_OP:
		return " ANDOP"
	case ASSIGN:
		return "  ASGN"
	case BASICTYPE:
		return " BASIC"
	case BITANDEQ:
		return "BANDEQ"
	case BITCLEAR_OP:
		return "BCLROP"
	case BITOREQ:
		return " BOREQ"
	case BITOR_OP:
		return " BOROP"
	case BITXOREQ:
		return "BXOREQ"
	case BITXOR_OP:
		return "BXOROP"
	case BOOL:
		return "  BOOL"
	case BOOLEAN_LITERAL:
		return "BOOLLT"
	case BREAK:
		return " BREAK"
	case BYTE:
		return "  BYTE"
	case BYTE_LITERAL:
		return "BYTELT"
	case CAFF:
		return "  CAFF"
	case CASE:
		return "  CASE"
	case CASSIGN:
		return "CASSGN"
	case CLAUSES:
		return "CLAUSE"
	case COLON:
		return " COLON"
	case COMMA:
		return " COMMA"
	case COMMENT:
		return "COMMNT"
	case CONST:
		return " CONST"
	case CONTINUE:
		return "CONTNU"
	case DEC_OP:
		return " DECOP"
	case DEF:
		return "   DEF"
	case DEFAULT:
		return "DFAULT"
	case DIVEQ:
		return " DIVEQ"
	case DIV_ASSIGN:
		return "DIVSET"
	case DIV_OP:
		return " DIVOP"
	case DOUBLE_LITERAL:
		return "DBLLIT"
	case DPROGRAM:
		return " DPROG"
	case DSTACK:
		return "DSTACK"
	case DSTATE:
		return "DSTATE"
	case ELSE:
		return "  ELSE"
	case ENUM:
		return "  ENUM"
	case EQUAL:
		return " EQUAL"
	case EQUALWORD:
		return "EQWORD"
	case EQ_OP:
		return "  EQOP"
	case EXP:
		return "   EXP"
	case EXPEQ:
		return " EXPEQ"
	case EXPR:
		return "  EXPR"
	case F32:
		return "   F32"
	case F64:
		return "   F64"
	case FIELD:
		return " FIELD"
	case FLOAT_LITERAL:
		return "FLOATL"
	case FOR:
		return "   FOR"
	case FUNC:
		return "  FUNC"
	case GE_OP:
		return "  GEOP"
	case GOTO:
		return "  GOTO"
	case GTEQ_OP:
		return "GTEQOP"
	case GTHANEQ:
		return "GTHNEQ"
	case GTHANWORD:
		return " GTHNW"
	case GT_OP:
		return "  GTOP"
	case I16:
		return "   I16"
	case I32:
		return "   I32"
	case I64:
		return "   I64"
	case I8:
		return "    I8"
	case IDENTIFIER:
		return " IDENT"
	case IF:
		return "    IF"
	case IMPORT:
		return "IMPORT"
	case INC_OP:
		return " INCOP"
	case INFER:
		return " INFER"
	case INPUT:
		return " INPUT"
	case INT_LITERAL:
		return "INTLIT"
	case LBRACE:
		return "LBRACE"
	case LBRACK:
		return "LBRACK"
	case LEFTSHIFT:
		return "LSHIFT"
	case LEFTSHIFTEQ:
		return " LSHEQ"
	case LEFT_ASSIGN:
		return "  LSET"
	case LEFT_OP:
		return "LEFTOP"
	case LE_OP:
		return "  LEOP"
	case LONG_LITERAL:
		return "LONGLT"
	case LPAREN:
		return "LPAREN"
	case LTEQ_OP:
		return "LTEQOP"
	case LTHANEQ:
		return "LTHNEQ"
	case LTHANWORD:
		return "LTHANW"
	case LT_OP:
		return "  LTOP"
	case MINUSEQ:
		return "MNUSEQ"
	case MINUSMINUS:
		return "MINUS2"
	case MOD_ASSIGN:
		return "MODSET"
	case MOD_OP:
		return " MODOP"
	case MULTEQ:
		return "MULTEQ"
	case MUL_ASSIGN:
		return "MULSET"
	case MUL_OP:
		return " MULOP"
	case NEG_OP:
		return " NEGOP"
	case NEW:
		return "   NEW"
	case NEWLINE:
		return "NEWLIN"
	case NE_OP:
		return "  NEOP"
	case NOT:
		return "   NOT"
	case OBJECT:
		return "OBJECT"
	case OBJECTS:
		return "OBJCTS"
	case OP:
		return "    OP"
	case OR:
		return "    OR"
	case OR_ASSIGN:
		return " ORSET"
	case OR_OP:
		return "  OROP"
	case OUTPUT:
		return "OUTPUT"
	case PACKAGE:
		return "PACKAG"
	case PERIOD:
		return "PERIOD"
	case PLUSEQ:
		return "PLUSEQ"
	case PLUSPLUS:
		return " PLUS2"
	case PSTEP:
		return " PSTEP"
	case PTR_OP:
		return " PTROP"
	case RBRACE:
		return "RBRACE"
	case RBRACK:
		return "RBRACK"
	case REF_OP:
		return " REFOP"
	case REM:
		return "   REM"
	case REMAINDER:
		return "REMNDR"
	case REMAINDEREQ:
		return "RMDREQ"
	case RETURN:
		return "RETURN"
	case RIGHTSHIFT:
		return "RSHIFT"
	case RIGHTSHIFTEQ:
		return " RSHEQ"
	case RIGHT_ASSIGN:
		return "  RSET"
	case RIGHT_OP:
		return "RGHTOP"
	case RPAREN:
		return "RPAREN"
	case SEMICOLON:
		return "SCOLON"
	case SFUNC:
		return " SFUNC"
	case SPACKAGE:
		return "SPACKG"
	case SSTRUCT:
		return "SSTRCT"
	case STEP:
		return "  STEP"
	case STR:
		return "   STR"
	case STRING_LITERAL:
		return "STRLIT"
	case STRUCT:
		return "STRUCT"
	case SUB_ASSIGN:
		return "SUBSET"
	case SUB_OP:
		return " SUBOP"
	case SWITCH:
		return "SWITCH"
	case TAG:
		return "   TAG"
	case TSTEP:
		return " TSTEP"
	case TYPE:
		return "  TYPE"
	case TYPSTRUCT:
		return "TSTRCT"
	case UI16:
		return "  UI16"
	case UI32:
		return "  UI32"
	case UI64:
		return "  UI64"
	case UI8:
		return "   UI8"
	case UNEQUAL:
		return "  UNEQ"
	case UNION:
		return " UNION"
	case VALUE:
		return " VALUE"
	case VAR:
		return "   VAR"
	case XOR_ASSIGN:
		return "XORSET"
	}
}
