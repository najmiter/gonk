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
	KindKeyword Kind = "keyword"
	KindPlain   Kind = "plain"
)

var kinds = map[string]Kind{
	"package": KindKeyword,
	"func":    KindKeyword,
	"import":  KindKeyword,
}

// tokens
type Token struct {
	name string
	kind Kind
}

func (t Token) String() string {
	return fmt.Sprintf("%s: %s", t.name, t.kind)
}
