package main

import "fmt"

const (
	Tab = iota
	NewLine
	WindowsThing
	Space
	Nil
)

// kind
type Kind string

const (
	KindKeyword  Kind = "keyword"
	KindPlain    Kind = "plain"
	KindOperator Kind = "operator"
	KindNumber   Kind = "number"
	KindGrammar  Kind = "grammar"
)

var kinds = map[string]Kind{
	// keywords
	"package":   KindKeyword,
	"func":      KindKeyword,
	"return":    KindKeyword,
	"import":    KindKeyword,
	"var":       KindKeyword,
	"range":     KindKeyword,
	"for":       KindKeyword,
	"defer":     KindKeyword,
	"switch":    KindKeyword,
	"case":      KindKeyword,
	"struct":    KindKeyword,
	"type":      KindKeyword,
	"interface": KindKeyword,

	// operators
	"=":  KindOperator,
	":=": KindOperator,
	"+":  KindOperator,
	"-":  KindOperator,
	"*":  KindOperator,
	"/":  KindOperator,
	"%":  KindOperator,
	"+=": KindOperator,
	"-=": KindOperator,
	"*=": KindOperator,
	"/=": KindOperator,
	"%=": KindOperator,
	"&":  KindOperator,
	".":  KindOperator,

	// grammar
	"{": KindGrammar,
	"}": KindGrammar,
	"(": KindGrammar,
	")": KindGrammar,
	"[": KindGrammar,
	"]": KindGrammar,
	";": KindGrammar,
	":": KindGrammar,
}

// tokens
type Token struct {
	name string
	kind Kind
}

func (t Token) String() string {
	return fmt.Sprintf("%s: %s", t.name, t.kind)
}
