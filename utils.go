package main

import (
	"errors"
)

func Parse(code string) ([][]Token, error) {
	if code == "" {
		return nil, errors.New("Code is empty")
	}

	word := ""
	tokens := []Token{}
	tokenses := [][]Token{}

	for _, char := range code {
		if space := IsSpace(char); space != Nil {
			kind := kinds[word]
			if kind == "" {
				kind = KindPlain
			}
			token := Token{name: word, kind: kind}
			switch space {
			case NewLine:
				if word == "" {
					continue
				}
				tokens = append(tokens, token)
				tokenses = append(tokenses, tokens)
				tokens = []Token{}
				word = ""
			case Space:
				tokens = append(tokens, token)
				word = ""
			case Tab:
				continue
			}
		} else {
			word += string(char)
		}
	}
	return tokenses, nil
}

func IsSpace(s rune) int {
	switch s {
	case 9: // "\t"
		return Tab
	case 10: // "\n"
		return NewLine
	case 13: // "\r"
		return WindowsThing
	case 32: // " "
		return Space
	}
	return Nil
}
