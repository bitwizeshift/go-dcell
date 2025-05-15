// Code generated from DCell.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DCell

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DCellParser struct {
	*antlr.BaseParser
}

var DCellParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dcellParserInit() {
	staticData := &DCellParserStaticData
	staticData.LiteralNames = []string{
		"", "'.'", "'['", "']'", "'is'", "'not'", "'in'", "'('", "')'", "'!'",
		"'~'", "'+'", "'-'", "'**'", "'*'", "'/'", "'//'", "'%'", "'&&'", "'and'",
		"'||'", "'or'", "'<->'", "'implies'", "'<<'", "'>>'", "'&'", "'^'",
		"'|'", "'<='", "'<'", "'>'", "'>='", "'=='", "'!='", "'?'", "':'", "'?:'",
		"'??'", "'as'", "','", "'true'", "'false'", "'null'", "'int'", "'float'",
		"'string'", "'bool'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER",
		"DECIMAL_INTEGER", "HEX_INTEGER", "OCTAL_INTEGER", "BINARY_INTEGER",
		"DECIMAL_FLOAT", "SCIENTIFIC_FLOAT", "SINGLE_QUOTE_STRING", "DOUBLE_QUOTE_STRING",
		"TRIPLE_QUOTE_STRING", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"program", "expression", "term", "invocation", "parameterList", "identifier",
		"index", "literal", "type", "list", "string", "integer", "float",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 192, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 42, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 108, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 114,
		8, 1, 10, 1, 12, 1, 117, 9, 1, 1, 2, 1, 2, 3, 2, 121, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 128, 8, 3, 1, 3, 1, 3, 3, 3, 132, 8, 3, 1, 4, 1,
		4, 1, 4, 5, 4, 137, 8, 4, 10, 4, 12, 4, 140, 9, 4, 1, 5, 1, 5, 1, 6, 3,
		6, 145, 8, 6, 1, 6, 1, 6, 3, 6, 149, 8, 6, 1, 6, 3, 6, 152, 8, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 160, 8, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 5, 9, 168, 8, 9, 10, 9, 12, 9, 171, 9, 9, 3, 9, 173, 8, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 180, 8, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 3, 11, 186, 8, 11, 1, 12, 1, 12, 3, 12, 190, 8, 12, 1, 12, 0, 1,
		2, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 12, 2, 0, 5, 5,
		9, 9, 1, 0, 11, 12, 1, 0, 14, 17, 1, 0, 18, 19, 1, 0, 20, 21, 1, 0, 22,
		23, 1, 0, 24, 25, 1, 0, 27, 28, 1, 0, 29, 32, 1, 0, 33, 34, 1, 0, 41, 42,
		1, 0, 44, 47, 224, 0, 26, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 120, 1, 0,
		0, 0, 6, 131, 1, 0, 0, 0, 8, 133, 1, 0, 0, 0, 10, 141, 1, 0, 0, 0, 12,
		151, 1, 0, 0, 0, 14, 159, 1, 0, 0, 0, 16, 161, 1, 0, 0, 0, 18, 163, 1,
		0, 0, 0, 20, 179, 1, 0, 0, 0, 22, 185, 1, 0, 0, 0, 24, 189, 1, 0, 0, 0,
		26, 27, 3, 2, 1, 0, 27, 28, 5, 0, 0, 1, 28, 1, 1, 0, 0, 0, 29, 30, 6, 1,
		-1, 0, 30, 42, 3, 4, 2, 0, 31, 32, 5, 7, 0, 0, 32, 33, 3, 2, 1, 0, 33,
		34, 5, 8, 0, 0, 34, 42, 1, 0, 0, 0, 35, 36, 7, 0, 0, 0, 36, 42, 3, 2, 1,
		18, 37, 38, 5, 10, 0, 0, 38, 42, 3, 2, 1, 17, 39, 40, 7, 1, 0, 0, 40, 42,
		3, 2, 1, 16, 41, 29, 1, 0, 0, 0, 41, 31, 1, 0, 0, 0, 41, 35, 1, 0, 0, 0,
		41, 37, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 42, 115, 1, 0, 0, 0, 43, 47, 10,
		20, 0, 0, 44, 48, 5, 6, 0, 0, 45, 46, 5, 5, 0, 0, 46, 48, 5, 6, 0, 0, 47,
		44, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 114, 3, 2,
		1, 21, 50, 51, 10, 15, 0, 0, 51, 52, 5, 13, 0, 0, 52, 114, 3, 2, 1, 16,
		53, 54, 10, 14, 0, 0, 54, 55, 7, 2, 0, 0, 55, 114, 3, 2, 1, 15, 56, 57,
		10, 13, 0, 0, 57, 58, 7, 1, 0, 0, 58, 114, 3, 2, 1, 14, 59, 60, 10, 12,
		0, 0, 60, 61, 7, 3, 0, 0, 61, 114, 3, 2, 1, 13, 62, 63, 10, 11, 0, 0, 63,
		64, 7, 4, 0, 0, 64, 114, 3, 2, 1, 12, 65, 66, 10, 10, 0, 0, 66, 67, 7,
		5, 0, 0, 67, 114, 3, 2, 1, 11, 68, 69, 10, 9, 0, 0, 69, 70, 7, 6, 0, 0,
		70, 114, 3, 2, 1, 10, 71, 72, 10, 8, 0, 0, 72, 73, 5, 26, 0, 0, 73, 114,
		3, 2, 1, 9, 74, 75, 10, 7, 0, 0, 75, 76, 7, 7, 0, 0, 76, 114, 3, 2, 1,
		8, 77, 78, 10, 6, 0, 0, 78, 79, 7, 8, 0, 0, 79, 114, 3, 2, 1, 7, 80, 81,
		10, 5, 0, 0, 81, 82, 7, 9, 0, 0, 82, 114, 3, 2, 1, 6, 83, 84, 10, 4, 0,
		0, 84, 85, 5, 35, 0, 0, 85, 86, 3, 2, 1, 0, 86, 87, 5, 36, 0, 0, 87, 88,
		3, 2, 1, 5, 88, 114, 1, 0, 0, 0, 89, 90, 10, 3, 0, 0, 90, 91, 5, 37, 0,
		0, 91, 114, 3, 2, 1, 4, 92, 93, 10, 2, 0, 0, 93, 94, 5, 38, 0, 0, 94, 114,
		3, 2, 1, 3, 95, 96, 10, 23, 0, 0, 96, 97, 5, 1, 0, 0, 97, 114, 3, 6, 3,
		0, 98, 99, 10, 22, 0, 0, 99, 100, 5, 2, 0, 0, 100, 101, 3, 12, 6, 0, 101,
		102, 5, 3, 0, 0, 102, 114, 1, 0, 0, 0, 103, 107, 10, 21, 0, 0, 104, 105,
		5, 4, 0, 0, 105, 108, 5, 5, 0, 0, 106, 108, 5, 4, 0, 0, 107, 104, 1, 0,
		0, 0, 107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 114, 3, 16, 8, 0,
		110, 111, 10, 1, 0, 0, 111, 112, 5, 39, 0, 0, 112, 114, 3, 16, 8, 0, 113,
		43, 1, 0, 0, 0, 113, 50, 1, 0, 0, 0, 113, 53, 1, 0, 0, 0, 113, 56, 1, 0,
		0, 0, 113, 59, 1, 0, 0, 0, 113, 62, 1, 0, 0, 0, 113, 65, 1, 0, 0, 0, 113,
		68, 1, 0, 0, 0, 113, 71, 1, 0, 0, 0, 113, 74, 1, 0, 0, 0, 113, 77, 1, 0,
		0, 0, 113, 80, 1, 0, 0, 0, 113, 83, 1, 0, 0, 0, 113, 89, 1, 0, 0, 0, 113,
		92, 1, 0, 0, 0, 113, 95, 1, 0, 0, 0, 113, 98, 1, 0, 0, 0, 113, 103, 1,
		0, 0, 0, 113, 110, 1, 0, 0, 0, 114, 117, 1, 0, 0, 0, 115, 113, 1, 0, 0,
		0, 115, 116, 1, 0, 0, 0, 116, 3, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 118,
		121, 3, 14, 7, 0, 119, 121, 3, 6, 3, 0, 120, 118, 1, 0, 0, 0, 120, 119,
		1, 0, 0, 0, 121, 5, 1, 0, 0, 0, 122, 132, 3, 10, 5, 0, 123, 132, 5, 14,
		0, 0, 124, 125, 3, 10, 5, 0, 125, 127, 5, 7, 0, 0, 126, 128, 3, 8, 4, 0,
		127, 126, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129,
		130, 5, 8, 0, 0, 130, 132, 1, 0, 0, 0, 131, 122, 1, 0, 0, 0, 131, 123,
		1, 0, 0, 0, 131, 124, 1, 0, 0, 0, 132, 7, 1, 0, 0, 0, 133, 138, 3, 2, 1,
		0, 134, 135, 5, 40, 0, 0, 135, 137, 3, 2, 1, 0, 136, 134, 1, 0, 0, 0, 137,
		140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 9, 1,
		0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 48, 0, 0, 142, 11, 1, 0, 0,
		0, 143, 145, 3, 2, 1, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 148, 5, 36, 0, 0, 147, 149, 3, 2, 1, 0, 148, 147,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 152, 3, 2,
		1, 0, 151, 144, 1, 0, 0, 0, 151, 150, 1, 0, 0, 0, 152, 13, 1, 0, 0, 0,
		153, 160, 3, 20, 10, 0, 154, 160, 3, 22, 11, 0, 155, 160, 3, 24, 12, 0,
		156, 160, 7, 10, 0, 0, 157, 160, 5, 43, 0, 0, 158, 160, 3, 18, 9, 0, 159,
		153, 1, 0, 0, 0, 159, 154, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 159, 156,
		1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 158, 1, 0, 0, 0, 160, 15, 1, 0,
		0, 0, 161, 162, 7, 11, 0, 0, 162, 17, 1, 0, 0, 0, 163, 172, 5, 2, 0, 0,
		164, 169, 3, 14, 7, 0, 165, 166, 5, 40, 0, 0, 166, 168, 3, 14, 7, 0, 167,
		165, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 172, 164, 1, 0,
		0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 5, 3, 0, 0,
		175, 19, 1, 0, 0, 0, 176, 180, 5, 55, 0, 0, 177, 180, 5, 56, 0, 0, 178,
		180, 5, 57, 0, 0, 179, 176, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 178,
		1, 0, 0, 0, 180, 21, 1, 0, 0, 0, 181, 186, 5, 49, 0, 0, 182, 186, 5, 50,
		0, 0, 183, 186, 5, 51, 0, 0, 184, 186, 5, 52, 0, 0, 185, 181, 1, 0, 0,
		0, 185, 182, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186,
		23, 1, 0, 0, 0, 187, 190, 5, 54, 0, 0, 188, 190, 5, 53, 0, 0, 189, 187,
		1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 25, 1, 0, 0, 0, 18, 41, 47, 107,
		113, 115, 120, 127, 131, 138, 144, 148, 151, 159, 169, 172, 179, 185, 189,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DCellParserInit initializes any static state used to implement DCellParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDCellParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DCellParserInit() {
	staticData := &DCellParserStaticData
	staticData.once.Do(dcellParserInit)
}

// NewDCellParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDCellParser(input antlr.TokenStream) *DCellParser {
	DCellParserInit()
	this := new(DCellParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DCellParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DCell.g4"

	return this
}

// DCellParser tokens.
const (
	DCellParserEOF                 = antlr.TokenEOF
	DCellParserT__0                = 1
	DCellParserT__1                = 2
	DCellParserT__2                = 3
	DCellParserT__3                = 4
	DCellParserT__4                = 5
	DCellParserT__5                = 6
	DCellParserT__6                = 7
	DCellParserT__7                = 8
	DCellParserT__8                = 9
	DCellParserT__9                = 10
	DCellParserT__10               = 11
	DCellParserT__11               = 12
	DCellParserT__12               = 13
	DCellParserT__13               = 14
	DCellParserT__14               = 15
	DCellParserT__15               = 16
	DCellParserT__16               = 17
	DCellParserT__17               = 18
	DCellParserT__18               = 19
	DCellParserT__19               = 20
	DCellParserT__20               = 21
	DCellParserT__21               = 22
	DCellParserT__22               = 23
	DCellParserT__23               = 24
	DCellParserT__24               = 25
	DCellParserT__25               = 26
	DCellParserT__26               = 27
	DCellParserT__27               = 28
	DCellParserT__28               = 29
	DCellParserT__29               = 30
	DCellParserT__30               = 31
	DCellParserT__31               = 32
	DCellParserT__32               = 33
	DCellParserT__33               = 34
	DCellParserT__34               = 35
	DCellParserT__35               = 36
	DCellParserT__36               = 37
	DCellParserT__37               = 38
	DCellParserT__38               = 39
	DCellParserT__39               = 40
	DCellParserT__40               = 41
	DCellParserT__41               = 42
	DCellParserT__42               = 43
	DCellParserT__43               = 44
	DCellParserT__44               = 45
	DCellParserT__45               = 46
	DCellParserT__46               = 47
	DCellParserIDENTIFIER          = 48
	DCellParserDECIMAL_INTEGER     = 49
	DCellParserHEX_INTEGER         = 50
	DCellParserOCTAL_INTEGER       = 51
	DCellParserBINARY_INTEGER      = 52
	DCellParserDECIMAL_FLOAT       = 53
	DCellParserSCIENTIFIC_FLOAT    = 54
	DCellParserSINGLE_QUOTE_STRING = 55
	DCellParserDOUBLE_QUOTE_STRING = 56
	DCellParserTRIPLE_QUOTE_STRING = 57
	DCellParserWS                  = 58
	DCellParserCOMMENT             = 59
)

// DCellParser rules.
const (
	DCellParserRULE_program       = 0
	DCellParserRULE_expression    = 1
	DCellParserRULE_term          = 2
	DCellParserRULE_invocation    = 3
	DCellParserRULE_parameterList = 4
	DCellParserRULE_identifier    = 5
	DCellParserRULE_index         = 6
	DCellParserRULE_literal       = 7
	DCellParserRULE_type          = 8
	DCellParserRULE_list          = 9
	DCellParserRULE_string        = 10
	DCellParserRULE_integer       = 11
	DCellParserRULE_float         = 12
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(DCellParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *DCellParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DCellParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.expression(0)
	}
	{
		p.SetState(27)
		p.Match(DCellParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LogicalNotExpressionContext struct {
	ExpressionContext
}

func NewLogicalNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalNotExpressionContext {
	var p = new(LogicalNotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalNotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type CoalesceExpressionContext struct {
	ExpressionContext
}

func NewCoalesceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoalesceExpressionContext {
	var p = new(CoalesceExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CoalesceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoalesceExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CoalesceExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type ShiftExpressionContext struct {
	ExpressionContext
}

func NewShiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ShiftExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type IsExpressionContext struct {
	ExpressionContext
}

func NewIsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsExpressionContext {
	var p = new(IsExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IsExpressionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

type PolarityExpressionContext struct {
	ExpressionContext
}

func NewPolarityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PolarityExpressionContext {
	var p = new(PolarityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PolarityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolarityExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type AdditiveExpressionContext struct {
	ExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type ContainsExpressionContext struct {
	ExpressionContext
}

func NewContainsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainsExpressionContext {
	var p = new(ContainsExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ContainsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ContainsExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type ParenthesisExpressionContext struct {
	ExpressionContext
}

func NewParenthesisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type MultiplicativeExpressionContext struct {
	ExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type LogicalOrExpressionContext struct {
	ExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type CastExpressionContext struct {
	ExpressionContext
}

func NewCastExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExpressionContext {
	var p = new(CastExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CastExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CastExpressionContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

type BitwiseOrExpressionContext struct {
	ExpressionContext
}

func NewBitwiseOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseOrExpressionContext {
	var p = new(BitwiseOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitwiseOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitwiseOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type InequalityExpressionContext struct {
	ExpressionContext
}

func NewInequalityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InequalityExpressionContext {
	var p = new(InequalityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InequalityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InequalityExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InequalityExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type InvocationExpressionContext struct {
	ExpressionContext
}

func NewInvocationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationExpressionContext {
	var p = new(InvocationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InvocationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvocationExpressionContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

type BitwiseNotExpressionContext struct {
	ExpressionContext
}

func NewBitwiseNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseNotExpressionContext {
	var p = new(BitwiseNotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitwiseNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseNotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type BitwiseAndExpressionContext struct {
	ExpressionContext
}

func NewBitwiseAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitwiseAndExpressionContext {
	var p = new(BitwiseAndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitwiseAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitwiseAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitwiseAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type LogicalAndExpressionContext struct {
	ExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type EqualityExpressionContext struct {
	ExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type ImplicationExpressionContext struct {
	ExpressionContext
}

func NewImplicationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImplicationExpressionContext {
	var p = new(ImplicationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ImplicationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplicationExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ImplicationExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type ExponentiationExpressionContext struct {
	ExpressionContext
}

func NewExponentiationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentiationExpressionContext {
	var p = new(ExponentiationExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExponentiationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentiationExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExponentiationExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type IndexExpressionContext struct {
	ExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

type ElvisExpressionContext struct {
	ExpressionContext
}

func NewElvisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElvisExpressionContext {
	var p = new(ElvisExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ElvisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElvisExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ElvisExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type TernaryExpressionContext struct {
	ExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type TermExpressionContext struct {
	ExpressionContext
}

func NewTermExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermExpressionContext {
	var p = new(TermExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TermExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermExpressionContext) Term() ITermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (p *DCellParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *DCellParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, DCellParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCellParserT__1, DCellParserT__13, DCellParserT__40, DCellParserT__41, DCellParserT__42, DCellParserIDENTIFIER, DCellParserDECIMAL_INTEGER, DCellParserHEX_INTEGER, DCellParserOCTAL_INTEGER, DCellParserBINARY_INTEGER, DCellParserDECIMAL_FLOAT, DCellParserSCIENTIFIC_FLOAT, DCellParserSINGLE_QUOTE_STRING, DCellParserDOUBLE_QUOTE_STRING, DCellParserTRIPLE_QUOTE_STRING:
		localctx = NewTermExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(30)
			p.Term()
		}

	case DCellParserT__6:
		localctx = NewParenthesisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.Match(DCellParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.expression(0)
		}
		{
			p.SetState(33)
			p.Match(DCellParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserT__4, DCellParserT__8:
		localctx = NewLogicalNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DCellParserT__4 || _la == DCellParserT__8) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(36)
			p.expression(18)
		}

	case DCellParserT__9:
		localctx = NewBitwiseNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.Match(DCellParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.expression(17)
		}

	case DCellParserT__10, DCellParserT__11:
		localctx = NewPolarityExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(39)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DCellParserT__10 || _la == DCellParserT__11) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(40)
			p.expression(16)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewContainsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				p.SetState(47)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case DCellParserT__5:
					{
						p.SetState(44)
						p.Match(DCellParserT__5)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case DCellParserT__4:
					{
						p.SetState(45)
						p.Match(DCellParserT__4)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(46)
						p.Match(DCellParserT__5)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}
				{
					p.SetState(49)
					p.expression(21)
				}

			case 2:
				localctx = NewExponentiationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(51)
					p.Match(DCellParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(52)
					p.expression(16)
				}

			case 3:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(53)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(54)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(55)
					p.expression(15)
				}

			case 4:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(57)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__10 || _la == DCellParserT__11) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(58)
					p.expression(14)
				}

			case 5:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(60)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__17 || _la == DCellParserT__18) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.expression(13)
				}

			case 6:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(63)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__19 || _la == DCellParserT__20) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(64)
					p.expression(12)
				}

			case 7:
				localctx = NewImplicationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(66)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__21 || _la == DCellParserT__22) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(67)
					p.expression(11)
				}

			case 8:
				localctx = NewShiftExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(69)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__23 || _la == DCellParserT__24) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(70)
					p.expression(10)
				}

			case 9:
				localctx = NewBitwiseAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(71)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(72)
					p.Match(DCellParserT__25)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(73)
					p.expression(9)
				}

			case 10:
				localctx = NewBitwiseOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(75)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__26 || _la == DCellParserT__27) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(76)
					p.expression(8)
				}

			case 11:
				localctx = NewInequalityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(78)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(79)
					p.expression(7)
				}

			case 12:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(81)
					_la = p.GetTokenStream().LA(1)

					if !(_la == DCellParserT__32 || _la == DCellParserT__33) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(82)
					p.expression(6)
				}

			case 13:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(84)
					p.Match(DCellParserT__34)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(85)
					p.expression(0)
				}
				{
					p.SetState(86)
					p.Match(DCellParserT__35)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(87)
					p.expression(5)
				}

			case 14:
				localctx = NewElvisExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(90)
					p.Match(DCellParserT__36)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(91)
					p.expression(4)
				}

			case 15:
				localctx = NewCoalesceExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(93)
					p.Match(DCellParserT__37)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)
					p.expression(3)
				}

			case 16:
				localctx = NewInvocationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.Match(DCellParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)
					p.Invocation()
				}

			case 17:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(99)
					p.Match(DCellParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(100)
					p.Index()
				}
				{
					p.SetState(101)
					p.Match(DCellParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 18:
				localctx = NewIsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(103)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				p.SetState(107)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(104)
						p.Match(DCellParserT__3)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(105)
						p.Match(DCellParserT__4)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case 2:
					{
						p.SetState(106)
						p.Match(DCellParserT__3)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				case antlr.ATNInvalidAltNumber:
					goto errorExit
				}
				{
					p.SetState(109)
					p.Type_()
				}

			case 19:
				localctx = NewCastExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DCellParserRULE_expression)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(111)
					p.Match(DCellParserT__38)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(112)
					p.Type_()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyAll(ctx *TermContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LiteralTermContext struct {
	TermContext
}

func NewLiteralTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralTermContext {
	var p = new(LiteralTermContext)

	InitEmptyTermContext(&p.TermContext)
	p.parser = parser
	p.CopyAll(ctx.(*TermContext))

	return p
}

func (s *LiteralTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTermContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

type InvocationTermContext struct {
	TermContext
}

func NewInvocationTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationTermContext {
	var p = new(InvocationTermContext)

	InitEmptyTermContext(&p.TermContext)
	p.parser = parser
	p.CopyAll(ctx.(*TermContext))

	return p
}

func (s *InvocationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationTermContext) Invocation() IInvocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (p *DCellParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DCellParserRULE_term)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCellParserT__1, DCellParserT__40, DCellParserT__41, DCellParserT__42, DCellParserDECIMAL_INTEGER, DCellParserHEX_INTEGER, DCellParserOCTAL_INTEGER, DCellParserBINARY_INTEGER, DCellParserDECIMAL_FLOAT, DCellParserSCIENTIFIC_FLOAT, DCellParserSINGLE_QUOTE_STRING, DCellParserDOUBLE_QUOTE_STRING, DCellParserTRIPLE_QUOTE_STRING:
		localctx = NewLiteralTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Literal()
		}

	case DCellParserT__13, DCellParserIDENTIFIER:
		localctx = NewInvocationTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Invocation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_invocation
	return p
}

func InitEmptyInvocationContext(p *InvocationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_invocation
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) CopyAll(ctx *InvocationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WildcardInvocationContext struct {
	InvocationContext
}

func NewWildcardInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WildcardInvocationContext {
	var p = new(WildcardInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *WildcardInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

type FunctionInvocationContext struct {
	InvocationContext
}

func NewFunctionInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionInvocationContext {
	var p = new(FunctionInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *FunctionInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionInvocationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionInvocationContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

type MemberInvocationContext struct {
	InvocationContext
}

func NewMemberInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberInvocationContext {
	var p = new(MemberInvocationContext)

	InitEmptyInvocationContext(&p.InvocationContext)
	p.parser = parser
	p.CopyAll(ctx.(*InvocationContext))

	return p
}

func (s *MemberInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberInvocationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (p *DCellParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DCellParserRULE_invocation)
	var _la int

	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMemberInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Identifier()
		}

	case 2:
		localctx = NewWildcardInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(DCellParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunctionInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Identifier()
		}
		{
			p.SetState(125)
			p.Match(DCellParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&287964294337814180) != 0 {
			{
				p.SetState(126)
				p.ParameterList()
			}

		}
		{
			p.SetState(129)
			p.Match(DCellParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *DCellParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DCellParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.expression(0)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DCellParserT__39 {
		{
			p.SetState(134)
			p.Match(DCellParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.expression(0)
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(DCellParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *DCellParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DCellParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(DCellParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) CopyAll(ctx *IndexContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SliceIndexContext struct {
	IndexContext
}

func NewSliceIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceIndexContext {
	var p = new(SliceIndexContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *SliceIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceIndexContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SliceIndexContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

type ExpressionIndexContext struct {
	IndexContext
}

func NewExpressionIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionIndexContext {
	var p = new(ExpressionIndexContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *ExpressionIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionIndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (p *DCellParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DCellParserRULE_index)
	var _la int

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSliceIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&287964294337814180) != 0 {
			{
				p.SetState(143)
				p.expression(0)
			}

		}
		{
			p.SetState(146)
			p.Match(DCellParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&287964294337814180) != 0 {
			{
				p.SetState(147)
				p.expression(0)
			}

		}

	case 2:
		localctx = NewExpressionIndexContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullLiteralContext struct {
	LiteralContext
}

func NewNullLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullLiteralContext {
	var p = new(NullLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *NullLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

type StringLiteralContext struct {
	LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

type ListLiteralContext struct {
	LiteralContext
}

func NewListLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListLiteralContext {
	var p = new(ListLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

type IntegerLiteralContext struct {
	LiteralContext
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

type FloatLiteralContext struct {
	LiteralContext
}

func NewFloatLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) Float() IFloatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

type BooleanLiteralContext struct {
	LiteralContext
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (p *DCellParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DCellParserRULE_literal)
	var _la int

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCellParserSINGLE_QUOTE_STRING, DCellParserDOUBLE_QUOTE_STRING, DCellParserTRIPLE_QUOTE_STRING:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.String_()
		}

	case DCellParserDECIMAL_INTEGER, DCellParserHEX_INTEGER, DCellParserOCTAL_INTEGER, DCellParserBINARY_INTEGER:
		localctx = NewIntegerLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Integer()
		}

	case DCellParserDECIMAL_FLOAT, DCellParserSCIENTIFIC_FLOAT:
		localctx = NewFloatLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Float()
		}

	case DCellParserT__40, DCellParserT__41:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(156)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DCellParserT__40 || _la == DCellParserT__41) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case DCellParserT__42:
		localctx = NewNullLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(157)
			p.Match(DCellParserT__42)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserT__1:
		localctx = NewListLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(158)
			p.List()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *DCellParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DCellParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&263882790666240) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLiteral() []ILiteralContext
	Literal(i int) ILiteralContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllLiteral() []ILiteralContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteralContext); ok {
			len++
		}
	}

	tst := make([]ILiteralContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteralContext); ok {
			tst[i] = t.(ILiteralContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Literal(i int) ILiteralContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *DCellParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DCellParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(DCellParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&287682819361079300) != 0 {
		{
			p.SetState(164)
			p.Literal()
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DCellParserT__39 {
			{
				p.SetState(165)
				p.Match(DCellParserT__39)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(166)
				p.Literal()
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(174)
		p.Match(DCellParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) CopyAll(ctx *StringContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleQuoteStringContext struct {
	StringContext
}

func NewSingleQuoteStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleQuoteStringContext {
	var p = new(SingleQuoteStringContext)

	InitEmptyStringContext(&p.StringContext)
	p.parser = parser
	p.CopyAll(ctx.(*StringContext))

	return p
}

func (s *SingleQuoteStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleQuoteStringContext) SINGLE_QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(DCellParserSINGLE_QUOTE_STRING, 0)
}

type DoubleQuoteStringContext struct {
	StringContext
}

func NewDoubleQuoteStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleQuoteStringContext {
	var p = new(DoubleQuoteStringContext)

	InitEmptyStringContext(&p.StringContext)
	p.parser = parser
	p.CopyAll(ctx.(*StringContext))

	return p
}

func (s *DoubleQuoteStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleQuoteStringContext) DOUBLE_QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(DCellParserDOUBLE_QUOTE_STRING, 0)
}

type TripleQuoteStringContext struct {
	StringContext
}

func NewTripleQuoteStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TripleQuoteStringContext {
	var p = new(TripleQuoteStringContext)

	InitEmptyStringContext(&p.StringContext)
	p.parser = parser
	p.CopyAll(ctx.(*StringContext))

	return p
}

func (s *TripleQuoteStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TripleQuoteStringContext) TRIPLE_QUOTE_STRING() antlr.TerminalNode {
	return s.GetToken(DCellParserTRIPLE_QUOTE_STRING, 0)
}

func (p *DCellParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DCellParserRULE_string)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCellParserSINGLE_QUOTE_STRING:
		localctx = NewSingleQuoteStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(DCellParserSINGLE_QUOTE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserDOUBLE_QUOTE_STRING:
		localctx = NewDoubleQuoteStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(DCellParserDOUBLE_QUOTE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserTRIPLE_QUOTE_STRING:
		localctx = NewTripleQuoteStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(178)
			p.Match(DCellParserTRIPLE_QUOTE_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) CopyAll(ctx *IntegerContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HexIntegerContext struct {
	IntegerContext
}

func NewHexIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexIntegerContext {
	var p = new(HexIntegerContext)

	InitEmptyIntegerContext(&p.IntegerContext)
	p.parser = parser
	p.CopyAll(ctx.(*IntegerContext))

	return p
}

func (s *HexIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexIntegerContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(DCellParserHEX_INTEGER, 0)
}

type DecimalIntegerContext struct {
	IntegerContext
}

func NewDecimalIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalIntegerContext {
	var p = new(DecimalIntegerContext)

	InitEmptyIntegerContext(&p.IntegerContext)
	p.parser = parser
	p.CopyAll(ctx.(*IntegerContext))

	return p
}

func (s *DecimalIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalIntegerContext) DECIMAL_INTEGER() antlr.TerminalNode {
	return s.GetToken(DCellParserDECIMAL_INTEGER, 0)
}

type OctalIntegerContext struct {
	IntegerContext
}

func NewOctalIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OctalIntegerContext {
	var p = new(OctalIntegerContext)

	InitEmptyIntegerContext(&p.IntegerContext)
	p.parser = parser
	p.CopyAll(ctx.(*IntegerContext))

	return p
}

func (s *OctalIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OctalIntegerContext) OCTAL_INTEGER() antlr.TerminalNode {
	return s.GetToken(DCellParserOCTAL_INTEGER, 0)
}

type BinaryIntegerContext struct {
	IntegerContext
}

func NewBinaryIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryIntegerContext {
	var p = new(BinaryIntegerContext)

	InitEmptyIntegerContext(&p.IntegerContext)
	p.parser = parser
	p.CopyAll(ctx.(*IntegerContext))

	return p
}

func (s *BinaryIntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryIntegerContext) BINARY_INTEGER() antlr.TerminalNode {
	return s.GetToken(DCellParserBINARY_INTEGER, 0)
}

func (p *DCellParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DCellParserRULE_integer)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCellParserDECIMAL_INTEGER:
		localctx = NewDecimalIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(181)
			p.Match(DCellParserDECIMAL_INTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserHEX_INTEGER:
		localctx = NewHexIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Match(DCellParserHEX_INTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserOCTAL_INTEGER:
		localctx = NewOctalIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(183)
			p.Match(DCellParserOCTAL_INTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserBINARY_INTEGER:
		localctx = NewBinaryIntegerContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(184)
			p.Match(DCellParserBINARY_INTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatContext is an interface to support dynamic dispatch.
type IFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFloatContext differentiates from other interfaces.
	IsFloatContext()
}

type FloatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatContext() *FloatContext {
	var p = new(FloatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_float
	return p
}

func InitEmptyFloatContext(p *FloatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DCellParserRULE_float
}

func (*FloatContext) IsFloatContext() {}

func NewFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatContext {
	var p = new(FloatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCellParserRULE_float

	return p
}

func (s *FloatContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatContext) CopyAll(ctx *FloatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecimalFloatContext struct {
	FloatContext
}

func NewDecimalFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalFloatContext {
	var p = new(DecimalFloatContext)

	InitEmptyFloatContext(&p.FloatContext)
	p.parser = parser
	p.CopyAll(ctx.(*FloatContext))

	return p
}

func (s *DecimalFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalFloatContext) DECIMAL_FLOAT() antlr.TerminalNode {
	return s.GetToken(DCellParserDECIMAL_FLOAT, 0)
}

type ScientificFloatContext struct {
	FloatContext
}

func NewScientificFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScientificFloatContext {
	var p = new(ScientificFloatContext)

	InitEmptyFloatContext(&p.FloatContext)
	p.parser = parser
	p.CopyAll(ctx.(*FloatContext))

	return p
}

func (s *ScientificFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScientificFloatContext) SCIENTIFIC_FLOAT() antlr.TerminalNode {
	return s.GetToken(DCellParserSCIENTIFIC_FLOAT, 0)
}

func (p *DCellParser) Float() (localctx IFloatContext) {
	localctx = NewFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DCellParserRULE_float)
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DCellParserSCIENTIFIC_FLOAT:
		localctx = NewScientificFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(DCellParserSCIENTIFIC_FLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DCellParserDECIMAL_FLOAT:
		localctx = NewDecimalFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(DCellParserDECIMAL_FLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *DCellParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DCellParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
